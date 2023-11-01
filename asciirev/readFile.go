package asciirev

import (
	"io/ioutil"
	"log"
	"strings"
)

func ReadFile(s string) []string {
	banner, err := ioutil.ReadFile(s + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	banner = banner[1:]
	data := strings.Split(strings.ReplaceAll(string(banner), "\r", ""), "\n\n")
	return data
}
