package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Verses return multiple verses
func Verses(upperBound, lowerBound int) (string, error) {
	if upperBound <= lowerBound {
		return "", errors.New("start less than stop")
	}
	builder := strings.Builder{}
	for i := upperBound; i >= lowerBound; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		builder.WriteString(verse)
		builder.WriteString("\n")
	}
	return builder.String(), nil
}

// Verse return single verse
func Verse(n int) (string, error) {
	if n < 0 {
		return "", errors.New("invalid stop")
	}
	if n > 99 {
		return "", errors.New("invalid start")
	}
	if n == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	}
	if n == 1 {
		return "1 bottle of beer on the wall, 1 bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", nil
	}
	if n == 2 {
		return "2 bottles of beer on the wall, 2 bottles of beer.\nTake one down and pass it around, 1 bottle of beer on the wall.\n", nil
	}
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
}

// Song return the lyrics: 99 Bottles of Beer on the Wall
func Song() string {
	actual, _ := Verses(99, 0)
	return actual
}
