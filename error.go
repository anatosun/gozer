package gozer

import (
	"encoding/json"
	"fmt"
)

type DeezerError struct {
	Content struct {
		Type    string `json:"type"`
		Message string `json:"message"`
		Code    uint64 `json:"code"`
	} `json:"error"`
}

func newDeezerError(bodyBytes []byte) *DeezerError {

	deezerError := &DeezerError{}
	json.Unmarshal(bodyBytes, deezerError)
	if deezerError.Content.Type == "" {
		return nil
	}
	return deezerError
}

func (e *DeezerError) Error() string {

	return fmt.Sprintf("an \"%s\" has occurred with message \"%s\" (code %d)", e.Content.Type, e.Content.Message, e.Content.Code)
}
