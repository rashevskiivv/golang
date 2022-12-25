package dump

import (
	"fmt"
	"path/filepath"
	"strings"
)

func filePathsFunc() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))
	a, b := filepath.Split(p)
	fmt.Printf("Split(p): %v, %v\n", a, b)

	fmt.Println(filepath.IsAbs("dir/file"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)
	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
