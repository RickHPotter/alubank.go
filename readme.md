# Getting Started

This is part of an [Alura Course in Golang](https://cursos.alura.com.br/course/go-lang-oo).

## Concepts

> Go Programming Language for Dummies

### Interface

Interfaces serve two important purposes: they make your code more versatile, and they force you to adopt *code encapsulation* (the pratice of hiding the implementation of your methods).

An interface is an abstract type that describes all methods that a type can implement, but it only provides the method signatures. There's defining in an interface, but no declaring.

```golang
    // models/shapes.go

    type Circle struct {
        radius float64
        name string
    }
    type Square struct {
        length float64
        name string
    }

    type Shape interface {
        Area() float64
    }

    // Circle implements Shape
    func (c Circle) Area() float64 {
        return math.Pi * math.Pow(c.radius, 2)
    }
    
    // Square implements Shape
    func (s Square) Area() float64 {
        return math.Pow(s.length, 2)
    }

    // main.go

    func main() {
        c1 := Circle{radius: 5, name: "c"}
        s1 := Circle{length: 7, name: "s"}

        shapes := []Shape{c1, s1}
        
        for _, shape := range shapes {
            fmt.Println("Area of shape is:", shape.Area())
        }
    }
```

You can also use an empty interface{} for the same purpose of `any` in some other languages, and you can assign multiple interfaces to a struct or assign methods that do not take part of an interface to a struct that's using an interface.

## Trailing

There was use of interface to mix operations between Basic Bank Accounts and Current Accounts, but one method that could not be replicated was transfer, given that there's no overload in Golang.

What could be done (with the knowledge I have at this time) is to refactor Transfer() methods into onesided operations, so we'd create TransferFrom() and TransferTo() without referencing the other struct, whichever it would be. TransferFrom() would be like Withdraw(), but this time without fee, and TransferTo() would be like Deposit().
