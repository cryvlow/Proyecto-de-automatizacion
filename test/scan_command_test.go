package tests

import (
	"bytes"
	"strings"
	"testing"

	"scout-cli/internal/service"
	"scout-cli/internal/usecase"
)

type scanMockRunner struct {
	calledName string
	calledArgs []string
	output     string
	err        error
}

func (m *scanMockRunner) CombinedOutput(name string, args ...string) (string, error) {
	m.calledName = name
	m.calledArgs = append([]string{}, args...)
	return m.output, m.err
}

func TestScanCommandExecute(t *testing.T) {
	mock := &scanMockRunner{
		output: "scan ok",
	}

	scoutService := service.NewScoutService(mock)
	var out bytes.Buffer
	cmd := usecase.NewScanCommand(&out, scoutService)

	err := cmd.Execute([]string{"-image", "demo:latest"})
	if err != nil {
		t.Fatalf("no se esperaba error, pero llegó: %v", err)
	}

	if mock.calledName != "docker" {
		t.Fatalf("se esperaba docker, llegó %s", mock.calledName)
	}

	if len(mock.calledArgs) == 0 || mock.calledArgs[0] != "scout" {
		t.Fatalf("se esperaba docker scout")
	}

	if !strings.Contains(out.String(), "Imagen analizada: demo:latest") {
		t.Fatalf("salida inesperada: %s", out.String())
	}
}