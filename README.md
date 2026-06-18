# Go Learning Journey

This repository documents the learning progress of Go programming language through practical examples and a booking application project.

## Table of Contents
1. [Basic Syntax](#basic-syntax)
2. [Data Structures](#data-structures)
3. [Functions and Pointers](#functions-and-pointers)
4. [Structs and Methods](#structs-and-methods)
5. [Booking Application Project](#booking-application-project)
6. [Key Concepts Learned](#key-concepts-learned)

## Basic Syntax

Started with fundamental Go syntax including:
- Package declaration and imports
- Variable declaration (`var`, `:=`)
- Constants (`const`)
- Basic data types (int, string, bool, float)
- Control flow (if-else statements)
- Formatted printing with `fmt.Println()` and `fmt.Printf()`

Example from `01_basics/main.go`:
```go
func main() {
    fmt.Println("Hello Go!");
    age, name:=20, "dhiraj";
    const ab=23;
    fmt.Println(age);
    fmt.Println("square of 5",math.Pow(5,2));
    fmt.Println("name:",name);
    
    islogged:=true;
    issubscription:=!false
    fmt.Println(islogged && issubscription);
}
```

## Data Structures

### Arrays
Fixed-size collections with explicit length:
```go
var arr[3] int
fmt.Println(len(arr)) // 3
arr[1]=20
```

### Slices
Dynamic arrays that can grow and shrink:
```go
arr1 := []int{10, 20, 30}
arr2 := []int{40, 50}
arr3 := append(arr2, arr1...) // [40 50 10 20 30]
```

### Maps
Key-value pairs for efficient lookups:
```go
person:=map[string]string{
    "name":"Dhiraj",
    "addr":"Siraha",
    "clz":"wrc",
}
// Checking if key exists
vn,okn:=person["name"] // vn="Dhiraj", okn=true
vs,oks:=person["abc"]  // vs="", oks=false (key doesn't exist)

// Maps with different value types
details:=map[string]interface{}{
    "name":"Dhiraj",
    "addr":"Siraha",
    "class":10,
    "roll":11,
    "islogged":false,
    "marks": []int{20,30,90},
}
```

## Functions and Pointers

### Functions
Reusable blocks of code:
```go
func main() {
    // function body
}
```

### Pointers
References to memory addresses:
```go
var x int = 5
p := &x           // p holds the address of x
fmt.Println(p)    // prints the memory address
fmt.Println(*p)   // prints the value at that address (5)
*p = *p + 10      // modifies the value at the address
fmt.Println(*p)   // prints 15
fmt.Println(x)    // also prints 15 (x was modified through pointer)
```

## Structs and Methods

### Structs
Custom data types that group related fields:
```go
type Address struct {
    name string
    address string
    education string
    percentage int
    roll int
}

// Creating instances
var a = Address{"Dhiraj","wrc","bachelor",80,11}
a2:=Address{"yubraj","trichandra","bachelor",62,0}
a3:=Address{name:"yp",address:"bishnu"}
```

### Methods
Functions associated with structs:
```go
type Person struct {
    name string
    age int
    address string
    weight float64
    isMarried bool
    hobby  []string
}

// Value receiver (doesn't modify original)
func (p Person) displayValue() {
    fmt.Println(p.name)
    fmt.Println(p.weight)
    // ...
}

// Pointer receiver (can modify original)
func (p *Person) changeAttribute(newname string , age int){
    p.name=newname
    p.age=age
}
```

Usage:
```go
p:=Person{
    name:"Dhiraj",
    address:"kathmandu",
    age:20,
    isMarried:false,
    hobby:[]string{"playing cricket","watching movies"},
}
p.displayValue();
p.changeAttribute("Niraj",22) // modifies original p
p.displayValue(); // shows updated values
```

## Booking Application Project

A practical application demonstrating real-world Go concepts:

### Project Structure
```
Booking_app/
├── go.mod
├── main.go
├── validate.go
└── greatuser/
    └── greatuser.go
```

### Key Features Demonstrated

#### 1. Multi-package Organization
- `main` package: Application entry point
- `greatuser` package: Reusable functionality for user greetings

#### 2. Concurrency with Goroutines
```go
go sendTicket(firstName, lastName, email, userTickets)
// Creates a concurrent thread for sending tickets
```

#### 3. Input Handling
```go
func getUserInput() (string, string, string, int) {
    var firstName, lastName, email string
    var userTickets int
    fmt.Println("Enter your first Name:")
    fmt.Scan(&firstName)
    // ... similar for other inputs
    return firstName, lastName, email, userTickets
}
```

#### 4. Input Validation
```go
func validateUserInput(firstName, lastName, email string, userTickets int) (bool, bool, bool) {
    isValidUserName := len(firstName) >= 2 && len(lastName) >= 2
    isValidEmail := strings.Contains(email, "@")
    isValidUserTickets := userTickets > 0 && userTickets <= 50
    return isValidUserName, isValidEmail, isValidUserTickets
}
```

#### 5. String Manipulation
```go
firstNames := []string{}
for _, value := range bookings {
    fname := strings.Fields(value) // splits by whitespace
    firstNames = append(firstNames, fname[0])
}
```

#### 6. Time Handling
```go
time.Sleep(10 * time.Second) // Simulates delay in sending ticket
```

## Key Concepts Learned

### 1. Go Fundamentals
- Strong typing with type inference (`:=`)
- Explicit variable declaration when needed
- Package-based organization
- Import conventions

### 2. Memory Management
- Value types vs reference types
- Pointers for modifying variables in functions
- Slices as dynamic views over arrays
- Maps as hash tables for O(1) lookups

### 3. Object-Oriented Concepts (Go-style)
- Structs as composite types
- Methods with value and pointer receivers
- Encapsulation through package-level visibility
- Composition over inheritance

### 4. Concurrency
- Goroutines for lightweight concurrent execution
- Channels (not shown in current code but implied learning path)
- Synchronization patterns

### 5. Practical Application Development
- Project structure and modularity
- Input validation and sanitization
- Error handling patterns
- Standard library usage (fmt, strings, time)

### 6. Best Practices Observed
- Clear, readable code with comments
- Proper naming conventions
- Modular function design
- Separation of concerns (validation vs business logic)

## Learning Progression (Git Commit History)

1. **first commit** (dc33f53): Array basics, loops, main program, slice introduction
2. **map basics** (6a44465): Map operations and interface{} usage
3. **pointer** (296d4ad): Functions, pointers, and slice manipulation
4. **struct basics** (0a1031e): Basic struct definition and usage
5. **struct** (5966b05): Advanced struct methods with value and pointer receivers

This learning journey demonstrates progression from basic syntax to building a concurrent application, covering all essential Go features needed for real-world development.

## HTTP Web Server Examples

This directory contains practical examples of building HTTP servers in Go using the standard `net/http` package, located in the `httpmodule/` directory.

### Examples

#### 1. Basic HTTP Server (`basic_http_server`)
A simple HTTP server that listens on port 8000 and responds to GET requests at the `/hello` endpoint with a plain text message.

**Key concepts demonstrated:**
- Basic HTTP server setup with `http.ListenAndServe`
- Registering handlers with `http.HandleFunc`
- Simple text response writing
- Method validation (only allowing GET requests)

**To run:**
```bash
go run main.go
# Then visit http://localhost:8000/hello
```

#### 2. Multiple Routes (`http_multiple_routes`)
An HTTP server that demonstrates handling multiple routes and serving static files.

**Key concepts demonstrated:**
- Handling multiple routes (`/` and `/test`)
- Serving static files with `http.ServeFile`
- Processing URL query parameters
- Creating HTML forms that submit to the server
- Method validation for different endpoints

**To run:**
```bash
go run main.go
# Then visit http://localhost:8001
# Try submitting the form to see query parameter handling
```

#### 3. JSON Encoding/Decoding (`json_encoder`)
An HTTP server that demonstrates encoding and decoding JSON data in HTTP requests and responses.

**Key concepts demonstrated:**
- Setting response headers (Content-Type: application/json)
- Encoding Go structs to JSON with `json.NewEncoder`
- Decoding JSON request bodies into Go structs with `json.NewDecoder`
- Handling different HTTP methods (GET and POST)
- Proper HTTP status codes and headers

**To run:**
```bash
go run main.go
# Then visit:
# - GET http://localhost:8002/ (returns JSON greeting)
# - POST http://localhost:8002/decode (with JSON body to echo back)
```

### Common Patterns

All examples follow these common patterns:
- Method validation to ensure only allowed HTTP methods are processed
- Proper error handling with `http.Error`
- Clear separation of concerns with dedicated handler functions
- Use of standard Go `net/http` package without external dependencies

### Learning Outcomes

By studying these examples, you will learn how to:
1. Set up basic HTTP servers in Go
2. Handle different routes and HTTP methods
3. Process query parameters and form data
4. Serve static files (HTML, CSS, etc.)
5. Work with JSON in HTTP requests and responses
6. Set appropriate HTTP status codes and headers
7. Structure HTTP handlers for maintainability

Each example is self-contained and can be run independently.

---

*Learning completed on: 2026-06-18*