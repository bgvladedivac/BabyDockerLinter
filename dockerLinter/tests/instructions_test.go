package test

import (
	"linter/handlers"
	"linter/utils"
	"os"
	"testing"
)

var validDockerInstructions, invalidDockerInstructions map[string]string

func setup() {
	validDockerInstructions = utils.GetDockerfileInstructions(utils.FullPathToValidDockerfile)
	invalidDockerInstructions = utils.GetDockerfileInstructions(utils.FullPathToInvalidDockerfile)
}

func TestMain(m *testing.M) {
	setup()
	exitVal := m.Run()
	os.Exit(exitVal)
}

//
//func TestTagExists(t *testing.T) {
//	t.Parallel()
//
//	want := true
//	got := handlers.EvaluateFromInstruction(validDockerInstructions["FROM"])
//	if want != got {
//		t.Error("Expecting a tag, but does not have one.")
//	}
//}
//
//func TestTagNotExisting(t *testing.T) {
//	t.Parallel()
//	want := false
//	got := handlers.EvaluateFromInstruction(invalidDockerInstructions["FROM"])
//
//	if want != got {
//		t.Error("Not expecting a tag, but have one.")
//	}
//}
//
//func TestExposeInsideRange(t *testing.T) {
//	t.Parallel()
//	want := true
//	got := handlers.EvaluteExposeInstruction(validDockerInstructions["EXPOSE"])
//
//	if want != got {
//		t.Error("Expecting the port in the Docker file to be in the range of allowed ports, but it's not")
//	}
//}
//
//func TestExposeOutsideOfRange(t *testing.T) {
//	t.Parallel()
//	want := false
//	got := handlers.EvaluteExposeInstruction(invalidDockerInstructions["EXPOSE"])
//
//	if want != got {
//		t.Error("Expecting the port to not be in the valid ranges, but it's")
//	}
//}
//
func TestCMDInstructionDeclaredJustOnce(t *testing.T) {
	t.Parallel()
	want := true
	got := handlers.EvaluateCMDInstruction(utils.FullPathToValidDockerfile)
	if want != got {
		t.Error("Expecting the CMD to be declared just once - flagged with true, but it's declared multiple times.")
	}
}

func TestCMDInstructionDeclatedMultipleTimes(t *testing.T) {
	t.Parallel()
	want := false
	got := handlers.EvaluateCMDInstruction(utils.FullPathToInvalidDockerfile)

	if want != got {
		t.Error("Expceting the CMD to be declared multiple times, but it's not.")
	}
}
