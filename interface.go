program1

package main

import "fmt"

type Article struct {
    Title string
    Author string
}

func (a Article) String() string {
    return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func main() {
    a := Article{
        Title: "Understanding Interfaces in Go",
        Author: "Sammy Shark",
    }
    fmt.Println(a.String())
}
output:
The "Understanding Interfaces in Go" article was written by Sammy Shark.
Program exited.





program2

package main

import (  
    "fmt"
)

type Worker interface {  
    Work()
}

type Person struct {  
    name string
}

func (p Person) Work() {  
    fmt.Println(p.name, "is working")
}

func describe(w Worker) {  
    fmt.Printf("Interface type %T value %v\n", w, w)
}

func main() {  
    p := Person{
name: "prathyusha",
    }
    var w Worker = p
    describe(w)
    w.Work()
}
output:
Interface type main.Person value {prathyusha}
prathyusha is working

Program exited.






program3

package main 
import "fmt"
type tank interface { 
    Tarea() float64 
    Volume() float64 
} 
func main() { 
    var t tank 
    fmt.Println("Value of the tank interface is: ", t) 
    fmt.Printf("Type of the tank interface is: %T ", t) 
} 
$go run main.go
Value of the tank interface is:  <nil>
Type of the tank interface is: <nil> 





program4

package main
import "fmt"
type Guitarist interface {
    
    PlayGuitar()
}

type BaseGuitarist struct {
    Name string
}

type AcousticGuitarist struct {
    Name string
}

func (b BaseGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
    fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
    var player BaseGuitarist
    player.Name = "prathyusha"
    player.PlayGuitar()

    var player2 AcousticGuitarist
    player2.Name = "Ringo"
    player2.PlayGuitar()
}
$go run main.go
prathyusha plays the Bass Guitar
Ringo plays the Acoustic Guitar






program5

package main
import (
    "fmt"
)
func myFunc(a interface{}) {
    fmt.Println(a)
}
func main() {
    var my_age int
    my_age = 20

    myFunc(my_age)
}
$go run main.go
20





program6

package main 
import "fmt"
func myfun(a interface{}) { 
    value, ok := a.(float64) 
    fmt.Println(value, ok) 
} 
func main() { 
  
    var a1 interface { 
    } = 96.09
  
    myfun(a1) 
  
    var a2 interface { 
    } = "supraja"
  
    myfun(a2) 
} 
$go run main.go
96.09 true
0 false




program7
package main 
import "fmt"
func myfun(a interface{}) { 
    val := a.(string) 
    fmt.Println("Value: ", val) 
} 
func main() { 
  
    var val interface { 
    } = "supraja"
      
    myfun(val) 
} 
$go run main.go
Value:  supraja