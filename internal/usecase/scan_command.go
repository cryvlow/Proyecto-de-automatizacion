package usecase

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"strings"
	"time"

	"scout-cli/internal/domain"
	"scout-cli/internal/service"
)

type ScanCommand struct {
	out   io.Writer
	scout *service.ScoutService
}

func NewScanCommand(out io.Writer, scout *service.ScoutService) *ScanCommand {
	return &ScanCommand{
		out:   out,
		scout: scout,
	}
}

func (c *ScanCommand) Name() string {
	return "scan"
}

func (c *ScanCommand) Description() string {
	return "Escanea una imagen con Docker Scout CLI"
}

func (c *ScanCommand) Execute(args []string) error {
	fs := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	fs.SetOutput(c.out)

	image := fs.String("image", "", "Imagen a analizar")
	severities := fs.String("severities", "critical,high", "Severidades separadas por coma")
	asJSON := fs.Bool("json", false, "Mostrar salida en JSON")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if strings.TrimSpace(*image) == "" {
		return fmt.Errorf("debes indicar -image")
	}

	severityList := splitCSV(*severities)
	output, err := c.scout.ScanCves(*image, severityList)

	result := domain.ScanResult{
		Image:           *image,
		GeneratedAt:     time.Now(),
		RawOutput:       output,
		Messages:        []string{"Escaneo ejecutado con Docker Scout CLI"},
		ExecutionStatus: "success",
	}

	if err != nil {
		result.ExecutionStatus = "failed"
		result.Messages = append(result.Messages, "El comando devolvió error")
	}

	if *asJSON {
		enc := json.NewEncoder(c.out)
		enc.SetIndent("", "  ")
		if encodeErr := enc.Encode(result); encodeErr != nil {
			return encodeErr
		}
	} else {
		_, _ = fmt.Fprintln(c.out, "Imagen analizada:", result.Image)
		_, _ = fmt.Fprintln(c.out, "Estado:", result.ExecutionStatus)
		if result.RawOutput != "" {
			_, _ = fmt.Fprintln(c.out, "Salida Docker Scout:")
			_, _ = fmt.Fprintln(c.out, result.RawOutput)
		}
	}

	return err
}