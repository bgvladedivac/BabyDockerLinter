package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const dataDirectoryPath string = "../data"
const validDockerFileParentDir string = "/valid"
const invalidDockerFilePartenDir string = "/invalid"
const FullPathToValidDockerfile string = dataDirectoryPath + validDockerFileParentDir
const FullPathToInvalidDockerfile string = dataDirectoryPath + invalidDockerFilePartenDir

func GetDockerfileInstructions(dockerfileDirectory string) map[string]string {
	fullDockerFilePath := dockerfileDirectory + "/Dockerfile"
	f, err := os.Open(fullDockerFilePath)

	if err != nil {
		log.Fatal("Can't open file over", fullDockerFilePath)
	}

	result := make(map[string]string, 0)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		lineTokens := strings.Split(line, " ")
		result[lineTokens[0]] = lineTokens[1]
	}

	return result
}
