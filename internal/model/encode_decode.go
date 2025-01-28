package model

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Validator interface {
	Valid(ctx context.Context) (problems map[string]string)
}

func Decode[T Validator](w http.ResponseWriter, r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode body: %w", err)
	}

	if problems := v.Valid(r.Context()); len(problems) != 0 {
		return v, fmt.Errorf("invalid %T, %d problems", v, len(problems))
	}

	return v, nil
}

func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8") // we set it to json event it is not
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("Encode Json %w", err)
	}
	return nil
}
