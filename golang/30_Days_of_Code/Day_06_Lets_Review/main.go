package main
import (
    "fmt"
	"bufio"
	"os"

	"strconv"
	"strings"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 reader := bufio.NewReader(os.Stdin)
 line, _  := reader.ReadString('\n')
 lineNum, _ := strconv.Atoi(strings.Trim(line, "\n"))
 
 var output [10]string
 for i := 0; i < lineNum; i++  {
	
	text, _ := reader.ReadString('\n')
		
		var even string
		var odd string
	
		for index, x := range strings.Trim(text, "\n") {
			if index % 2 == 0 {
				even += string(x)
			} else {
				odd += string(x)
			}
		}
		output[i] = fmt.Sprintf("%s %s\n", even, odd)
 }

 for i := 0; i < lineNum; i++   {
	 fmt.Printf("%s", output[i])
 }
}
