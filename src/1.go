// You can edit this code!
// Click here and start typing.
package main

import "fmt"

//import ("fmt"
//    "sync"
//    "os"
//)
//
//func main() {
//    wg := &sync.WaitGroup{}
//    wg.Add(1)
//    go func(){
//        defer wg.Done()
//        go func(){
//            fmt.Println("gogo---Hello, 世界")
//        }()
//        fmt.Println("go--Hello, 世界")
//        os.Exit(1)
//    }()
//
//    fmt.Println("Hello, 世界")
//
//    wg.Wait()
//}

func main() {

    ret := make(chan string, 1)

    ret <- "1"

    a := <-ret
    fmt.Println(a)



}