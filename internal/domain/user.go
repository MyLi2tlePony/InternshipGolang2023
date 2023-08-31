package domain

type ChangeUserSegments struct {
	InsertSegments []Segment `json:"insertSegments"`
	DeleteSegments []Segment `json:"deleteSegments"`
}
