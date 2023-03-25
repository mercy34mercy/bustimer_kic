package domain

import "context"

type Busstop struct {
	Name string `json:"name"`
}

type BusstopUsecase interface {
	GetBusstop(c context.Context, userID string) (*Busstop,error)
}