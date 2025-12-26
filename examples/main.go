package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bitfield/script"
)

func main() {
	for {
		fmt.Println("started")
		//script.Echo("hi\n").Stdout()
		//script.File("servers.txt").ExecForEach("dig +short {{.}} | head -n1").Stdout()
		//script.File("servers.txt").EachLine(func(url string, out *strings.Builder) {
		//	script.Exec(fmt.Sprintf("dig +short %s", url)).First(1).Stdout()
		//})
		script.File("servers.txt").FilterLine(func(url string) string {
			ip, _ := script.Exec(fmt.Sprintf("dig +short %s", url)).First(1).String()
			return strings.TrimSpace(ip)
		}).Stdout()

		// idea - get ip rout get default and arp -n on a container which could timeout

		time.Sleep(time.Second)
	}
}
