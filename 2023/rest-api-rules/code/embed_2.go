package main

import (
	"embed"
)

//go:embed assets/*.json
var folder embed.FS

func main() {
	content2, _ := folder.ReadFile("folder/file2.json")
    print(string(content2))
}