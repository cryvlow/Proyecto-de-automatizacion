package usecase

import (
	"fmt"
	"io"

	"scout-cli/internal/adapters"
	"scout-cli/internal/domain"
	"scout-cli/internal/service"
)

type App struct {
	commands map[string]domain.Command
	out      io.Writer
	errOut   io.Writer
}

func NewApp(out io.Writer, errOut io.Writer) *App {
	return NewAppWithRunner(out, errOut, adapters.ExecRunner{})
}

func NewAppWithRunner(out io.Writer, errOut io.Writer, runner service.Runner) *App {
	dockerService := service.NewDockerService(runner)
	scoutService := service.NewScoutService(runner)
	reportService := service.NewReportService()

	cmds := map[string]domain.Command{
		"help":    NewHelpCommand(out),
		"build":   NewBuildCommand(out, dockerService),
		"scan":    NewScanCommand(out, scoutService),
		"analyze": NewAnalyzeCommand(out, scoutService),
		"report":  NewReportCommand(out, scoutService, reportService),
	}

	return &App{
		commands: cmds,
		out:      out,
		errOut:   errOut,
	}
}

func (a *App) Run(args []string) error {
	if len(args) == 0 {
		return a.commands["help"].Execute(nil)
	}

	commandName := args[0]
	command, ok := a.commands[commandName]
	if !ok {
		_, _ = fmt.Fprintln(a.errOut, "comando no encontrado:", commandName)
		_, _ = fmt.Fprintln(a.errOut, "usa: scout-cli help")
		return fmt.Errorf("comando no encontrado: %s", commandName)
	}

	return command.Execute(args[1:])
}