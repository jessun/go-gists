package main

import "time"

func Retry_V2(body func() (RetrySignal, error), maxRetries int, minDelay time.Duration, maxDelay time.Duration) error {
	currentDelay := minDelay
	ticker := time.NewTicker(currentDelay)
	defer ticker.Stop()

	retries := 0

	for {
		select {
		case <-ticker.C:
			response, err := body()
			if err != nil {
				ticker.Stop()
				return err
			}

			switch response {
			case FuncSuccess:
				currentDelay = minDelay
				ticker.Stop()
				ticker = time.NewTicker(currentDelay)
				retries = 0
			case FuncFailure:
				currentDelay = minDuration(currentDelay*2, maxDelay)
				ticker.Stop()
				ticker = time.NewTicker(currentDelay)
				retries++
			case FuncComplete:
				return nil
			}
		}

		if retries >= maxRetries {
			return nil
		}
	}
}
