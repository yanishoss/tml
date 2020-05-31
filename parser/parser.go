package parser

import (
	"errors"
	"fmt"
	"github.com/yanishoss/tml/lexer"
	"strconv"
)

var (
	ErrInvalidRPE  = errors.New("invalid RPE has been provided")
	ErrInvalidUnit = errors.New("invalid unit has been provided")
)

type Row struct {
	Sets   int      `json:"sets"`
	Reps   int      `json:"reps"`
	RPE    *float64 `json:"rpe,omitempty"`
	Weight float64  `json:"weight"`
	Unit   string   `json:"unit"`
}

type Exercise struct {
	Name    string  `json:"name"`
	Rows    []Row   `json:"rows"`
	Comment *string `json:"comment,omitempty"`
}

type Workout struct {
	Comment   *string    `json:"comment,omitempty"`
	Exercises []Exercise `json:"exercises"`
}

type Parser interface {
	Parse() (*Workout, error)
}

type Config struct {
	DefaultUnit string
	ValidUnits  []string
	RPERange    [2]float64
}

type parser struct {
	l lexer.Lexer
	Config
}

func New(l lexer.Lexer, conf Config) Parser {
	return &parser{l, conf}
}

func (p *parser) isValidUnit(u string) bool {
	for _, unit := range p.ValidUnits {
		if unit == u {
			return true
		}
	}

	return false
}

func (p *parser) isValidRPE(rpe float64) bool {
	return rpe >= p.RPERange[0] && rpe <= p.RPERange[1]
}

func (p *parser) parseExercise() (Exercise, error) {
	e := Exercise{}

	name := p.l.NextToken()

	if name.Type != lexer.IDENT {
		return e, errors.New(fmt.Sprintf("bad token: Token{%s, %s}, before: Token{%s, %s}", name.Type, name.Literal, p.l.PeekToken().Type, p.l.PeekToken().Literal))
	}

	e.Name = name.Literal

	// Skip DELIMITER
	p.l.NextToken()

	rows, err := p.parseRows()

	if err != nil {
		return e, err
	}

	e.Rows = rows

	tok := p.l.NextToken()

	for tok.Type == lexer.DELIMITER {
		tok = p.l.NextToken()
	}

	if tok.Type == lexer.TEXT {
		e.Comment = &tok.Literal
	}

	return e, nil
}

func (p *parser) parseRows() ([]Row, error) {
	rows := make([]Row, 0)

	for true {
		for p.l.PeekToken().Type == lexer.DELIMITER {
			p.l.NextToken()
		}

		if p.l.PeekToken().Type == lexer.TEXT {
			return rows, nil
		}

		if p.l.PeekToken().Type == lexer.HASH {
			return rows, nil
		}

		if p.l.PeekToken().Type == lexer.EOF {
			return rows, nil
		}

		row, err := p.parseRow()

		if err != nil {
			return rows, err
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func (p *parser) parseRow() (Row, error) {
	var err error
	row := Row{}
	row.Sets = 1

	if !p.isValidUnit(p.DefaultUnit) {
		return row, ErrInvalidUnit
	}

	row.Unit = p.DefaultUnit

	tok := p.l.NextToken()

	if tok.Type != lexer.NUMBER {
		return row, errors.New(fmt.Sprintf("bad token: Token{%s, %s}, before: Token{%s, %s}", tok.Type, tok.Literal, p.l.PeekToken().Type, p.l.PeekToken().Literal))
	}

	row.Weight, err = strconv.ParseFloat(tok.Literal, 64)

	if err != nil {
		return row, err
	}

	tok = p.l.NextToken()

	if tok.Type == lexer.LABEL {
		if !p.isValidUnit(tok.Literal) {
			return row, ErrInvalidUnit
		}

		row.Unit = tok.Literal
		tok = p.l.NextToken()
	}

	if tok.Type == lexer.CROSS {
		tok = p.l.NextToken()
	} else {
		return row, errors.New(fmt.Sprintf("bad token: Token{%s, %s}, before: Token{%s, %s}", tok.Type, tok.Literal, p.l.PeekToken().Type, p.l.PeekToken().Literal))
	}

	if tok.Type != lexer.NUMBER {
		return row, errors.New(fmt.Sprintf("bad token: Token{%s, %s}, before: Token{%s, %s}", tok.Type, tok.Literal, p.l.PeekToken().Type, p.l.PeekToken().Literal))
	}

	row.Reps, err = strconv.Atoi(tok.Literal)

	if err != nil {
		return row, err
	}

	tok = p.l.NextToken()

	if tok.Type == lexer.CROSS {
		tok = p.l.NextToken()

		if tok.Type != lexer.NUMBER {
			return row, errors.New(fmt.Sprintf("bad token: Token{%s, %s}, before: Token{%s, %s}", tok.Type, tok.Literal, p.l.PeekToken().Type, p.l.PeekToken().Literal))
		}

		row.Sets, err = strconv.Atoi(tok.Literal)

		if err != nil {
			return row, err
		}
		tok = p.l.NextToken()
	} else if tok.Type == lexer.AT {
		tok = p.l.NextToken()

		if tok.Type != lexer.NUMBER {
			return row, errors.New(fmt.Sprintf("bad token: Token{%s, %s}, before: Token{%s, %s}", tok.Type, tok.Literal, p.l.PeekToken().Type, p.l.PeekToken().Literal))
		}

		rpe, err := strconv.ParseFloat(tok.Literal, 64)

		if err != nil {
			return row, err
		}

		if !p.isValidRPE(rpe) {
			return row, ErrInvalidRPE
		}

		row.RPE = &rpe
		tok = p.l.NextToken()
	}

	if tok.Type == lexer.AT {
		tok = p.l.NextToken()

		if tok.Type != lexer.NUMBER {
			return row, errors.New(fmt.Sprintf("bad token: Token{%s, %s}, before: Token{%s, %s}", tok.Type, tok.Literal, p.l.PeekToken().Type, p.l.PeekToken().Literal))
		}

		rpe, err := strconv.ParseFloat(tok.Literal, 64)

		if err != nil {
			return row, err
		}

		if !p.isValidRPE(rpe) {
			return row, ErrInvalidRPE
		}

		row.RPE = &rpe

		tok = p.l.NextToken()
	}

	if tok.Type != lexer.DELIMITER {
		return row, errors.New(fmt.Sprintf("bad token: Token{%s, %s}, before: Token{%s, %s}", tok.Type, tok.Literal, p.l.PeekToken().Type, p.l.PeekToken().Literal))
	}

	if row.Unit == "count" {
		row.Sets = row.Reps
		row.Reps = int(row.Weight)
		row.Weight = 1
	}

	return row, nil
}

func (p *parser) Parse() (*Workout, error) {
	w := &Workout{
		Comment:   nil,
		Exercises: make([]Exercise, 0),
	}

	for true {
		tok := p.l.NextToken()

		if tok.Type == lexer.EOF {
			break
		}

		switch tok.Type {
		case lexer.DELIMITER:
			continue
		case lexer.TEXT:
			w.Comment = &tok.Literal
			break
		case lexer.HASH:
			e, err := p.parseExercise()

			if err != nil {
				return w, err
			}

			w.Exercises = append(w.Exercises, e)
			break
		default:
			return w, errors.New(fmt.Sprintf("bad token: Token{%s, %s}, before: Token{%s, %s}", tok.Type, tok.Literal, p.l.PeekToken().Type, p.l.PeekToken().Literal))
		}
	}

	return w, nil
}
