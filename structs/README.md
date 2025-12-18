# Go Fundamentals Cheat Sheet

## 1. Structs (Value Types)

A struct is used to group related data together:

```go
type Person struct {
    name string
    age  int
}
```

### Structs are Value Types
- Passing a struct to a function creates a **copy**
- Modifying inside the function does **not change the original**

```go
func update(p Person) {
    p.name = "New" // modifies copy
}
```

### Use Pointers to Modify Original Struct

```go
func update(p *Person) {
    p.name = "New"
}
```

---

## 2. Slices (Reference Types)

A slice internally contains:
- Pointer to underlying array  
- Length  
- Capacity  

Passing a slice to a function shares the same underlying data.

```go
func change(s []int) {
    s[0] = 10
}
```

### No pointer needed to modify original slice.

Pointers only required when:
- replacing entire slice  
- resetting  
- reassigning  

---

## 3. Maps (Reference Types)

Maps store pointers to internal hash tables.

Passing a map to a function shares original map.

```go
func change(m map[string]int) {
    m["x"] = 100
}
```

No pointer needed to modify the original map.

---

## 4. Pointer Concept

- `&x` → address of x  
- `*p` → value at p  

Use pointers to modify original struct.

Not required for slices/maps.

---

## 5. Method Receivers

### Value Receiver
```go
func (p Person) print() {}
```

### Pointer Receiver
```go
func (p *Person) updateName(name string) {}
```

---

## Automatic Conversion in Go

Go automatically converts between pointer and value when calling methods.

### You CAN call pointer receiver using value:

```go
p.updateName("A") // Go converts to (&p).updateName("A")
```

### You CAN call value receiver using pointer:

```go
(&p).print() // Go converts to p.print()
```

---

## 6. When to Use Pointer Receivers vs Value Receivers

### Pointer Receiver (`*T`)
Use when:
- modifying original struct  
- struct is large  
- maintaining consistency  

### Value Receiver (`T`)
Use when:
- not modifying original  
- struct is small  
- need copy semantics  

---

## 7. Behavior Summary Table

| Type     | Passed As | Modifies Original? | Needs Pointer? |
|----------|-----------|---------------------|----------------|
| Struct   | Value     | ❌ No               | ✔ Yes          |
| Slice    | Reference | ✔ Yes              | ❌ No          |
| Map      | Reference | ✔ Yes              | ❌ No          |
| Array    | Value     | ❌ No               | ✔ Yes          |
| String   | Value     | ❌ No (immutable)   | ❌ No          |

---

## 8. Memory Rules (Easy to Remember)

### Struct = Box  
Copying struct → new box  
Pointer → edit the same box

### Slice = Remote Control  
Passing slice → still controls same underlying array  
No pointer needed

### Map = Shared Storage  
Everyone sees same storage  
No pointer needed

---

## 9. Key Takeaways

- Structs need pointers for modification  
- Slices & maps modify original automatically  
- Methods auto-convert pointer/value on calls  
- Pointer receivers are about **behavior**, not calling syntax  
- Go simplifies method calling, not method behavior  

---

