package service

import (
	"errors"
	"time"
)

type DaysService struct{}

func NewDaysService() *DaysService {
	return &DaysService{}
}

const jan = time.January

func (ds *DaysService) Left(now time.Time) (int, error) {
	until := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	daysLeft := int(until.Sub(time.Now()).Hours() / 24)

	if daysLeft < 0 {
		return 0, errors.New("time left!")
	}

	return daysLeft, nil
}
