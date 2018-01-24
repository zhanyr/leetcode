package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func solution(line string) string {
    // 在此处理单行数据
    lineSlice := strings.Split(line, " - ")
    aBytes := []byte(lineSlice[0])
    bBytes := []byte(lineSlice[1])

    lenA := len(aBytes)
    lenB := len(bBytes)

    ret := make([]byte, lenA)
    j := lenB -1
    borrow := false
    for i := lenA - 1; i > 0; i-- {
        a := aBytes[i]
        if j >= 0 {
            if borrow {
                a = a -1
            }

            if a < bBytes[j] {
                borrow = true
                ret[i] = a + 10 - bBytes[j]
            } else {
                borrow = false
                ret[i] = a - bBytes[j]
            }
        } else {
            if borrow {
                if aBytes[i] < 1 {
                    ret[i] = aBytes[i] + 10 - 1
                    borrow = true
                } else {
                    ret[i] = aBytes[i] - 1
                    borrow = false
                }
            } else {
                ret[i] = aBytes[i]
                borrow = false
            }

        }
        fmt.Println("i %v, ret, %v, a %v", i, ret[i], aBytes[i])

        j--
    }

    result := strings.TrimLeft(string(ret), "0")
    // 返回处理后的结果
    return result
}

func main() {
    r := bufio.NewReader(os.Stdin)
    for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
        fmt.Println(solution(string(line)))
    }
}