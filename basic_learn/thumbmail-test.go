package main

import (
	// "gopl.io/ch8/thumbnail"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopl.io/ch8/thumbnail"
)

func getDirectory(path string) (files []string, err error) {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if !fi.IsDir() {
			ok := strings.HasSuffix(fi.Name(), ".jpg")
			if ok {
				files = append(files, path+PthSep+fi.Name())
			} else {
				// fmt.Printf("略过%s\n", fi.Name())
			}
		}
	}
	return files, nil
}

// makeThumbnails5 makes thumbnails for the specified files in parallel.
// It returns the generated file names in an arbitrary order,
// or an error if any step failed.
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}
func main() {
	files, err := getDirectory("/home/adl/图片")
	if err != nil {
		log.Fatal(err)
	} else {
		thumbfiles, err := makeThumbnails5(files)
		if err != nil {
			log.Fatal(err)
		} else {
			for _, file := range thumbfiles {
				fmt.Println(file)
			}
		}
	}
	// thumbnail.ImageFile("/home/adl/图片/os2.jpg")
}
