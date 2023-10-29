package tasks

import (
	"log"
)

func installRecover() {
	if err := recover(); err != nil {
		log.Printf("panic: %v", err)
	}
}

func GoSafe(fn JobFunc) {
	go func() {
		defer installRecover()
		fn()
	}()
}

func GoSyncSafe(fn JobFunc) error {
	done := make(chan error)
	go func() {
		defer installRecover()

		done <- fn()
	}()
	return <-done
}
