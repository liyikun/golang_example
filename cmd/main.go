package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	dataCmd := exec.Command("mkdir", "-p", "./test")

	fmt.Println(dataCmd)
	output, err := dataCmd.Output()

	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))

	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	binary, lookErr := exec.LookPath("ls")

	if lookErr != nil {
		panic(lookErr)
	}

	fmt.Println(binary)

	args := []string{"ls", "-a", "-l"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}
}
