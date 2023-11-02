package constants

import (
	"github.com/RoaringBitmap/roaring"
)

type Severity string
type Health string
type QueryError string

type (
	AttributeID      string
	AttributeValueID string
	AttributeType    string

	// Locale like en or en_us
	Locale          string
	LocalizedString map[Locale]string
)

type Attributes map[AttributeID]AttributeDefinition

// AttributeDefinition describes an attribute.
type AttributeDefinition struct {
	ID          AttributeID                        `json:"id"`
	Type        AttributeType                      `json:"type"`
	EnumStrings map[AttributeValueID]*string       `json:"enumStrings,omitempty"`
	Meta        AttributeMeta                      `json:"meta,omitempty"`
	MetaValues  map[AttributeValueID]AttributeMeta `json:"metaValues,omitempty"`
	StepSize    int                                `json:"stepSize,omitempty"`
	Mandatory   bool                               `json:"mandatory,omitempty"`
}

// AttributeMeta models meta information for an attribute.
type AttributeMeta struct {
	Label       LocalizedString   `json:"label,omitempty"`
	Description LocalizedString   `json:"description,omitempty"`
	Custom      map[string]string `json:"custom,omitempty"`
	SortingRank map[string]int    `json:"sortingRank,omitempty"`
}

// Query structure
type Query struct {
	Explanation string          `json:"explanation,omitempty"`
	Operation   Operation       `json:"operation"`
	Elements    []*QueryElement `json:"elements"`
}

// Operation defines how to compare bitmaps
type Operation string

// QueryElement structure
type QueryElement struct {
	Matcher *Matcher `json:"matcher,omitempty"`
	Query   *Query   `json:"query,omitempty"`
}

// Matcher structure
// Identifies a bitmap
// the result of each match operation is a bitmap of entity ids
type Matcher struct {
	Attribute   AttributeID `json:"attribute,omitempty"`
	Explanation string      `json:"explanation,omitempty"`

	// strings
	StringIn        *StringIn        `json:"stringIn,omitempty"`
	StringAllIn     *StringAllIn     `json:"stringAllIn,omitempty"`
	StringNotIn     *StringNotIn     `json:"stringNotIn,omitempty"`
	StringEquals    *StringEquals    `json:"stringEquals,omitempty"`
	StringNotEquals *StringNotEquals `json:"stringNotEquals,omitempty"`

	// integers
	IntInRange   *IntInRange   `json:"intInRange,omitempty"`
	IntFrom      *IntFrom      `json:"intFrom,omitempty"`
	IntTo        *IntTo        `json:"intTo,omitempty"`
	IntEquals    *IntEquals    `json:"intEquals,omitempty"`
	IntNotEquals *IntNotEquals `json:"intNotEquals,omitempty"`

	// booleans
	BoolEquals *BoolEquals `json:"boolEquals,omitempty"`

	// bitmap
	Bitmap *Bitmap `json:"bitmap,omitempty"`
}

// StringIn matches if the input value equals any of the strings specified
type StringIn struct {
	Values []string `json:"values"`
}

// StringAllIn matches if all values appear in the input
type StringAllIn struct {
	Values []string `json:"values"`
}

// StringNotIn matches if the input value does not equal any of the strings specified
type StringNotIn struct {
	Values []string `json:"values"`
}

// StringEquals matches strings that DO equal the supplied value
type StringEquals struct {
	Value string `json:"value"`
}

// StringNotEquals matches strings that DO NOT equal the supplied value
type StringNotEquals struct {
	Value string `json:"value"`
}

// IntInRange matches integers in the given range
type IntInRange struct {
	From int `json:"from"`
	To   int `json:"to"`
}

// IntFrom matches integers starting from the given value (>=)
type IntFrom struct {
	From int `json:"from"`
}

// IntTo matches integers until the given value (<=)
type IntTo struct {
	To int `json:"to"`
}

// IntEquals matches integers exactly (==)
type IntEquals struct {
	Value int `json:"value"`
}

// IntNotEquals matches integers that do not equal the given value (!=)
type IntNotEquals struct {
	Value int `json:"value"`
}

// BoolEquals matches booleans exactly (==)
type BoolEquals struct {
	Value bool `json:"value"`
}

// Bitmap allows to use a *roaring.Bitmap directly for a matcher
type Bitmap struct {
	// value is private to hide it from gotsrpc
	value *roaring.Bitmap `json:"-"`

	FacetValue string `json:"facetValue"`
}
