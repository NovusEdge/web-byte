package main

import (
	wb "github.com/NovusEdge/web-byte/src"
	wbhandlers "github.com/NovusEdge/web-byte/common-handlers"
	"log"
)

func main() {
	bs, _ := wb.MakeByteServer(42069, nil)

	bs.AddHandler("/", wbhandlers.Welcome)
	log.Fatal(bs.Serve())
}
