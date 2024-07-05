package service

import (
	timetracker "time-tracker"
	"time-tracker/pkg/repository"
)

type Authorisation interface {
	CreateUser(user timetracker.User) (int, error)
	GetAllUsers() ([]timetracker.User, error)
	DeleteUser(id int) error
	UpdateUser(id int, input timetracker.UpdateUserInput) error
	Search(nameFilter, passportNumberFilter string, page, pageSize int) ([]timetracker.User, error)
}

type TimeTrackerItem interface {
	CreateItem(item timetracker.TimeTrackerItem) (int, error)
	GetItemsById(userId int) ([]timetracker.TimeTrackerItem, error)
	DeleteItem(id int) error
	UpdateItem(id int, input timetracker.UpdateItemInput) error
	UpdateItemTime(id int, flag bool) error
	GetItemsByDate(userId int, datePeriod timetracker.DatePeriod) ([]timetracker.GetItemsByDate, error)
}

type Service struct {
	Authorisation
	TimeTrackerItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation:   NewAuthService(repos.Authorisation),
		TimeTrackerItem: NewTimeTrackerItemService(repos.TimeTrackerItem),
	}
}
