package utils

import (
	"fmt"
	"testing"
)

func TestGetUnixTimestamp(t *testing.T) {
	timestamp := GetUnixTimestamp()
	fmt.Println(timestamp)
}
