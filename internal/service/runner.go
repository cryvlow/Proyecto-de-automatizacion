package service

type Runner interface {
	CombinedOutput(name string, args ...string) (string, error)
}