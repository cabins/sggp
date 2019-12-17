package util

import "fmt"

const version = "0.0.1"

func Version() {
	fmt.Println("ghc", version, "(c) Cabins, 2019-2020")
}