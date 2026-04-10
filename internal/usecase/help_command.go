package usecase

import (
	"fmt"
	"io"
)

type HelpCommand struct {
	out io.Writer
}

func NewHelpCommand(out io.Writer) *HelpCommand {
	return &HelpCommand{out: out}
}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Description() string {
	return "Muestra la ayuda del programa"
}

func (c *HelpCommand) Execute(_ []string) error {
	_, err := fmt.Fprint(c.out, `Scout CLI

Uso:
  scout-cli <comando> [opciones]

Comandos:
  help     Muestra esta ayuda
  build    Construye la imagen Docker
  scan     Ejecuta Docker Scout sobre una imagen
  analyze  Compara una imagen con otra base
  report   Genera un reporte Markdown o JSON

Ejemplos:
  scout-cli build -image scout-cli:latest
  scout-cli scan -image scout-cli:latest -severities critical,high
  scout-cli analyze -image scout-cli:latest -base scout-cli:old
  scout-cli report -image scout-cli:latest -out docs/report.md -format md
`)
	return err
}