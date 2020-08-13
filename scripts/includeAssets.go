package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
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

const (
`)
	check(err)
	for _, file := range files {
		_, err = assetsFile.WriteString("	" + formatName(file.Name()) + " = `") // Write const var name
		check(err)
		fileO, err := os.Open(assetsFolder + "/" + file.Name())
		check(err)
		io.Copy(assetsFile, fileO)
		fileO.Close()
		_, err = assetsFile.WriteString("`\n")
	}
	_, err = assetsFile.WriteString(")\n")
	check(err)
	fmt.Println("Successfully included assets !")
}
