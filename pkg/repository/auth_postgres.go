package repository

import (
	"fmt"
	"strconv"
	"strings"
	timetracker "time-tracker"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user timetracker.User) (int, error) {
	logrus.Debug("Creating user")
	var id int
	query := fmt.Sprintf("INSERT INTO %s (pasportNumber, name) values ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.PasportNumber, user.Name)
	if err := row.Scan(&id); err != nil {
		logrus.WithError(err).Error("Failed to create user")
		return 0, err
	}
	logrus.WithField("id", id).Info("User created successfully")
	return id, nil
}

func (r *AuthPostgres) GetAllUsers() ([]timetracker.User, error) {
	logrus.Debug("Fetching all users")
	var userList []timetracker.User
	query := fmt.Sprintf("SELECT * FROM %s", usersTable)
	err := r.db.Select(&userList, query)
	if err != nil {
		logrus.WithError(err).Error("Failed to fetch all users")
		return nil, err
	}
	logrus.WithField("count", len(userList)).Info("Fetched all users successfully")
	return userList, err
}

func (r *AuthPostgres) DeleteUser(id int) error {
	logrus.WithField("id", id).Debug("Deleting user")
	query := fmt.Sprintf("DELETE FROM %s where id = $1", usersTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		logrus.WithError(err).Error("Failed to delete user")
		return err
	}
	logrus.WithField("id", id).Info("User deleted successfully")
	return nil
}

func (r *AuthPostgres) UpdateUser(id int, input timetracker.UpdateUserInput) error {
	logrus.WithField("id", id).Debug("Updating user")
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.PasportNumber != nil {
		setValues = append(setValues, fmt.Sprintf("pasportnumber=$%d", argId))
		args = append(args, *input.PasportNumber)
		argId++
	}

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if argId > 1 {
		setQuery := strings.Join(setValues, ", ")
		query := fmt.Sprintf("update %s set %s where id = $%d", usersTable, setQuery, argId)
		args = append(args, id)
		_, err := r.db.Exec(query, args...)
		if err != nil {
			logrus.WithError(err).Error("Failed to update user")
			return err
		}
		logrus.WithField("id", id).Info("User updated successfully")
	}

	return nil
}

func (r *AuthPostgres) Search(nameFilter, passportNumberFilter string, page, pageSize int) ([]timetracker.User, error) {
	logrus.Debug("Searching users")
	var users []timetracker.User
	query := "SELECT * FROM users WHERE 1=1"
	var args []interface{}

	if nameFilter != "" {
		query += " AND name ILIKE '%' || $1 || '%'"
		args = append(args, nameFilter)
	}

	if passportNumberFilter != "" {
		if nameFilter == "" {
			query += " AND pasportNumber ILIKE '%' || $1 || '%'"
		} else {
			query += " AND pasportNumber ILIKE '%' || $2 || '%'"
		}
		args = append(args, passportNumberFilter)
	}

	offset := (page - 1) * pageSize
	query += " ORDER BY id LIMIT $"
	args = append(args, pageSize)
	query += strconv.Itoa(len(args))
	query += " OFFSET $"
	args = append(args, offset)
	query += strconv.Itoa(len(args))

	err := r.db.Select(&users, query, args...)
	if err != nil {
		logrus.WithError(err).Error("Failed to search users")
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"count":    len(users),
		"page":     page,
		"pageSize": pageSize,
	}).Info("Users search completed successfully")
	return users, nil
}
