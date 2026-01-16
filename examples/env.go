package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bitfield/script"
)

func main() {
	fmt.Printf("Started\n")
	env := []string{`MY_ENV=foo`}
	out1, _ := script.NewPipe().WithEnv(env).Exec(`echo $MY_ENV`).String()

	cmd := exec.Command("sh", "-c", `echo $MY_ENV`)
	cmd.Env = env
	out2, _ := cmd.Output()

	out3, _ := exec.Command("sh", "-c", `MY_ENV=foo; echo $MY_ENV`).Output()
	fmt.Printf("out1 = '%s'\n", out1)
	fmt.Printf("out2 = '%s'\n", out2)
	fmt.Printf("out3 = '%s'\n", out3)

	os.Setenv("MY_ENV", "foo")
	out4, _ := script.NewPipe().WithEnv(env).Exec(`echo $MY_ENV`).String()
	fmt.Printf("(After setting env) out4 = '%s'\n", out4)

}
