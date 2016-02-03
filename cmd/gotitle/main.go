package main

import (
	"fmt"
	"os"
	"time"

	"github.com/itang/gotang"
	gtime "github.com/itang/gotang/time"
	"github.com/itang/gotitle"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请指定url.")
	}

	gotang.Time(func() {
		for _, url := range os.Args[1:] {
			fmt.Printf(">http get '%v'...\n", url)

			title, err := gotitle.GetTitleFromUrl(url)
			gotang.AssertNoError(err, "http请求出错!")

			fmt.Printf("\t->>title: %v\n", title)

			displayResult(url, title)
		}
	})
}

func displayResult(url, title string) {
	fmt.Println()
	tpl := `rs << Read.new "%s",
    title: "%s",
    created_at: "%s"`
	fmt.Printf(tpl, url, title, gtime.FormatDefault(time.Now()))
	fmt.Print("\n\n")
}
