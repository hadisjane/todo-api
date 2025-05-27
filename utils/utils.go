package utils

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func ExtractID(r *http.Request) (int, error) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		return 0, errors.New("invalid URL")
	}
	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid id")
	}
	return id, nil
}
