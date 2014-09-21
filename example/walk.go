/**
 * http://d.hatena.ne.jp/taknb2nch/20140110/1389280260
 *
 *
 *
 */

package main

import "fmt"
import "os"
import "path/filepath"

func main() {

	//root := "~/WORK"
	root := "C:\\Learning"

	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {

			rel, err := filepath.Rel(root, path)
			fmt.Println(rel)

			return nil
		})

	if err != nil {
		//fmt.Println(1, err)
	}
}
