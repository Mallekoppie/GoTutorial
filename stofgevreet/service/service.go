package service

import (
	"net/http"

	"tutorial/stofgevreet/model"
	"tutorial/stofgevreet/repository"

	"github.com/Mallekoppie/goslow/platform"
	"go.uber.org/zap"
)

func SaveScan(w http.ResponseWriter, r *http.Request) {
	platform.Logger.Info("Executing save scan")

	input := model.Scan{}

	err := platform.JsonMarshaller.ReadJsonRequest(r.Body, &input)
	if err != nil {
		platform.Logger.Error("Unable to read request for SaveScan", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = repository.SaveScan(input)
	if err != nil {
		platform.Logger.Error("Error saving scan to DB", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SavePoint(w http.ResponseWriter, r *http.Request) {
	platform.Logger.Info("Executing save point")

	input := model.Point{}

	err := platform.JsonMarshaller.ReadJsonRequest(r.Body, &input)
	if err != nil {
		platform.Logger.Error("Unable to read request for SavePoint", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = repository.SavePoint(input)
	if err != nil {
		platform.Logger.Error("Error saving scan to DB for SavePoint", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func SaveStopwatch(w http.ResponseWriter, r *http.Request) {
	platform.Logger.Info("Executing save stopwatch")

	input := model.Stopwatch{}

	err := platform.JsonMarshaller.ReadJsonRequest(r.Body, &input)
	if err != nil {
		platform.Logger.Error("Unable to read request for SaveStopwatch", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = repository.SaveStopwatch(input)
	if err != nil {
		platform.Logger.Error("Error saving scan to DB for SaveStopwatch", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
