package main

import "os"
import "bufio"
import "fmt"
import "strings"
import "io"

type InputData struct {

}

func defineInputData() {
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadBytes('\n')

		if err == io.EOF {
				break
		}
		if err != nil {
				panic(err)
		}

		fmt.Println(strings.TrimRight(string(line),"\n"))
	}
}

func main() {
	defineInputData()
}
