package helper

import (
	"strings"
	"strconv"
)

// Trim string space
func TrimStringSpace(str string) string {
	return strings.Trim(str, " ")
}

// string to int64
func String2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// string to int32
func String2Int32(str string) int32 {
	int64Str, _ := strconv.Atoi(str)

	return int32(int64Str)
}

// map map[int]string
// map 的 key 必须从 0 开始，顺序排列
// sep is placed between elements in the resulting string.
func ImplodeMapIntString(maps map[int]string, sep string) (str string) {
	mapLen := len(maps)

	for i := 0; i < mapLen; i++ {
		str += maps[i] + sep
	}

	str = strings.TrimRight(str, sep)

	return
}

// map map[int32]string
// map 的 key 必须从 0 开始，顺序排列
// sep is placed between elements in the resulting string.
func ImplodeMapInt32String(maps map[int32]string, sep string) (str string) {
	mapLen := int32(len(maps))

	for i := int32(0); i < mapLen; i++ {
		str += maps[i] + sep
	}

	str = strings.TrimRight(str, sep)

	return
}