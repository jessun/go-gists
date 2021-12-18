package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		fmt.Printf("%s\n", path)
		fmt.Printf("%s\n", info.Name())
		return nil
	})
}
