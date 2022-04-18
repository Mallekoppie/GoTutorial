package repository

import (
	"fmt"
	"testing"
)

func TestCallHelloWorld(t *testing.T) {
	result, err := CallHelloWorld()
	if err != nil {
		fmt.Println("Error calling hello world", err.Error())
		t.FailNow()
	}

	fmt.Println("Result: ", result)

}
