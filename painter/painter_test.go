package painter

import (
	"testing"
)

func TestDrawWordField(t *testing.T) {
	lettersNumber := 10

	result := DrawWordField(lettersNumber)

	expectations := "__________"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}

func TestAddLetter(t *testing.T) {
	index := 4
	letter := "л"
	wordField := "рамб_ер"

	result := AddLetter(index, letter, wordField)

	expectations := "рамблер"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}

func TestDrawGallows(t *testing.T) {
	result := DrawGallows()

	expectations := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||\n" +
		" ||\n" +
		" ||\n" +
		" ||\n" +
		" ||\n" +
		" /\\\n" +
		"/  \\\n"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}

func TestAddHangmanElement(t *testing.T) {
	mistakesNumber := 3

	result := AddHangmanElement(mistakesNumber)

	expectations := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||         { }\n" +
		" ||        / |\n" +
		" ||       /  |\n" +
		" ||\n" +
		" ||\n" +
		" /\\\n" +
		"/  \\\n"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}

func TestAddHangmanHead(t *testing.T) {
	result := addHangmanHead()

	expectations := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||         { }\n" +
		" ||\n" +
		" ||\n" +
		" ||\n" +
		" ||\n" +
		" /\\\n" +
		"/  \\\n"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}

func TestAddHangmanBody(t *testing.T) {
	result := addHangmanBody()

	expectations := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||         { }\n" +
		" ||          |\n" +
		" ||          |\n" +
		" ||\n" +
		" ||\n" +
		" /\\\n" +
		"/  \\\n"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}

func TestAddHangmanLeftHand(t *testing.T) {
	result := addHangmanLeftHand()

	expectations := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||         { }\n" +
		" ||        / |\n" +
		" ||       /  |\n" +
		" ||\n" +
		" ||\n" +
		" /\\\n" +
		"/  \\\n"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}

func TestAddHangmanRightHand(t *testing.T) {
	result := addHangmanRightHand()

	expectations := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||         { }\n" +
		" ||        / | \\\n" +
		" ||       /  |  \\\n" +
		" ||\n" +
		" ||\n" +
		" /\\\n" +
		"/  \\\n"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}

func TestAddHangmanLeftLeg(t *testing.T) {
	result := addHangmanLeftLeg()

	expectations := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||         { }\n" +
		" ||        / | \\\n" +
		" ||       /  |  \\\n" +
		" ||         /\n" +
		" ||        /\n" +
		" /\\       /\n" +
		"/  \\\n"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}

func TestAddHangmanRightLeg(t *testing.T) {
	result := addHangmanRightLeg()

	expectations := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||         { }\n" +
		" ||        / | \\\n" +
		" ||       /  |  \\\n" +
		" ||         / \\\n" +
		" ||        /   \\\n" +
		" /\\       /     \\\n" +
		"/  \\\n"

	if result != expectations {
		t.Fatalf("\nresult: %s\nexpectatios: %s", result, expectations)
	}
}
