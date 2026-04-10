package usecase

import (
	"flag"
	"fmt"
	"io"
	"time"

	"scout-cli/internal/adapters"
	"scout-cli/internal/domain"
	"scout-cli/internal/service"
)

type ReportCommand struct {
	out    io.Writer
	scout  *service.ScoutService
	report *service.ReportService
	fs     adapters.FileSystem
}

func NewReportCommand(out io.Writer, scout *service.ScoutService, report *service.ReportService) *ReportCommand {
	return &ReportCommand{
		out:    out,
		scout:  scout,
		report: report,
		fs:     adapters.FileSystem{},
	}
}

func (c *ReportCommand) Name() string {
	return "report"
}

func (c *ReportCommand) Description() string {
	return "Genera un reporte Markdown o JSON"
}

func (c *ReportCommand) Execute(args []string) error {
	fs := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	fs.SetOutput(c.out)

	image := fs.String("image", "", "Imagen a analizar")
	outFile := fs.String("out", "docs/report.md", "Ruta del archivo de salida")
	format := fs.String("format", "md", "Formato: md o json")
	severities := fs.String("severities", "critical,high", "Severidades separadas por coma")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *image == "" {
		return fmt.Errorf("debes indicar -image")
	}

	scanOutput, err := c.scout.ScanCves(*image, splitCSV(*severities))

	result := domain.ScanResult{
		Image:           *image,
		GeneratedAt:     time.Now(),
		RawOutput:       scanOutput,
		Messages:        []string{"Reporte generado desde Docker Scout CLI"},
		ExecutionStatus: "success",
	}

	if err != nil {
		result.ExecutionStatus = "failed"
		result.Messages = append(result.Messages, "Hubo un error al ejecutar el escaneo")
	}

	switch *format {
	case "json":
		content, buildErr := c.report.BuildJSON(result)
		if buildErr != nil {
			return buildErr
		}
		if err := c.fs.WriteFile(*outFile, content); err != nil {
			return err
		}
	default:
		content := c.report.BuildMarkdown(result)
		if err := c.fs.WriteFile(*outFile, []byte(content)); err != nil {
			return err
		}
	}

	_, _ = fmt.Fprintln(c.out, "Reporte generado en:", *outFile)
	return err
}