package entity

import "time"

type Segment struct {
	ID        int
	Name      string
	DeletedAt *time.Time
}
