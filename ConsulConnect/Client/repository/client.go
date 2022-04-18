package repository

import (
	"client/model"
	"errors"
	"fmt"
	"github.com/Mallekoppie/goslow/platform"
	"go.uber.org/zap"
	io "io/ioutil"
	"net/http"
)

var (
	client   *http.Client
	basePath string

	ErrIncorrectResponseCode = errors.New("incorrect response code")
)

func init() {
	client = platform.CreateHttpClient("default")
	config := model.Config{}
	err := platform.GetComponentConfiguration("componentconfigexample", &config)
	if err != nil {
		platform.Logger.Fatal("Repository config loading failed", zap.Error(err))
	}

	basePath = config.BasePath
}

func CallHelloWorld() (result string, err error) {

	response, err := client.Get(fmt.Sprintf("%s/", basePath))
	if err != nil {
		platform.Logger.Error("Error calling Hello World", zap.Error(err))
		return "", err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		platform.Logger.Error("Error reading Hello World response body", zap.Error(err))
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		platform.Logger.Error("Incorrect Response Code from Server",
			zap.Int("expected", http.StatusOK),
			zap.Int("actual", response.StatusCode))

		return "", ErrIncorrectResponseCode
	}

	return string(data), nil
}
