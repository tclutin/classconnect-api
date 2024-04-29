package schedule

type UploadScheduleRequest struct {
	Weeks []Week `json:"weeks" binding:"required"`
}

type Week struct {
	IsEven bool  `json:"is_even" binding:"required"`
	Days   []Day `json:"days" binding:"required"`
}

type Day struct {
	WeekNumber int       `json:"week_number" binding:"required"`
	Subjects   []Subject `json:"subjects" binding:"required"`
}

type Subject struct {
	Name        string `json:"name" binding:"required"`
	Cabinet     string `json:"cabinet" binding:"required"`
	Teacher     string `json:"teacher" binding:"required"`
	Description string `json:"description" binding:"required"`
	StartTime   string `json:"start_time" binding:"required"`
	EndTime     string `json:"end_time" binding:"required"`
}
