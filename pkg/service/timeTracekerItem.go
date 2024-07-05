package service

import (
	timetracker "time-tracker"
	"time-tracker/pkg/repository"
)

type ItemService struct {
	repo repository.TimeTrackerItem
}

func NewTimeTrackerItemService(repo repository.TimeTrackerItem) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) CreateItem(item timetracker.TimeTrackerItem) (int, error) {
	return s.repo.CreateItem(item)
}

func (s *ItemService) GetItemsById(userId int) ([]timetracker.TimeTrackerItem, error) {
	return s.repo.GetItemsById(userId)
}

func (s *ItemService) DeleteItem(id int) error {
	return s.repo.DeleteItem(id)
}

func (s *ItemService) UpdateItem(id int, input timetracker.UpdateItemInput) error {
	return s.repo.UpdateItem(id, input)
}

func (s *ItemService) UpdateItemTime(id int, flag bool) error {
	return s.repo.UpdateItemTime(id, flag)
}

func (s *ItemService) GetItemsByDate(userId int, datePeriod timetracker.DatePeriod) ([]timetracker.GetItemsByDate, error) {
	return s.repo.GetItemsByDate(userId, datePeriod)
}
