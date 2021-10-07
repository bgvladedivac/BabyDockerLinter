package main

import (
	"fmt"
	"linter/handlers"
	"linter/utils"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	dockerFileLocation := args[0]

	if _, err := os.Stat(dockerFileLocation); os.IsNotExist(err) {
		log.Fatal("Sorry, the provided Absolute Path does not exists.")
	}

	m := utils.GetDockerfileInstructions(dockerFileLocation)

	for key, value := range m {
		switch instruction := key; instruction {
		case "FROM":
			handlers.EvaluateFromInstruction(value)
		case "EXPOSE":
			handlers.EvaluteExposeInstruction(value)
		case "MAINTAINER":
			handlers.EvaluateMaintainerInstruction()
		case "CMD":
			handlers.EvaluateCMDInstruction(dockerFileLocation)
		default:
			fmt.Println("At this stage the instruction is not supported.")
		}
	}
}
