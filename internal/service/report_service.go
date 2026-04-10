package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"scout-cli/internal/domain"
)

type ReportService struct{}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (r *ReportService) BuildMarkdown(result domain.ScanResult) string {
	var b strings.Builder

	b.WriteString("# Reporte Docker Scout CLI\n\n")
	b.WriteString(fmt.Sprintf("- Imagen analizada: `%s`\n", result.Image))
	b.WriteString(fmt.Sprintf("- Fecha: `%s`\n", result.GeneratedAt.Format(time.RFC3339)))
	b.WriteString(fmt.Sprintf("- Estado: `%s`\n\n", result.ExecutionStatus))

	if len(result.Messages) > 0 {
		b.WriteString("## Mensajes\n")
		for _, msg := range result.Messages {
			b.WriteString("- " + msg + "\n")
		}
		b.WriteString("\n")
	}

	b.WriteString("## Salida cruda\n")
	b.WriteString("```text\n")
	b.WriteString(result.RawOutput)
	b.WriteString("\n```\n")

	return b.String()
}

func (r *ReportService) BuildJSON(result domain.ScanResult) ([]byte, error) {
	return json.MarshalIndent(result, "", "  ")
}