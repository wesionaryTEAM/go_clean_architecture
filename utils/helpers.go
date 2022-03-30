package utils

// StatusInList -> checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// ItemInList --> checks if the given item is in the list
func ItemInList[T comparable](item T, itemList []T) bool {
	for _, i := range itemList {
		if i == item {
			return true
		}
	}
	return false
}

// RemoveDuplicate --> remove duplicate items from  array or slice of comparable datatypes
func RemoveDuplicate[K comparable](s []K) (result []K) {
	inResultMap := make(map[K]bool)

	for _, iter := range s {
		if _, ok := inResultMap[iter]; !ok {
			inResultMap[iter] = true
			result = append(result, iter)
		}
	}
	return result
}
