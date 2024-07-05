package timetracker

import "time"

type TimeTrackerItem struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	TimeStart time.Time `json:"timeStart" db:"timestart"`
	TimeStop  time.Time `json:"timeStop" db:"timestop"`
	TimeUsed  string    `json:"timeUsed" db:"timeused"`
	UserId    int       `json:"userId" binding:"required"`
}

type UpdateItemInput struct {
	Name *string `json:"name"`
}
type GetItemsByDate struct {
	Name     string `json:"name"`
	TimeUsed string `json:"timeUsed"`
}
type DatePeriod struct {
	TimeStart time.Time `json:"timeStart"`
	TimeStop  time.Time `json:"timeStop"`
}

type CreateItemInput struct {
	Name   string `json:"name" binding:"required" example:"Завтрак"`
	UserId int    `json:"userId" binding:"required"`
}
