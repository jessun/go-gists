package main

import "time"

func Retry_ORIGIN(body func() (RetrySignal, error), maxRetries int, minDelay time.Duration, maxDelay time.Duration) error {
	currentDelay := minDelay
	ticker := time.NewTicker(currentDelay)
	defer ticker.Stop()

	retries := 0
	for range ticker.C {
		response, err := body()
		if err != nil {
			return err
		}

		switch response {
		case FuncSuccess:
			currentDelay = minDelay
			ticker.Reset(currentDelay)
			retries = 0
		case FuncFailure:
			currentDelay = minDuration(currentDelay*2, maxDelay)
			ticker.Reset(currentDelay)
			retries++
		case FuncComplete:
			return nil
		}

		if retries >= maxRetries {
			return nil
		}
	}

	return nil
}
