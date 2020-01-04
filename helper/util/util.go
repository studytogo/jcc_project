package util

import (
	"fmt"
	"strconv"
)

//保留i位小数,i是要保留位数, 结果会四舍五入
func Decimal(value float64, i int) string {
	return fmt.Sprintf(`%.`+strconv.Itoa(i)+`f`, value)
}
