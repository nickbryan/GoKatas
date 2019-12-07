package internal

import (
	"fmt"
	"io"

	"github.com/nickbryan/GoKatas/datamunging"
)

type App struct {
	Input                io.Reader
	Output               io.Writer
	IndexCol, ACol, BCol string
}

func NewApp(input io.Reader, output io.Writer, i, a, b string) *App {
	return &App{Input: input, Output: output, IndexCol: i, ACol: a, BCol: b}
}

func (a *App) Run() error {
	var err error

	rows := datamunging.Rows{}
	if err = rows.ReadFrom(a.Input, a.IndexCol, a.ACol, a.BCol); err != nil {
		return fmt.Errorf("unable to parse file: %w", err)
	}

	if err = datamunging.WriteMinSpread(&rows, a.Output); err != nil {
		return fmt.Errorf("unable to write to Output: %w", err)
	}

	return nil
}
