package slices_utils

import (
	frameUtils "github.com/qor/qor/utils"
	"my_qor_test/config/consts"
	"strings"
)

//noinspection ALL
const INDEX_NOT_FOUND = -1

// region int
func IndexOfInt(slice []int, valueToFind int, startIndex int) int {
	if slice == nil || len(slice) == 0 {
		return INDEX_NOT_FOUND
	}
	if startIndex < 0 {
		startIndex = 0
	}
	sliceLength := len(slice)
	for i := startIndex; i < sliceLength; i++ {
		if valueToFind == slice[i] {
			return i
		}
	}
	return INDEX_NOT_FOUND
}

func ContainsInt(slice []int, valueToFind int) bool {
	return IndexOfInt(slice, valueToFind, 0) != INDEX_NOT_FOUND
}

func SubtractInt(originalSlice []int, comparedSlice []int) []int {
	var resultSlice []int
	for _, s := range originalSlice {
		if IndexOfInt(comparedSlice, s, 0) == -1 {
			resultSlice = append(resultSlice, s)
		}
	}
	return resultSlice
}
func SubtractUint(originalSlice []uint, comparedSlice []uint) []uint {
	var resultSlice []uint
	for _, s := range originalSlice {
		if IndexOfUint(comparedSlice, s, 0) == -1 {
			resultSlice = append(resultSlice, s)
		}
	}
	return resultSlice
}

func JoinInt(slice []int, sep string) string {
	switch len(slice) {
	case 0:
		return consts.EMPTY
	case 1:
		return frameUtils.ToString(slice[0])
	}
	n := len(sep) * (len(slice) - 1)
	for i := 0; i < len(slice); i++ {
		n += len(frameUtils.ToString(slice[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(frameUtils.ToString(slice[0]))
	for _, s := range slice[1:] {
		b.WriteString(sep)
		b.WriteString(frameUtils.ToString(s))
	}
	return b.String()
}

// endregion

// region uint
func IndexOfUint(slice []uint, valueToFind uint, startIndex int) int {
	if slice == nil || len(slice) == 0 {
		return INDEX_NOT_FOUND
	}
	if startIndex < 0 {
		startIndex = 0
	}
	sliceLength := len(slice)
	for i := startIndex; i < sliceLength; i++ {
		if valueToFind == slice[i] {
			return i
		}
	}
	return INDEX_NOT_FOUND
}

func ContainsUint(slice []uint, valueToFind uint) bool {
	return IndexOfUint(slice, valueToFind, 0) != INDEX_NOT_FOUND
}

func JoinUint(slice []uint, sep string) string {
	switch len(slice) {
	case 0:
		return consts.EMPTY
	case 1:
		return frameUtils.ToString(slice[0])
	}
	n := len(sep) * (len(slice) - 1)
	for i := 0; i < len(slice); i++ {
		n += len(frameUtils.ToString(slice[i]))
	}

	var b strings.Builder
	b.Grow(n)
	b.WriteString(frameUtils.ToString(slice[0]))
	for _, s := range slice[1:] {
		b.WriteString(sep)
		b.WriteString(frameUtils.ToString(s))
	}
	return b.String()
}

// endregion

// region string
func IndexOfString(slice []string, valueToFind string, startIndex int) int {
	if slice == nil || len(slice) == 0 {
		return INDEX_NOT_FOUND
	}
	if startIndex < 0 {
		startIndex = 0
	}
	sliceLength := len(slice)
	for i := startIndex; i < sliceLength; i++ {
		if valueToFind == slice[i] {
			return i
		}
	}
	return INDEX_NOT_FOUND
}

func IndexOfStringFold(slice []string, valueToFind string, startIndex int) int {
	if slice == nil || len(slice) == 0 {
		return INDEX_NOT_FOUND
	}
	if startIndex < 0 {
		startIndex = 0
	}
	sliceLength := len(slice)
	for i := startIndex; i < sliceLength; i++ {
		if strings.EqualFold(valueToFind, slice[i]) {
			return i
		}
	}
	return INDEX_NOT_FOUND
}

func ContainsString(slice []string, valueToFind string) bool {
	return IndexOfString(slice, valueToFind, 0) != INDEX_NOT_FOUND
}

func ContainsStringFold(slice []string, valueToFind string) bool {
	return IndexOfStringFold(slice, valueToFind, 0) != INDEX_NOT_FOUND
}

func SubtractString(originalSlice []string, comparedSlice []string) []string {
	var resultSlice []string
	for _, s := range originalSlice {
		if IndexOfString(comparedSlice, s, 0) == -1 {
			resultSlice = append(resultSlice, s)
		}
	}
	return resultSlice
}

func SubtractStringFold(originalSlice []string, comparedSlice []string) []string {
	var resultSlice []string
	for _, s := range originalSlice {
		if IndexOfStringFold(comparedSlice, s, 0) == -1 {
			resultSlice = append(resultSlice, s)
		}
	}
	return resultSlice
}

func IntersectString(originalSlice []string, comparedSlice []string) []string {
	var resultSlice []string
	for _, s := range originalSlice {
		if IndexOfString(comparedSlice, s, 0) >= 0 {
			resultSlice = append(resultSlice, s)
		}
	}
	return resultSlice
}

func IntersectStringFold(originalSlice []string, comparedSlice []string) []string {
	var resultSlice []string
	for _, s := range originalSlice {
		if IndexOfStringFold(comparedSlice, s, 0) >= 0 {
			resultSlice = append(resultSlice, s)
		}
	}
	return resultSlice
}

func TrimSpace(slice []string) {
	for i := range slice {
		slice[i] = strings.TrimSpace(slice[i])
	}
}

// endregion

func Expansion(slice *[]string, maxLength int) {
	sliceLength := len(*slice)
	if sliceLength < maxLength {
		for i, v := range *slice {
			(*slice)[i] = strings.TrimSpace(v)
		}
		lengthDiff := maxLength - sliceLength
		for i := 0; i < lengthDiff; i++ {
			*slice = append(*slice, consts.EMPTY)
		}
	}
}

func ExpansionAndTrimSpace(slice *[]string, length int) {
	Expansion(slice, length)
	TrimSpace(*slice)
}

func FormatExcelRows(rows [][]string, startRowIndex int, maxCellLength int) [][]string {
	var formattedExcelRows [][]string
	rows = rows[startRowIndex:]
	for rowIndex := range rows {
		Expansion(&rows[rowIndex], maxCellLength)
		var nonemptyRow bool
		for cellIndex := 0; cellIndex < maxCellLength; cellIndex++ {
			rows[rowIndex][cellIndex] = strings.TrimSpace(rows[rowIndex][cellIndex])
			if rows[rowIndex][cellIndex] != consts.EMPTY && !nonemptyRow {
				nonemptyRow = true
			}
		}
		if nonemptyRow {
			formattedExcelRows = append(formattedExcelRows, rows[rowIndex])
		}
	}
	return formattedExcelRows
}

// 去除 字段左右空格以及单引号
func FormatCSVExcelRows(rows [][]string) [][]string {
	for i, row := range rows {
		for x, v := range row {
			if v != "" {
				rows[i][x] = strings.Trim(strings.TrimSpace(v), "'")
			}
			continue
		}
	}
	return rows
}

// 去除重复元素
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// 去除重复元素
func RemoveRepeatedElementUint(arr []uint) (newArr []uint) {
	newArr = make([]uint, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// 去除空元素
func RemoveEmptyElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] != consts.EMPTY {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

//	取交集
func GetIntersection(a []string, b []string) (inter []string) {
	// interacting on the smallest list first can potentailly be faster...but not by much, worse case is the same
	low, high := a, b
	if len(a) > len(b) {
		low = b
		high = a
	}

	done := false
	for i, l := range low {
		for j, h := range high {
			// get future index values
			f1 := i + 1
			f2 := j + 1
			if l == h {
				inter = append(inter, h)
				if f1 < len(low) && f2 < len(high) {
					// if the future values aren't the same then that's the end of the intersection
					if low[f1] != high[f2] {
						done = true
					}
				}
				// we don't want to interate on the entire list everytime, so remove the parts we already looped on will make it faster each pass
				high = high[:j+copy(high[j:], high[j+1:])]
				break
			}
		}
		// nothing in the future so we are done
		if done {
			break
		}
	}
	return
}

//	根据 每组数值大小 分组
func GetArrayGroupBy(stu []interface{}, num int64) [][]interface{} {
	max := int64(len(stu))
	//判断数组大小是否小于等于指定分割大小的值，是则把原数组放入二维数组返回
	if max <= num {
		return [][]interface{}{stu}
	}
	//获取应该数组分割为多少份
	var quantity int64
	if max%num == 0 {
		quantity = max / num
	} else {
		quantity = (max / num) + 1
	}
	//声明分割好的二维数组
	var segments = make([][]interface{}, 0)
	//声明分割数组的截止下标
	var start, end, i int64
	for i = 1; i <= quantity; i++ {
		end = i * num
		if i != quantity {
			segments = append(segments, stu[start:end])
		} else {
			segments = append(segments, stu[start:])
		}
		start = i * num
	}
	return segments
}
