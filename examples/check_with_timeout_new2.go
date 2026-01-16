//go:build ignore

package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bitfield/script"
)

func main() {

	for {
		script.File("servers.txt").FilterLine(func(url string) string {
			ip, err := script.ExecWithTimeout(fmt.Sprintf("dig +short %s", url), 200*time.Millisecond).
				First(1).String()
			res := strings.TrimSpace(ip)

			if err != nil {
				res = err.Error()
			}
			return fmt.Sprintf("[%v] %s -> %s", time.Now(), url, res)
		}).Stdout()
		time.Sleep(time.Second)
	}
}
