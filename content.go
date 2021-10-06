package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type link struct {
	path  string
	title string
}

func main() {
	curDir := "."

	os.Remove("./README.md")
	readmeF, err := os.Create("./README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer readmeF.Close()

	links := listAllReadme(curDir)
	fmt.Fprintln(readmeF, "# LeetCode-go")
	for i, link := range links {
		dir, _ := filepath.Split(link.path)
		fmt.Fprintf(readmeF, "%d. [%s](%s)   \n", i+1, link.title, dir)
	}
}

func listAllReadme(dir string) []*link {
	links := make([]*link, 0)
	walkDir(dir, func(filename string) {
		path := strings.Split(filename, "/")
		suffix := path[len(path)-1]
		if suffix != "readme.md" {
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
			path:  filename,
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
