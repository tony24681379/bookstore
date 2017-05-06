package tools

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

//Highlight the value for sales stauts and inventory status
func Highlight(value string, low string, high string, len int) string {
	v, _ := strconv.Atoi(value)
	s, _ := strconv.Atoi(low)
	h, _ := strconv.Atoi(high)
	strlen := "%" + strconv.Itoa(len) + "s"
	value = fmt.Sprintf(strlen, value)
	if v < s {
		return color.RedString(value)
	} else if v > h {
		return color.GreenString(value)
	}
	return value
}

func TruncateString(str string, num int) string {
	bnoden := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		bnoden = str[0:num] + "..."
	}
	return bnoden
}
