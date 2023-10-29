package tasks

import (
	"log"
)

func GoSafe(fn JobFunc) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
			}
		}()
		fn()
	}()
}
