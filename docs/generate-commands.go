package main

import (
	"log"

	"github.com/open-feature/cli/cmd"
)

const docPath = "./docs/commands"

// GenerateDoc generates cobra docs of the cmd
func main() {
	if err := cmd.GenerateDoc(docPath); err != nil {
		log.Fatal(err)
	}
}
