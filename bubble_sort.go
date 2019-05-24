package main
import "fmt"
func Bubble(arr []int) {
	size := len(arr)
	for i:=0; i<size;i++ {
		for j:= 0; j< (size - 1-i);j++{
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	for i:=0; i<size;i++ {
		fmt.Print(arr[i])
	}
}

func main(){
	list := []int{5,6,4,9,2}
	Bubble(list)
	
}
