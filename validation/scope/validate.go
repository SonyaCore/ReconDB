package scope

import (
	"ReconDB/config"
	"ReconDB/database"
	"ReconDB/models"
	"ReconDB/utils"
	"ReconDB/validation"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"regexp"
)

// errors
var CompanyNotRegister = "Scope are not registered for this company"
var DuplicateEntry = "Duplicate Entry"

// ValidateScopeType host the typeassert of scope with validation.Scopes
// if it was not in scopes []string c.Abort with error message otherwise c.Next()
func ValidateScopeType(c *gin.Context) {
	var Scope models.Scopes

	rawBody, err := utils.ReadBuffer(c)
	if err != nil {
		log.Println(err)
		return
	}

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	for i, _ := range validation.Scopes {
		if Scope.ScopeType == validation.Scopes[i] {
			c.Next()
			return
		}
		continue
	}
	c.JSON(http.StatusFailedDependency, gin.H{
		"error":       "Scope Type is not valid",
		"valid_types": validation.Scopes,
		"status":      http.StatusFailedDependency,
	})
	c.Abort()
	return
}

func OutScopeCheck(c *gin.Context) {
	var Scope models.Scopes
	var configuration config.RouterURI
	var results int64

	configuration, _ = config.RouterConfig()

	rawBody, err := utils.ReadBuffer(c)

	// Unmarshal rawBody to Scope
	err = json.Unmarshal(rawBody, &Scope)
	if err != nil {
		log.Printf(err.Error())
		return
	}

	ScopeQuery := bson.M{
		"companyname": Scope.CompanyName,
		"scopetype":   Scope.ScopeType,
		"scope": primitive.Regex{
			Pattern: "^" + regexp.QuoteMeta(Scope.Scope) + "$",
			Options: "i",
		},
	}

	// only use this section if request uri was /api/scope
	if c.Request.RequestURI == configuration.API+"/"+configuration.Scope {
		companyQuery := bson.M{
			"companyname": Scope.CompanyName,
		}
		results, err = CompanyCheck(companyQuery)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.CompanyName,
				"result": err.Error(),
				"status": http.StatusNotAcceptable,
			})
		}
		results, err = DuplicateCheck(ScopeQuery)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"input":  Scope.Scope,
				"result": err.Error(),
				"status": http.StatusNotAcceptable,
			})
			return
		}
	}

	results, err = database.CountDocuments("OutofScopes", ScopeQuery)
	if err != nil {
		log.Print(err.Error())
	}

	if results >= 1 {
		if c.Request.RequestURI == configuration.API+"/"+configuration.Scope {
			c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
				"companyname": Scope.CompanyName,
				"result":      "Out of Scope",
				"status":      http.StatusNotAcceptable,
			})
			return
		}

		// outscope uri
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"scope":  Scope.Scope,
			"result": DuplicateEntry,
			"status": http.StatusNotAcceptable,
		})
		return
	}
	c.Next()
}

func CompanyCheck(query bson.M) (int64, error) {
	count, err := database.CountDocuments("Company", query)
	if err != nil {
		return 0, err
	}
	if count == 0 {
		return 0, fmt.Errorf(CompanyNotRegister)
	}
	return 1, nil
}

func DuplicateCheck(query bson.M) (int64, error) {
	count, err := database.CountDocuments("Scopes", query)
	if err != nil {
		return 1, err
	}
	if count >= 1 {
		return count, fmt.Errorf("duplicate Entry")
	}

	return 0, nil
}
