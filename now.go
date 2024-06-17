package gohelpers

import (
	"time"
)

func Now() string {
	now := time.Now()

	return now.Format("02.01.2006 15:04:05")

}
