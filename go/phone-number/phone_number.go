package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var re = regexp.MustCompile("[^0-9]")

// Number return a valid NANP phone number
func Number(nanpNum string) (string, error) {
	nanpNum = re.ReplaceAllString(nanpNum, "")
	if len(nanpNum) > 11 {
		return nanpNum, errors.New("more than 11 digits")
	}
	if len(nanpNum) <= 9 {
		return nanpNum, errors.New("incorrect number of digits")
	}

	if len(nanpNum) == 11 {
		if nanpNum[0] != '1' {
			return nanpNum, errors.New("11 digits must start with 1")
		}
		nanpNum = nanpNum[1:]
	}

	area, exchange := nanpNum[:3], nanpNum[3:6]

	if area[0] == '0' {
		return nanpNum, errors.New("area code cannot start with zero")
	}
	if area[0] == '1' {
		return nanpNum, errors.New("area code cannot start with one")
	}
	if exchange[0] == '0' {
		return nanpNum, errors.New("exchange code cannot start with zero")
	}
	if exchange[0] == '1' {
		return nanpNum, errors.New("exchange code cannot start with one")
	}

	return nanpNum, nil
}

// AreaCode return the area code of phone number
func AreaCode(phone string) (string, error) {
	nanpNum, err := Number(phone)
	return nanpNum[:3], err
}

// Format return the phone number with NANP format, (NXX)-NXX-XXXX.
func Format(phone string) (string, error) {
	nanpNum, err := Number(phone)
	return fmt.Sprintf("(%s) %s-%s", nanpNum[:3], nanpNum[3:6], nanpNum[6:]), err
}
