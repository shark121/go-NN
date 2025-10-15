package helpers

import (
	"os"
	"path/filepath"
)

func UbyteReader(path string) ([]os.DirEntry, []string) {

	files, error := os.ReadDir(path)

	if error != nil {
		println(error)
	}

	docs := []os.DirEntry{}

	fullPaths := []string{}

	for _, v := range files {
		if v.IsDir() {
			path, fullPath := UbyteReader(filepath.Join(path, v.Name()))

			println(path)

			fullPaths = append(fullPaths, fullPath...)

			docs = append(docs, path...)
		} else {
			docs = append(docs, v)
			fullPaths = append(fullPaths, filepath.Join(path, v.Name()))
		}
	}

	return docs, fullPaths
}
