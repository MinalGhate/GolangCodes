# Golang
## Introduction
Fast compilation
cluster-> collection of server
To develop cloud services
Pointer arithmetic , inheritence etc not supported
## Why
Cross plaform compatibility
Fast run and compilation time

## Execution

Run hello.go  => execution, compile and run desnt make exe
Type hello.go => write code
Build hello.go  => make .exe file
Now by clicking .exe file execution will be shown (text based integration so just open and imdiately close screen so do it on command prompt)


## Variables
Can start with _ and alphabet
Snake case ex: my_name=”minal”, my_variable_name
A,b,c = 10,20.5,’c’

## Constant
const pi = 3.14
Cannot be change
Non mutable
If try to change will give error (\basic.go:8:2: cannot assign to g (neither addressable nor a map index expression)
Unchangeable and read onle
- 
### Types
#### Typed constant

const a int =1
While declaring const giving type

#### Untyped constant
const pi = 3.14
Not declaring type

const g int = 10    //typed
    const gravity = 9.8 //untyped
    fmt.Println("Gravity = ", g, gravity)
    //g = g + 5 //we cannot change const variable (error)
    fmt.Println("updated Gravity = ", g)


### Multiple constant type
package main


import "fmt"


func main() {


    const (
        g    int    = 10
        name        = "gravity"
        n    string = "it is constant"
        grav        = 9.8
    )


    fmt.Println(name)
    fmt.Println(n)
    fmt.Println(g)
    fmt.Println(grav)


}




# Keywords
Reserved word
Cannot be used as a variable



