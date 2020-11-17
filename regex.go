program1

package main
import (
	"fmt"
	"regexp"
)

func main() {
  re := regexp.MustCompile("prathyu$")
  fmt.Println(re.FindString("shaprathyu"))
}
output:
prathyu

Program exited.





program2

package main
import (
  "fmt"
  "regexp"
)

func main() {
  re := regexp.MustCompile("et$")
  
  fmt.Println(re.FindString("cricket"))
  fmt.Println(re.FindString("hacked"))
  fmt.Println(re.FindString("racket"))
}
output:
et

et

Program exited.





program3

package main
import (
  "fmt"
  "regexp"
)

func main() {  
  re := regexp.MustCompile("tel")
  fmt.Println(re.FindStringIndex("telephone"))
  fmt.Println(re.FindStringIndex("carpet"))
  fmt.Println(re.FindStringIndex("cartel"))
}
output:
[0 3]
[]
[3 6]






program4

package main
import (
  "fmt"
  "regexp"
)

func main() { 
  re := regexp.MustCompile("p([a-z]+)ch")
  fmt.Println(re.FindStringSubmatch("peach punch"))
  fmt.Println(re.FindStringSubmatch("cricket"))
}
output:
[peach ea]
[]







program5

package main

import (
    "fmt"
    "log"
    "regexp"
)

func main() {

    words := [...]string{"Seven", "even", "Maven", "Amen", "eleven"}

    for _, word := range words {

        found, err := regexp.MatchString(".even", word)

        if err != nil {
            log.Fatal(err)
        }

        if found {

            fmt.Printf("%s matches\n", word)
        } else {

            fmt.Printf("%s does not match\n", word)
        }
    }
}
output:
Seven matches
even does not match
Maven does not match
Amen does not match
eleven matches

Program exited.







program6

package main

import (
    "fmt"
    "log"
    "regexp"
    "strconv"
)

func main() {

    var data = `22, 1, 3, 4, 5, 17, 4, 3, 21, 4, 5, 1, 48, 9, 42`

    sum := 0

    re := regexp.MustCompile(",\\s*")

    vals := re.Split(data, -1)

    for _, val := range vals {

        n, err := strconv.Atoi(val)

        sum += n

        if err != nil {
            log.Fatal(err)
        }
    }

    fmt.Println(sum)
}
output:
189

Program exited.




program7

package main

import (
    "fmt"
    "regexp"
)

func main() {

    websites := [...]string{"webcode.me", "zetcode.com", "freebsd.org", "netbsd.org"}

    re := regexp.MustCompile("(\\w+)\\.(\\w+)")

    for _, website := range websites {

        parts := re.FindStringSubmatch(website)

        for i, _ := range parts {
            fmt.Println(parts[i])
        }

        fmt.Println("---------------------")
    }
}
output:
webcode.me
webcode
me
---------------------
zetcode.com
zetcode
com
---------------------
freebsd.org
freebsd
org
---------------------
netbsd.org
netbsd
org
---------------------

Program exited.







program8

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "regexp"
    "strings"
)

func main() {

    resp, err := http.Get("http://webcode.me")

    if err != nil {
        log.Fatal(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {

        log.Fatal(err)
    }

    content := string(body)

    re := regexp.MustCompile("<[^>]*>")
    replaced := re.ReplaceAllString(content, "")

    fmt.Println(strings.TrimSpace(replaced))
}
output:
connect: no route to host
Program exited: status 1.






program9

package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {

    content := "an old eagle"

    re := regexp.MustCompile(`[^aeiou]`)

    fmt.Println(re.ReplaceAllStringFunc(content, strings.ToUpper))
}
output:

aN oLD eaGLe

Program exited.





package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {

    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    r, _ := regexp.Compile("p([a-z]+)ch")

    fmt.Println(r.MatchString("peach"))

    fmt.Println(r.FindString("peach punch"))

    fmt.Println(r.FindStringIndex("peach punch"))

    fmt.Println(r.FindStringSubmatch("peach punch"))

    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    fmt.Println(r.FindAllString("peach punch pinch", -1))

    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    fmt.Println(r.FindAllString("peach punch pinch", 2))

    fmt.Println(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}