package main
import "fmt"
func insertion(arr []int){
	size := len(arr)
	for i:=1; i<size; i++ {
		for j:=i; j>0 && arr[j]< arr[j-1]; j--{
			arr[j],arr[j-1] = arr[j-1],arr[j]
		}
	}
	for i:=0;i<size;i++ {
		fmt.Print(arr[i]," ")
	}
}

func main(){
	list := []int{15,40,16,50,20,41,6}
	insertion(list)
}