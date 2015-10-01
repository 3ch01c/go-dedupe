package main

import (
  "fmt"
  //"io/ioutil"
  "flag"
  //"path"
  "path/filepath"
  "os"
  //"crypto/md5"
)

func main() {
  flag.Parse()

  rootDirs := flag.Args()

  for _, rootDir := range rootDirs {
    filepath.Walk(rootDir, func(path string, fi os.FileInfo, err error) (e error) {
      fmt.Printf("%+v %+v %+v\n",path,fi.Size(),fi.ModTime())
      return nil
    })
  }
}
