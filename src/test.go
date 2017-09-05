package main
import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please inout your name:")
	input,err := inputReader.ReadStriging('\n')
	if err != nil {
		fmt.Printf("Found an error:%s\n",err)
	}else{
		input = input[:len(input)-1]
		fmt.Printf("Hello, %s\n",input)
	}
}