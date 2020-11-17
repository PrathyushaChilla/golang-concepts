program1

package main   
import (   
    "fmt"
    "sync" // to import sync later on 
) 
var GFG  = 0 
func worker(wg *sync.WaitGroup) {   
    GFG = GFG + 1 
      wg.Done() 
} 
func main() {  
      var w sync.WaitGroup 
   for i := 0; i < 1000; i++ { 
        w.Add(1)         
        go worker(&w) 
    } 
       w.Wait() 
    fmt.Println("Value of x", GFG) 
} 
output:
Value of x 953

Program exited.





program2

package main   
import (   
    "fmt"
    "sync" // to import sync later on 
) 
var GFG  = 0 
func worker(wg *sync.WaitGroup, m *sync.Mutex) {  
       m.Lock()  
    GFG = GFG + 1 
    m.Unlock() 
      wg.Done() 
} 
func main() {  
       var w sync.WaitGroup 
  var m sync.Mutex 
      for i := 0; i < 1000; i++ { 
        w.Add(1)         
        go worker(&w, &m) 
    } 
       w.Wait() 
    fmt.Println("Value of x", GFG) 
} 
output:
Value of x 1000

Program exited.






program3

package main

import (
    "fmt"
    "sync"
)

var (
    mutex   sync.Mutex
    balance int
)

func init() {
    balance = 1000
}

func deposit(value int, wg *sync.WaitGroup) {
    mutex.Lock()
    fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
    balance += value
    mutex.Unlock()
    wg.Done()
}

func withdraw(value int, wg *sync.WaitGroup) {
    mutex.Lock()
    fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
    balance -= value
    mutex.Unlock()
    wg.Done()
}

func main() {
    fmt.Println("Go Mutex Example")

	var wg sync.WaitGroup
	wg.Add(2)
    go withdraw(700, &wg)
    go deposit(500, &wg)
    wg.Wait()

    fmt.Printf("New Balance %d\n", balance)
}
output:

Go Mutex Example
Depositing 500 to account with balance: 1000
Withdrawing 700 from account with balance: 1500
New Balance 800

Program exited.





program4

package main  
import (  
    "fmt"
    "sync"
    )
var x  = 0  
func increment(wg *sync.WaitGroup, m *sync.Mutex) {  
    m.Lock()
    x = x + 1
    m.Unlock()
    wg.Done()   
}
func main() {  
    var w sync.WaitGroup
    var m sync.Mutex
    for i := 0; i < 500; i++ {
        w.Add(1)        
        go increment(&w, &m)
    }
    w.Wait()
 fmt.Println("final value of x", x)
}
output:
final value of x 500

Program exited.






program5

package main  
import (  
   "sync"  
   "time"  
   "math/rand"  
   "fmt"  
)  
var wait sync.WaitGroup  
var count int  
var mutex sync.Mutex  
func  increment(s string)  {  
   for i :=0;i<3;i++ {  
      mutex.Lock()  
      x := count  
      x++;  
      time.Sleep(time.Duration(rand.Intn(10))*time.Millisecond)  
      count = x;  
      fmt.Println(s, i,"Count: ",count)  
      mutex.Unlock()  
        
   }  
   wait.Done()  
     
}  
func main(){  
   wait.Add(2)  
   go increment("foo: ")  
   go increment("bar: ")  
   wait.Wait()  
   fmt.Println("last count value " ,count)  
}  
output:
foo:  0 Count:  1
foo:  1 Count:  2
foo:  2 Count:  3
bar:  0 Count:  4
bar:  1 Count:  5
bar:  2 Count:  6
last count value  6








