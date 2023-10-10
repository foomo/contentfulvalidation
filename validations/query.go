package validations

import (
	catvo "github.com/bestbytes/catalogue/vo"
	"github.com/foomo/contentfulvalidation/constants"
)

func ValidateQuery(query *catvo.Query, attributes catvo.Attributes) []constants.QueryError {
	errors := []constants.QueryError{}

	isValueExpired := func(value string, def catvo.AttributeDefinition) {
		if len(value) < 1 {
			errors = append(errors, constants.MissingQueryFieldValues)
		} else {
			if _, ok := def.EnumStrings[catvo.AttributeValueID(value)]; !ok {
				errors = append(errors, constants.QueryValueExpired)
			}
		}
	}

	areValuesExpired := func(values []string, def catvo.AttributeDefinition) {
		if len(values) < 1 {
			errors = append(errors, constants.MissingQueryFieldValues)
		}
		for _, v := range values {
			isValueExpired(v, def)
		}
	}

	for _, e := range query.Elements {
		if e.Matcher != nil {
			if def, ok := attributes[e.Matcher.Attribute]; ok {
				switch {
				case e.Matcher.StringIn != nil:
					areValuesExpired(e.Matcher.StringIn.Values, def)
				case e.Matcher.StringAllIn != nil:
					areValuesExpired(e.Matcher.StringAllIn.Values, def)
				case e.Matcher.StringNotIn != nil:
					areValuesExpired(e.Matcher.StringNotIn.Values, def)
				case e.Matcher.StringEquals != nil:
					isValueExpired(e.Matcher.StringEquals.Value, def)
				case e.Matcher.StringNotEquals != nil:
					isValueExpired(e.Matcher.StringNotEquals.Value, def)
				default:
					errors = append(errors, constants.MissingQueryCondition)
				}
			} else {
				errors = append(errors, constants.MissingQueryField)
			}
		}
	}

	return errors
}
