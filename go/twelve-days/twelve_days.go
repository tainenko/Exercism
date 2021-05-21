package twelve

import (
	"fmt"
	"strings"
)

var num2Ordinal = []string{
	"none",
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth"}

var gifts = []string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"and a Partridge in a Pear Tree",
}

// Song return the lyrics to 'The Twelve Days of Christmas'.
func Song() string {
	var s []string
	for i := 1; i <= 12; i++ {
		s = append(s, Verse(i))
	}
	return strings.Join(s[:], "\n")
}

func giftList(num int) string {
	if num == len(gifts)-1 {
		return "a Partridge in a Pear Tree"
	}
	return strings.Join(gifts[num:], ", ")
}

// Verse return nth sentence of 'The Twelve Days of Christmas'.
func Verse(num int) string {
	ord := num2Ordinal[num]
	gift := giftList(len(gifts) - num)
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", ord, gift)
}
