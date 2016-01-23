package main

import (
	"fmt"
	"os"

	"github.com/itang/gotang"
	"github.com/itang/gotitle"
)

func main() {
	url := os.Args[1]
	if url == "" {
		fmt.Errorf("请指定url.")
	}

	fmt.Printf(">http get '%v'...\n", url)

	title, err := gotitle.GetTitleFromUrl(url)
	gotang.AssertNoError(err, "http请求出错!")

	fmt.Printf("\t->>title: %v\n", title)
}
