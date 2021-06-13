package misc

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func StrArr2IntArr(strArr []string) ([]int, error) {
	var intArr []int
	for _, value := range strArr {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		intArr = append(intArr, intValue)
	}
	return intArr, nil
}

func Str2Int(str string) int {
	intValue, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return intValue
}

//func IntArr2StrArr(intArr []int) []string {
//	var strArr []string
//	for _, value := range intArr {
//		strValue := strconv.Itoa(value)
//		strArr = append(strArr, strValue)
//	}
//	return strArr
//}

func Int2Str(Int int) string {
	return strconv.Itoa(Int)
}

func IsInStrArr(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func IsInIntArr(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func ReadLine(fileName string, handler func(string, bool)) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = FixLine(line)
		handler(line, true)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

func FixLine(line string) string {
	line = strings.Replace(line, "\r", "", -1)
	line = strings.Replace(line, " ", "", -1)
	line = strings.Replace(line, "\t", "", -1)
	line = strings.Replace(line, "\r", "", -1)
	line = strings.Replace(line, "\n", "", -1)
	return line
}

func UniStrAppend(slice []string, elems ...string) []string {
	for _, elem := range elems {
		if IsInStrArr(slice, elem) {
			continue
		} else {
			slice = append(slice, elem)
		}
	}
	return slice
}

func FileIsExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func Xrange(args ...int) []int {
	var start, stop int
	var step = 1
	var r []int
	switch len(args) {
	case 1:
		stop = args[0]
		start = 0
	case 2:
		start, stop = args[0], args[1]
	case 3:
		start, stop, step = args[0], args[1], args[2]
	default:
		return nil
	}
	if start > stop {
		return nil
	}
	if step < 0 {
		return nil
	}

	for i := start; i <= stop; i += step {
		r = append(r, i)
	}
	return r
}

func FilterPrintStr(s string) string {
	// 将字符串转换为rune数组
	srcRunes := []rune(s)
	// 创建一个新的rune数组，用来存放过滤后的数据
	dstRunes := make([]rune, 0, len(srcRunes))
	// 过滤不可见字符，根据上面的表的0-32和127都是不可见的字符
	for _, c := range srcRunes {
		if c >= 0 && c <= 31 {
			continue
		}
		if c == 127 {
			continue
		}
		dstRunes = append(dstRunes, c)
	}

	return string(dstRunes)
}

func SprintStringMap(stringMap map[string]string) string {
	var rArr []string
	var assistArr []string
	for key, value := range stringMap {
		if value == "" {
			continue
		}
		if IsInStrArr(assistArr, value) == false {
			assistArr = append(assistArr, value)
			rArr = append(rArr, fmt.Sprintf("%s:%s", key, value))
		}
	}
	return strings.Join(rArr, "、")
}

func MustLength(s string, i int) string {
	if len(s) > i {
		return s[:i]
	}
	return s
}

func Percent(int1 int, int2 int) string {
	float1 := float64(int1)
	float2 := float64(int2)
	f := 1 - float1/float2
	f = f * 100
	return strconv.FormatFloat(f, 'f', 2, 64)
}

func StrRandomCut(s string, length int) string {
	sRune := []rune(s)
	if len(sRune) > length {
		i := rand.Intn(len(sRune) - length)
		return string(sRune[i : i+length])
	} else {
		return s
	}
}
