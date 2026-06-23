package internal

import (
	"encoding/csv"
	"fmt"
	"os"
)

type OutputFormat interface {
	Header(data []string) error
	Write(data []interface{}) error
	Close()
}

type CSVFormatter struct {
	w *csv.Writer
}

func NewCSVFormatter(sep rune) *CSVFormatter {
	w := csv.NewWriter(os.Stdout)
	w.Comma = sep
	return &CSVFormatter{w: w}
}

func (s *CSVFormatter) Close() {
	s.w.Flush()
	s.w = nil
}

func (s *CSVFormatter) Header(data []string) error {
	if s.w == nil {
		return fmt.Errorf("csv formatter is closed")
	}

	return s.w.Write(data)
}

func (s *CSVFormatter) Write(data []interface{}) error {
	if s.w == nil {
		return fmt.Errorf("csv formatter is closed")
	}

	dataStr := make([]string, len(data))
	for i, d := range data {
		switch v := d.(type) {
		case nil:
			dataStr[i] = ""
		case []byte:
			dataStr[i] = string(v)
		default:
			dataStr[i] = fmt.Sprint(v)
		}
	}
	return s.w.Write(dataStr)
}
