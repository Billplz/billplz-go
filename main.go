package main

import (
  "os"
  billplz "billplz/pkg/billplz"
)

func main () {
  billplz := billplz.Init{os.Args[1]}
  return billplz.GetCollection(os.Args[2])
}
