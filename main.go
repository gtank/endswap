package main

import (
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	decoded, err := hex.DecodeString(os.Args[1])
	if err != nil {
	    fmt.Println(err)
	    os.Exit(1)
        }
	for i := 0; i < len(decoded)/2; i++ {
		decoded[i], decoded[len(decoded) - 1 - i] = decoded[len(decoded)-1-i], decoded[i]
	}
	fmt.Println(hex.EncodeToString(decoded))
}
