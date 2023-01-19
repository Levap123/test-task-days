package service

import "time"

type Service struct {
	Days
}

type Days interface {
	Left(now time.Time) (int, error)
}

func NewService() *Service {
	return &Service{
		Days: NewDaysService(),
	}
}
