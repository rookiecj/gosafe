package tasks

import (
	"log"
)

func installRecover() {
	if err := recover(); err != nil {
		log.Printf("panic: %v", err)
	}
}

// GoSafe runs a function in a goroutine, and recovers from any panics.
func GoSafe(fn JobFunc) {
	go func() {
		defer installRecover()
		fn()
	}()
}

// GoSyncSafe runs a function in a goroutine synchronously, and recovers from any panics.
func GoSyncSafe(fn JobFunc) error {
	done := make(chan error)
	go func() {
		defer installRecover()

		done <- fn()
	}()
	return <-done
}
