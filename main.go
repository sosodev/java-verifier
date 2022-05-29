package main

import (
	"crypto/sha256"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: ./java-verifier <path-to-original-jar-directory> <path-to-comparison-jar-directory>")
		os.Exit(22)
	}

	originalDirectoryPath := os.Args[1]
	comparisonDirectoryPath := os.Args[2]

	_, err := os.Stat(originalDirectoryPath)
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(comparisonDirectoryPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Checking the bytecode for any differences")

	diffDetected := false
	err = filepath.Walk(comparisonDirectoryPath, func(path string, info fs.FileInfo, err error) error {
		filename := filepath.Base(path)
		if !strings.HasSuffix(filename, ".class") {
			return nil
		}
		originalPath := strings.Replace(path, comparisonDirectoryPath, originalDirectoryPath, 1)

		comparisonFile, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		originalFile, err := ioutil.ReadFile(originalPath)
		if err != nil {
			return err
		}

		if sha256.Sum256(comparisonFile) != sha256.Sum256(originalFile) {
			fmt.Printf("difference detected for %s\n", filename)
			diffDetected = true
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	if !diffDetected {
		fmt.Println("No differences detected! :)")
	}
}
