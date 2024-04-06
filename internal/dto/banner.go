package dto

type GetBanner struct {
	TagId     int
	FeatureId int
	Limit     int
	Offset    int
}

type Banner struct {
	TagId     int
	FeatureId int
	Content   any
}
