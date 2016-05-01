package main

import (
	"github.com/gocarina/gocsv"
	"testing"
)

type Type struct {
	Headerless string `csv:""`
	Headered   string `csv:"header"`
}

func TestNoHeader(t *testing.T) {
	types := []*Type{}
	csv := `,header
foo,bar
baz,quux`

	if err := gocsv.UnmarshalString(csv, &types); err != nil {
		t.Fatal(err)
	}

	for i, tc := range []struct {
		actual   string
		expected string
	}{
		{types[0].Headerless, "foo"},
		{types[0].Headered, "bar"},
		{types[1].Headerless, "baz"},
		{types[1].Headered, "quux"},
	} {
		if tc.actual != tc.expected {
			t.Errorf("Test case %d: Got \"%s\", expected \"%s\"", i, tc.actual, tc.expected)
		}
	}
}
