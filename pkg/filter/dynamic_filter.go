package filter

type FilterOperator string

const (
	OpContains           FilterOperator = "contains"
	OpNotContains        FilterOperator = "notContains"
	OpEquals             FilterOperator = "equals"
	OpNotEqual           FilterOperator = "notEqual"
	OpStartsWith         FilterOperator = "startsWith"
	OpEndsWith           FilterOperator = "endsWith"
	OpLessThan           FilterOperator = "lessThan"
	OpLessThanOrEqual    FilterOperator = "lessThanOrEqual"
	OpGreaterThan        FilterOperator = "greaterThan"
	OpGreaterThanOrEqual FilterOperator = "greaterThanOrEqual"
	OpInRange            FilterOperator = "inRange"
)

type FilterDataType string

const (
	DataTypeText   FilterDataType = "text"
	DataTypeNumber FilterDataType = "number"
	DataTypeDate   FilterDataType = "date"
)

type Sort struct {
	Column string `json:"column"`
	Order  string `json:"orderBy"`
}

type Filter struct {
	Type       FilterOperator `json:"type"`
	From       string         `json:"from"`
	To         string         `json:"to"`
	FilterType FilterDataType `json:"filterType"`
}

type DynamicFilter struct {
	Sort   []Sort            `json:"sort"`
	Search string            `json:"search" form:"search"`
	Filter map[string]Filter `json:"filter"`
}
