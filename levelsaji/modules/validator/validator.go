package validator

import (
	"errors"
	"strconv"
)

type Validator struct{}

func (validator *Validator) CheckString(name string, value string) (string, error) {
	if len(value) <= 0 {
		return "", errors.New(name + "を入力してください")
	}
	return value, nil
}

func (validator *Validator) CheckNumeric(name string, value string, options ...int) (int, error) {
	if len(value) <= 0 {
		return 0, errors.New(name + "を入力してください")
	} else {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return 0, errors.New(name + "は数値のみで入力してください")
		} else {
			if len(options) >= 1 {
				min := options[0]
				if i < int64(min) {
					return 0, errors.New(name + "は" + strconv.Itoa(min) + "以上で入力してください")
				}
			}
			if len(options) >= 2 {
				max := options[1]
				if i > int64(max) {
					return 0, errors.New(name + "は" + strconv.Itoa(max) + "以下で入力してください")
				}
			}
		}
	}
	i, _ := strconv.ParseInt(value, 10, 64)
	return int(i), nil
}
