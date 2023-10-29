package tasks

import (
	"log"
)

func installRecover() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %v", err)
		}
	}()
}

func GoSafe(fn JobFunc) {
	go func() {
		installRecover()
		_ = fn()
	}()
}

func GoSyncSafe(fn JobFunc) error {
	done := make(chan error)
	go func() {
		installRecover()

		done <- fn()
	}()
	return <-done
}
