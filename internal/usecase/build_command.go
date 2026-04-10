package usecase

import (
	"flag"
	"fmt"
	"io"

	"scout-cli/internal/service"
)

type BuildCommand struct {
	out    io.Writer
	docker *service.DockerService
}

func NewBuildCommand(out io.Writer, docker *service.DockerService) *BuildCommand {
	return &BuildCommand{
		out:    out,
		docker: docker,
	}
}

func (c *BuildCommand) Name() string {
	return "build"
}

func (c *BuildCommand) Description() string {
	return "Construye la imagen Docker del proyecto"
}

func (c *BuildCommand) Execute(args []string) error {
	fs := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	fs.SetOutput(c.out)

	image := fs.String("image", "scout-cli:latest", "Nombre de la imagen Docker")
	contextDir := fs.String("context", ".", "Directorio de contexto para Docker build")

	if err := fs.Parse(args); err != nil {
		return err
	}

	output, err := c.docker.BuildImage(*image, *contextDir)
	if output != "" {
		_, _ = fmt.Fprint(c.out, output)
	}
	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(c.out, "Imagen construida correctamente:", *image)
	return err
}