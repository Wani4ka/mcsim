package util

import (
	"fmt"
	"strings"
)

func FormatFloat(f float32, places int) string {
	res := strings.TrimRight(fmt.Sprintf(fmt.Sprintf("%%.%df", places), f), "0.")
	if res == "" {
		return "0"
	}
	return res
}
