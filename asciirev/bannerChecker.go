package asciirev

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func BannerChecker(s string) bool {
	regB := regexp.MustCompile(`standard|shadow|thinkertoy`)
	if !regB.MatchString(s) {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return false
	}
	if !CheckFileHashing(s) {

		fmt.Println("ERROR: the file was changed!")
		return false
	}
	return true
}

func CheckFileHashing(fileName string) bool {
	fileName += ".txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, 30*1024)
	sha256 := sha256.New()
	for {
		n, err := file.Read(buf)
		if n > 0 {
			_, err := sha256.Write(buf[:n])
			if err != nil {
				fmt.Println(err.Error())
				return false
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf("Read %d bytes: %v", n, err)
			break
		}
	}

	sum1 := sha256.Sum(nil)
	sum := hex.EncodeToString(sum1)

	shadow := "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoy := "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
	standard := "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"

	if fileName == "standard.txt" && string(sum) == standard {
		return true
	} else if fileName == "thinkertoy.txt" && string(sum) == thinkertoy {
		return true
	} else if fileName == "shadow.txt" && string(sum) == shadow {
		return true
	}
	return false
}
