package util

import "github.com/google/uuid"

func Uuid() string {
	return uuid.New().String()
}
