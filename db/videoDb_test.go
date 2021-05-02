package db

import (
	"fmt"
	"testing"
)

func TestGetVideosBySinceId(t *testing.T) {
	sinceId := "0"
	videos, err := GetVideosBySinceId(sinceId, 10)
	if err != nil {
		t.Error(err)
	}
	for _, video := range videos {
		fmt.Printf("%#v\n", video)
	}
}
