package testb

import (
    "fmt"
    "github.com/melodyshu/testc"
)

func Sum(i1 int, i2 int) int {
    return i1 + i2
}

func Hello(s string) string {
    res := fmt.Sprintf("你好,%s", s)
    res = testc.Hello("sam")
    return res
}
