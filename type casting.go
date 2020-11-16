program1

package main
 import (
    "fmt"
)
 func main() {
    var a int = 42
    f := float64(a)
    fmt.Println(f)       
}
output:
42
Program exited.




program2

package main 
  import "fmt"
  func main() { 
          var totalsum int = 846 
    var number int = 19 
    var avg float32 
          avg = float32(totalsum) / float32(number)   
    fmt.Printf("Average = %f\n", avg) 
} 
output:
Average = 44.526318

Program exited.



program3

package main
import "fmt"
func main() {
    var x int = 19
    var y int = 21
    var mul float32
    mul = float32(x) * float32(y)
    fmt.Printf("Multiplication = %f\n", mul)
}
output:
Multiplication = 399.000000

Program exited.






program4

package main
 import (
    "fmt"
    "strconv"
)
 func main() {
    var s string = "29"
    v, _ := strconv.Atoi(s)       
        fmt.Println(v)    
    var i int = 29
    str := strconv.Itoa(i)         
    fmt.Println(str) 
}
output:
29
29
Program exited.





program5

package main
 import (
    "fmt"
)
 func main() {
    f := 12.34567
    i := int(f)  
    fmt.Println(i)           
    ii := 34
    ff := float64(ii)
    fmt.Println(ff)     
}
output:
12
34

Program exited.





program6

package main
import (
    "fmt"
)
func main() {
    var s string = "Hello World"
    var b []byte = []byte(s)     
    fmt.Println(b)  
    ss := string(b)              
    fmt.Println(ss)    
}
output:
[72 101 108 108 111 32 87 111 114 108 100]
Hello World

Program exited.







program7

package main
 import (
    "fmt"
)
 func main() {
    a := 6/3        
    f := 6.3/3       
    fmt.Println(a, f)     
}
output:
2 2.1
Program exited.






program8

package main
import (
	"fmt"
	"strconv"
)
func main() {
	s := "3.1415926535"
	f, _ := strconv.ParseFloat(s, 8)
	fmt.Printf("%T, %v\n", f, f)
	s1 := "-3.141"
	f1, _ := strconv.ParseFloat(s1, 8)
	fmt.Printf("%T, %v\n", f1, f1)
	s2 := "-3.141"
	f2, _ := strconv.ParseFloat(s2, 32)
	fmt.Printf("%T, %v\n", f2, f2)
}
output:
float64, 3.1415926535
float64, -3.141
float64, -3.1410000324249268

Program exited.





program9

package main
import (
	"fmt"
	"reflect"
	"strconv"
)
func main() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))
	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}
output:
bool
string
Program exited.






program10

package main
import (
	"fmt"
	"reflect"
	"strconv"
)
func main() {
	var i int64 = -654
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)
	var s string = strconv.FormatInt(i, 10)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}
output:
int64
-654
string
-654
Program exited.