package utils

import (
	"strings"

	"github.com/google/uuid"
)

// UUID uuid
func UUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(u.String(), "-", ""), nil
}

// MustUUID return uuid without error
func MustUUID() string {
	return strings.ReplaceAll(uuid.NewString(), "-", "")
}
