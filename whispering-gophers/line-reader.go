/**
 * refs: http://qiita.com/jpshadowapps/items/ae7274ec0d40882d76b5
 *
 *
 *
 */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := fromFile("aaa.txt")
	fmt.Println(line)
}

func fromFile(filePath string) []string {

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", filePath, err)
		os.Exit(1)
	}

	defer f.Close()

	lines := make([]string, 0, 100)

	scanner = bufio.NewReader(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if s_err := scanner.Err(); s_err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not scan: %v\n", filePath, s_err)
	}

	return lines
}
