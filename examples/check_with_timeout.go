//go:build ignore

package main

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/bitfield/script"
)

func ExecuteWithTimeout(f func() string, timeout time.Duration) (string, error) {
	result := make(chan string, 1)
	go func() {
		result <- f()
	}()
	select {
	case msg := <-result:
		return msg, nil
	case <-time.After(timeout):
		return "", errors.New("timeout")
	}
}

func main() {
	for {
		script.File("servers.txt").FilterLine(func(url string) string {
			res, err := ExecuteWithTimeout(func() string {
				ip, _ := script.Exec(fmt.Sprintf("dig +short %s", url)).First(1).String()
				return strings.TrimSpace(ip)
			}, time.Millisecond*200)

			if err != nil {
				res = err.Error()
			}
			return fmt.Sprintf("[%v] %s -> %s", time.Now(), url, res)
		}).Stdout()
		time.Sleep(time.Second)
	}
}
