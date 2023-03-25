package domain

import "context"

type Profile struct {
	Campus string `json:"campus"`
}

type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*Profile, error)
}