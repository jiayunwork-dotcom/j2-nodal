package table

import (
	"testing"

	"j2-nodal/internal/j2"
)

func TestFromRates(t *testing.T) {
	rate := j2.Rates{A: 7000000, E: 0, I: 90, RAANDot: 1, ArgPeriDot: 2}
	table := FromRates([]j2.Rates{rate})
	if Rows(table) != 1 || Columns(table) != 5 {
		t.Fatalf("table=%+v", table)
	}
}

func TestColumn(t *testing.T) {
	table := Table{Header: []string{"a", "e", "i", "RAAN_dot", "arg"},
		Rows: [][]string{{"1", "2", "3", "4", "5"}, {"2", "3", "4", "6", "7"}}}
	values, err := Column(table, 3)
	if err != nil {
		t.Fatal(err)
	}
	if values[1] != 6 {
		t.Fatalf("values=%v", values)
	}
}

func TestMaxColumn(t *testing.T) {
	table := Table{Header: []string{"a"}, Rows: [][]string{{"1"}, {"3"}}}
	max, err := MaxColumn(table, 0)
	if err != nil {
		t.Fatal(err)
	}
	if max != 3 {
		t.Fatalf("max=%g", max)
	}
}

func TestValidate(t *testing.T) {
	table := Table{Header: []string{"a"}, Rows: [][]string{{"1"}}}
	if err := Validate(table); err != nil {
		t.Fatal(err)
	}
}

func TestTextAndSummary(t *testing.T) {
	table := Table{Header: []string{"a"}, Rows: [][]string{{"1"}}}
	if Text(table) == "" || Summary(table) == "" {
		t.Fatal("empty output")
	}
}
