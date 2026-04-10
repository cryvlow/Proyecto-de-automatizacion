package tests

import (
	"bytes"
	"testing"

	"scout-cli/internal/usecase"
)

type mockRunner struct {
	output string
}

func (m *mockRunner) CombinedOutput(name string, args ...string) (string, error) {
	return m.output, nil
}

func TestAppHelp(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	a := usecase.NewAppWithRunner(&out, &errOut, &mockRunner{output: "ok"})
	if err := a.Run([]string{}); err != nil {
		t.Fatalf("no se esperaba error: %v", err)
	}

	if out.Len() == 0 {
		t.Fatal("la ayuda no imprimió nada")
	}
}

func TestAppUnknownCommand(t *testing.T) {
	var out bytes.Buffer
	var errOut bytes.Buffer

	a := usecase.NewAppWithRunner(&out, &errOut, &mockRunner{output: "ok"})
	if err := a.Run([]string{"inexistente"}); err == nil {
		t.Fatal("se esperaba error por comando inexistente")
	}
}