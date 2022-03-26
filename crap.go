package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"os"
)

func writeCRAPReport(vv []violation) error {
	out := os.Stdout
	if filename := os.Getenv("CRAP_REPORT_OUTPUT_FILE"); filename != "" {
		var err error
		out, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			return fmt.Errorf("could not open file '%s': %w", filename, err)
		}
	}

	format := "xml"
	if os.Getenv("CRAP_REPORT_OUTPUT_FORMAT") != "" {
		format = os.Getenv("CRAP_REPORT_OUTPUT_FORMAT")
	}

	if format == "xml" {
		return writeCRAPReportXML(out, vv)
	} else if format == "csv" {
		return writeCRAPReportCSV(out, vv)
	} else {
		return fmt.Errorf("unsupported output format: %s", format)
	}
}

const xmlReportTmpl = `<CrapReport>
{{range . }}<Violation Filename="{{ .Filename }}" Level="{{ .Level }}">{{ .Message }}</Violation>
{{end}}</CrapReport>`

func writeCRAPReportXML(out io.Writer, vv []violation) error {
	t, err := template.New("CRAP Report").Parse(xmlReportTmpl)
	if err != nil {
		return fmt.Errorf("couldn't parse xml template: %w", err)
	}

	err = t.Execute(out, vv)
	if err != nil {
		return fmt.Errorf("couldn't execute template: %w", err)
	}

	return nil
}

func writeCRAPReportCSV(out io.Writer, vv []violation) error {
	w := csv.NewWriter(out)

	err := w.Write([]string{"Filename", "Message", "Level"})
	if err != nil {
		return fmt.Errorf("could not write csv header: %w", err)
	}

	for _, v := range vv {
		err := w.Write([]string{v.Filename, v.Message, v.Level})
		if err != nil {
			return fmt.Errorf("could not write csv line: %w", err)
		}
	}

	w.Flush()
	err = w.Error()
	if err != nil {
		return fmt.Errorf("could not flush csv writer: %w", err)
	}

	return nil
}
