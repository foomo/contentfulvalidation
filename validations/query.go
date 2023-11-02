package validations

import (
	"github.com/foomo/contentfulvalidation/constants"
)

func ValidateQuery(query *constants.Query, attributes constants.Attributes) []constants.QueryError {
	errors := []constants.QueryError{}

	isValueExpired := func(value string, def constants.AttributeDefinition) {
		if len(value) < 1 {
			errors = append(errors, constants.MissingQueryFieldValues)
		} else {
			if _, ok := def.EnumStrings[constants.AttributeValueID(value)]; !ok {
				errors = append(errors, constants.QueryValueExpired)
			}
		}
	}

	areValuesExpired := func(values []string, def constants.AttributeDefinition) {
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
				case e.Matcher.IntInRange != nil:
				case e.Matcher.IntFrom != nil:
				case e.Matcher.IntTo != nil:
				case e.Matcher.IntEquals != nil:
				case e.Matcher.IntNotEquals != nil:
				case e.Matcher.BoolEquals != nil:
				case e.Matcher.Bitmap != nil:
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
