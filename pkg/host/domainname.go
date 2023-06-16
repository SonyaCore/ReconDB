package host

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// CheckDomain returns an error if the domain input is not valid
// See https://tools.ietf.org/html/rfc1034#section-3.5 and
// https://tools.ietf.org/html/rfc1123#section-2.
func CheckDomain(input string) error {
	if strings.Contains(input, ":") {
		parts := strings.Split(input, ":")
		input = parts[0]
		port := parts[1]
		if err := Port(port); err != nil {
			return fmt.Errorf("%v", err.Error())
		}
	}
	switch {
	case len(input) == 0:
		return nil // an empty domain input will result in a cookie without a domain restriction
	case len(input) > 255:
		return fmt.Errorf("cookie domain: input length is %d, can't exceed 255", len(input))
	}
	var l int
	for i := 0; i < len(input); i++ {
		b := input[i]
		if b == '.' {
			// check domain labels validity
			switch {
			case i == l:
				return fmt.Errorf("cookie domain: invalid character '%c' at offset %d: label can't begin with a period", b, i)
			case i-l > 63:
				return fmt.Errorf("cookie domain: byte length of label '%s' is %d, can't exceed 63", input[l:i], i-l)
			case input[l] == '-':
				return fmt.Errorf("cookie domain: label '%s' at offset %d begins with a hyphen", input[l:i], l)
			case input[i-1] == '-':
				return fmt.Errorf("cookie domain: label '%s' at offset %d ends with a hyphen", input[l:i], l)
			}
			l = i + 1
			continue
		}
		// test label character validity, note: tests are ordered by decreasing validity frequency
		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b == '-' || b >= 'A' && b <= 'Z') {
			// show the printable unicode character starting at byte offset i
			c, _ := utf8.DecodeRuneInString(input[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("cookie domain: invalid rune at offset %d", i)
			}
			return fmt.Errorf("cookie domain: invalid character '%c' at offset %d", c, i)
		}
	}
	// check top level domain validity
	switch {
	case l == len(input):
		return fmt.Errorf("cookie domain: missing top level domain, domain can't end with a period")
	case len(input)-l > 63:
		return fmt.Errorf("cookie domain: byte length of top level domain '%s' is %d, can't exceed 63", input[l:], len(input)-l)
	case input[l] == '-':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d begins with a hyphen", input[l:], l)
	case input[len(input)-1] == '-':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d ends with a hyphen", input[l:], l)
	case input[l] >= '0' && input[l] <= '9':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d begins with a digit", input[l:], l)
	}
	return nil
}

func CheckWildCard(input string) error {
	var l int
	switch {
	case len(input) == 0:
		return nil // an empty domain input will result in a cookie without a domain restriction
	case len(input) > 255:
		return fmt.Errorf("cookie domain: input length is %d, can't exceed 255", len(input))
	}
	for i := 0; i < len(input); i++ {
		b := input[i]
		if b == '.' {
			// check domain labels validity
			switch {
			case i == l:
				return fmt.Errorf("cookie domain: invalid character '%c' at offset %d: label can't begin with a period", b, i)
			case i-l > 63:
				return fmt.Errorf("cookie domain: byte length of label '%s' is %d, can't exceed 63", input[l:i], i-l)
			case input[l] == '-':
				return fmt.Errorf("cookie domain: label '%s' at offset %d begins with a hyphen", input[l:i], l)
			case input[i-1] == '-':
				return fmt.Errorf("cookie domain: label '%s' at offset %d ends with a hyphen", input[l:i], l)
			}
			l = i + 1
			continue
		}
		// test label character validity, note: tests are ordered by decreasing validity frequency
		if !(b >= 'a' && b <= 'z' || b >= '0' && b <= '9' || b == '-' || b == '*' || b >= 'A' && b <= 'Z') {
			// show the printable unicode character starting at byte offset i
			c, _ := utf8.DecodeRuneInString(input[i:])
			if c == utf8.RuneError {
				return fmt.Errorf("cookie domain: invalid rune at offset %d", i)
			}
			return fmt.Errorf("cookie domain: invalid character '%c' at offset %d", c, i)
		}
	}

	// check top level domain validity
	switch {
	case l == len(input):
		return fmt.Errorf("cookie domain: missing top level domain, domain can't end with a period")
	case len(input)-l > 63:
		return fmt.Errorf("cookie domain: byte length of top level domain '%s' is %d, can't exceed 63", input[l:], len(input)-l)
	case input[l] == '-':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d begins with a hyphen", input[l:], l)
	case input[len(input)-1] == '-':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d ends with a hyphen", input[l:], l)
	case input[l] >= '0' && input[l] <= '9':
		return fmt.Errorf("cookie domain: top level domain '%s' at offset %d begins with a digit", input[l:], l)
	}
	// test wildcard contain in string
	if !strings.Contains(input, "*") {
		return fmt.Errorf("not a valid wildcard")
	}
	return nil
}
