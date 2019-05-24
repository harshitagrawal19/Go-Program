package main
import "fmt"
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
	list := []int{12,50,11,40,65,03,11,25}
	quick(list)
}
