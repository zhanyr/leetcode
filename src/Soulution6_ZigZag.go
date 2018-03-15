package main

//import "fmt"

//func main() {
//    s:= "PAYPALISHIRING"
//    fmt.Println(convert(s, 3))
//}
//锯齿形矩阵按行输出
//1.循环间隔numRows + numRows -2
//2.除了第一行和最后一行，多了一组循环，循环与上一个之间的间隔是 (2n-2i-2)
func convert(s string, numRows int) string {
    sLen := len(s)

    if sLen == 0 || numRows < 2 {
        return s
    }

    sRune := []rune(s)
    ret := make([]rune, 0, sLen)

    step := 2*numRows - 2
    for i:=0; i<numRows; i++ {
        for j := i; j <sLen; j=j+step {
            ret = append(ret, sRune[j])

            if i != 0 && i < numRows-1 {
                step2 := j+2*numRows-2*i-2
                if step2 < sLen {
                    ret = append(ret, sRune[step2])
                }
            }
        }

    }
    return string(ret)
}