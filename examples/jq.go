//go:build ignore
package main

import (
	"fmt"
	"strconv"

	"github.com/bitfield/script"
)

func JQRaw(s string) string {
	res, _ := strconv.Unquote(s)
	return res
}

func main() {
	s := `{"json":"{\"in\":\"json\"}"}`

	fmt.Printf("echo '%s' | jq '.json' ->\n", s)
	_, err := script.Echo(s).JQ(".json").Stdout()
	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("echo '%s' | jq -r '.json' ->\n", s)
	_, err = script.Echo(s).JQ(".json").FilterLine(JQRaw).Stdout()
	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("echo '%s' | jq -r '.json' | jq .in' ->\n", s)
	_, err = script.Echo(s).JQ(".json").FilterLine(JQRaw).JQ(".in").Stdout()
	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
}
