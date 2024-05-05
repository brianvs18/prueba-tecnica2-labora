package common

import "time"

func Sleep(amount time.Duration) {
	time.Sleep(amount * time.Second)
}
