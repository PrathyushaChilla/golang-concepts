program1

package main
import "fmt"
func factorial(i int)int {
   if(i <= 1) {
      return 1
   }
   return i * factorial(i - 1)
}
func main() { 
   var i int = 10
   fmt.Printf("Factorial of %d is %d", i, factorial(i))
}
output:
Factorial of 10 is 3628800
Program exited.




program2

package main
import "fmt"
func fibonaci(i int) (ret int) {
   if i == 0 {
      return 0
   }
   if i == 1 {
      return 1
   }
   return fibonaci(i-1) + fibonaci(i-2)
}
func main() {
   var i int
   for i = 0; i < 10; i++ {
      fmt.Printf("%d ", fibonaci(i))
   }
}
output:
0 1 1 2 3 5 8 13 21 34 
Program exited.




program3

package main  
  import (  
 "fmt"  
)  
  func recursivefunction() {  
 fmt.Println("welcome")  
 recursivefunction()  
}  
  func main() {  
 recursivefunction()  
}  
output: 
welcome  
welcome  
welcome  
.....  
infite times 






program4
package main  
  
import (  
 "fmt"  
)  
  
var count = 0  
  
func recursivemethod() {  
 count++  
 if count <= 10 {  
  fmt.Println(count)  
  recursivemethod()  
 }  
}  
  
func main() {  
 recursivemethod()  
}  
output:
1
2
3
4
5
6
7
8
9
10
Program exited. 






program5

package main  
  import "fmt"  
  func main() {  
 var myfunction func()  
 myfunction = func() {  
  fmt.Println("Anonymous function example")  
  myfunction()  
  
 }  
 myfunction()  
  
}  
output:
prathyusha
prathyusha
prathyusha
...  
...  
...  
infinite times  






program6

package main 
  
import ( 
    "fmt"
) 
func factorial_calc(number int) int { 
    if number == 0 || number == 1 { 
        return 1 
    } 
        if number < 0 { 
        fmt.Println("Invalid number") 
        return -1 
    } 
      
     return number*factorial_calc(number - 1) 
} 
func main() { 
  
        answer1 := factorial_calc(0) 
    fmt.Println(answer1) 
       answer2 := factorial_calc(5) 
    fmt.Println(answer2) 
      
       answer3 := factorial_calc(-1) 
    fmt.Println(answer3) 
      
   answer4 := factorial_calc(10) 
    fmt.Println(answer4) 
} 
output:
1
120
Invalid number
-1
3628800
Program exited.






program7
package main 
  
import ( 
    "fmt"
)
func print_num(n int) {     
        if n > 0 { 
        fmt.Println(n) 
                      print_num(n-1) 
    } 
}  
func main() {  
        print_num(5) 
} 
output:
5
4
3
2
1
Program exited.






program8

package main 
  
import ( 
    "fmt"
)
func print_hello() { 
    fmt.Println("GeeksforGeeks") 
    print_hello() 
} 
func main() { 
    print_hello() 
} 
output:
GeeksforGeeks
GeeksforGeeks
..... infinite times


