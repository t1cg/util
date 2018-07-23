package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	path := os.Getenv("SWAGGERTAGS_DIR")

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	sm := SwaggerModels{
		DirectoryPath: path,
		Files:         files,
	}

	sm.processAllFiles()
}

type SwaggerModels struct {
	DirectoryPath string
	Files         []os.FileInfo
}

func (sm SwaggerModels) processAllFiles() {
	for _, f := range sm.Files {
		sm.processFile(f.Name())
	}
}

func (sm SwaggerModels) processFile(name string) {
	file, err := os.Open(sm.DirectoryPath + "/" + name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newFile := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if sm.isMissingBson(line) {
			newFile += sm.addBSON(line) + "\n"
		} else {
			newFile += line + "\n"
		}
	}

	err = ioutil.WriteFile(sm.DirectoryPath+"/"+name, []byte(newFile), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (sm SwaggerModels) addBSON(line string) string {
	newLine := ""

	split := strings.Split(line, "`")

	newLine += split[0] + " `" + split[1]
	newLine += " bson:\"" + strings.Split(split[1], "\"")[1] + "\"`"

	return newLine
}

func (sm SwaggerModels) isMissingBson(line string) bool {
	hasJSON := strings.Contains(line, "`json:\"")
	hasBSON := strings.Contains(line, "bson:\"")

	return hasJSON && !hasBSON
}
