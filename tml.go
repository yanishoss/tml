package tml

import (
	"github.com/yanishoss/tml/lexer"
	"github.com/yanishoss/tml/parser"
)

func Parse(input string, conf parser.Config) (*parser.Workout, error) {
	l := lexer.New(input)

	p := parser.New(l, conf)

	return p.Parse()
}

func WithDefaultConfig() parser.Config {
	return parser.Config{
		DefaultUnit: "kg",
		ValidUnits:  []string{"kg", "lbs", "s", "min", "count"},
		RPERange:    [2]float64{0, 11},
	}
}
