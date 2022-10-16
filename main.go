package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/eyedeekay/goSam"
	tr "github.com/snakesel/libretranslate"
)

func main() {
	stdout := os.Stdout
	os.Stdout = os.NewFile(0, os.DevNull)
	os.Stderr = os.NewFile(0, os.DevNull)
	lang := os.Getenv("TRANSLATE_LANG")
	if lang == "" {
		lang = "fr"
	}
	URL := os.Getenv("TRANSLATE_SERVICE")
	if URL == "" {
		URL = "http://w62j277kjls7agmctbtjzuthvsaiz7zzjthmahdk7pgweditlfzq.b32.i2p"
	}
	flag.StringVar(&lang, "lang", lang, "Set the target language")
	flag.StringVar(&URL, "service", URL, "Service to use to translate the text")
	flag.Parse()
	sam, err := goSam.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	defer sam.Close()
	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: sam.Dial,
		},
	}
	http.DefaultClient = httpClient
	translate := tr.New(tr.Config{
		Url: URL,
	})
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	os.Stdout = stdout
	for _, line := range lines {
		if line != "" && !strings.HasPrefix(line, "```") {
			trtext, err := translate.Translate(line, "auto", lang)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s\n", trtext)
		} else {
			fmt.Printf("%s\n", line)
		}
	}
}
