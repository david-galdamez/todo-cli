package utils

import "strconv"

func ParseUint(id string) (uint, error) {
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, err
	}
	todoId := uint(parsedId)

	return todoId, nil
}
