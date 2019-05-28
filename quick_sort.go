package main
import (
	"fmt"
	"os"
	"strconv"
)
func quick(arr []int){
	if len(arr) > 1{
		pivot := len(arr) / 2
		var left_side = []int{}
		var right_side = []int{}
		for i:= range arr {
			val := arr[i]
			if i != pivot {
				if val <arr[pivot]{
					left_side = append(left_side, val)
				}else{
					right_side = append(right_side, val)
				}
			}
		}
		quick(left_side)
		quick(right_side)

		var merge []int = append(append(append([]int{},left_side...),[]int{arr[pivot]}...),right_side...)

		for j:= 0; j<len(arr);j++ {
			arr[j] = merge[j]
		}

		fmt.Println(arr)
	}

	fmt.Println(arr)
}

func main(){
	list := strconv.ParseInt(os.Args, 10, 64)
	arg := os.Args[1:]
	var argi = []int{}
    for _, i := range arg {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        argi = append(argi, j)
    }
	quick(argi)
}
