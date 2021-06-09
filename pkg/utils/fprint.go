package utils

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type tableWriter struct {
	w *tabwriter.Writer
}

func (tw *tableWriter) Print(items []string) {
	tw.w.Init(os.Stdout, 0, 8, 1, '\t', 0)

	for _, item := range items {
		_, err := fmt.Fprintln(tw.w, item)
		if err != nil {
		}
	}

	err := tw.w.Flush()
	if err != nil {
	}
}

func Fprint(items []string) {
	tw := &tableWriter{w: new(tabwriter.Writer)}
	tw.Print(items)
}
