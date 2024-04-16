package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	d   string
	n   int
	ext string
)

func init() {
	flag.StringVar(&d, "d", "", "start directory for file type search")
	flag.IntVar(&n, "n", 1, "number of files for this type to search for")
	flag.StringVar(&ext, "ext", "", "the file type/extension to search for")
	flag.Parse()
}

func main() {
	err := findFilesByType(d, ext)

	if err == io.EOF {
		err = nil
	}
	if err != nil {
		fmt.Println("Error searching files:", err)
	}
}

func findFilesByType(rootPath string, fileType string) error {
	filesFound := 0
	err := func() error {
		err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				if os.IsPermission(err) {
					return nil
				}
				return err
			}
			if !info.IsDir() {
				fileInfo, _ := os.Stat(path)
				if filepath.Ext(path) == fileType {
					filesFound++
					fmt.Println(path, fileInfo.Size())
					if filesFound == n {
						return io.EOF
					}
				}

			}
			return nil
		})
		return err
	}()
	return err
}
