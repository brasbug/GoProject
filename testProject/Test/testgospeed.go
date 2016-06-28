
package main
import "fmt"
import "time"
func main(){
	j := 10001
	l := 0
	fmt.Printf("%s\n", time.Now())
	for d:=1; d<j; d = d + 1 {
		for i:=1; i<j; i = i + 1 {
			l = l + 1
		}
	}
	fmt.Printf("%s\n", time.Now())
	fmt.Printf("%d\n", l)
}