package table

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"j2-nodal/internal/j2"
)

type Table struct {
	Header []string   `json:"header"`
	Rows   [][]string `json:"rows"`
}

func FromRates(rates []j2.Rates) Table {
	table := Table{Header: []string{"a", "e", "i", "RAAN_dot", "arg_peri_dot"}}
	for _, rate := range rates {
		table.Rows = append(table.Rows, []string{
			strconv.FormatFloat(rate.A, 'g', -1, 64),
			strconv.FormatFloat(rate.E, 'g', -1, 64),
			strconv.FormatFloat(rate.I, 'g', -1, 64),
			strconv.FormatFloat(rate.RAANDot, 'g', -1, 64),
			strconv.FormatFloat(rate.ArgPeriDot, 'g', -1, 64),
		})
	}
	return table
}

func SaveCSV(path string, table Table) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	if err := writer.Write(table.Header); err != nil {
		return err
	}
	for _, row := range table.Rows {
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func Text(table Table) string {
	var b strings.Builder
	for i, header := range table.Header {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(header)
	}
	b.WriteString("\n")
	for _, row := range table.Rows {
		for i, cell := range row {
			if i > 0 {
				b.WriteString(" ")
			}
			b.WriteString(cell)
		}
		b.WriteString("\n")
	}
	return b.String()
}

func Summary(table Table) string {
	return fmt.Sprintf("%d rows x %d columns", len(table.Rows), len(table.Header))
}

func Validate(table Table) error {
	if len(table.Header) == 0 {
		return fmt.Errorf("empty header")
	}
	for i, row := range table.Rows {
		if len(row) != len(table.Header) {
			return fmt.Errorf("row %d width mismatch", i)
		}
	}
	return nil
}

func Column(table Table, index int) ([]float64, error) {
	values := make([]float64, 0, len(table.Rows))
	for _, row := range table.Rows {
		if index >= len(row) {
			return nil, fmt.Errorf("row too short")
		}
		value, err := strconv.ParseFloat(row[index], 64)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, nil
}

func MaxColumn(table Table, index int) (float64, error) {
	values, err := Column(table, index)
	if err != nil {
		return 0, err
	}
	if len(values) == 0 {
		return 0, fmt.Errorf("empty column")
	}
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max, nil
}

func MinColumn(table Table, index int) (float64, error) {
	values, err := Column(table, index)
	if err != nil {
		return 0, err
	}
	if len(values) == 0 {
		return 0, fmt.Errorf("empty column")
	}
	min := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	return min, nil
}

func Copy(table Table) Table {
	out := Table{Header: append([]string(nil), table.Header...)}
	for _, row := range table.Rows {
		out.Rows = append(out.Rows, append([]string(nil), row...))
	}
	return out
}

func Rows(table Table) int {
	return len(table.Rows)
}

func Columns(table Table) int {
	return len(table.Header)
}
