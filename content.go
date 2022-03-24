//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	flag "github.com/spf13/pflag"
)

// -d(--dir) 要生成目录的文件夹
// -o(--output) 输出文件位置
var dirPtr *string = flag.StringP("dir", "d", ".", "folder to generate the catalog")
var outputPtr *string = flag.StringP("output", "o", ".", "output file location")

type link struct {
	path  string
	title string
}

func main() {
	flag.Parse()

	var dir, output string
	var err error
	if dir, err = filepath.Abs(*dirPtr); err != nil {
		log.Fatal(err)
	}

	if output, err = filepath.Abs(*outputPtr); err != nil {
		log.Fatal(err)
	}

	links := listAllReadme(dir) // 从dir获取links

	_ = os.Remove(filepath.Join(output, "/README.md")) // 删除已有的readme.md

	readmeF, err := os.Create(filepath.Join(output, "/README.md")) // 创建新的readme.md
	if err != nil {
		log.Fatal(err)
	}
	defer readmeF.Close()

	writeReadme(readmeF, links, filepath.Base(dir))
}

func writeReadme(w io.Writer, links []*link, title string) {
	fmt.Fprintln(w, fmt.Sprintf("## %s", title)) // 写入标题
	for i, link := range links {                 // 写入目录
		fmt.Fprintf(w, "%d. [%s](./%s%s)\n", i+1, link.title, link.path, "readme.md")
	}
}

func listAllReadme(dir string) []*link {
	links := make([]*link, 0)
	walkDir(dir, func(filename string) {
		filedir, file := filepath.Split(filename)
		if file != "readme.md" {
			return
		}

		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		bf := bufio.NewReader(f)
		title, err := bf.ReadBytes('\n')
		if err != nil {
			log.Fatalf("error read file %v", filename)
		}

		title = title[2 : len(title)-1] // remove "# " and "\n"
		links = append(links, &link{
			path:  strings.Replace(filedir, dir+"/", "", -1),
			title: string(title),
		})
	})
	return links
}

func walkDir(dir string, fn func(filename string)) {
	fileEntrys, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, entry := range fileEntrys {
		child := fmt.Sprintf("%s/%s", dir, entry.Name())
		if entry.IsDir() {
			walkDir(child, fn)
		} else {
			fn(child)
		}
	}
}
