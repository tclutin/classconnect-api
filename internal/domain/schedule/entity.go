package schedule

import "time"

type Schedule struct {
	ID        uint64
	SubjectID *uint64
	DayNumber int
	IsEven    bool
}

type Subject struct {
	ID          uint64
	Teacher     string
	Name        string
	Cabinet     string
	Description string
	TimeStart   time.Time
	TimeEnd     time.Time
}
