package repository

import (
	timetracker "time-tracker"

	"github.com/jmoiron/sqlx"
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
	GetItemsByDate(userId int, DatePeriod timetracker.DatePeriod) ([]timetracker.GetItemsByDate, error)
}

type Repository struct {
	Authorisation
	TimeTrackerItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation:   NewAuthPostgres(db),
		TimeTrackerItem: NewItemPostgres(db),
	}
}
