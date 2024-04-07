package dto

import "time"

type GetBanner struct {
	TagId     int
	FeatureId int
	Limit     int
	Offset    int
}

type Banner struct {
	BannerId  int
	TagIds    []int
	FeatureId int
	Content   any
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PostBanner struct {
	TagIds    []int
	FeatureId int
	Content   any
	IsActive  bool
}
