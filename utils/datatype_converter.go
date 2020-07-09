package utils

import (
	"strconv"
)

func ConvertStringToInt(value string) (int, error) {
	converted, err := strconv.Atoi(value)

	if err != nil {
		return 0, err
	}

	return converted, nil
}
