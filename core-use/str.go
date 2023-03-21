package core_use

import (
	"fmt"
	"strconv"
	"strings"
)

var str = "hello go"

// 要导出函数， 首字母必须大写， 否则就是包内私有的
func StringUsage() {
	fmt.Printf("%t \n", strings.HasPrefix(str, "he"))
}

// 字符串类型转换
func Strconvrt() {
	var an int
	var orig = "666"
	an, _ = strconv.Atoi(orig)
	an += 5
	news := strconv.Itoa(an)
	fmt.Println(news)
}
