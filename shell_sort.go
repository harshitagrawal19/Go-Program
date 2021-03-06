package main

import "fmt"

func main() {
   list := []int{90,-80,70,31,-60,50,-40,30,-20,10,31,33,0}
   fmt.Println("before:", list)
   shellsort(list)
   fmt.Println("after: ", list)
}

func shellsort(list []int) {
   for inc := len(list) / 2; inc > 0; inc = (inc + 1) * 5 / 11 {
      for i := inc; i < len(list); i++ {
         j, temp := i, list[i]
         for ; j >= inc && list[j-inc] > temp; j -= inc {
            list[j] = list[j-inc]
         }
         list[j] = temp
      }
   }
}
