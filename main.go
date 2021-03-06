package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: md5check [file] [provided md5 sum]")
		os.Exit(1)
	}

	filename := args[0]
	knownSum, err := hex.DecodeString(args[1])
	if err != nil {
		fmt.Println("Error decoding provided md5 hash: ", err)
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading input file: ", err)
		file.Close()
		os.Exit(1)
	}

	sum := md5.New()
	_, err = io.Copy(sum, file)
	if err != nil {
		fmt.Println("Error computing md5 sum of provided file: ", err)
		file.Close()
		os.Exit(1)
	}

	file.Close()
	fileSum := sum.Sum(nil)

	if bytes.Compare(knownSum, fileSum) != 0 {
		fmt.Printf("MISMATCH! MD5 Sum of given file was: %x\n", fileSum)
		os.Exit(1)
	} else {
		fmt.Println("OK")
	}
}
