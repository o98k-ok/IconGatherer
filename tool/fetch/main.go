package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// parse https://icons8.com/icon/pGqqobAPSa_u/touchpad
// to https://img.icons8.com/?size=250&id=pGqqobAPSa_u&format=png
// to https://img.icons8.com/?size=250&id=pGqqobAPSa_u&format=png&color=FFFFFF
func GenAllURLs(url string) map[string]string {
	base := "https://img.icons8.com/?size=250&id="
	fields := strings.Split(url, "/")
	id := fields[len(fields)-2]

	format := "&format=png"
	color := "&color=FFFFFF"

	urls := map[string]string{
		"black": base + id + format,
		"white": base + id + format + color,
	}
	return urls
}

func Download(urls map[string]string, filePrfix string) {
	for color, url := range urls {
		response, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		os.WriteFile(fmt.Sprintf("%s_%s.png", filePrfix, color), body, 0o644)
	}
}

func GenAllURLsV2(url string) map[string]string {
	base := "https://img.icons8.com/?size=250&id="
	fields := strings.Split(url, "/")
	id := fields[len(fields)-2]

	format := "&format=png"
	urls := map[string]string{
		"color": base + id + format,
	}
	return urls
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run main.go <url> <output_prefix>")
	}

	raw := os.Args[1]
	urls := GenAllURLsV2(raw)
	Download(urls, os.Args[2])
}
