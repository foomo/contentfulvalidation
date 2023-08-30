package validations

import (
	"fmt"

	catvo "github.com/bestbytes/catalogue/vo"
)

func IsAttributeExpired(query *catvo.Query, attributes catvo.Attributes) bool {

	// TODO is this ok?
	isValueExpired := func(value string, def catvo.AttributeDefinition) bool {
		if _, ok := def.EnumStrings[catvo.AttributeValueID(value)]; !ok {
			// thow error
			// provide contex on value with this id ....
			fmt.Println("Attribute NOT found: ", ok)
			return true

		} else {
			fmt.Println("Attribute found: ", ok)
			return false
		}
	}

	areValuesExpired := func(values []string, def catvo.AttributeDefinition) bool {
		expired := false
		for _, v := range values {
			if isValueExpired(v, def) {
				expired = true
			}
		}
		return expired
	}

	for _, e := range query.Elements {
		fmt.Println("THE Matcher: ", e.Matcher)

		// @TODO validate if there is even an attribute set, or is empty string
		if e.Matcher != nil {
			if def, ok := attributes[e.Matcher.Attribute]; ok {
				switch {
				case e.Matcher.StringIn != nil:
					return areValuesExpired(e.Matcher.StringIn.Values, def)
				case e.Matcher.StringAllIn != nil:
					return areValuesExpired(e.Matcher.StringAllIn.Values, def)
				case e.Matcher.StringNotIn != nil:
					return areValuesExpired(e.Matcher.StringNotIn.Values, def)
				case e.Matcher.StringEquals != nil:
					return isValueExpired(e.Matcher.StringEquals.Value, def)
				case e.Matcher.StringNotEquals != nil:
					return isValueExpired(e.Matcher.StringNotEquals.Value, def)
				}
			} else {
				// throw error
				fmt.Println("NO attribute within ALL ATTRIBUTES: ")
				// TODO uncomment once catalogue attr are in, maybe a different validation for this
				// return true
			}

		} else {
			fmt.Println("MATCHER is NIL e.matcher: ")
			return true
		}
	}
	return false
}
