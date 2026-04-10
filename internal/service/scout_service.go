package service

import (
	"fmt"
	"strings"
)

type ScoutService struct {
	Runner Runner
}

func NewScoutService(r Runner) *ScoutService {
	return &ScoutService{Runner: r}
}

func (s *ScoutService) ScanCves(image string, severities []string) (string, error) {
	if image == "" {
		return "", fmt.Errorf("la imagen no puede estar vacía")
	}

	args := []string{"scout", "cves", image}
	if len(severities) > 0 {
		args = append(args, "--only-severities", strings.Join(severities, ","))
	}

	return s.Runner.CombinedOutput("docker", args...)
}

func (s *ScoutService) Compare(image string, base string) (string, error) {
	if image == "" || base == "" {
		return "", fmt.Errorf("las imágenes para comparar no pueden estar vacías")
	}

	args := []string{"scout", "compare", image, "--to", base}
	return s.Runner.CombinedOutput("docker", args...)
}