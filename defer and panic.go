program1

package main
import (
	"fmt"
)
func test() {
	fmt.Println("Testing Golang defer Function")
}
func main() {

	fmt.Println("Start")
	defer test()
	fmt.Println("Third")

}
output:
Start
Third
Testing Golang defer Function
Program exited.




program2
package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Golang Panic Occurred!")
	fmt.Println("end")

}
output:
Start
2009/11/10 23:00:00 Error: Golang Panic Occurred!

Program exited.





program3

package main 
  
import "fmt"
func add(a1, a2 int) int { 
    res := a1 + a2 
    fmt.Println("Result: ", res) 
    return 0 
} 
 func main() { 
  
    fmt.Println("Start") 
      defer fmt.Println("End") 
    defer add(34, 56) 
    defer add(10, 10) 
} 
output:
Start
Result:  20
Result:  90
End

Program exited.






program4

package main 
  
import "fmt"
 func mul(a1, a2 int) int { 
  
    res := a1 * a2 
    fmt.Println("Result: ", res) 
    return 0 
} 
  
func show() { 
    fmt.Println("Hello!, prathyusha") 
} 

func main() { 
  
    mul(23, 45) 
  
    defer mul(23, 56) 
      show() 
} 
output:
Result:  1035
Hello!, prathyusha
Result:  1288

Program exited.






program5
package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Golang Panic Occurred!")
	fmt.Println("end")

}

output:
Start





program6

package main 
  
import "fmt"

func main() { 
      var myarr [3]string 
      myarr[0] = "GFG"
    myarr[1] = "GeeksforGeeks"
    myarr[2] = "Geek"
  fmt.Println("Elements of Array:") 
    fmt.Println("Element 1: ", myarr[0]) 
      fmt.Println("Element 2: ", myarr[5]) 
  
} 
output:








program7
package main 
  
import "fmt"
 
func entry(lang *string, aname *string) { 
      if lang == nil { 
        panic("Error: Language cannot be nil") 
    } 
       if aname == nil { 
        panic("Error: Author name cannot be nil") 
    } 
      fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname) 
} 
 func main() { 
  
    A_lang := "GO Language"
      entry(&A_lang, nil) 
	  }
output:
	  panic: Error: Author name cannot be nil


