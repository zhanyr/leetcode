package main

//import (
//    "bufio"
//    "fmt"
//    "os"
//    "strings"
//    "strconv"
//)
//
//func solution(line string) string {
//    // 在此处理单行数据
//    lineSlice := strings.Split(line, " ")
//    var a int64
//    for i, item := range lineSlice {
//        itemInt, _ := strconv.ParseInt(item, 10, 64)
//        if i == 0 {
//            a = itemInt
//        } else {
//            a ^= itemInt
//        }
//    }
//
//    // 返回处理后的结果
//    return strconv.FormatInt(a, 10)
//}
//
//func main() {
//    r := bufio.NewReader(os.Stdin)
//    for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
//        fmt.Println(solution(string(line)))
//    }
//}