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
			ip, _ := script.Exec(fmt.Sprintf("dig +short %s", url)).First(1).String()
			ip = strings.TrimSpace(ip)
			return fmt.Sprintf("[%v] %s -> %s", time.Now(), url, ip)
		}).Stdout()
		time.Sleep(time.Second)
	}
}
