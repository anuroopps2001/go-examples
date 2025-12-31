# Go Concurrency – Select, `done`, and `context.Context`

This document summarizes **everything we covered today**, in the correct mental order. It is meant to be **re-read**, not memorized.

---

## 1. The Core Problem

Goroutines do **not** stop automatically.

Without extra control:

* A goroutine can block forever
* `WaitGroup` can only *wait*, not *stop*
* Closing data channels is not always possible or correct

We need a way to:

* stop goroutines **immediately**
* stop **multiple goroutines at once**
* stop even when no data is flowing

---

## 2. The `done` Channel

### What `done` is

```go
done := make(chan struct{})
```

`done` is:

* a **signal channel**
* not used for data
* closed to **broadcast stop**

```go
close(done) // sends stop signal to everyone
```

### What `done` means

> **"Stop whatever you are doing and exit now."**

---

## 3. Why `WaitGroup` Is NOT Enough

`WaitGroup` answers only one question:

> "Have all goroutines finished?"

It **cannot**:

* cancel work
* interrupt blocking goroutines
* force shutdown

### Correct roles

| Tool        | Purpose                     |
| ----------- | --------------------------- |
| `WaitGroup` | Wait for goroutines to exit |
| `done`      | Tell goroutines to stop     |

They are **complementary**, not interchangeable.

---

## 4. `select` – The Control Point

`select` lets a goroutine wait on **multiple events**.

### Basic pattern

```go
select {
case job := <-jobs:
    process(job)
case <-done:
    return
}
```

This means:

* work when work exists
* exit immediately when stop signal arrives

---

## 5. Why `default` Can Be Dangerous

```go
select {
case v := <-ch:
    fmt.Println(v)
default:
    fmt.Println("runs immediately")
}
```

Key rule:

> **If `default` exists, `select` never blocks.**

Use `default` only for **non-blocking checks**, not shutdown logic.

---

## 6. Why Producers and Workers Both Listen to `done`

Stopping only the producer is **not enough** because:

* workers may be blocked
* channels may be idle
* shutdown may be required immediately

### Correct rule

> **Every goroutine that sends OR receives must respect `done`.**

---

## 7. `close(data)` vs `done`

These mean **different things**.

| Action                | Meaning                                         |
| --------------------- | ----------------------------------------------- |
| `close(data)`         | No more data will be sent (graceful completion) |
| `done` / `ctx.Done()` | Stop immediately (forced shutdown)              |

### Important rule

> **Do NOT use `close(data)` to mean cancellation.**

---

## 8. When Producers Can Close Data Channels

A producer may `close(data)` **only if**:

* it is the **only sender**
* it knows no more data will be sent

If there are multiple producers:

* none of them close the channel
* a coordinator does

---

## 9. Replacing `done` with `context.Context`

### Conceptual equivalence

| `done`                | `context.Context`    |
| --------------------- | -------------------- |
| `make(chan struct{})` | `context.WithCancel` |
| `close(done)`         | `cancel()`           |
| `<-done`              | `<-ctx.Done()`       |

### Key idea

> **`ctx.Done()` is a standardized, propagating `done` channel.**

---

## 10. Why Context Is Preferred

`context.Context` adds:

* cancellation propagation
* timeouts
* deadlines
* standard API contracts

It is mandatory in:

* servers
* background workers
* Kubernetes controllers

---

## 11. Producer vs Worker Separation

### Producer

* creates work
* sends to channel
* never processes

### Worker

* receives work
* processes it
* never creates work

> **Producers and workers never call each other directly.**

Channels provide decoupling.

---

## 12. The Canonical Worker Loop (Final Mental Model)

```go
for {
    select {
    case <-ctx.Done():
        return
    case job := <-jobs:
        process(job)
    }
}
```

You do NOT memorize this.
You **recognize** it.

---

## 13. One-Sentence Summary 

> **`done` / `ctx.Done()` is a broadcast stop signal that allows multiple goroutines to exit immediately, regardless of data flow.**

---

