package service

import "fmt"

type DockerService struct {
	Runner Runner
}

func NewDockerService(r Runner) *DockerService {
	return &DockerService{Runner: r}
}

func (d *DockerService) BuildImage(tag string, contextDir string) (string, error) {
	if tag == "" {
		return "", fmt.Errorf("la etiqueta de la imagen no puede estar vacía")
	}
	if contextDir == "" {
		contextDir = "."
	}
	return d.Runner.CombinedOutput("docker", "build", "-t", tag, contextDir)
}

func (d *DockerService) RunImage(tag string, args ...string) (string, error) {
	if tag == "" {
		return "", fmt.Errorf("la imagen no puede estar vacía")
	}
	cmdArgs := append([]string{"run", "--rm", tag}, args...)
	return d.Runner.CombinedOutput("docker", cmdArgs...)
}