package main

import (
    "fmt"
    "time"
)

func main() {
    go start()
    fmt.Println("Started")
    time.Sleep(1 * time.Second)
    fmt.Println("Finished")
}

func start() {
    fmt.Println("In Goroutine")
}
output:
Started
In Goroutine
Finished

Program exited.





program2
package main

import (
    "fmt"
    "time"
)

func main() {
    go start()
    fmt.Println("Started")
    time.Sleep(1 * time.Second)
    fmt.Println("Finished")
}

func start() {
    go start2()
    fmt.Println("In Goroutine")
}
func start2() {
    fmt.Println("In Goroutine2")
}
output:
Started
In Goroutine
In Goroutine2
Finished

Program exited.







program3
package main

import (
    "fmt"
    "time"
)

func execute(id int) {
    fmt.Printf("id: %d\n", id)
}

func main() {
    fmt.Println("Started")
    for i := 0; i < 10; i++ {
        go execute(i)
    }
    time.Sleep(time.Second * 2)
    fmt.Println("Finished")
}
output:
Started
id: 4
id: 9
id: 7
id: 8
id: 5
id: 1
id: 2
id: 3
id: 6
id: 0
Finished





program4
package main

import "fmt"

func f(msg string) {
   fmt.Println(msg)
}

func main() {
   go f("go routine")
   f("function")
   fmt.Scanln()
}
output:
function

Program exited.






program5

package main
import (
    "fmt"
    "time"
)
func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}
func main() {
f("direct")
go f("goroutine")
go func(msg string) {
fmt.Println(msg)
}("going")
time.Sleep(time.Second)
fmt.Println("done")
}
output:
direct : 0
direct : 1
direct : 2
goroutine : 0
goroutine : 1
goroutine : 2
going
done



program6

package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
output:
hello
world
world
hello
world
hello
world
hello
hello

Program exited.






program 7


package main
import (
    "fmt"
)
func main() {
    go start()
    fmt.Println("Started")
    fmt.Println("Finished")
}
func start() {
    fmt.Println("In Goroutine")
}
output:
Started
Finished

Program exited.