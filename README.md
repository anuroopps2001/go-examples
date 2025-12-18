# Go Fundamentals – Clean & Practical Notes

This document summarizes key Go concepts with clear explanations and practical examples.  
No emojis, no noise — just clean, easy‑to‑read notes.

---

## 1. Structs (Value Types)

Structs group related data into a single custom type:

```go
type Person struct {
    name string
    age  int
}
```

### Structs Are Value Types

When you pass a struct to a function, Go copies it:

```go
func update(p Person) {
    p.name = "New" // modifies only the copy
}

p := Person{name: "Old", age: 25}
update(p)
fmt.Println(p.name) // Output: Old
```

The original does not change.

### Modify the Original Struct Using a Pointer

```go
func update(p *Person) {
    p.name = "New"
}

p := Person{name: "Old", age: 25}
update(&p)
fmt.Println(p.name) // Output: New
```

---

## 2. Slices (Reference Types)

A slice contains:
- a pointer to an array  
- length  
- capacity  

Passing a slice to a function shares the same underlying array:

```go
func modify(s []int) {
    s[0] = 99
}

nums := []int{1, 2, 3}
modify(nums)
fmt.Println(nums) // Output: [99 2 3]
```

No pointer needed to modify the original slice.

---

## 3. Maps (Reference Types)

Maps point to a shared underlying hash table.

```go
func modify(m map[string]int) {
    m["score"] = 100
}

grades := map[string]int{"score": 50}
modify(grades)
fmt.Println(grades) // Output: map[score:100]
```

Again, no pointer required.

---

## 4. Pointers – Basic Concept

A pointer stores the memory address of a value.

- `&x` → address of x  
- `*p` → value at address p  

Example:

```go
x := 10
p := &x
fmt.Println(*p) // Output: 10
```

Pointers are needed mainly when working with structs to avoid copying.

---

## 5. Method Receivers

### Value Receiver

```go
func (p Person) display() {
    fmt.Println(p.name)
}
```

Works on a copy. Original struct is unchanged.

### Pointer Receiver

```go
func (p *Person) updateName(newName string) {
    p.name = newName
}
```

Modifies the original struct.

---

## 6. Automatic Conversion in Method Calls

Go automatically converts value ↔ pointer when calling methods.

### Value → Pointer Conversion

```go
p.updateName("A")
// Go internally does: (&p).updateName("A")
```

### Pointer → Value Conversion

```go
(&p).display()
// Go internally does: p.display()
```

This behavior makes methods easy to call.

---

## 7. When to Use Pointer vs Value Receivers

Use pointer receivers when:
- modifying the original struct  
- the struct is large  
- the type has methods that already use pointer receivers (consistency)

Use value receivers when:
- the method does not modify the struct  
- the struct is small  
- copy semantics are desired  

---

## 8. Behavior Summary Table

| Type     | Passed As | Modifies Original? | Needs Pointer? |
|----------|-----------|---------------------|----------------|
| Struct   | Value     | No                  | Yes            |
| Slice    | Reference | Yes                 | No             |
| Map      | Reference | Yes                 | No             |
| Array    | Value     | No                  | Yes            |
| String   | Value     | No (immutable)      | No             |

---

## 9. Memory Behavior – Simple Models

### Struct  
Like a box of data. Passing it copies the entire box.

### Slice  
Like a remote control pointing to a TV. Passing it gives another remote pointing to the same TV.

### Map  
Like shared storage. Everyone sees the same underlying data.

---

## 10. Complete Example

```go
package main

import "fmt"

type Person struct {
    name string
    age  int
}

// Value receiver (does not modify original)
func (p Person) printInfo() {
    fmt.Println("Name:", p.name, "Age:", p.age)
}

// Pointer receiver (modifies original)
func (p *Person) setName(newName string) {
    p.name = newName
}

func main() {
    p := Person{name: "Anuroop", age: 25}

    // Calling value receiver
    p.printInfo() // Output: Name: Anuroop Age: 25

    // Calling pointer receiver using value (auto converted)
    p.setName("Updated")
    p.printInfo() // Output: Name: Updated Age: 25
}
```

---

## 11. Key Takeaways

- Go is **always pass-by-value**.  
- Structs require pointers for modification.  
- Slices and maps automatically modify the original.  
- Method calls automatically convert between value and pointer.  
- Pointer receivers describe behavior, not calling syntax.

