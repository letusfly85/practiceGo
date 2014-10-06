/**
 * https://gobyexample.com/execing-processes
 *
 */

package main

import "syscall"
import "os"
import "os/exec"

func main() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	env := os.Environ()

	args := []string{"ls", "-a", "-l", "-h"}
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	args = []string{"echo", "a"}
	execErr = syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
