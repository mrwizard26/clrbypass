// Used to do a quick and dirty xor on the file

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	_33d9ccf0 := byte(195)

	_7b3acddd := flag.String("file", "", "Path Of the Shellcode")
	flag.Parse()

	_d2359f93 := *_7b3acddd

	_d4178102, _a30f3ccf := os.ReadFile(_d2359f93)
	if _a30f3ccf != nil {
		fmt.Println("Error Opening file")
		fmt.Println(_a30f3ccf.Error())
	}

	var _33c4de43 []byte

	for _0dd0a21d := 0; _0dd0a21d < len(_d4178102); _0dd0a21d++ {
		_33c4de43 = append(_33c4de43, _d4178102[_0dd0a21d]^_33d9ccf0)
	}

	_e965f019 := "file.enc"

	_a30f3ccf = os.WriteFile(_e965f019, _33c4de43, 0644)
	if _a30f3ccf != nil {
		fmt.Println("Failed to write file:", _a30f3ccf)
		return
	}
}
