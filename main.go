package main

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var age = flag.Duration("age", 7*24*time.Hour, "older will be archived")

func main() {
	flag.Parse()
	if err := Main(); err != nil {
		panic(err)
	}
}

func Main() (err error) {
	dirs, err := ioutil.ReadDir(flag.Arg(0))
	if err != nil {
		return err
	}

	for _, info := range dirs {
		if strings.HasPrefix(info.Name(), ".") {
			continue
		}

		if time.Since(info.ModTime()) > *age {
			err := archive(flag.Arg(0), info)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func archive(base string, info os.FileInfo) (err error) {
	j := func(path ...string) string {
		return filepath.Join(base, filepath.Join(path...))
	}

	archive_folder := info.ModTime().Format(".2006-01-02")
	if err := os.MkdirAll(j(archive_folder), 0755); err != nil {
		return err
	}

	return os.Rename(j(info.Name()), j(archive_folder, info.Name()))
}
