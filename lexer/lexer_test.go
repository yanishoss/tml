package lexer

import "testing"

func TestLexer_NextToken(t *testing.T) {
	l := New(`
some comments about the workout
# Squat
60kg x 3 x 5 @ 6.5
60kg x 3 x 5 @ 6
good squat session

# Deadlift
60kg x 3 x 5 @ 6
60kg x 3 x 5 @ 6.4
good deadlift session
`)

	correctTokens := []Token{
		{DELIMITER, "\n"},
		{TEXT, "some comments about the workout\n"},
		{HASH, "#"},
		{IDENT, "Squat"},
		{DELIMITER, "\n"},
		{NUMBER, "60"},
		{LABEL, "kg"},
		{CROSS, "x"},
		{NUMBER, "3"},
		{CROSS, "x"},
		{NUMBER, "5"},
		{AT, "@"},
		{NUMBER, "6.5"},
		{DELIMITER, "\n"},
		{NUMBER, "60"},
		{LABEL, "kg"},
		{CROSS, "x"},
		{NUMBER, "3"},
		{CROSS, "x"},
		{NUMBER, "5"},
		{AT, "@"},
		{NUMBER, "6"},
		{DELIMITER, "\n"},
		{TEXT, "good squat session\n"},
		{DELIMITER, "\n"},
		{HASH, "#"},
		{IDENT, "Deadlift"},
		{DELIMITER, "\n"},
		{NUMBER, "60"},
		{LABEL, "kg"},
		{CROSS, "x"},
		{NUMBER, "3"},
		{CROSS, "x"},
		{NUMBER, "5"},
		{AT, "@"},
		{NUMBER, "6"},
		{DELIMITER, "\n"},
		{NUMBER, "60"},
		{LABEL, "kg"},
		{CROSS, "x"},
		{NUMBER, "3"},
		{CROSS, "x"},
		{NUMBER, "5"},
		{AT, "@"},
		{NUMBER, "6.4"},
		{DELIMITER, "\n"},
		{TEXT, "good deadlift session\n"},
		{EOF, ""},
	}

	tokens := make([]Token, 0)

	for true {
		tk := l.NextToken()

		tokens = append(tokens, tk)

		t.Logf("Token: Type: %s Literal: %s\n", tk.Type, tk.Literal)

		if tk.Type == EOF {
			break
		}
	}

	if len(correctTokens) != len(tokens) {
		t.Fatalf("incorrect len of tokens: expected:%d, got:%d\n", len(correctTokens), len(tokens))
	}

	for i, tk := range tokens {
		if tk.Type != correctTokens[i].Type {
			t.Fatalf("incorrect token type: expected:%s, got:%s\n", correctTokens[i].Type, tk.Type)
		}

		if tk.Literal != correctTokens[i].Literal {
			t.Fatalf("incorrect token literal: expected:%s, got:%s\n", correctTokens[i].Literal, tk.Literal)
		}
	}
}
