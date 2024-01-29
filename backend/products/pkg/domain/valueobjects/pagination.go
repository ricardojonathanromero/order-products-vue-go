package valueobjects

type Filters struct {
	Limit    int64  `query:"limit" validate:"max=20" example:"10"`
	Offset   int64  `query:"offset" validate:"omitempty" example:"0"`
	FilterBy string `query:"filter_by" validate:"max=100" example:"catalog=soda"`
	SortBy   string `query:"sort_by" validate:"max=100" example:"-_id"`
}
