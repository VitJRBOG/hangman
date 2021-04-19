package judgeman

import (
	"testing"
)

func TestCheckLetter(t *testing.T) {
	letter := "р"
	word := "рамблер"

	result := CheckLetter(letter, word)

	expectations := []int{0, 6}

	if len(result) != len(expectations) {
		t.Fatalf("\nresult: %d\nexpected: %d\n", result, expectations)
	}

	for i := 0; i < len(result); i++ {
		if result[i] != expectations[i] {
			t.Fatalf("\nresult: %d\nexpected: %d\n", result, expectations)
		}
	}
}

func TestCheckWord(t *testing.T) {
	wordField := "рам_лер"

	result := CheckWord(wordField)

	expectations := false

	if result != expectations {
		t.Fatalf("\nresult: %v\nexpected: %v\n", result, expectations)
	}
}

func TestCheckHangmanImage(t *testing.T) {
	hangmanImage := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||         { }\n" +
		" ||        / | \\\n" +
		" ||       /  |  \\\n" +
		" ||         / \\\n" +
		" ||        /   \\\n" +
		" /\\       /     \\\n" +
		"/  \\\n"

	result := CheckHangmanImage(hangmanImage)

	expectations := true

	if result != expectations {
		t.Fatalf("\nresult: %v\nexpected: %v\n", result, expectations)
	}
}
