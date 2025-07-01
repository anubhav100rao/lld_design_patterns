package decorator

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Operation func() error

func RetryDecorator(op Operation, maxAttempts int, backoff time.Duration) Operation {
	return func() error {
		var err error
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			err = op()
			if err == nil {
				return nil
			}
			time.Sleep(backoff)
		}
		return errors.New("operation failed after retries: " + err.Error())
	}
}

func unreliableNetworkCall() error {
	if rand.Intn(2) == 0 {
		return errors.New("network timeout")
	}
	return nil
}

func RunRetryDecoratorDemo() {
	retryingCall := RetryDecorator(unreliableNetworkCall, 3, 500*time.Millisecond)
	if err := retryingCall(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Call succeeded")
}
