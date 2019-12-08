package internal

import (
	"fmt"
	"io"

	"github.com/nickbryan/GoKatas/datamunging"
)

// App is responsible for the handling of input and output within each application.
type App struct {
	Input                io.Reader
	Output               io.Writer
	IndexCol, ACol, BCol string
}

// NewApp will create a new App instance and provides an easy way to set the required data for the App to run.
func NewApp(input io.Reader, output io.Writer, i, a, b string) *App {
	return &App{Input: input, Output: output, IndexCol: i, ACol: a, BCol: b}
}

// Run creates a new Rows set and reads the data from the given Input into Rows. It will then write the minimum spread
// to the Output Writer.
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
