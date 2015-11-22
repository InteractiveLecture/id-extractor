package idextractor

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Extractor func(r *http.Request) (string, error)

func MuxIdExtractor(fieldName string) Extractor {
	return func(r *http.Request) (string, error) {
		result, ok := mux.Vars(r)[fieldName]
		if !ok {
			return "", errors.New("Could not extract " + fieldName)
		}
		return result, nil
	}

}
