program1

package main

import (
    "flag"
    "fmt"
)

func main() {

    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
output:

word: foo
numb: 42
fork: false
svar: bar
tail: []

Program exited.




program2

package main 
  
import ( 
    "fmt"
    "os"
) 
  
func main() { 
  
    // The first argument 
    // is always program name 
    myProgramName := os.Args[0] 
      
    // it will display  
    // the program name 
    fmt.Println(myProgramName) 
} 

output:

/tmpfs/play

Program exited.





program3

package main
import (
    "fmt"
    "os"
)
func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
You can get individual args with normal indexing.

    arg := os.Args[3]
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
output:
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c






program4

package main

import (
    "flag"
    "fmt"
)

func main() {

    wordPtr := flag.String("word", "default value", "a string for description")
    flag.Parse()
    fmt.Println("word:", *wordPtr)

}
output:
go run main.go -word=welcome
 word : welcome
 
 
 

program5

package main

import (
    "flag"
    "fmt"
)

func main() {

    wordPtr := flag.String("word", "default value", "a string for description")
    flag.Parse()
    fmt.Println("word:", *wordPtr)

}
output:
go run main.go -word=hello
 word: hello
 
 
 
 
 program6
 
 package main

import ("fmt"
        "os"
)

func main() {
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
    arg := os.Args[3]
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
output:
$ go run test.go 1 2 3 4 5
[/tmp/go-build162373819/command-line-arguments/_obj/exe/modbus 1 2 3 4 5]
[1 2 3 4 5]
3




program7

package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(len(os.Args), os.Args)
}

$go run main.go
1 [/tmp/go-build486280812/command-line-arguments/_obj/exe/main]


