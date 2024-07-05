package repository

import (
	"fmt"
	"strings"
	"time"
	timetracker "time-tracker"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ItemPostgres struct {
	db *sqlx.DB
}

func NewItemPostgres(db *sqlx.DB) *ItemPostgres {
	return &ItemPostgres{db: db}
}

func (r *ItemPostgres) CreateItem(item timetracker.TimeTrackerItem) (int, error) {
	logrus.Debug("Creating item")
	var id int
	item.TimeStart = time.Now()
	item.TimeStop = time.Now()
	item.TimeUsed = "0 часов 0 минут"
	query := fmt.Sprintf("INSERT INTO %s (name, userId, timeStart, timeStop, timeUsed) values ($1, $2, $3, $4, $5) RETURNING id", timeTrackerItemsTable)
	row := r.db.QueryRow(query, item.Name, item.UserId, item.TimeStart, item.TimeStop, "")
	if err := row.Scan(&id); err != nil {
		logrus.WithError(err).Error("Failed to create item")
		return 0, err
	}
	logrus.WithField("id", id).Info("Item created successfully")
	return id, nil
}

func (r *ItemPostgres) GetItemsById(userId int) ([]timetracker.TimeTrackerItem, error) {
	logrus.WithField("userId", userId).Debug("Fetching items by user ID")
	var items []timetracker.TimeTrackerItem
	query := fmt.Sprintf("select * from %s where userId = $1", timeTrackerItemsTable)
	err := r.db.Select(&items, query, userId)
	if err != nil {
		logrus.WithError(err).Error("Failed to fetch items by user ID")
		return nil, err
	}
	logrus.WithField("count", len(items)).Info("Fetched items by user ID successfully")
	return items, err
}

func (r *ItemPostgres) DeleteItem(id int) error {
	logrus.WithField("id", id).Debug("Deleting item")
	query := fmt.Sprintf("DELETE FROM %s where id = $1", timeTrackerItemsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete item")
		return err
	}
	logrus.WithField("id", id).Info("Item deleted successfully")
	return nil
}

func (r *ItemPostgres) UpdateItem(id int, input timetracker.UpdateItemInput) error {
	logrus.WithField("id", id).Debug("Updating item")
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if argId > 1 {
		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("update %s set %s where id = $%d", timeTrackerItemsTable, setQuery, argId)
		args = append(args, id)
		_, err := r.db.Exec(query, args...)
		if err != nil {
			logrus.WithError(err).Error("Failed to update item")
			return err
		}
		logrus.WithField("id", id).Info("Item updated successfully")
	}
	return nil
}

func (r *ItemPostgres) UpdateItemTime(id int, flag bool) error {
	logrus.WithFields(logrus.Fields{"id": id, "flag": flag}).Debug("Updating item time")
	if flag {
		query := fmt.Sprintf("update %s set timestart = $1 where id = $2", timeTrackerItemsTable)
		_, err := r.db.Exec(query, time.Now(), id)
		if err != nil {
			logrus.WithError(err).Error("Failed to update item start time")
			return err
		}
		logrus.WithField("id", id).Info("Item start time updated successfully")
	} else {
		query := fmt.Sprintf(`UPDATE %s 
		SET timestop = $1, 
			timeUsed = (
				(EXTRACT(EPOCH FROM AGE($1, timestart)) / 3600)::int || ' часов ' || 
				((EXTRACT(EPOCH FROM AGE($1, timestart)) / 60)::int %% 60) || ' минут'
			)
		WHERE id = $2`, timeTrackerItemsTable)

		_, err := r.db.Exec(query, time.Now(), id)
		if err != nil {
			logrus.WithError(err).Error("Failed to update item stop time")
			return err
		}
		logrus.WithField("id", id).Info("Item stop time updated successfully")
	}
	return nil
}

func (r *ItemPostgres) GetItemsByDate(userId int, datePeriod timetracker.DatePeriod) ([]timetracker.GetItemsByDate, error) {
	logrus.WithFields(logrus.Fields{"userId": userId, "datePeriod": datePeriod}).Debug("Fetching items by date")
	var items []timetracker.GetItemsByDate
	query := fmt.Sprintf("select name, timeused from %s where userId = $1 and timeStart between $2 and $3 order by  (substring(timeUsed from '(\\d+) часов')::int * 60) + substring(timeUsed from '(\\d+) минут')::int DESC;", timeTrackerItemsTable)
	err := r.db.Select(&items, query, userId, datePeriod.TimeStart, datePeriod.TimeStop)
	if err != nil {
		logrus.WithError(err).Error("Failed to fetch items by date")
		return nil, err
	}
	logrus.WithFields(logrus.Fields{"count": len(items), "userId": userId}).Info("Fetched items by date successfully")
	return items, err
}
