program1

package main
import (
	"fmt"
)
func main() {
	c := make(chan int)
	go func() {
		c <- 29
	}()
	fmt.Println(<-c)
}
output:
29

Program exited.





program2
package main

import (
	"fmt"
)

func main() {
	chat := make(chan string, 5)
	chat <- "Hey!"
	chat <- "How's it going?"
	chat <- "Doing well"
	chat <- "I started watching tv"
	chat <- "It's so good"
	fmt.Println(<- chat)
	fmt.Println(<- chat)
	fmt.Println(<- chat)
	fmt.Println(<- chat)
	fmt.Println(<- chat)
}
output:
Hey!
How's it going?
Doing well
I started watching Dark
It's so good

Program exited.






program3
package main

import (
	"fmt"
)
func main() {
	c := make(chan int, 1)
	send(c)
	receive(c)
}
func send(c chan<- int) {
	c <- 30
}
func receive(c <-chan int) {
	fmt.Println(<-c)
}
output:
30

Program exited.






progrm4
package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}
func main() {
ch := make(chan int)
wg.Add(2)
go func() {
i:= <- ch
	fmt.Println(i)
wg.Done()
}()
go func() {
ch <- 42
wg.Done()
}()
wg.Wait()
}
output:
42
Program exited.






program5

package main

import (
	"fmt"
	"sync"
)
var wg = sync.WaitGroup{}
func main() {
ch := make(chan int)
wg.Add(2)
go func() {
i:= <- ch
	fmt.Println(i)
	ch <- 111
wg.Done()
}()
go func() {
ch <- 222
fmt.Println(<-ch)
wg.Done()
}()
wg.Wait()
}
output:
222
111

Program exited.






program6

package main 
  
import "fmt"
  
func main() {     var mychannel chan int
    fmt.Println("Value of the channel: ", mychannel) 
    fmt.Printf("Type of the channel: %T ", mychannel) 
    mychannel1 := make(chan int) 
    fmt.Println("\nValue of the channel1: ", mychannel1) 
    fmt.Printf("Type of the channel1: %T ", mychannel1) 
} 
output:
Value of the channel:  <nil>
Type of the channel: chan int 
Value of the channel1:  0xc000100060
Type of the channel1: chan int 
Program exited.






program7

package main 
  
import "fmt"
  
func myfunc(ch chan int) { 
  
    fmt.Println(23 + <-ch) 
} 
func main() { 
    fmt.Println("start Main method") 
        ch := make(chan int) 
  
    go myfunc(ch) 
    ch <- 23 
    fmt.Println("End Main method") 
} 
output:
start Main method
46
End Main method

Program exited.






program8

package main 
  
import "fmt"
func main() { 
      mychnl := make(chan string, 5) 
    mychnl <- "GFG"
    mychnl <- "gfg"
    mychnl <- "Geeks"
    mychnl <- "GeeksforGeeks"
  mychnl <- "prathyusha"
      fmt.Println("Length of the channel is: ", len(mychnl)) 
output:
Length of the channel is:  5

Program exited.





program9

package main 
  
import "fmt"
 func main() { 
      mychnl := make(chan string) 
          go func() { 
        mychnl <- "damu"
        mychnl <- "bavitha"
        mychnl <- "supraja"
        mychnl <- "prathyusha"
        close(mychnl) 
    }() 
for res := range mychnl { 
        fmt.Println(res) 
    } 
} 
output:
damu
bavitha
supraja
prathyusha

Program exited.






program10

package main 
  
import "fmt"

func myfun(mychnl chan string) { 
  
    for v := 0; v < 4; v++ { 
        mychnl <- "morning"
    } 
    close(mychnl) 
} 
func main() { 
    c := make(chan string) 
      go myfun(c) 
      for { 
res, ok := <-c 
        if ok == false { 
            fmt.Println("Channel Close ", ok) 
            break
        } 
        fmt.Println("Channel Open ", res, ok) 
    } 
} 
output:
Channel Open  morning true
Channel Open  morning true
Channel Open  morning true
Channel Open  morning true
Channel Close  false

Program exited.














