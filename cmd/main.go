package main

import (
	"fmt"

	"iw-mapviewer.cmh.platform-dev.evinternal.net/api/test_2024_09_20"
)

func main() {
	req := test_2024_09_20.RankImagesRequest{}
	req.Hello = "world"
	fmt.Printf("hello %s", req.Hello)
}
