package repository

import (
	"log"
	"testing"
	"time"
	"tutorial/stofgevreet/model"
)

func TestSavePoint(t *testing.T) {
	input := model.Point{
		Car:        "A001",
		User:       "Fanie user",
		Method:     "start",
		Scantime:   time.Now(),
		Points:     "123",
		Checkpoint: "some checkpoint",
	}

	err := SavePoint(input)
	if err != nil {
		log.Println("Save point failed: ", err.Error())
		t.Fail()
	}
}

func TestSaveScan(t *testing.T) {
	input := model.Scan{
		Car:        "A001",
		User:       "Fanie user",
		Method:     "start",
		Scantime:   time.Now(),
		Checkpoint: "some checkpoint",
	}

	err := SaveScan(input)
	if err != nil {
		log.Println("Save point failed: ", err.Error())
		t.Fail()
	}
}

func TestSaveStopwatch(t *testing.T) {
	input := model.Stopwatch{
		Car:      "A001",
		User:     "Fanie user",
		Method:   "start",
		Scantime: time.Now(),
		Lap:      "some value",
	}

	err := SaveStopwatch(input)
	if err != nil {
		log.Println("Save point failed: ", err.Error())
		t.Fail()
	}
}
