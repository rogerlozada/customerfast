package common

//ContainsValue return true if has a value
func ContainsValue(maps map[string]interface{}) bool {
	return len(maps) > 0
}
