package utils

import (
	"errors"
	"net/http"
)

func ResponseError(isErr bool, hr http.ResponseWriter, eMessage string, statusCode int) error{
	if isErr {
		http.Error(hr, eMessage, statusCode)
		return errors.New(eMessage)
	}

	return nil
}