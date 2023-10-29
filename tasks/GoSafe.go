package safe

import (
	"log"

	"./tasks"
)

func GoSafe(fn tasks.JobFunc) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %v", err)
			}
		}()
		fn()
	}()
}
