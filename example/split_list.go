/**
 * http://golang.org/pkg/path/filepath/#example_SplitList
 *
 */

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	//windows系であれば、パスの区切り文字には;を利用する。
	fmt.Println("On Windows:", filepath.SplitList("c:\\Learning;c:\\build"))

	//unix系であれば、パスの区切り文字には:を利用する。
	//fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin"))
}
