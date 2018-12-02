package main

import(
    "fmt"
    "io/ioutil" 
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    dat, err := ioutil.ReadFile("day1/day1.txt")
    check(err)
    input := string(dat)[:(len(string(dat))-1)]
    length := len(input)
    total := 0
    for i := 0; i < length; i++ {
        if int(input[i]) == int(input[(i+1)%length]){
            num, _ := strconv.Atoi(string(input[i]))
            total += num
        }
    }

    fmt.Println("Part 1: ", total)

    total = 0
    for i := 0; i < length; i++ {
        if int(input[i]) == int(input[(i+(length/2))%length]){
            num, _ := strconv.Atoi(string(input[i]))
            total += num
        }
    }

    fmt.Println("Part 2: ", total)

}