package models

import "time"

type EventType = int

const (
	Login EventType = iota
	Logout
)

type Event struct {
	Id   int       `json:"id" gorm:"primaryKey"`
	Time time.Time `json:"time" sql:"type:timestamp without time zone"`
	User string    `json:"user"`
	Type EventType `json:"type"`
	Desc string    `json:"desc"`
}
