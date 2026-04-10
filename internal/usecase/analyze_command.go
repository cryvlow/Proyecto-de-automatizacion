package usecase

import (
	"flag"
	"fmt"
	"io"

	"scout-cli/internal/service"
)

type AnalyzeCommand struct {
	out   io.Writer
	scout *service.ScoutService
}

func NewAnalyzeCommand(out io.Writer, scout *service.ScoutService) *AnalyzeCommand {
	return &AnalyzeCommand{
		out:   out,
		scout: scout,
	}
}

func (c *AnalyzeCommand) Name() string {
	return "analyze"
}

func (c *AnalyzeCommand) Description() string {
	return "Analiza una imagen comparándola con una base"
}

func (c *AnalyzeCommand) Execute(args []string) error {
	fs := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	fs.SetOutput(c.out)

	image := fs.String("image", "", "Imagen nueva a comparar")
	base := fs.String("base", "scout-cli:latest", "Imagen base de referencia")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *image == "" {
		return fmt.Errorf("debes indicar -image")
	}

	output, err := c.scout.Compare(*image, *base)
	if output != "" {
		_, _ = fmt.Fprint(c.out, output)
	}
	return err
}