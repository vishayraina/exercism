package phonenumber

import (
    "regexp"
    "fmt"
    "strings"
    "errors"
)

func Number(phoneNumber string) (string, error) {
	re := regexp.MustCompile(`\d`)
    numList := re.FindAllString(phoneNumber, -1)
    if len(numList) < 10 || len(numList) > 11 {
        return "", errors.New("")
    }
    numString := strings.Join(numList, "")
    if string(numString[0]) == "1" {
        numString = numString[1:]
    }
    if len(numString) != 10 {
        return "", errors.New("")
    } else if string(numString[0]) == "1" || string(numString[0]) == "0" {
        return "", errors.New("")
    } else if string(numString[3]) == "1" || string(numString[3]) == "0" {
        return "", errors.New("")
    }
    return numString, nil
}

func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
    if err != nil {
        return "", errors.New("")
    }
    return num[:3], nil
}

func Format(phoneNumber string) (string, error) {
    num, errNum := Number(phoneNumber)
    ac, errAc := AreaCode(phoneNumber)
    
    if errNum == nil && errAc == nil {
        return fmt.Sprintf("(%s) %s-%s", ac, num[3:6], num[6:]), nil
    }
	return "", errors.New("")
}
