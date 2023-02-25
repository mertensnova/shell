package main

import "encoding/csv"

func RemoveDuplicates(header []string) []string {
	check := make(map[string]int)
	res := make([]string, 0)
	for _, val := range header {
		check[val] = 1
	}

	for letter := range check {
		res = append(res, letter)
	}

	return res
}

type Writer struct {
	w      *csv.Writer
	header []string
}

func NewWriter(w *csv.Writer, header []string) *Writer {
	return &Writer{
		w:      w,
		header: header,
	}
}

func (w *Writer) WriteHeader() error {
	return w.w.Write(w.header)
}

func (w *Writer) Write(record map[string]string) error {
	s := make([]string, len(w.header))
	for i, name := range w.header {
		s[i] = record[name]
	}
	return w.w.Write(s)
}

func (w *Writer) WriteAll(records []map[string]string) error {
	for _, record := range records {
		err := w.Write(record)
		if err != nil {
			return err
		}
	}
	w.w.Flush()
	return w.w.Error()
}
