package basic

import "time"

// Retry 超时重试 interval间隔时间
func Retry(callback func() error, maxRetries int, interval time.Duration) error {
	var err error
	for i := One; i <= maxRetries; i++ {
		if err = callback(); err != nil {
			time.Sleep(interval)
			continue
		}
		return nil
	}
	return err
}
