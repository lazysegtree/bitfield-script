//go:build ignore

package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bitfield/script"
)

func main() {

	for {
		script.File("servers.txt").FilterLine(func(url string) string {
			ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
			defer cancel()
			
			ip, err := script.ExecWithContext(ctx, fmt.Sprintf("dig +short %s", url)).First(1).String()
			res := strings.TrimSpace(ip)

			if err != nil {
				res = err.Error()
			}
			return fmt.Sprintf("[%v] %s -> %s", time.Now(), url, res)
		}).Stdout()
		time.Sleep(time.Second)
	}
}
