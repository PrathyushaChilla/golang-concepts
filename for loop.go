program1

package main
import (
	"fmt"
)
func main() {
for i := 0; i<5; i= i+2 {

	fmt.Println(i)
}
}
output:
0
2
4
Program exited.




program2

package main
import (
	"fmt"
)
func main() {
for i,j :=0, 0; i<5; i,j= i+1,j+1 {

	fmt.Println(i)
}
}
output:
0
1
2
3
4
Program exited.





program3

package main
import (
	"fmt"
)
func main() {
for i,j :=0, 0; i<5; i,j= i+1 ,j+2{
	fmt.Println(i,j)
}
}
output:
0 0
1 2
2 4
3 6
4 8
Program exited.




program4

package main
import (
	"fmt"
)
func main() {
for i :=0;i<5; i++ {
	fmt.Println(i)
	if i%2 == 0 {
	i/=2
	}else{
	i= 2*i +1
}
}
}
output:
0
1
4
3
Program exited.





program5

package main
import (
	"fmt"
)
// for loop with two varaibales
func main() {
	for i := 1; i < 7; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
output:
*
**
***
****
*****





program6

package main
import (
	"fmt"
)
// for loop with three varaibles
func main() {
	for i := 2; i < 5; i++ {
		for j := 0; j < i; j++ {
			for k := 0; k < j; k++ {
				fmt.Printf("/")
			}
			fmt.Println("*")
		}
		fmt.Println()
		}
		}
output:
*
/*

*
/*
//*

*
/*
//*
///*




program7

package main
import (
	"fmt"
)
func main() {
s:= []int{1,2,3}
for k,v := range s{
	fmt.Println(k,v)
	}
}
output:
0 1
1 2
2 3
Program exited.






program8

package main
import (
	"fmt"
)
func main() {
s:= "prathyusha!"
for k,v := range s{
	fmt.Println(k,string(v))
	}
}
output:
0 p
1 r
2 a
3 t
4 h
5 y
6 u
7 s
8 h
9 a
10 !
Program exited.





program9

package main
import (
	"fmt"
)
func main() {
s:= "supraja!"
for k,v := range s{
	fmt.Println(k,v)
	}
}
output:
0 112
1 114
2 97
3 116
4 104
5 121
6 117
7 115
8 104
9 97
10 33
Program exited.




program10

package main
import (
	"fmt"
)
func main() {
statepopulation := map[string]int{
"andhra pradesh": 3234,
"tamil nadu": 56789,
"kerala": 12345,
}
for _,v := range statepopulation {
	fmt.Println(v)
}
}
output:
3234
56789
12345
Program exited.
