package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const assetsFolder = "./assets"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func formatName(fileName string) (varName string) {
	varName = strings.ToUpper(strings.ReplaceAll(strings.ReplaceAll(fileName, ".", "_"), "-", "_"))
	return
}

func main() {
	files, err := ioutil.ReadDir(assetsFolder)
	check(err)
	assetsFile, err := os.Create("./cli/assets.go")
	check(err)
	_, err = assetsFile.WriteString(`package cli

var (
`)
	check(err)
	for _, file := range files {
		content, err := ioutil.ReadFile(assetsFolder + "/" + file.Name())
		check(err)
		_, err = assetsFile.WriteString("	" + formatName(file.Name()) + " = []byte{") // Write const var name
		check(err)
		var sb strings.Builder
		for i, b := range content {
			sb.WriteString(strconv.Itoa(int(b)))
			if i != len(content)-1 {
				sb.WriteRune(',')
			}
		}
		_, err = assetsFile.WriteString(sb.String())
		check(err)
		_, err = assetsFile.WriteString("}\n")
		check(err)
	}
	_, err = assetsFile.WriteString(")\n")
	check(err)
	fmt.Println("Successfully included assets !")
}
