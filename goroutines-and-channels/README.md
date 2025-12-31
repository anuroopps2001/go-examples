# Go Concurrency – Consolidated Notes (Channels, WaitGroups, Worker Pools, Fan‑in)

---

## 1. Goroutines

* A goroutine is a lightweight thread managed by Go runtime
* Started using `go func()` or `go someFunction()`
* Goroutines run **concurrently**, not necessarily in parallel

Key rule:

> Goroutines do nothing special by themselves. They must interact (usually via channels or shared state) to be useful.

---

## 2. Channels – Core Concept

> A **channel is a pipe**, not a collection.

* Used to pass data between goroutines
* Can be **unbuffered** or **buffered**
* Channel operations are **blocking by default**

```go
ch := make(chan int)
```

---

## 3. Unbuffered Channels (Most Important)

### Behavior

* `send` blocks until a `receive` happens
* `receive` blocks until a `send` happens
* Sender and receiver must be **running concurrently**

```go
ch <- 1   // blocks until someone receives
<-ch      // blocks until someone sends
```

Key rule:

> With unbuffered channels, sender and receiver must be alive at the same time.

This is why goroutines are required.

---

## 4. Buffered Channels

```go
ch := make(chan int, 5)
```

* Sender can send up to `buffer size` without blocking
* Receiver still blocks if buffer is empty
* Buffering affects **performance**, not **correctness**

Key rule:

> Buffering changes *when* goroutines block, not *who* owns close().

---

## 5. `range` on Channels

```go
for v := range ch {
    fmt.Println(v)
}
```

What it means:

* Receive values **one by one** as they arrive
* Loop exits **only when channel is closed**

Important distinction:

| range slice         | range channel          |
| ------------------- | ---------------------- |
| Fixed size          | Unknown length         |
| Data exists upfront | Data arrives over time |

Key rule:

> If you use `range ch`, someone **must** close the channel.

---

## 6. `close(ch)` – What It Really Means

`close(ch)`:

* Does NOT delete the channel
* Does NOT stop goroutines
* Signals **end‑of‑stream**

After close:

* Receivers can still read remaining buffered values
* Further sends cause panic

Key rule:

> Closing a channel means “no more values will ever be sent”.

---

## 7. Who Should Close a Channel?

### The One Rule (Final Form)

> **The goroutine that knows that no more sends will ever happen must close the channel.**

### Practical table

| Situation         | Who closes  |
| ----------------- | ----------- |
| Single sender     | That sender |
| Multiple senders  | Coordinator |
| Workers receiving | Producer    |
| Receiver          | ❌ Never     |

---

## 8. WaitGroup (`sync.WaitGroup`)

Purpose:

* Tracks **lifecycle of goroutines**
* NOT related to data transfer

```go
wg.Add(n)
wg.Done()
wg.Wait()
```

Important clarifications:

* `wg.Add()` does NOT start goroutines
* It only declares how many `Done()` calls are expected

---

## 9. Worker Pools

### What is a worker pool?

> A fixed number of goroutines reused to process many jobs

### Core structure

* One **jobs channel**
* Multiple **worker goroutines**
* Workers `range` over jobs

```go
for job := range jobs {
    process(job)
}
```

Why worker pools exist:

* Limit concurrency
* Avoid spawning unlimited goroutines

---

## 10. Worker Pool WITHOUT Fan‑in (Jobs Only)

### Characteristics

* Workers **receive only**
* Only main sends jobs
* Only main closes jobs channel

```go
close(jobs)
wg.Wait()
```

Why this works:

* Workers exit when jobs channel closes
* Main can safely wait
* No fan‑in involved

---

## 11. Fan‑in (Critical Concept)

### Definition

> **Fan‑in = multiple goroutines send data into ONE channel, and one goroutine receives.**

Visual:

```
worker1 ─┐
worker2 ─┼──▶ results ───▶ main
worker3 ─┘
```

Fan‑in exists **only if**:

* Multiple goroutines SEND to the same channel

---

## 12. Why Fan‑in Needs a Coordinator

Problem:

* Multiple senders
* No sender knows when *all* others are done

Solution:

* Use `WaitGroup`
* Close channel after `wg.Wait()`

```go
go func() {
    wg.Wait()
    close(results)
}()
```

Key rule:

> If more than one goroutine sends on a channel, none of them should close it.

---

## 13. When Fan‑in Is NOT Needed

| Scenario                      | Fan‑in? |
| ----------------------------- | ------- |
| Workers only receive jobs     | ❌ No    |
| Single worker sends results   | ❌ No    |
| Multiple workers send results | ✅ Yes   |
| API calls with no data        | ❌ No    |

Fan‑in is about **data flow**, not goroutine count.

---

## 14. Deadlocks – Common Causes

### 1. Sending without a receiver

```go
ch <- 1   // no goroutine receiving
```

### 2. Waiting before closing a channel

```go
wg.Wait()  // workers blocked on range
close(ch) // never reached
```

### 3. `range ch` without `close(ch)`

Key rule:

> Deadlocks come from circular waiting, not from missing syntax.

---

## 15. Why Programs Sometimes Run “Too Fast”

Reasons:

* No real work (no sleep, no I/O)
* Unbuffered channel handshakes are very fast
* Printing is buffered by OS

Concurrency ≠ slowness

---

## 16. Mental Checklist (Use This Always)

Before writing concurrent Go code, ask:

1. Who sends?
2. Who receives?
3. Who closes the channel?
4. Who waits?
5. Is fan‑in involved?

If you can answer these, your design is correct.

---

## 17. How This Connects to Kubernetes Controllers

Controllers are built from:

* Worker pools
* Fan‑in (event sources)
* Queues
* Coordinators
* Long‑running goroutines

What’s left to learn next:

* `select`
* `context.Context`
* graceful shutdown

---