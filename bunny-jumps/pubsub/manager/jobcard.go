package main

import (
	"github.com/google/uuid"
)

type JobCard struct {
	UUID     uuid.UUID `json:"uuid"`
	Title    string    `json:"title"`
	Duration int       `json:"duration"`
}

func NewJobCard(title string, duration int) JobCard {
	return JobCard{
		UUID:     uuid.New(),
		Title:    title,
		Duration: duration,
	}
}
