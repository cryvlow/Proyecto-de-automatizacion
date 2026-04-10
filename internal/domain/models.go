package domain

import "time"

type Vulnerability struct {
	ID       string `json:"id"`
	Severity string `json:"severity"`
	Package  string `json:"package"`
	Version  string `json:"version"`
	FixedIn  string `json:"fixedIn,omitempty"`
}

type ScanResult struct {
	Image           string          `json:"image"`
	GeneratedAt     time.Time       `json:"generatedAt"`
	Findings        []Vulnerability `json:"findings"`
	RawOutput       string          `json:"rawOutput"`
	Messages        []string        `json:"messages"`
	ExecutionStatus string          `json:"executionStatus"`
}