package format

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type (
	// Printer prints objects
	Printer interface {
		Print(obj interface{}) error
	}

	// Printable is an object that can format itself
	Printable interface {
		Print(writer io.Writer, format OutputType) error
	}

	// StdoutPrinter is a printer that prints to os.Stdout
	StdoutPrinter struct {
		Format OutputType
	}

	// OutputType represents the type of output, JSON, XML, TSV, etc.
	OutputType string
)

var (
	// JSONFormat tell the printer to print json
	JSONFormat OutputType = "json"
)

// Print prints an object to os.Stdout
func (stdPrinter StdoutPrinter) Print(obj interface{}) error {
	if printable, ok := obj.(Printable); ok {
		return printable.Print(os.Stdout, stdPrinter.Format)
	}

	switch stdPrinter.Format {
	case JSONFormat:
		return printJSON(os.Stdout, obj)
	default:
		return fmt.Errorf("unable to print %v as type %s", obj, stdPrinter.Format)
	}
}

func printJSON(writer io.Writer, obj interface{}) error {
	bits, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	_, err = fmt.Fprint(writer, string(bits))
	return err
}
