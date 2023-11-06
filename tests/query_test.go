package tests

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/foomo/contentfulvalidation/constants"
	"github.com/foomo/contentfulvalidation/validations"

	// testingx "github.com/foomo/go/testing"
	"github.com/stretchr/testify/assert"
)

//go:embed data/query.json
var queryData []byte

//go:embed data/attributes.json
var attributesData []byte

func getMockAttributes() constants.Attributes {
	var attributes constants.Attributes
	err := json.Unmarshal(attributesData, &attributes)
	if err != nil {
		return constants.Attributes{}
	}
	return attributes
}

func getMockQuery() constants.Query {
	var query constants.Query
	err := json.Unmarshal(queryData, &query)
	if err != nil {
		return constants.Query{}
	}
	return query
}

func Test_Query(t *testing.T) {
	// testingx.Tags(t)

	query := getMockQuery()
	testResult := []constants.QueryError{"Query Field is empty", "Query field value is expired", "Missing field values", "Missing query condition", "Query Field is empty"}
	queryErrors := validations.ValidateQuery(&query, getMockAttributes())
	fmt.Println("queryErrors", queryErrors)

	assert.Equal(t, testResult, queryErrors)
}
