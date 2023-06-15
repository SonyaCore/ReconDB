package company

import (
	"ReconDB/database"
	"ReconDB/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func AddCompany(c *gin.Context) {
	var Company models.Company

	c.ShouldBindJSON(&Company)

	// checking empty fields for company or program typeassert
	if Company.CompanyName == "" || Company.ProgramType == "" {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"error":  "empty fields not allowed",
			"status": http.StatusFailedDependency,
		})
		return
	}

	// insert company to db
	collection := database.Collection("Company")
	result, err := collection.InsertOne(database.Ctx, Company)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{
			"error":  err.Error(),
			"status": http.StatusFailedDependency,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"collectionID": result.InsertedID,
		"message":      "company added",
		"result":       Company,
		"status":       http.StatusOK,
	})
}

func GetAllCompanies(c *gin.Context) {
	var ctx = context.TODO()
	var Company []bson.M

	collection := database.Collection("Company")
	results, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Print(err.Error())
	}

	if err = results.All(ctx, &Company); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"companies": Company,
		"status":    http.StatusOK,
	})
}

func GetCompany(c *gin.Context) {
	var Param = c.Param("companyname")
	var ctx = context.TODO()
	var Company []bson.M

	collection := database.Collection("Company")
	filter, err := collection.Find(ctx, bson.M{"companyname": Param})
	if err != nil {
		log.Print(err.Error())
	}

	if err = filter.All(ctx, &Company); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"companies": Company,
		"status":    http.StatusOK,
	})
}

func DeleteCompany(c *gin.Context) {
	var Param = c.Param("companyname")
	var ctx = context.TODO()

	collection := database.Collection("Company")
	filter, err := collection.DeleteMany(ctx, bson.M{"companyname": Param})
	if err != nil {
		log.Print(err.Error())
	}

	// deleting scopes related to company name
	scope := database.Collection("Scopes")
	deleteScopes, _ := scope.DeleteMany(ctx, bson.M{"companyname": Param})

	// deleting out of scopes relate to company name
	Outscope := database.Collection("OutofScopes")
	deleteOutScopes, _ := Outscope.DeleteMany(ctx, bson.M{"companyname": Param})

	// deleting assets relate to company name
	Assets := database.Collection("Assets")
	deleteAssets, _ := Assets.DeleteMany(ctx, bson.M{"companyname": Param})

	c.JSON(http.StatusOK, gin.H{
		"company_name":    Param,
		"deleted_count":   filter.DeletedCount,
		"scopes_count":    deleteScopes.DeletedCount,
		"out_Scope_count": deleteOutScopes.DeletedCount,
		"assets_count":    deleteAssets.DeletedCount,
		"status":          http.StatusOK,
	})
}
