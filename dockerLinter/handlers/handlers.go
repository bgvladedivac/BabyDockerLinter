package handlers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func EvaluateFromInstruction(instructionContent string) bool {
	tagDelimeter := ":"
	tagDelimiterIndex := strings.Index(instructionContent, tagDelimeter)

	if tagDelimiterIndex == -1 {
		fmt.Println("A tag is missing after FROM instruction.")
		return false
	}

	splitted := strings.Split(instructionContent, tagDelimeter)
	fmt.Printf("The tag is set as %s in the FROM instruction, good job!\n", splitted[1])
	return true
}

func EvaluteExposeInstruction(instructionContent string) bool {
	validStartPort, validEndPort := 1, 65535
	exposeDelimiter := "/"

	port, err := strconv.Atoi(strings.Split(instructionContent, exposeDelimiter)[0])

	if err != nil {
		log.Fatal("There's an error in your EXPOSE instruction, no separator for port/protocol.")
	}

	if port >= validStartPort && port <= validEndPort {
		fmt.Printf("Port %s is between %s and %s inside EXPOSE instruction, good job!\n", strconv.Itoa(port), strconv.Itoa(validEndPort), strconv.Itoa(validEndPort))
		return true
	}

	fmt.Println("The port inside EXPOSE instruction is not supported.")
	return false
}

func EvaluateMaintainerInstruction() {
	fmt.Println("MAINTAINER instruction is deprecated, please remove it from the Dockerfile.")
}

func EvaluateCMDInstruction(dockerfileDirectory string) bool {
	f, err := os.Open(dockerfileDirectory + "/Dockerfile")

	if err != nil {
		log.Fatal("Can't open file over", dockerfileDirectory)
	}
	scanner := bufio.NewScanner(f)

	cmdCounter := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineTokens := strings.Split(line, " ")
		if strings.Compare(lineTokens[0], "CMD") == 0 {
			cmdCounter += 1
		}
	}

	if cmdCounter == 1 {
		fmt.Println("CMD instruction is met just once, good job buddy!")
		return true
	}

	fmt.Println("CMD is met more than once in your Dockerfile, there can only be one CMD instruction in a Dockerfile.")
	return false
}
