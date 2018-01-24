package main
//
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
//    item0, _ := strconv.Atoi(lineSlice[0])
//    item1, _ := strconv.Atoi(lineSlice[1])
//
//    // 返回处理后的结果
//    return strconv.Itoa(item0 + item1)
//}
//
//func main() {
//    r := bufio.NewReader(os.Stdin)
//    for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
//        fmt.Println(solution(string(line)))
//    }
//}