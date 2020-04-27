package main

import (
	"flag"
	"fmt"
	"github.com/gtinside/gosearch/internal/pkg/scrapper"
	"os"
)

const domain = "https://api.duckduckgo.com/"
func main() {
	fmt.Println("Utility to get top links for a given keyword :")
	keyword := flag.String("keyword", "golang", " Enter the keyword that you want to search")
	num := flag.Int("pages", 5, "Enter the number of links you want to get back")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	output := scrapper.Retrieve(*keyword, *num, domain)
	for _, s := range output {
		fmt.Println(s)
	}
}
