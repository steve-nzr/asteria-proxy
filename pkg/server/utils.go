package server

import "time"

func timeNow() int64 {
	return time.Now().UnixNano()
}

func find(list []string, a string) int {
	for i, n := range list {
		if a == n {
			return i
		}
	}
	return -1
}
