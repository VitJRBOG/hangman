package painter

import "strings"

// DrawWordField строит поле для слова из символов
func DrawWordField(lettersNumber int) string {
	var wordField string
	for i := 0; i < lettersNumber; i++ {
		wordField += "_"
	}
	return wordField
}

// AddLetter добавляет букву в поле
func AddLetter(index int, letter, wordField string) string {
	letters := strings.Split(wordField, "")
	letters[index] = letter
	wordField = strings.Join(letters, "")
	return wordField
}

// DrawGallows строит виселицу из символов
func DrawGallows() string {
	gallowsImage := `------------
                    ||          |
                    ||
                    ||
                    ||
                    ||
                    ||
                    /\
                   /  \`
	return gallowsImage
}

// AddHangmanElement добавляет элемент висельника к висельнице
func AddHangmanElement(mistakesNumber int) string {
	var hangmanImage string
	switch mistakesNumber {
	case 1:
		hangmanImage = addHangmanHead()
	case 2:
		hangmanImage = addHangmanBody()
	case 3:
		hangmanImage = addHangmanLeftHand()
	case 4:
		hangmanImage = addHangmanRightHand()
	case 5:
		hangmanImage = addHangmanLeftLeg()
	case 6:
		hangmanImage = addHangmanRightLeg()
	}
	return hangmanImage
}

// addHangmanHead добавляет голову висельника
func addHangmanHead() string {
	hangmanImage := `------------
                    ||          |
                    ||         { }
                    ||
                    ||
                    ||
                    ||
                    /\
                   /  \`
	return hangmanImage
}

// addHangmanBody добавляет тело висельника
func addHangmanBody() string {
	hangmanImage := `------------
                    ||          |
                    ||         { }
                    ||          |
                    ||          |
                    ||
                    ||
                    /\
                   /  \`
	return hangmanImage
}

// addHangmanLeftHand добавляет левую руку висельника
func addHangmanLeftHand() string {
	hangmanImage := `------------
                    ||          |
                    ||         { }
                    ||        / |
                    ||       /  |
                    ||
                    ||
                    /\
                   /  \`
	return hangmanImage
}

// addHangmanRightHand добавляет правую руку висельника
func addHangmanRightHand() string {
	hangmanImage := `------------
                    ||          |
                    ||         { }
                    ||        / | \
                    ||       /  |  \
                    ||
                    ||
                    /\
                   /  \`
	return hangmanImage
}

// addHangmanLeftLeg добавляет левую ногу висельника
func addHangmanLeftLeg() string {
	hangmanImage := `------------
                    ||          |
                    ||         { }
                    ||        / | \
                    ||       /  |  \
                    ||         /
                    ||        /
                    /\       /
                   /  \`
	return hangmanImage
}

// addHangmanRightLeg добавляет правую ногу висельника
func addHangmanRightLeg() string {
	hangmanImage := `------------
                    ||          |
                    ||         { }
                    ||        / | \
                    ||       /  |  \
                    ||         / \
                    ||        /   \
                    /\       /     \
                   /  \`
	return hangmanImage
}
