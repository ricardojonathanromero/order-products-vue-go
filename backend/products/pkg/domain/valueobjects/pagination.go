package valueobjects

type Filters struct {
	Limit    int64  `param:"limit"`
	Offset   int64  `param:"offset"`
	FilterBy string `param:"filter_by"`
	SortBy   string `param:"sort_by"`
	OrderBy  string `param:"order_by"`
}
