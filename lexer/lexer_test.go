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

# Push-up
5 x 3
`)

	correctTokens := []Token{
		{DELIMITER, "\n", 0, 0},
		{TEXT, "some comments about the workout\n", 0, 0},
		{HASH, "#", 0, 0},
		{IDENT, "Squat", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{NUMBER, "60", 0, 0},
		{LABEL, "kg", 0, 0},
		{CROSS, "x", 0, 0},
		{NUMBER, "3", 0, 0},
		{CROSS, "x", 0, 0},
		{NUMBER, "5", 0, 0},
		{AT, "@", 0, 0},
		{NUMBER, "6.5", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{NUMBER, "60", 0, 0},
		{LABEL, "kg", 0, 0},
		{CROSS, "x", 0, 0},
		{NUMBER, "3", 0, 0},
		{CROSS, "x", 0, 0},
		{NUMBER, "5", 0, 0},
		{AT, "@", 0, 0},
		{NUMBER, "6", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{TEXT, "good squat session\n", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{HASH, "#", 0, 0},
		{IDENT, "Deadlift", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{NUMBER, "60", 0, 0},
		{LABEL, "kg", 0, 0},
		{CROSS, "x", 0, 0},
		{NUMBER, "3", 0, 0},
		{CROSS, "x", 0, 0},
		{NUMBER, "5", 0, 0},
		{AT, "@", 0, 0},
		{NUMBER, "6", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{NUMBER, "60", 0, 0},
		{LABEL, "kg", 0, 0},
		{CROSS, "x", 0, 0},
		{NUMBER, "3", 0, 0},
		{CROSS, "x", 0, 0},
		{NUMBER, "5", 0, 0},
		{AT, "@", 0, 0},
		{NUMBER, "6.4", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{TEXT, "good deadlift session\n", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{HASH, "#", 0, 0},
		{IDENT, "Push-up", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{NUMBER, "5", 0, 0},
		{CROSS, "x", 0, 0},
		{NUMBER, "3", 0, 0},
		{DELIMITER, "\n", 0, 0},
		{EOF, "", 0, 0},
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
