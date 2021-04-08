package main

import "fmt"

func main() {
	// goto 跳转到任意lable标签的位置
	for i := 0; i < 10; i++ {
		if i == 4 {
			goto End
		}
	}
End:
	fmt.Printf("End")
}
