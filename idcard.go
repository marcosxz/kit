package kit

import (
	"errors"
	"strconv"
)

type sex int8

const (
	SexUnknown sex = iota
	SexMan
	SexWoman
)

var InvalidIdCardNumber = errors.New("invalid id card number")

// 通过身份证号获取性别
// 15位的身份证号：最后一位奇数表示男，偶数表示女
// 18位的身份证号：倒数第二位奇数表示男，偶数表示女
func IdCardSex(number string) (sex, error) {

	numberLen := len(number)
	if numberLen == 0 {
		return SexUnknown, InvalidIdCardNumber
	}

	var tag string
	if numberLen == 15 {
		tag = number[numberLen-1:]
	} else if numberLen == 18 {
		tag = number[numberLen-2 : numberLen-1]
	} else {
		return SexUnknown, InvalidIdCardNumber
	}

	num, err := strconv.ParseInt(tag, 10, 64)
	if err != nil {
		return SexUnknown, err
	}

	if num%2 == 0 {
		return SexWoman, nil
	} else {
		return SexMan, nil
	}
}
