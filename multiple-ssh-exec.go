/**
 * http://kukuruku.co/hub/golang/ssh-commands-execution-on-hundreds-of-servers-via-go
 * http://www.golangpatterns.info/concurrency/parallel-for-loop
 *
 *
 */

package main

import "code.google.com/p/go.crypto/ssh"
import "fmt"
import "bytes"
import "strconv"
import "sync"
import "os"

type password string

func main() {
	counter, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic("Failed to convert: " + err.Error())
	}

	execute_commands := make([]string, counter)
	//http://stackoverflow.com/questions/8539551/dynamically-initialize-array-size-in-go
	for i := 0; i < counter; i++ {
		execute_commands[i] = "echo 'hello\t" + strconv.Itoa(i) + "'"
	}

	var wg sync.WaitGroup
	for _, command := range execute_commands {
		wg.Add(1)

		go func(command string) {

			config := &ssh.ClientConfig{
				User: "USERNAME",
				Auth: []ssh.AuthMethod{
					ssh.Password("PASSWORD"),
				},
			}
			client, err := ssh.Dial("tcp", "HOSTNAME:PORT", config)
			if err != nil {
				panic("Failed to dial: " + err.Error())
			}

			session, err := client.NewSession()
			if err != nil {
				panic("Failed to create session: " + err.Error())
			}
			defer session.Close()

			var b bytes.Buffer
			session.Stdout = &b
			if err := session.Run(command); err != nil {
				panic("Failed to run: " + err.Error())
			}
			fmt.Println(b.String())

			wg.Done()
		}(command)
	}
	wg.Wait()
}
