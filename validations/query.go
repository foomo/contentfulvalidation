package validations

import (
	catvo "github.com/bestbytes/catalogue/vo"
	"github.com/foomo/contentfulvalidation/constants"
)

func ValidateQuery(query *catvo.Query, attributes catvo.Attributes) (constants.QueryError, bool) {

	isValueExpired := func(value string, def catvo.AttributeDefinition) (constants.QueryError, bool) {
		if len(value) < 1 {
			return constants.MissingQueryFieldValues, true
		}
		if _, ok := def.EnumStrings[catvo.AttributeValueID(value)]; !ok {
			return constants.QueryValueExpired, true
		} else {
			return "", false
		}
	}

	areValuesExpired := func(values []string, def catvo.AttributeDefinition) (constants.QueryError, bool) {
		if len(values) < 1 {
			return constants.MissingQueryFieldValues, true
		}
		for _, v := range values {
			if res, ok := isValueExpired(v, def); ok {
				return res, true
			}
		}
		return "", false
	}

	for _, e := range query.Elements {
		errorMessage := constants.QueryError("")
		hasError := false

		if e.Matcher != nil {
			if def, ok := attributes[e.Matcher.Attribute]; ok {
				switch {
				case e.Matcher.StringIn != nil:
					errorMessage, hasError = areValuesExpired(e.Matcher.StringIn.Values, def)
				case e.Matcher.StringAllIn != nil:
					errorMessage, hasError = areValuesExpired(e.Matcher.StringAllIn.Values, def)
				case e.Matcher.StringNotIn != nil:
					errorMessage, hasError = areValuesExpired(e.Matcher.StringNotIn.Values, def)
				case e.Matcher.StringEquals != nil:
					errorMessage, hasError = isValueExpired(e.Matcher.StringEquals.Value, def)
				case e.Matcher.StringNotEquals != nil:
					errorMessage, hasError = isValueExpired(e.Matcher.StringNotEquals.Value, def)
				default:
					errorMessage, hasError = constants.MissingQueryCondition, true
				}
			} else {
				return constants.MissingQueryField, true
			}
		}

		if hasError {
			return errorMessage, hasError
		}
	}
	return "", false
}
