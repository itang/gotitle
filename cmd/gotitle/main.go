package main

import (
	"fmt"
	"os"

	"github.com/itang/gotang"
	"github.com/itang/gotitle"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请指定url.")
	}

	for _, url := range os.Args[1:] {
		fmt.Printf(">http get '%v'...\n", url)

		title, err := gotitle.GetTitleFromUrl(url)
		gotang.AssertNoError(err, "http请求出错!")

		fmt.Printf("\t->>title: %v\n", title)
	}
}
