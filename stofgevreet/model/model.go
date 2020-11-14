package model

import (
	"time"
)

type Point struct {
	Car        string    `json:"car"`
	Scantime   time.Time `json:"scantime"`
	Method     string    `json:"method"`
	User       string    `json:"user"`
	Points     int       `json:"points"`
	Checkpoint string    `json:"checkpoint"`
}

type Scan struct {
	Car        string    `json:"car"`
	Scantime   time.Time `json:"scantime"`
	Method     string    `json:"method"`
	User       string    `json:"user"`
	Checkpoint string    `json:"checkpoint"`
}

type Stopwatch struct {
	Car      string    `json:"car"`
	Scantime time.Time `json:"scantime"`
	Method   string    `json:"method"`
	User     string    `json:"user"`
	Lap      string    `json:"lap"`
}
