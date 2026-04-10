package tests

import (
	"strings"
	"testing"
	"time"

	"scout-cli/internal/domain"
	"scout-cli/internal/service"
)

func TestBuildMarkdown(t *testing.T) {
	rs := service.NewReportService()

	md := rs.BuildMarkdown(domain.ScanResult{
		Image:           "demo:latest",
		GeneratedAt:     time.Date(2026, 4, 5, 12, 0, 0, 0, time.UTC),
		RawOutput:       "No vulnerabilities found",
		ExecutionStatus: "success",
		Messages:        []string{"Todo correcto"},
	})

	if !strings.Contains(md, "demo:latest") {
		t.Fatalf("el markdown debe incluir el nombre de la imagen")
	}

	if !strings.Contains(md, "No vulnerabilities found") {
		t.Fatalf("el markdown debe incluir la salida cruda")
	}
}