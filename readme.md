# INTERMEDIATE GOLANG CONCEPTS🔵
_Curated with 💖 by [Soumadip "Skyy" Banerjee 👨🏻‍💻](https://www.instagram.com/iamskyy666/)_



## **1. What is a Closure in Go?**

In Go, a **closure** is a function value that **captures variables from the scope in which it was defined**, even after that scope has finished executing.

That means:

* The **inner function** remembers the environment (variables) from where it was created.
* These captured variables **live on in memory** as long as the returned function value exists.
* Each returned function has its **own copy** of those variables.

---

### **Analogy**

Think of closures like:

> A backpack 🎒 that a function carries with it wherever it goes, containing the variables from the place where it was born.

---

## **2. The `adder()` example**

```go
func adder() func() int {
    i := 0
    fmt.Println("Previous value of i:", i)

    return func() int {
        i++
        fmt.Println("Adding 1 to i")
        return i
    }
}
```

### Step-by-Step:

1. When `adder()` is called, it creates a **local variable** `i := 0`.
2. It returns an **anonymous function** (`func() int { ... }`) that:

   * Accesses and modifies `i` even after `adder()` has finished running.
3. This works because the **closure** keeps a reference to `i` in memory (on the heap, not the stack).

---

### How `main()` uses it:

```go
sequence := adder()

fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
```

**Execution flow:**

* `sequence := adder()`

  * Calls `adder()` → sets `i = 0` → returns a closure that can increment `i`.
  * `sequence` now holds that closure **and the "backpack" with i**.
* First `sequence()` → increments `i` to 1 → prints and returns 1.
* Second `sequence()` → increments `i` to 2 → prints and returns 2.
* Third `sequence()` → increments `i` to 3 → prints and returns 3.
* Fourth → increments to 4.

---

### Important: `sequence2 := adder()`

* This **calls `adder()` again**, creating a **new and independent `i`** starting from 0.
* That’s why `sequence2()` starts counting from 1 again.
* Each closure has **its own copy** of `i`.

---

## **3. The `multiplier` example inside `main`**

```go
multiplier := func() func(int) int {
    product := 2
    return func(x int) int {
        product *= x
        return product
    }
}()
```

### What’s happening:

1. We define an **anonymous function** that:

   * Declares `product := 2`
   * Returns a closure `func(x int) int` that multiplies `product` by `x`.
2. The `()` at the end **calls it immediately**, so `multiplier` becomes:

   * The closure `func(int) int` with `product` in its backpack.

**Usage:**

```go
fmt.Println(multiplier(5)) // product = 2 * 5 = 10
fmt.Println(multiplier(3)) // product = 10 * 3 = 30
```

* Notice that `product` **persists between calls**, because it’s stored in the closure’s environment.

---

## **4. Memory Perspective (Heap vs Stack)**

* Normally, a local variable like `i` or `product` would be on the **stack**, and vanish after the function returns.
* But closures need those variables **after the function exits**.
* Go’s compiler detects this and **allocates them on the heap** instead.
* The closure stores a reference to them → as long as the closure is alive, the variables stay alive in memory.

---

## **5. Why Closures are Useful**

* **Stateful functions** without global variables.
* **Factory functions** (like `adder` and `multiplier`) that produce customized behavior.
* **Callbacks** that need to remember context.
* **Encapsulation**: we hide variables from the outside world but keep them accessible to the inner function.

---

## **1. What is Recursion?**

Recursion is when a function **calls itself**, either **directly** or **indirectly**, to solve a problem.

A recursive function:

1. Has a **base case** — a condition where it stops calling itself.
2. Has a **recursive case** — where it calls itself with smaller/simpler input.

---

### **General Syntax in Go**

```go
func functionName(args) returnType {
    if baseCondition {
        return baseValue
    }
    // recursive case
    return functionName(smallerArgs)
}
```

---

## **2. Example — Factorial**

Mathematically:

```
n! = n × (n-1) × (n-2) × ... × 1
0! = 1 (base case)
```

**In Go:**

```go
package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1 // base case
    }
    return n * factorial(n-1) // recursive call
}

func main() {
    fmt.Println(factorial(5)) // Output: 120
}
```

**How it works:**

```
factorial(5) → 5 * factorial(4)
factorial(4) → 4 * factorial(3)
factorial(3) → 3 * factorial(2)
factorial(2) → 2 * factorial(1)
factorial(1) → 1 * factorial(0)
factorial(0) → 1  (base case reached)

Now it returns:
5 * 4 * 3 * 2 * 1 = 120
```

---

## **3. Example — Fibonacci**

Fibonacci sequence:

```
0, 1, 1, 2, 3, 5, 8, ...
fib(n) = fib(n-1) + fib(n-2)
```

**In Go:**

```go
func fibonacci(n int) int {
    if n <= 1 {
        return n // base cases: fib(0)=0, fib(1)=1
    }
    return fibonacci(n-1) + fibonacci(n-2) // recursive case
}
```

---

## **4. Important Points in Go Recursion**

* **Base case is mandatory** — without it, recursion runs forever and crashes with a stack overflow.
* Go optimizes tail recursion *sometimes*, but not as aggressively as some languages (like Scheme).
* Recursion is more readable for problems like:

  * Tree traversal
  * Divide and conquer algorithms
  * Mathematical definitions
* But **iteration is often faster** in Go due to lower function call overhead.

---

## **5. Stack vs Heap in Recursion**

* Every recursive call adds a **stack frame** (variables, return address).
* Deep recursion can cause **stack overflow** if too many calls happen before hitting the base case.
* That’s why we must ensure **small inputs progress toward the base case**.

---

# What is a pointer?

A **pointer** is a value that holds the **memory address** of another value. In Go a pointer type is written `*T` — “pointer to `T`”.

```go
var a int = 42
var p *int = &a   // p is a pointer to a
fmt.Println(*p)   // 42   (dereference)
```

* `&` — **address-of** operator: `&a` returns `*int`.
* `*` — **dereference** operator: `*p` gives the value stored at that address.

# Why use pointers?

* **Modify caller’s value**: functions receive copies by default; pass a pointer to let the callee change the original.
* **Avoid large copies**: pass pointer to big structs instead of copying them.
* **Shared mutable state** when multiple parts of code should see updates.

# Basic examples

**Modify an int via pointer**

```go
func inc(x *int) {
    *x++         // change the caller's variable
}
n := 5
inc(&n)
fmt.Println(n) // 6
```

**Swap two values**

```go
func swap(a, b *int) {
    *a, *b = *b, *a
}
```

# nil pointers and panics

* A pointer’s zero value is `nil`.
* **Dereferencing `nil`** causes a runtime panic: `panic: runtime error: invalid memory address or nil pointer dereference`.

```go
var p *int
fmt.Println(p == nil) // true
// fmt.Println(*p) // panic
```

Always check for `nil` before dereferencing if it might be nil.

# Pointers & composite types

* **Pointer to struct**

  ```go
  type User struct { Name string }
  u := &User{Name: "Skyy"} // shorthand, allocates and returns *User
  ```

* **Pointer to array**

  ```go
  arr := [3]int{1,2,3}
  var pa *[3]int = &arr
  fmt.Println(pa[1]) // 2  (pointer-to-array supports indexing)
  ```

* **Slices, maps, channels** are *reference types* already (they contain internal pointers). You rarely need `*[]T` or `*map[K]V`. Use pointer to slice only if you must replace the slice header itself (length/capacity) in the caller.

# `new` vs `&T{}` vs `make`

* `new(T)` returns `*T` with zeroed T (allocates memory, may be on heap).

  ```go
  p := new(int) // p *int, *p == 0
  ```
* `&T{}` or `&MyStruct{Field:...}` is common and idiomatic — often the compiler allocates on the heap if needed (escape analysis).
* `make` is for **slices, maps, channels** (it returns an initialized value, not a pointer).

# Returning pointer to local variable — safe

Go performs **escape analysis**. If a local variable is returned via pointer, the compiler will move it to the heap — that’s safe.

```go
func newCounter() *int {
    i := 0
    return &i   // safe: i escapes to heap
}
```

# Pointer receivers vs value receivers (methods)

* **Pointer receiver** `func (p *T) M()`:

  * Method can modify the receiver.
  * Avoids copying large structs.
* **Value receiver** `func (t T) M()`:

  * Method gets a copy; cannot change original.
  * Good for small, immutable receivers.

**Note:** you can call pointer-receiver methods on addressable values — the compiler will take the address automatically (`v.M()` becomes `(&v).M()` if `v` is addressable). But an interface holding a `T` value does not automatically become `*T`.

```go
type Big struct{ buf [1024]byte }

func (b *Big) Zero() { ... }  // pointer receiver to avoid copying
```

# Pointer equality and comparison

* You can compare pointers: `p == q`, `p != nil`.
* Two pointers are equal when they point to the same memory location.

# Pointer to pointer

```go
var x int = 1
var p *int = &x
var pp **int = &p
fmt.Println(**pp) // 1
```

# No pointer arithmetic

Go **does not** allow pointer arithmetic like C. That makes programs safer and more portable.

# Concurrency considerations

Pointers are frequently shared across goroutines. **Sharing mutable state requires synchronization** (e.g., `sync.Mutex`, `sync/atomic`) to avoid race conditions.

```go
// Danger: race if two goroutines modify shared *int without sync
```

Use the Go race detector (`go run -race`) in development.

# Performance trade-offs

* Passing a pointer avoids copying big values — saves CPU and memory.
* But pointers may force heap allocation (GC overhead) and hurt cache locality. For small structs, prefer value receivers and passing by value.
* Measure (benchmarks) rather than guessing.

# Unsafe pointers

* `unsafe.Pointer` exists and can do conversions or mimic pointer arithmetic — **dangerous** and platform-dependent. Use only when necessary and with full understanding.

# Practical idioms & patterns

* **Factory returning pointer**

  ```go
  func NewUser(name string) *User { return &User{Name: name} }
  ```

* **Mutable method via pointer receiver**

  ```go
  func (u *User) SetName(n string) { u.Name = n }
  ```

* **Avoid pointer-to-slice/map unless needed** — slices/maps/channels already behave like references.

* **Use pointer for large structs** or when mutation is required.

# Example: using pointer receiver vs value receiver

```go
type Counter struct { n int }

func (c *Counter) Inc() { c.n++ }      // modifies receiver
func (c Counter) Value() int { return c.n } // read-only

c := Counter{}
c.Inc()             // works; compiler takes &c
fmt.Println(c.Value())
```

# Debugging tips

* `fmt.Printf("%p\n", p)` prints pointer address.
* `go vet`, `go test -race` find misuse in concurrent code.
* Be mindful of `nil` checks before deref.

# Quick checklist for using pointers

* Want to **modify** the argument? → pass pointer.
* Passing **large struct**? → consider pointer to avoid copying.
* Returning early-created value? → `&value` is fine (escape analysis handles heap).
* Using slices/maps/channels? → usually **no pointer needed**.
* Sharing across goroutines? → **sync** or use channels.

---

## Short summary

* Pointers (`*T`) hold addresses; `&` & `*` are address-of and dereference.
* They let you mutate caller’s data and avoid copying large values.
* Nil pointer dereferences panic — always guard if needed.
* Go forbids pointer arithmetic; `unsafe` exists for low-level cases.
* Use pointers for mutable/large data, but measure performance and be mindful of GC and concurrency.

Here’s the benchmark result for simulating a large struct in Python to compare **pass-by-value** vs **pass-by-reference**:

* **Pass-by-value simulation** (copy-like behavior): **\~0.02948 seconds**
* **Pass-by-reference simulation** (pointer-like behavior): **\~0.000096 seconds**

📌 **Interpretation in the context of Go:**

* In Go, when we pass large structs **by value**, it copies all the fields to a new memory location (on the stack or heap), which is **slower** and uses **more memory**.
* When we pass them **by pointer** (e.g., `*MyStruct`), we’re only passing an **8-byte memory address**, so it’s much **faster** and avoids extra memory allocations.
* This is exactly why in real-world Go projects, for large structs, functions often take `*Struct` instead of `Struct`.
---

## **1️⃣ What is a String in Go?**

In Go:

```go
str := "Hello"
```

* A **string** is **immutable** — we cannot change its contents once created.
* Internally, a `string` is just a **read-only slice of bytes**:

  * **Pointer to the underlying byte array**
  * **Length of the string**

We can visualize:

```
"Hello"
[Pointer] ───▶ [72 101 108 108 111]  // ASCII byte values
Length = 5
```

📌 **Key points:**

* Strings are stored as **UTF-8 encoded bytes**.
* Since UTF-8 can take 1–4 bytes per character, indexing a string directly (`str[i]`) gives **bytes**, **not characters**.

---

## **2️⃣ UTF-8 and Why `str[i]` Might Not Work for All Characters**

Example:

```go
str := "Hé"
fmt.Println(len(str))  // 3
fmt.Println(str[0])    // 72 ('H')
fmt.Println(str[1])    // 195 (part of 'é')
fmt.Println(str[2])    // 169 (rest of 'é')
```

Here, `"é"` takes **2 bytes** (`0xC3 0xA9`), so:

* Index `1` gives the **first byte** of `é`, not the whole character.

---

## **3️⃣ What is a Rune in Go?**

* A **rune** is just an alias for `int32` in Go.
* It represents a **Unicode code point** — the numeric value of a character.
* Unlike a byte, a rune **can store the entire character**, even if it’s multi-byte in UTF-8.

Example:

```go
var r rune = 'é'
fmt.Println(r)         // 233 (Unicode code point)
fmt.Printf("%c\n", r)  // é
```

---

## **4️⃣ Converting Strings to Runes**

If we want to process actual **characters** instead of raw bytes:

```go
str := "Hé😊"
runes := []rune(str)
fmt.Println(len(runes)) // 3 (characters)
fmt.Println(runes)      // [72 233 128522]
```

Here:

* `"H"` → `72`
* `"é"` → `233`
* `"😊"` → `128522` (U+1F60A)

---

## **5️⃣ Iterating Over a String Correctly**

🔴 **Wrong way** (byte by byte):

```go
for i := 0; i < len(str); i++ {
    fmt.Printf("%c ", str[i]) // May print broken chars for multi-byte
}
```

✅ **Correct way** (range over runes):

```go
for i, r := range str {
    fmt.Printf("Index %d: %c\n", i, r)
}
```

* `range` automatically decodes UTF-8 into runes.
* `i` will be the **byte index**, not rune index.

---

## **6️⃣ Memory & Performance Notes**

| Type     | Size     | Meaning                            | Use Case                    |
| -------- | -------- | ---------------------------------- | --------------------------- |
| `byte`   | 1 byte   | ASCII value / raw byte             | Low-level byte handling     |
| `rune`   | 4 bytes  | Unicode code point (UTF-8 decoded) | Proper character handling   |
| `string` | 16 bytes | Read-only byte slice (pointer+len) | Storing immutable text data |

* Strings are **immutable**, so modifying them means creating a **new string**.
* Runes are stored in UTF-8 form inside strings, so multi-byte characters are **not** directly indexable.

---

## **7️⃣ Real-world Example**

```go
package main

import "fmt"

func main() {
    str := "Go💙"

    fmt.Println("Length in bytes:", len(str))       // 6
    fmt.Println("Length in runes:", len([]rune(str))) // 3

    fmt.Println("\nIterating over runes:")
    for i, r := range str {
        fmt.Printf("Byte index: %d, Rune: %c, CodePoint: %U\n", i, r, r)
    }
}
```

**Output:**

```
Length in bytes: 6
Length in runes: 3

Iterating over runes:
Byte index: 0, Rune: G, CodePoint: U+0047
Byte index: 1, Rune: o, CodePoint: U+006F
Byte index: 2, Rune: 💙, CodePoint: U+1F499
```

---

✅ **In short:**

* **String** = Immutable UTF-8 bytes.
* **Rune** = Unicode code point (`int32`).
* Use `[]byte` for raw byte processing, `[]rune` for character-safe processing.
* Always remember UTF-8 characters can be multi-byte — `str[i]` is not always a full character.

---

## **1️⃣ What Are Formatting Verbs?**

In Go, **formatting verbs** are special placeholders we use inside `fmt` functions like `fmt.Printf`, `fmt.Sprintf`, etc., to control **how data is printed**.

Example:

```go
name := "Skyy"
age := 29
fmt.Printf("Name: %s, Age: %d\n", name, age)
```

Output:

```
Name: Skyy, Age: 29
```

Here:

* `%s` → prints a string
* `%d` → prints a decimal integer

---

## **2️⃣ Common Categories of Formatting Verbs**

Let’s break them down:

---

### **A. General Verbs**

| Verb  | Meaning                                     |
| ----- | ------------------------------------------- |
| `%v`  | Default format (auto-chooses based on type) |
| `%+v` | Like `%v`, but with struct field names      |
| `%#v` | Go syntax representation of the value       |
| `%T`  | Type of the value                           |
| `%%`  | Literal `%` character                       |

Example:

```go
type Person struct {
    Name string
    Age  int
}
p := Person{"Alice", 30}

fmt.Printf("%v\n", p)   // {Alice 30}
fmt.Printf("%+v\n", p)  // {Name:Alice Age:30}
fmt.Printf("%#v\n", p)  // main.Person{Name:"Alice", Age:30}
fmt.Printf("%T\n", p)   // main.Person
fmt.Printf("%%\n")      // %
```

---

### **B. Boolean**

| Verb | Meaning       |
| ---- | ------------- |
| `%t` | true or false |

```go
fmt.Printf("%t\n", true) // true
```

---

### **C. Integer**

| Verb | Meaning                             |
| ---- | ----------------------------------- |
| `%d` | Decimal (base 10)                   |
| `%b` | Binary                              |
| `%o` | Octal                               |
| `%x` | Hexadecimal (lowercase)             |
| `%X` | Hexadecimal (uppercase)             |
| `%c` | Character (from Unicode code point) |
| `%q` | Quoted character literal            |

Example:

```go
n := 65
fmt.Printf("%d\n", n)  // 65
fmt.Printf("%b\n", n)  // 1000001
fmt.Printf("%o\n", n)  // 101
fmt.Printf("%x\n", n)  // 41
fmt.Printf("%X\n", n)  // 41
fmt.Printf("%c\n", n)  // A
fmt.Printf("%q\n", n)  // 'A'
```

---

### **D. Floating Point & Complex Numbers**

| Verb | Meaning                                 |
| ---- | --------------------------------------- |
| `%f` | Decimal point, no exponent              |
| `%F` | Same as `%f`                            |
| `%e` | Scientific notation (e.g., `1.23e+03`)  |
| `%E` | Scientific notation (uppercase)         |
| `%g` | `%e` or `%f`, whichever is more compact |
| `%G` | `%E` or `%F`, whichever is more compact |
| `%x` | Hexadecimal fraction (lowercase)        |
| `%X` | Hexadecimal fraction (uppercase)        |

Example:

```go
pi := 3.1415926535
fmt.Printf("%f\n", pi)  // 3.141593
fmt.Printf("%.2f\n", pi) // 3.14
fmt.Printf("%e\n", pi)  // 3.141593e+00
fmt.Printf("%g\n", pi)  // 3.14159
```

---

### **E. String & Slice of Bytes**

| Verb | Meaning                          |
| ---- | -------------------------------- |
| `%s` | Plain string                     |
| `%q` | Double-quoted string (Go syntax) |
| `%x` | Hexadecimal (lowercase)          |
| `%X` | Hexadecimal (uppercase)          |

Example:

```go
s := "GoLang"
fmt.Printf("%s\n", s)  // GoLang
fmt.Printf("%q\n", s)  // "GoLang"
fmt.Printf("%x\n", s)  // 476f4c616e67
fmt.Printf("%X\n", s)  // 476F4C616E67
```

---

### **F. Pointer**

| Verb | Meaning        |
| ---- | -------------- |
| `%p` | Pointer in hex |

Example:

```go
x := 42
fmt.Printf("%p\n", &x) // e.g., 0xc000014090
```

---

## **3️⃣ Width, Precision, and Flags**

We can fine-tune formatting:

* **Width**: Minimum number of characters (`%6d` → pad with spaces if less than 6 digits)
* **Precision**: Digits after decimal for floats (`%.2f` → 2 decimal places)
* **Flags**:

  * `-` → Left-justify
  * `+` → Always show sign
  * `0` → Pad with zeros

Example:

```go
n := 42
pi := 3.14159

fmt.Printf("|%6d|\n", n)    // |    42|
fmt.Printf("|%-6d|\n", n)   // |42    |
fmt.Printf("|%06d|\n", n)   // |000042|
fmt.Printf("%.2f\n", pi)    // 3.14
fmt.Printf("%+f\n", pi)     // +3.141590
```

---

## **4️⃣ Summary Table of Most Common Verbs**

| Verb        | Usage                   |
| ----------- | ----------------------- |
| `%v`        | Default value           |
| `%+v`       | Struct with field names |
| `%#v`       | Go syntax               |
| `%T`        | Type                    |
| `%t`        | Boolean                 |
| `%d`        | Decimal integer         |
| `%b`        | Binary                  |
| `%o`        | Octal                   |
| `%x` / `%X` | Hexadecimal             |
| `%c`        | Unicode character       |
| `%q`        | Quoted string/char      |
| `%f`        | Decimal float           |
| `%e` / `%E` | Scientific notation     |
| `%g` / `%G` | Compact float           |
| `%p`        | Pointer address         |

---

### **String Comparison in GoLang** (Lexicographical Order)

In GoLang, when we compare two strings using relational operators (`<`, `>`, `==`, etc.), the comparison is done **lexicographically** — meaning it follows the order of Unicode code points, very similar to how words are sorted in a dictionary.

#### **How It Works**

* Go compares strings **byte by byte**, starting from the first character of both strings.
* It uses the **Unicode code point** values (UTF-8 encoded) for comparison.
* The comparison stops as soon as it finds the first differing byte between the two strings.
* If all compared characters are equal and the strings are of different lengths, the shorter string is considered "less" than the longer one.

---

#### **Example**

```go
str1 := "Apple 🍎"
str2 := "Banana 🍌"
str3 := "App 📱"
str4 := "ab 🔤"

fmt.Println(str1 < str2) // true
fmt.Println(str3 < str1) // true
fmt.Println(str4 < str1) // false
```

**Step-by-step reasoning:**

1. `"Apple 🍎" < "Banana 🍌"` → We compare `'A'` (65) with `'B'` (66). Since `65 < 66`, result is `true`.
2. `"App 📱" < "Apple 🍎"` → First three characters `"App"` match. Next, space `' '` (32) is less than `'l'` (108), so result is `true`.
3. `"ab 🔤" < "Apple 🍎"` → First characters: `'a'` (97) vs `'A'` (65). Since `97 > 65`, result is `false`.

---

💡 **Important Notes:**

* Go is **case-sensitive** in string comparisons because Unicode code points for uppercase and lowercase letters are different.
* Emojis and non-ASCII characters also follow their Unicode code point order.
* If we want case-insensitive comparison, we can use functions from the `strings` package like `strings.EqualFold()`.

---

Here’s our **Go Formatting Verbs Cheat Sheet** so we can quickly reference them whenever we print or format values in Go.
We’ll group them by category for clarity.

---

## **📌 1. General Verbs**

| Verb  | Meaning                  | Example Output                  |
| ----- | ------------------------ | ------------------------------- |
| `%v`  | Default format           | `fmt.Printf("%v", 123)` → `123` |
| `%+v` | Struct with field names  | `"{Name:Skyy Age:29}"`          |
| `%#v` | Go-syntax representation | `"[]int{1, 2, 3}"`              |
| `%T`  | Type of the value        | `"int"`, `"string"`             |
| `%%`  | Literal percent sign     | `"50%"`                         |

---

## **📌 2. Boolean**

| Verb | Meaning           | Example  |
| ---- | ----------------- | -------- |
| `%t` | `true` or `false` | `"true"` |

---

## **📌 3. Integer**

| Verb | Meaning                   | Example      |
| ---- | ------------------------- | ------------ |
| `%b` | Base 2 (binary)           | `1010`       |
| `%c` | Character from code point | `'A'` for 65 |
| `%d` | Base 10 (decimal)         | `123`        |
| `%o` | Base 8 (octal)            | `173`        |
| `%q` | Quoted character          | `'A'`        |
| `%x` | Base 16 lowercase         | `7b`         |
| `%X` | Base 16 uppercase         | `7B`         |
| `%U` | Unicode format            | `U+0041`     |

---

## **📌 4. Floating-Point & Complex Numbers**

| Verb        | Meaning                                            | Example      |
| ----------- | -------------------------------------------------- | ------------ |
| `%b`        | Exponent as power of two                           | `132p-2`     |
| `%e`        | Scientific notation (lowercase)                    | `1.23e+03`   |
| `%E`        | Scientific notation (uppercase)                    | `1.23E+03`   |
| `%f` / `%F` | Decimal point, no exponent                         | `123.456000` |
| `%g`        | Compact representation (smallest between %e or %f) | `1.23`       |
| `%G`        | Compact representation (%E or %F)                  | `1.23`       |

---

## **📌 5. String & Slice of Bytes**

| Verb | Meaning                   | Example       |
| ---- | ------------------------- | ------------- |
| `%s` | Plain string              | `"Hello"`     |
| `%q` | Quoted string (Go syntax) | `"\"Hello\""` |
| `%x` | Hexadecimal (lowercase)   | `68656c6c6f`  |
| `%X` | Hexadecimal (uppercase)   | `68656C6C6F`  |

---

## **📌 6. Pointer**

| Verb | Meaning               | Example        |
| ---- | --------------------- | -------------- |
| `%p` | Pointer address (hex) | `0xc0000140a0` |

---

## **📌 7. Width & Precision**

* **Width**: Minimum number of characters to print (pads with spaces by default)

  ```go
  fmt.Printf("|%6d|", 45) // |    45|
  ```
* **Precision**: Digits after decimal for floats / max chars for strings

  ```go
  fmt.Printf("|%.2f|", 3.14159) // |3.14|
  fmt.Printf("|%.3s|", "Golang") // |Gol|
  ```
* **Width + Precision**:

  ```go
  fmt.Printf("|%6.2f|", 3.14159) // |  3.14|
  ```

---

## **📌 8. Flags**

| Flag | Meaning                                    | Example             |
| ---- | ------------------------------------------ | ------------------- |
| `-`  | Left align                                 | `%-6d` → `"45    "` |
| `+`  | Always show sign                           | `%+d` → `+45`       |
| `0`  | Pad with zeros                             | `%06d` → `000045`   |
| `#`  | Alternate form (e.g., `0x` prefix for hex) | `%#x` → `0x7b`      |

---

Let’s dive into the **`fmt` package** in Go and understand it from our perspective as a team learning Golang deeply.

---

## **1. Overview of `fmt`**

The **`fmt` package** in Go (short for *format*) is part of Go’s standard library and is used for **formatted I/O** — meaning it helps us print to the console, read from input, and format data into strings.

It’s somewhat similar to:

* `printf` / `scanf` in C
* `console.log` + `String.format` combined in JavaScript/Java

The `fmt` package covers:

* **Printing** data to the console or other outputs
* **Formatting** data into strings
* **Scanning** data from inputs (stdin, strings, etc.)

---

## **2. Three Main Groups of Functions**

### **A. Printing Functions**

These are for output (console or any writer).

| Function                                                        | Description                                                      |
| --------------------------------------------------------------- | ---------------------------------------------------------------- |
| **`fmt.Print(a ...interface{})`**                               | Prints arguments in their default format (no newline).           |
| **`fmt.Println(a ...interface{})`**                             | Prints arguments separated by spaces, then adds a newline.       |
| **`fmt.Printf(format string, a ...interface{})`**               | Prints using a **format string** with verbs like `%d`, `%s`.     |
| **`fmt.Fprint(w io.Writer, a ...interface{})`**                 | Prints to any `io.Writer` (like a file, buffer) without newline. |
| **`fmt.Fprintln(w io.Writer, a ...interface{})`**               | Prints to `io.Writer` with newline.                              |
| **`fmt.Fprintf(w io.Writer, format string, a ...interface{})`** | Prints formatted string to `io.Writer`.                          |

---

### **B. Formatting into Strings**

Instead of printing to the console, these return the formatted data as a **string**.

| Function                                    | Description                                         |
| ------------------------------------------- | --------------------------------------------------- |
| **`fmt.Sprint(a ...interface{})`**          | Returns a string with arguments in default format.  |
| **`fmt.Sprintln(a ...interface{})`**        | Returns string with space-separated args + newline. |
| **`fmt.Sprintf(format, a ...interface{})`** | Returns formatted string with format verbs.         |

Example:

```go
name := "Bruce"
age := 29
msg := fmt.Sprintf("Name: %s, Age: %d", name, age)
fmt.Println(msg) // Name: Bruce, Age: 29
```

---

### **C. Scanning (Reading Input)**

Used for reading values from input.

| Function                                                       | Description                                                 |
| -------------------------------------------------------------- | ----------------------------------------------------------- |
| **`fmt.Scan(a ...interface{})`**                               | Reads space-separated values from **stdin** into variables. |
| **`fmt.Scanln(a ...interface{})`**                             | Like `Scan` but stops at newline.                           |
| **`fmt.Scanf(format string, a ...interface{})`**               | Reads input in a **specific format**.                       |
| **`fmt.Fscan(r io.Reader, a ...interface{})`**                 | Reads from any `io.Reader`.                                 |
| **`fmt.Fscanf(r io.Reader, format string, a ...interface{})`** | Reads from `io.Reader` using a format string.               |
| **`fmt.Sscan(str string, a ...interface{})`**                  | Reads from a string into variables.                         |
| **`fmt.Sscanf(str string, format string, a ...interface{})`**  | Reads from a string with specific format.                   |

Example:

```go
var name string
var age int
fmt.Scan(&name, &age) // Input: Bruce 29
fmt.Println(name, age) // Bruce 29
```

---

## **3. Formatting Verbs in `fmt.Printf`**

We already made a **big cheat sheet** earlier, but here’s the overview:

### **General:**

* `%v` → default format
* `%#v` → Go-syntax representation
* `%T` → type of value
* `%%` → prints a literal `%`

### **Integers:**

* `%d` → decimal
* `%+d` → decimal with sign
* `%b` → binary
* `%o` / `%O` → octal
* `%x` / `%X` → hexadecimal

### **Strings & Chars:**

* `%s` → plain string
* `%q` → quoted string
* `%x` → hex representation
* `%c` → character

### **Booleans:**

* `%t` → `true` / `false`

### **Floats:**

* `%f` → decimal (fixed-point)
* `%e` / `%E` → scientific notation
* `%g` / `%G` → compact

---

## **4. Special Features**

* **Width & Precision Control**

```go
fmt.Printf("%6.2f", 9.1) // "  9.10" → width=6, precision=2
```

* **Left / Right alignment**

```go
fmt.Printf("%-6s", "Hi") // "Hi    "
```

* **Zero Padding**

```go
fmt.Printf("%04d", 7) // 0007
```

---

## **5. Why `fmt` is Important in Go**

* We control **how** values are displayed.
* We can **format output** exactly as needed (important for logs, reports, etc.).
* We can **scan** and parse data from various sources.
* It works uniformly across types due to **Go’s interface system** (`interface{}`).

---

Let’s make our **Go `fmt` Master Sheet** with **all differences** clearly laid out so we can grab it any time without thinking.

---

# 📝 **Go `fmt` Master Sheet**

*(Printing, Formatting, and Scanning — with Differences)*

---

## **1. PRINTING FAMILY**

These **output directly** (usually to `stdout` unless we use `F*` variants).

| Function                        | Outputs to                       | Newline? | Formatting?                      | Returns            |
| ------------------------------- | -------------------------------- | -------- | -------------------------------- | ------------------ |
| `fmt.Print(a ...interface{})`   | Stdout                           | ❌        | Default format                   | (n int, err error) |
| `fmt.Println(a ...interface{})` | Stdout                           | ✅        | Default format (space-separated) | (n int, err error) |
| `fmt.Printf(format, a...)`      | Stdout                           | ❌        | ✅ Format verbs                   | (n int, err error) |
| `fmt.Fprint(w, a...)`           | `io.Writer` (file, buffer, etc.) | ❌        | Default format                   | (n int, err error) |
| `fmt.Fprintln(w, a...)`         | `io.Writer`                      | ✅        | Default format                   | (n int, err error) |
| `fmt.Fprintf(w, format, a...)`  | `io.Writer`                      | ❌        | ✅ Format verbs                   | (n int, err error) |

**Example:**

```go
fmt.Print("Hi")            // Hi
fmt.Println("Hi")          // Hi\n
fmt.Printf("%s %d", "Hi", 5) // Hi 5
```

---

## **2. STRING-FORMATTING FAMILY**

These **do not print** — they return a string.

| Function                    | Newline? | Formatting?    | Returns |
| --------------------------- | -------- | -------------- | ------- |
| `fmt.Sprint(a...)`          | ❌        | Default        | string  |
| `fmt.Sprintln(a...)`        | ✅        | Default        | string  |
| `fmt.Sprintf(format, a...)` | ❌        | ✅ Format verbs | string  |

**Example:**

```go
msg := fmt.Sprintf("User: %s, Age: %d", "Bruce", 29)
fmt.Println(msg) // User: Bruce, Age: 29
```

---

## **3. SCANNING FAMILY**

These **read** values into variables.

| Function                         | Reads from | Stops at      | Formatting? | Returns            |
| -------------------------------- | ---------- | ------------- | ----------- | ------------------ |
| `fmt.Scan(&a...)`                | stdin      | Space/newline | ❌           | (n int, err error) |
| `fmt.Scanln(&a...)`              | stdin      | Newline       | ❌           | (n int, err error) |
| `fmt.Scanf(format, &a...)`       | stdin      | Format rules  | ✅           | (n int, err error) |
| `fmt.Fscan(r, &a...)`            | io.Reader  | Space/newline | ❌           | (n int, err error) |
| `fmt.Fscanf(r, format, &a...)`   | io.Reader  | Format rules  | ✅           | (n int, err error) |
| `fmt.Sscan(str, &a...)`          | string     | Space/newline | ❌           | (n int, err error) |
| `fmt.Sscanf(str, format, &a...)` | string     | Format rules  | ✅           | (n int, err error) |

**Example:**

```go
var name string
var age int
fmt.Scan(&name, &age) // Input: Bruce 29
fmt.Println(name, age) // Bruce 29
```

---

## **4. FORMAT VERBS CHEAT SHEET**

### **General**

| Verb  | Meaning                     |
| ----- | --------------------------- |
| `%v`  | Default format              |
| `%+v` | Adds field names in structs |
| `%#v` | Go syntax representation    |
| `%T`  | Type of the value           |
| `%%`  | Literal `%`                 |

### **Boolean**

| Verb | Meaning      |
| ---- | ------------ |
| `%t` | true / false |

### **Integer**

| Verb  | Meaning              |
| ----- | -------------------- |
| `%d`  | Decimal              |
| `%+d` | Decimal with sign    |
| `%b`  | Binary               |
| `%o`  | Octal                |
| `%O`  | Octal with 0o prefix |
| `%x`  | Hex (lower)          |
| `%X`  | Hex (upper)          |
| `%#x` | Hex with 0x prefix   |
| `%c`  | Unicode character    |
| `%q`  | Quoted character     |

### **Float**

| Verb | Meaning                 |
| ---- | ----------------------- |
| `%f` | Decimal point           |
| `%e` | Scientific notation (e) |
| `%E` | Scientific notation (E) |
| `%g` | Compact representation  |
| `%G` | Compact (E form)        |

### **String & Slice of bytes**

| Verb | Meaning       |
| ---- | ------------- |
| `%s` | String        |
| `%q` | Quoted string |
| `%x` | Hex lower     |
| `%X` | Hex upper     |

---

## **5. WIDTH, PRECISION, ALIGNMENT**

We can control **width** (total space) and **precision** (digits after decimal).

```go
fmt.Printf("%6d", 45)      // "    45" (width=6)
fmt.Printf("%-6d", 45)     // "45    " (left-aligned)
fmt.Printf("%06d", 45)     // "000045" (zero-padded)
fmt.Printf("%.2f", 3.14159) // "3.14"  (precision=2)
fmt.Printf("%6.2f", 3.14)   // "  3.14" (width=6, prec=2)
```

---

## **6. Quick Difference Summary**

* **`Print*`** → writes to output
* **`Sprint*`** → returns string
* **`Fprint*`** → writes to `io.Writer`
* **`Scan*`** → reads into variables
* **`Sscan*`** → reads from a string
* **`Fscan*`** → reads from an `io.Reader`
* `*ln` → adds newline
* `*f` → supports format verbs

---

## **1. What is a Struct in Go?**

In Go, a **struct** (short for *structure*) is a composite data type that groups **zero or more fields** together under one name.

* Each field has a **name** and a **type**.
* Structs let us model **real-world entities** with multiple properties.
* They’re similar to classes in OOP languages, but **Go structs have no methods by default** (methods are attached separately).

---

## **2. Declaring a Struct**

```go
type Person struct {
    name string
    age  int
    city string
}
```

Here:

* `Person` is the struct type.
* `name`, `age`, and `city` are fields.

---

## **3. Creating Struct Values**

### **3.1 Using Field Names**

```go
p1 := Person{name: "Bruce Wayne", age: 29, city: "Gotham"}
```

✔ Safer — order doesn’t matter.
❌ Cannot skip fields (unless they have zero values).

---

### **3.2 Without Field Names**

```go
p2 := Person{"Clark Kent", 35, "Metropolis"}
```

✔ Shorter, but **field order must match struct definition**.
❌ Easy to break if struct definition changes.

---

### **3.3 Zero Value Struct**

```go
var p3 Person
fmt.Println(p3) // { 0 }
```

* Strings default to `""`, ints to `0`.

---

### **3.4 Using `new()`**

```go
p4 := new(Person)
p4.name = "Diana Prince"
p4.age = 28
```

* `new()` returns a **pointer** to the struct.
* Fields can be accessed via `p4.field` (Go automatically dereferences).

---

### **3.5 Anonymous Struct**

```go
p5 := struct {
    name string
    age  int
}{"Barry Allen", 27}
```

* Useful for quick, temporary data structures.

---

## **4. Accessing and Modifying Fields**

```go
fmt.Println(p1.name) // Bruce Wayne
p1.age = 30
```

* Dot `.` notation works for both value and pointer structs.

---

## **5. Structs with Pointers**

Struct fields themselves can hold **pointers**:

```go
type Address struct {
    city  string
    pincode *int
}
```

* Allows optional data without nil-checking primitives.

---

## **6. Nested Structs**

```go
type Employee struct {
    name string
    address Address
}
```

* Access via `emp.address.city`.

---

## **7. Promoted Fields (Embedding)**

Go supports **struct embedding** (composition over inheritance):

```go
type Contact struct {
    email string
    phone string
}

type Employee struct {
    name string
    Contact
}

e := Employee{name: "Alfred", Contact: Contact{"alfred@wayne.com", "123"}}
fmt.Println(e.email) // directly accessible
```

* Fields from embedded structs are **promoted**.

---

## **8. Structs and Methods**

We can attach methods to structs:

```go
func (p Person) greet() {
    fmt.Printf("Hello, my name is %s\n", p.name)
}
```

* `(p Person)` is the **receiver**.
* Can be value receiver or pointer receiver.

---

## **9. Struct Comparisons**

* Structs are comparable **if all fields are comparable**.

```go
p1 := Person{"Bruce", 30, "Gotham"}
p2 := Person{"Bruce", 30, "Gotham"}
fmt.Println(p1 == p2) // true
```

* **Slices, maps, and functions** in struct fields make it **non-comparable**.

---

## **10. Memory Layout**

* Struct fields are stored **contiguously** in memory.
* Go may insert **padding** for alignment, so **field order can affect memory size**.

---

## **✅ Do’s**

* ✅ Use **field names** when initializing — improves readability and reduces errors.
* ✅ Keep field names lowercase if we want them **private** to the package.
* ✅ Group related fields logically.
* ✅ Use **embedding** for reusability, not inheritance.
* ✅ Use **pointers** for large structs to avoid copying.

---

## **❌ Don’ts**

* ❌ Don’t rely on zero values without clear intent.
* ❌ Don’t export fields unnecessarily (respect package boundaries).
* ❌ Don’t embed too many structs — can lead to naming conflicts.
* ❌ Don’t use unnamed initialization unless we control the struct definition.
* ❌ Don’t assume struct size — padding can change it.

---

## **1. What Are Anonymous Struct Fields?**

In Go, **anonymous fields** are struct fields declared **without an explicit name**, where the field’s **type** itself acts as the name.

* These are also called **embedded fields**.
* They’re often used for **struct embedding** (composition) but can also exist standalone.

Example:

```go
type Person struct {
    string // anonymous field of type string
    int    // anonymous field of type int
}
```

Here:

* `string` and `int` are **field types** and **names** at the same time.
* We can access them directly by their type name:

```go
p := Person{"Bruce Wayne", 30}
fmt.Println(p.string) // Bruce Wayne
fmt.Println(p.int)    // 30
```

---

## **2. Why Do We Use Anonymous Fields?**

1. **Code Composition** – We can embed one struct into another so the fields of the embedded struct are accessible directly.
2. **Field Promotion** – When we embed a struct, all its exported fields and methods are “promoted” to the outer struct.
3. **Reuse Without Inheritance** – Go avoids classical inheritance, so anonymous fields give us a way to compose behavior.

---

## **3. Anonymous Struct Fields with Struct Embedding**

Instead of naming the field, we just write the type:

```go
type Contact struct {
    email string
    phone string
}

type Employee struct {
    name string
    Contact // anonymous field (embedded struct)
}
```

Now:

```go
e := Employee{
    name:    "Alfred Pennyworth",
    Contact: Contact{"alfred@wayne.com", "123456"},
}

fmt.Println(e.email)  // direct access — no e.Contact.email needed
fmt.Println(e.phone)  // promoted field
```

---

## **4. Method Promotion with Anonymous Fields**

If the embedded type has methods, they are promoted too:

```go
func (c Contact) showContact() {
    fmt.Println("Email:", c.email, "Phone:", c.phone)
}

e.showContact() // works directly
```

* We don’t need `e.Contact.showContact()` — the method is **promoted**.

---

## **5. Rules of Anonymous Fields**

* **Only one anonymous field of a given type** is allowed (or compilation fails due to ambiguity).
* If multiple embedded structs have the **same field name**, we must access it explicitly via the embedded struct’s name to avoid conflicts.
* The embedded type must be a **named type** (struct, interface, etc.) — we can’t embed primitives *unless* we want them to behave as fields with their type as name.

---

## **6. Anonymous Fields vs Named Fields**

| **Aspect**         | **Named Field**           | **Anonymous Field**                     |
| ------------------ | ------------------------- | --------------------------------------- |
| Access             | `struct.fieldName`        | `struct.TypeName` or promoted name      |
| Reusability        | Mostly standalone         | Encourages composition                  |
| Field Promotion    | ❌ No promotion            | ✅ Fields/methods promoted               |
| Readability        | More explicit             | Can be implicit but sometimes confusing |
| Multiple same type | Allowed (different names) | ❌ Not allowed (same type duplicates)    |

---

## ✅ **Do’s**

* ✅ Use anonymous fields for **struct composition** where we want to reuse and promote fields/methods.
* ✅ Keep embedded struct’s exported fields/methods public only if they need to be accessible.
* ✅ Use named fields for primitives unless embedding is intentional.

## ❌ **Don’ts**

* ❌ Don’t overuse anonymous fields for unrelated types — it makes the struct harder to understand.
* ❌ Don’t embed multiple types that have overlapping field names without expecting conflicts.
* ❌ Don’t confuse anonymous fields with inheritance — they are **composition**, not parent-child relationships.

---


Now let’s go step-by-step and deeply understand **interfaces in Go**, including how they work under the hood, their rules, and common do’s/don’ts.

---

## **1. What is an Interface in Go?**

In Go, an **interface** is a **type** that defines a set of **method signatures**.

* If a type has methods matching all of an interface’s method signatures, it **implicitly implements** that interface — no explicit declaration like `implements` is required (unlike Java, C#, etc.).
* Interfaces allow us to write **polymorphic** and **flexible** code without inheritance.

Example:

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

Here:

* Any type that has **both** `Area()` and `Perimeter()` methods (with matching signatures) is considered a `Shape`.

---

## **2. Implicit Implementation (No Keywords)**

Example:

```go
type Rectangle struct {
    width, height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.width + r.height)
}
```

We never say “`Rectangle implements Shape`”, but:

```go
var s Shape
s = Rectangle{10, 5}

fmt.Println(s.Area())      // 50
fmt.Println(s.Perimeter()) // 30
```

✅ Works because `Rectangle` has **all methods** required by `Shape`.

---

## **3. Empty Interface (`interface{}`)**

`interface{}` means **“zero methods”** — which **every** type implements.

* Useful for functions that take **any type** (like generics in older Go versions).

```go
func describe(i interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", i, i)
}

describe(42)
describe("Batman")
describe(true)
```

* In Go 1.18+, **generics** often replace the need for `interface{}` in many cases.

---

## **4. Type Assertions & Type Switches**

When we have an interface variable, we might want the **underlying concrete type**.

**Type assertion:**

```go
var i interface{} = "Bruce Wayne"
str, ok := i.(string)
if ok {
    fmt.Println("String:", str)
}
```

**Type switch:**

```go
switch v := i.(type) {
case string:
    fmt.Println("It's a string:", v)
case int:
    fmt.Println("It's an int:", v)
default:
    fmt.Println("Unknown type")
}
```

---

## **5. Interfaces with Pointers vs Values**

* If a method has a **pointer receiver**, only a pointer to the type implements the interface.
* If a method has a **value receiver**, both a value and a pointer to the type implement the interface.

Example:

```go
type Writer interface {
    Write() string
}

type Person struct {
    name string
}

func (p Person) Write() string { // value receiver
    return "Name: " + p.name
}

var w Writer = Person{"Clark"}  // works
var w2 Writer = &Person{"Kent"} // also works
```

---

## **6. Interfaces Inside Structs (Composition)**

We can use interfaces as struct fields:

```go
type Printer interface {
    Print()
}

type Document struct {
    p Printer
}

func (d Document) Run() {
    d.p.Print()
}
```

This allows **dependency injection** — we can pass any type that implements `Printer`.

---

## **7. Nil Interfaces vs Nil Concrete Types**

A nil interface means both:

* The **type** is `nil`
* The **value** is `nil`

Example:

```go
var i interface{}
fmt.Println(i == nil) // true

var p *Person = nil
var i2 interface{} = p
fmt.Println(i2 == nil) // false (type is *Person, value is nil)
```

⚠ This often confuses beginners — the interface is not nil if it has a type, even if the value is nil.

---

## **8. Interface Internals**

Under the hood, an interface variable stores:

1. **Type information** (concrete type)
2. **Value** (actual data)

Think of it as a **(type, value)** pair.

---

## ✅ Do’s

* ✅ Use small, focused interfaces (e.g., `io.Reader`, `io.Writer`) instead of large “fat” ones.
* ✅ Name interfaces with `-er` suffix when possible (`Reader`, `Writer`, `Closer`).
* ✅ Use interfaces for **behavior**, not just data grouping.

## ❌ Don’ts

* ❌ Don’t create an interface just to “mock” something if it’s only used once — mock concrete types instead.
* ❌ Don’t force a type to implement an interface it doesn’t logically need.
* ❌ Don’t pass interfaces where concrete types work better — interfaces are for abstraction, not always replacement.

---

Alright — here’s a **clear side-by-side comparison** of **interfaces vs structs** in Go:

---

## **Interfaces vs Structs in Go**

| Feature              | **Interface**                                                                                     | **Struct**                                                                                              |
| -------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| **Definition**       | A type that defines a set of **method signatures** (behavior contract).                           | A composite type that groups **fields (data)** together.                                                |
| **Purpose**          | To define **behavior** that multiple types can implement.                                         | To store and organize **data**.                                                                         |
| **Contains**         | Only **method definitions** (no fields, no implementation).                                       | Fields (variables) and methods (with receivers).                                                        |
| **Implementation**   | Implicit — a type implements an interface automatically if it has all the required methods.       | Explicit — we define fields and methods directly in the struct.                                         |
| **Instantiation**    | Cannot be instantiated directly (only holds values of types that implement it).                   | Can be instantiated with `struct{} {}` literal or `new`.                                                |
| **Memory**           | Stores a **type** and a **value** (runtime pairing).                                              | Stores the actual **data** in memory.                                                                   |
| **Polymorphism**     | Supports polymorphism (we can pass any type implementing the interface).                          | Does not support polymorphism by itself — but can be used as a concrete implementation of an interface. |
| **Nil Behavior**     | A nil interface means both type and value are nil; can be tricky if type is set but value is nil. | A nil struct pointer is straightforward — just nil.                                                     |
| **Example Use Case** | Defining `Shape` behavior for multiple shapes (`Circle`, `Rectangle`).                            | Defining `User` data with `name`, `email`, `age` fields.                                                |

---

### **Example**

**Interface:**

```go
type Shape interface {
    Area() float64
}
```

**Struct implementing it:**

```go
type Circle struct {
    radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}
```

---

## **Summary**

* **Interfaces** = Behavior (**what we can do**).
* **Structs** = Data + optional methods (**what we have and how we handle it**).
* Together, they give us **composition-based design** rather than inheritance.

---

## **What Are Embedded Structs in Go?**

In Go, we can place a struct **inside** another struct **without explicitly giving it a field name** — this is called **embedding**.

When we embed a struct:

* The inner struct’s fields and methods become **promoted** to the outer struct.
* We can access them **directly** from the outer struct as if they were its own.
* This is Go’s way of achieving a form of **composition** (instead of classical inheritance like in OOP).

---

### **Basic Syntax**

```go
package main

import "fmt"

// Inner struct
type Person struct {
    Name string
    Age  int
}

// Outer struct embedding Person
type Employee struct {
    Person   // Embedded struct (no field name)
    Company  string
}

func main() {
    // Creating an Employee instance
    emp := Employee{
        Person: Person{
            Name: "Bruce Wayne",
            Age:  35,
        },
        Company: "Wayne Enterprises",
    }

    // Accessing embedded struct fields directly
    fmt.Println(emp.Name)    // Bruce Wayne
    fmt.Println(emp.Age)     // 35
    fmt.Println(emp.Company) // Wayne Enterprises
}
```

---

### **How It Works**

* By writing `Person` inside `Employee` without a name, we make it an **embedded field**.
* We don’t need to write `emp.Person.Name` — we can directly write `emp.Name` because Go **promotes** the fields.
* This also works with **methods**.

---

### **Example with Methods**

```go
package main

import "fmt"

type Engine struct {
    HorsePower int
}

func (e Engine) Start() {
    fmt.Println("Engine started with", e.HorsePower, "HP")
}

type Car struct {
    Engine // Embedded struct
    Brand  string
}

func main() {
    c := Car{
        Engine: Engine{HorsePower: 250},
        Brand:  "Tesla",
    }

    // Directly calling method of embedded struct
    c.Start() // Engine started with 250 HP
}
```

---

## **Key Points to Remember**

1. **Field Promotion**

   * Embedded struct’s fields/methods become part of the outer struct’s namespace.
   * If there’s no naming conflict, we can access them directly.

2. **Field Name Conflicts**

   * If both the outer struct and embedded struct have the same field name, **we must use the full path**.

   ```go
   type A struct { Name string }
   type B struct {
       Name string
       A
   }
   b := B{Name: "Outer", A: A{Name: "Inner"}}
   fmt.Println(b.Name)    // Outer
   fmt.Println(b.A.Name)  // Inner
   ```

3. **Multiple Embeddings**

   * We can embed multiple structs, and Go will promote all unique fields/methods.

4. **Anonymous Structs Can Be Embedded Too**

   ```go
   type Outer struct {
       struct {
           X int
           Y int
       }
   }
   ```

---

## ✅ **Do’s**

* Use embedding for **composition** instead of inheritance.
* Use it to **reuse methods and fields** between structs.
* Keep field names unique to avoid ambiguity.

## ❌ **Don’ts**

* Don’t overuse embedding — it can make code confusing.
* Avoid name collisions between outer and inner structs unless necessary.
* Don’t treat it as classical OOP inheritance — Go’s philosophy is **composition over inheritance**.

---

Let’s dig deep into **errors and error handling in Go** 

---

## **1. What Is an Error in Go?**

In Go, an **error** is simply **a value** that represents a problem.

* Go’s standard library defines the `error` type as an **interface**:

```go
type error interface {
    Error() string
}
```

* This means an error is **any type** that has an `Error() string` method.
* Functions in Go often return an **`error`** as the **last return value**.
* If nothing went wrong, the function returns `nil` for the error.

---

## **2. Creating and Returning Errors**

We can create errors using the built-in **`errors`** package or **`fmt.Errorf`**.

### Example:

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // Error: cannot divide by zero
        return
    }
    fmt.Println("Result:", result)
}
```

---

## **3. Using `fmt.Errorf` for Formatted Errors**

```go
func login(user, pass string) error {
    if user != "admin" {
        return fmt.Errorf("user %s not found", user)
    }
    return nil
}
```

* `fmt.Errorf` allows string formatting inside error messages.
* It’s useful for including **dynamic data** in errors.

---

## **4. Best Practices for Error Handling**

In Go, we **check errors explicitly** rather than using exceptions.

```go
value, err := doSomething()
if err != nil {
    // handle error
    log.Println("Something went wrong:", err)
    return
}
// continue if no error
```

✅ **Why This Is Good in Go**

* Keeps flow **explicit**.
* We always know **exactly where** errors are handled.

---

## **5. Custom Error Types**

We can define our own error types to add **extra context**.

```go
type MyError struct {
    Code    int
    Message string
}

func (e MyError) Error() string {
    return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func riskyOperation() error {
    return MyError{Code: 404, Message: "Resource not found"}
}

func main() {
    err := riskyOperation()
    if err != nil {
        fmt.Println(err) // Code 404: Resource not found
    }
}
```

---

## **6. Wrapping and Unwrapping Errors** (Go 1.13+)

Go 1.13 introduced **error wrapping** with `%w` in `fmt.Errorf` and the `errors` package’s `Is` and `As` functions.

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func getData() error {
    return fmt.Errorf("database error: %w", ErrNotFound)
}

func main() {
    err := getData()

    // Check if the underlying error is ErrNotFound
    if errors.Is(err, ErrNotFound) {
        fmt.Println("The data was not found!")
    }
}
```

**Key Functions:**

* `errors.Is(err, target)` → checks if `err` wraps `target`.
* `errors.As(err, &target)` → checks if `err` can be cast to a specific type.

---

## **7. Panic vs Error**

* **Error** → Expected, recoverable problem (e.g., file not found).
* **Panic** → Unexpected, unrecoverable problem (e.g., nil pointer dereference).
* We should **avoid panic** for normal error handling — use it only for truly exceptional situations.

---

## ✅ **Do’s**

* Always check errors explicitly.
* Create meaningful error messages.
* Use `%w` to wrap errors for context.
* Use `errors.Is` and `errors.As` for type checking.

## ❌ **Don’ts**

* Don’t ignore errors (`_ = something()` is a red flag).
* Don’t use panic for regular errors.
* Don’t make vague error messages like `"failed"` — include context.

---

Now, let’s unpack **why `Error()` must be exactly that** and how Go’s interface satisfaction rules work.

---

## 1️⃣ The `error` interface in Go

In Go, **`error`** is just a **built-in interface type** defined in the standard library (`builtin.go`):

```go
type error interface {
    Error() string
}
```

That’s literally all it is — one method named `Error` that returns a `string`.

For a type to be considered an **`error`**, it must implement that method **exactly**:

* **Name**: `Error`
* **Return type**: `string`
* **Receiver**: Can be value or pointer receiver — depends on how we intend to use it.

---

## 2️⃣ How Go matches methods to interfaces

Go uses **structural typing** for interfaces:

> If a type has all the methods that an interface requires, in name, parameters, and return types, then it automatically implements that interface — no `implements` keyword needed.

But this matching is **exact** — the compiler checks:

* The **method name** must match **case-sensitively** (`Error` ≠ `Err` ≠ `error`).
* The **parameters** must match exactly (including number, order, and types).
* The **return type** must match exactly.

So if we write:

```go
func (mE *myErr) Err() string { ... }
```

That’s not the same as:

```go
func (mE *myErr) Error() string { ... }
```

and the compiler refuses to assign it to `error`.

---

## 3️⃣ Why Go enforces this strictness

Go enforces **explicit method contracts** so:

* We avoid accidental interface implementation with mismatched semantics.
* We keep errors consistent — any `error` can be converted to a string via `err.Error()`.
* The compiler can guarantee type safety **before runtime**.

---

## 4️⃣ Pointer vs Value Receivers with `Error()`

If we write:

```go
func (mE myErr) Error() string { ... }
```

Then both `myErr` and `*myErr` satisfy the `error` interface.

If we write:

```go
func (mE *myErr) Error() string { ... }
```

Then only `*myErr` satisfies `error` — `myErr` alone will not.
Most custom error types use **pointer receivers** so that they don’t copy data unnecessarily.

---

## 5️⃣ Example

```go
package main

import "fmt"

type myErr struct {
	msg string
}

func (e *myErr) Error() string {
	return fmt.Sprintf("🔴Error: %s", e.msg)
}

func doSomething() error {
	return &myErr{"something went wrong"}
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err) // Automatically calls Error()
	}
}
```

**Output:**

```
🔴Error: something went wrong
```

Here, `fmt.Println(err)` automatically calls `err.Error()` because `err` is of type `error`.

---

## 6️⃣ DO’s & DON’Ts for custom errors

**✅ Do:**

* Always implement `Error() string` exactly.
* Use pointer receivers for large structs to avoid copying.
* Return your custom type as `error` from functions.

**❌ Don’t:**

* Rename the method (`Err()`, `ErrorMsg()`, etc. won’t work for `error`).
* Change return type — must be `string`.
* Add parameters to `Error()` — the signature must be empty `()`.

---

`fmt.Errorf()` in Go is one of our best friends for **creating formatted errors**.

---

## **1️⃣ What is `fmt.Errorf()`?**

* It’s a function from Go’s `fmt` package.
* It **creates and returns** a new `error` value.
* It works just like `fmt.Sprintf()`, but instead of returning a `string`, it returns an `error`.

Signature:

```go
func Errorf(format string, a ...any) error
```

---

## **2️⃣ How it works**

Example:

```go
err := fmt.Errorf("file %s not found", "data.txt")
fmt.Println(err)
```

**Output:**

```
file data.txt not found
```

Here:

* `"file %s not found"` → format string
* `"data.txt"` → value to replace `%s`
* `fmt.Errorf()` internally formats the string and wraps it into an `error` object.

---

## **3️⃣ Why not just `errors.New()`?**

`errors.New()` only creates a **static error message**:

```go
err := errors.New("something went wrong") // No formatting
```

If we need **dynamic values** inside the message, `fmt.Errorf()` is better.

---

## **4️⃣ Adding context to errors**

One of the most important uses of `fmt.Errorf()` is to wrap an existing error for more context:

```go
func readFile() error {
    return errors.New("disk read failure")
}

func processFile() error {
    err := readFile()
    if err != nil {
        return fmt.Errorf("processFile failed: %w", err)
    }
    return nil
}

func main() {
    if err := processFile(); err != nil {
        fmt.Println(err)
    }
}
```

**Output:**

```
processFile failed: disk read failure
```

Here, `%w` is special — it **wraps** the original error so we can later check/unpack it using `errors.Is()` or `errors.As()`.

---

## **5️⃣ Special verbs in `fmt.Errorf()`**

We can use **all formatting verbs** (`%s`, `%d`, `%v`, `%w`, etc.) here, but `%w` is unique because:

* It wraps an **existing error** for chaining.
* Only **one `%w`** is allowed in the format string.

Example:

```go
err1 := errors.New("network timeout")
err2 := fmt.Errorf("login failed: %w", err1)
fmt.Println(err2) // login failed: network timeout
```

---

✅ **Key Takeaways:**

* `fmt.Errorf()` = `fmt.Sprintf()` + wrap into `error`.
* Use it for **dynamic** error messages.
* Use `%w` when we want to wrap another error and preserve its type for later checks.

---

Alright — let’s deep dive into **custom errors in Go** so we not only know *how* to create them, but also *why* and *when* to use them.
We’ll go from **basics → custom types → wrapping → best practices**.

---

## **1. Quick Refresher: Errors in Go**

In Go, an **error** is just a value of the built-in `error` type.

```go
err := errors.New("something went wrong")
if err != nil {
    fmt.Println(err)
}
```

* `error` is an **interface**:

```go
type error interface {
    Error() string
}
```

It just requires an **`Error()` method that returns a string**.

---

## **2. Custom Error Basics**

Since `error` is an interface, we can make **our own type** that satisfies it.

### Example:

```go
package main

import (
    "fmt"
)

type MyError struct {
    Code    int
    Message string
}

// Implement the Error() method
func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething() error {
    return &MyError{Code: 404, Message: "Resource not found"}
}

func main() {
    err := doSomething()
    if err != nil {
        fmt.Println(err) // Error 404: Resource not found
    }
}
```

---

### **Key Points Here**

* `MyError` is a struct holding **extra info** (`Code`, `Message`).
* We made it satisfy `error` by adding `Error() string`.
* We returned a pointer (`*MyError`) because that’s common when working with struct receivers in Go.

---

## **3. Using `errors.New()` vs `fmt.Errorf()`**

* `errors.New("msg")` → Quick way to make a basic error.
* `fmt.Errorf("...")` → Lets us format strings into errors.

Example:

```go
err := fmt.Errorf("user %s not found", "Skyy")
fmt.Println(err) // user Skyy not found
```

---

## **4. Wrapping Errors (Go 1.13+)**

Go allows **wrapping** an error with `%w` in `fmt.Errorf()`.

* This keeps the original error while adding context.
* We can later **unwrap** it to get the cause.

Example:

```go
import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func findUser() error {
    return fmt.Errorf("in findUser: %w", ErrNotFound)
}

func main() {
    err := findUser()

    // Check if it's our specific error
    if errors.Is(err, ErrNotFound) {
        fmt.Println("User not found!")
    }
}
```

---

## **5. Adding More Context with Custom Errors**

We can mix **custom types** with wrapping.

Example:

```go
type DatabaseError struct {
    Op  string
    Err error
}

func (e *DatabaseError) Error() string {
    return fmt.Sprintf("database %s error: %v", e.Op, e.Err)
}

func queryDB() error {
    return &DatabaseError{
        Op:  "SELECT",
        Err: errors.New("connection refused"),
    }
}

func main() {
    err := queryDB()
    fmt.Println(err)
}
```

Output:

```
database SELECT error: connection refused
```

---

## **6. Why Use Custom Errors?**

* **Add metadata** (e.g., error code, operation name).
* **Check specific types** using `errors.As()`:

```go
var dbErr *DatabaseError
if errors.As(err, &dbErr) {
    fmt.Println("This is a DatabaseError with op:", dbErr.Op)
}
```

* **Better debugging** with more context.
* **Clear separation** between different error types in large projects.

---

## **7. Best Practices**

1. **Keep them simple** — only add fields you actually need.
2. **Use sentinel errors** (`var ErrXYZ = errors.New("...")`) for known conditions.
3. **Wrap errors** when adding context, don’t replace them.
4. **Use `errors.Is` & `errors.As`** for checking, not string matching.
5. Avoid creating **too many** custom error types unless necessary.

---

## **8. TL;DR Custom Error Flow**

1. Define a struct for extra fields.
2. Implement `Error() string` method.
3. Return it where needed.
4. Use `errors.Is` / `errors.As` to check it.
5. Wrap with `%w` if you need to add context.

---

Alright — this is one of those Go error-handling details that trips people up at first, so let’s make **`errors.Is`** and **`errors.As`** crystal clear.

---

## **1. `errors.Is`**

**Purpose:**
Checks if an error **is or wraps a specific target error** (compares by value, like equality).

**Think:**

> “Is this error equal to, or caused by, this known error?”

**Example:**

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func findUser() error {
    return fmt.Errorf("failed to find user: %w", ErrNotFound)
}

func main() {
    err := findUser()
    if errors.Is(err, ErrNotFound) {
        fmt.Println("User not found!") // ✅ True, because err wraps ErrNotFound
    }
}
```

**When to use:**

* We have **sentinel errors** (predeclared constants like `ErrNotFound`, `ErrUnauthorized`).
* We want to know if the chain of wrapped errors contains that specific one.

---

## **2. `errors.As`**

**Purpose:**
Checks if an error **is or wraps a specific type**, and if yes, assigns it to a variable.

**Think:**

> “Is this error of this custom type? If yes, give me access to it.”

**Example:**

```go
package main

import (
    "errors"
    "fmt"
)

type MyError struct {
    Code int
}

func (e *MyError) Error() string {
    return fmt.Sprintf("code: %d", e.Code)
}

func doSomething() error {
    return fmt.Errorf("something failed: %w", &MyError{Code: 404})
}

func main() {
    err := doSomething()

    var myErr *MyError
    if errors.As(err, &myErr) {
        fmt.Println("Got MyError with code:", myErr.Code) // ✅ 404
    }
}
```

**When to use:**

* We have **custom error types** with extra fields.
* We need to **extract and use** that extra info.

---

## **3. Quick Visual**

| Feature     | `errors.Is`               | `errors.As`                             |
| ----------- | ------------------------- | --------------------------------------- |
| Compares by | **Value**                 | **Type**                                |
| Target      | Predefined error variable | Pointer to a variable of the error type |
| Returns     | `bool`                    | `bool` (and assigns)                    |
| Use Case    | “Does this wrap ErrX?”    | “Is this a MyError? Give it to me”      |

---

## **4. Both Together Example**

```go
if errors.Is(err, ErrNotFound) {
    fmt.Println("We hit the not found case")
}

var myErr *MyError
if errors.As(err, &myErr) {
    fmt.Println("Error code:", myErr.Code)
}
```

Here:

* `errors.Is` → checks for a **specific instance**.
* `errors.As` → pulls out a **specific type**.

---

**all the important string functions in Go**, so we can have a full reference whenever we need it.

We’ll focus mainly on the **`strings` package**, plus a few `fmt`, `strconv`, and `bytes` helpers that are commonly used with strings.

---

## **📚 1. Importing the package**

```go
import "strings"
```

All functions are **case-sensitive** and operate on UTF-8 encoded strings.

---

## **🔍 2. Search & Check Functions**

| Function                         | Description                                                | Example                                        |
| -------------------------------- | ---------------------------------------------------------- | ---------------------------------------------- |
| `strings.Contains(s, substr)`    | Checks if `substr` is inside `s`. Returns `bool`.          | `strings.Contains("Hello", "ell") // true`     |
| `strings.ContainsAny(s, chars)`  | True if any Unicode code point from `chars` exists in `s`. | `strings.ContainsAny("team", "aeiou") // true` |
| `strings.ContainsRune(s, r)`     | True if rune `r` exists in `s`.                            | `strings.ContainsRune("golang", 'g') // true`  |
| `strings.HasPrefix(s, prefix)`   | True if `s` starts with `prefix`.                          | `strings.HasPrefix("GoLang", "Go") // true`    |
| `strings.HasSuffix(s, suffix)`   | True if `s` ends with `suffix`.                            | `strings.HasSuffix("GoLang", "Lang") // true`  |
| `strings.Count(s, substr)`       | Counts non-overlapping occurrences of `substr` in `s`.     | `strings.Count("banana", "na") // 2`           |
| `strings.Index(s, substr)`       | Returns first index of `substr`, or `-1` if not found.     | `strings.Index("banana", "na") // 2`           |
| `strings.LastIndex(s, substr)`   | Last occurrence of `substr`.                               | `strings.LastIndex("banana", "na") // 4`       |
| `strings.IndexAny(s, chars)`     | Index of first occurrence of any char from `chars`.        | `strings.IndexAny("team", "aeiou") // 1`       |
| `strings.LastIndexAny(s, chars)` | Last occurrence of any char from `chars`.                  | `strings.LastIndexAny("team", "aeiou") // 3`   |
| `strings.IndexRune(s, r)`        | Index of first rune `r`.                                   | `strings.IndexRune("golang", 'g') // 0`        |

---

## **✂️ 3. Modify / Transform Strings**

| Function                                           | Description                                | Example                                          |
| -------------------------------------------------- | ------------------------------------------ | ------------------------------------------------ |
| `strings.ToUpper(s)`                               | Converts to uppercase.                     | `"go" → "GO"`                                    |
| `strings.ToLower(s)`                               | Converts to lowercase.                     | `"GO" → "go"`                                    |
| `strings.Title(s)` (deprecated, use `cases.Title`) | Title-case each word.                      | `"hello world" → "Hello World"`                  |
| `strings.ToTitle(s)`                               | Unicode title mapping.                     | `"hello" → "HELLO"` (differs in special locales) |
| `strings.Repeat(s, count)`                         | Repeat string `count` times.               | `Repeat("ab", 3) → "ababab"`                     |
| `strings.Replace(s, old, new, n)`                  | Replace first `n` occurrences.             | `Replace("foo", "o", "0", -1) → "f00"`           |
| `strings.ReplaceAll(s, old, new)`                  | Replace all occurrences.                   | `"foo" → "f00"`                                  |
| `strings.Trim(s, cutset)`                          | Remove leading/trailing runes in `cutset`. | `Trim("!!hi!!", "!") → "hi"`                     |
| `strings.TrimSpace(s)`                             | Removes all leading/trailing whitespace.   | `"  hi  " → "hi"`                                |
| `strings.TrimPrefix(s, prefix)`                    | Removes prefix if present.                 | `"GoLang" → "Lang"`                              |
| `strings.TrimSuffix(s, suffix)`                    | Removes suffix if present.                 | `"GoLang" → "Go"`                                |

---

## **🪓 4. Splitting & Joining**

| Function                         | Description                     | Example                                                  |   |              |
| -------------------------------- | ------------------------------- | -------------------------------------------------------- | - | ------------ |
| `strings.Split(s, sep)`          | Splits into slice by separator. | `Split("a,b,c", ",") → ["a","b","c"]`                    |   |              |
| `strings.SplitN(s, sep, n)`      | At most `n` substrings.         | `SplitN("a,b,c", ",", 2) → ["a","b,c"]`                  |   |              |
| `strings.SplitAfter(s, sep)`     | Keep separator in result.       | `"a,b,c" → ["a,","b,","c"]`                              |   |              |
| `strings.SplitAfterN(s, sep, n)` | Keep separator, limit results.  |                                                          |   |              |
| `strings.Fields(s)`              | Split by any whitespace.        | `"a b\tc\n" → ["a","b","c"]`                             |   |              |
| `strings.FieldsFunc(s, f)`       | Custom split based on function. | \`FieldsFunc("a-b\_c", func(r rune) bool { return r=='-' |   | r=='\_' })\` |
| `strings.Join(slice, sep)`       | Join slice into single string.  | `Join([]string{"a","b"}, "-") → "a-b"`                   |   |              |

---

## **🧠 5. Mapping & Functional Helpers**

| Function                      | Description                       | Example                    |
| ----------------------------- | --------------------------------- | -------------------------- |
| `strings.Map(mappingFunc, s)` | Apply function to each rune.      | Convert all vowels to `*`. |
| `strings.Compare(a, b)`       | Lexicographic compare (-1, 0, 1). | `"a" vs "b"` → -1          |
| `strings.EqualFold(s1, s2)`   | Case-insensitive compare.         | `"Go" vs "gO"` → true      |

---

## **🔢 6. Conversion Helpers (Other Packages)**

| Package                     | Function        | Example           |
| --------------------------- | --------------- | ----------------- |
| `strconv.Itoa(int)`         | Int → string.   | `Itoa(42) → "42"` |
| `strconv.Atoi(str)`         | String → int.   | `Atoi("42") → 42` |
| `strconv.FormatFloat(...)`  | Float → string. |                   |
| `strconv.ParseBool("true")` | String → bool.  |                   |

---

## **📌 7. Notes & Gotchas**

* Strings in Go are **immutable** — every transformation returns a **new** string.
* For **performance** with many concatenations, use `strings.Builder` or `bytes.Buffer`.
* Be careful with Unicode — `len(s)` returns **bytes**, not runes. Use `utf8.RuneCountInString(s)` for character count.

---

## **1. What is Regex in Go?**

A **Regular Expression (Regex)** is a string pattern used to match, search, or manipulate text.
In Go, we use the **`regexp`** package from the standard library.

**Import:**

```go
import "regexp"
```

Regex in Go is **Perl-style** (similar to other languages), but Go’s API is stricter and **must be compiled before use**.

---

## **2. Creating Regex in Go**

We have **two main functions** for creating a Regex:

### **a. `regexp.Compile()`**

* Compiles a regex pattern and returns a `*Regexp` object.
* Can return an error if the pattern is invalid.

```go
re, err := regexp.Compile(`\d+`)
if err != nil {
	fmt.Println("Invalid regex:", err)
}
```

### **b. `regexp.MustCompile()`**

* Same as `Compile()` but panics if the pattern is invalid.
* Common when the pattern is constant and guaranteed valid.

```go
re := regexp.MustCompile(`\d+`) // no error handling
```

---

## **3. Common Regex Methods in Go**

Let’s use:

```go
re := regexp.MustCompile(`\d+`) // matches one or more digits
```

| Method                       | Purpose                                          | Example                                                                         | Output                        |
| ---------------------------- | ------------------------------------------------ | ------------------------------------------------------------------------------- | ----------------------------- |
| **`MatchString()`**          | Checks if pattern matches anywhere in the string | `re.MatchString("Age: 29")`                                                     | `true`                        |
| **`FindString()`**           | Returns **first** match                          | `re.FindString("Age: 29, Height: 180")`                                         | `"29"`                        |
| **`FindAllString()`**        | Returns **all** matches                          | `re.FindAllString("Age: 29, Height: 180", -1)`                                  | `["29", "180"]`               |
| **`FindStringIndex()`**      | Returns `[start, end]` index of first match      | `re.FindStringIndex("abc123xyz")`                                               | `[3, 6]`                      |
| **`FindAllStringIndex()`**   | Returns indexes of all matches                   | `re.FindAllStringIndex("abc123xyz456", -1)`                                     | `[[3, 6], [9, 12]]`           |
| **`ReplaceAllString()`**     | Replaces all matches with text                   | `re.ReplaceAllString("abc123xyz", "#")`                                         | `"abc#xyz"`                   |
| **`ReplaceAllStringFunc()`** | Replace all matches using a function             | `re.ReplaceAllStringFunc("abc123", func(s string) string { return "["+s+"]" })` | `"abc[123]"`                  |
| **`Split()`**                | Splits string by pattern                         | `re.Split("one1two2three3", -1)`                                                | `["one", "two", "three", ""]` |

---

## **4. Regex Syntax in Go**

Go follows **RE2 syntax** (safe, no catastrophic backtracking like some PCRE regexes).
Here’s a quick reference:

### **a. Character Classes**

| Pattern  | Meaning                         |
| -------- | ------------------------------- |
| `.`      | Any character except newline    |
| `[abc]`  | Any of a, b, or c               |
| `[^abc]` | Any char NOT a, b, or c         |
| `[a-z]`  | Range (lowercase letters)       |
| `\d`     | Digit (0-9)                     |
| `\D`     | Non-digit                       |
| `\w`     | Word char (letters, digits, \_) |
| `\W`     | Non-word char                   |
| `\s`     | Whitespace                      |
| `\S`     | Non-whitespace                  |

---

### **b. Quantifiers**

| Pattern | Meaning         |
| ------- | --------------- |
| `*`     | 0 or more       |
| `+`     | 1 or more       |
| `?`     | 0 or 1          |
| `{n}`   | Exactly n       |
| `{n,}`  | n or more       |
| `{n,m}` | Between n and m |

---

### **c. Anchors**

| Pattern | Meaning           |
| ------- | ----------------- |
| `^`     | Start of string   |
| `$`     | End of string     |
| `\b`    | Word boundary     |
| `\B`    | Non-word boundary |

---

### **d. Groups**

| Pattern         | Meaning               |
| --------------- | --------------------- |
| `(abc)`         | Capturing group       |
| `(?:abc)`       | Non-capturing group   |
| `(?P<name>abc)` | Named capturing group |

---

## **5. Example: Extracting Emails**

```go
package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Contact us at support@example.com or sales@example.org"
	re := regexp.MustCompile(`[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}`)
	emails := re.FindAllString(text, -1)
	fmt.Println(emails) // [support@example.com sales@example.org]
}
```

---

## **6. Tips & Gotchas**

✅ Use `MustCompile()` for fixed patterns.
✅ Always escape `\` as `\\` in Go string literals.
✅ Remember: Go regex doesn’t support lookbehind assertions.
✅ `(?i)` makes it case-insensitive:

```go
re := regexp.MustCompile(`(?i)hello`)
```

---

Alright — let’s dive deep into **`strings.Builder`** in Go, because it’s one of those tools that can make our string handling **much faster and more memory-efficient** if used correctly.

---

## **1. What is `strings.Builder`?**

`strings.Builder` is a type in Go’s `strings` package that lets us **build strings efficiently** by appending multiple parts together **without creating intermediate copies** each time.

Normally, if we do:

```go
s := ""
s += "Hello"
s += " "
s += "World"
```

Go has to create **a new string each time** (because strings are immutable in Go), which is inefficient for large concatenations.

With `strings.Builder`:

* We can append to an internal buffer.
* It only allocates new memory when needed.
* This avoids multiple intermediate allocations.

---

## **2. Importing**

```go
import "strings"
```

---

## **3. Creating a `strings.Builder`**

```go
var b strings.Builder // zero value is ready to use
// OR
b := strings.Builder{}
```

We don’t need to manually initialize it — the zero value works.

---

## **4. Common Methods**

Let’s explore the key methods (all from `strings.Builder`):

| Method                    | Purpose                              | Example                 |
| ------------------------- | ------------------------------------ | ----------------------- |
| **`Write([]byte)`**       | Appends byte slice to the builder    | `b.Write([]byte("Go"))` |
| **`WriteString(string)`** | Appends string                       | `b.WriteString("Lang")` |
| **`WriteRune(rune)`**     | Appends a single Unicode character   | `b.WriteRune('✓')`      |
| **`String()`**            | Returns the built string             | `result := b.String()`  |
| **`Len()`**               | Current length of built string       | `b.Len()`               |
| **`Cap()`**               | Current capacity (allocated memory)  | `b.Cap()`               |
| **`Grow(n int)`**         | Preallocates at least `n` more bytes | `b.Grow(50)`            |
| **`Reset()`**             | Clears the content for reuse         | `b.Reset()`             |

---

## **5. Example Usage**

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder

	b.WriteString("Hello")
	b.WriteRune(' ')
	b.WriteString("World")
	b.WriteByte('!')

	fmt.Println(b.String()) // Hello World!
	fmt.Println("Length:", b.Len())
	fmt.Println("Capacity:", b.Cap())
}
```

---

## **6. Why Use `strings.Builder`?**

* **Efficiency:** Avoids repeated allocations from string concatenation.
* **Reusability:** Can reset and reuse the builder without reallocation (in some cases).
* **Flexibility:** Works with strings, runes, and byte slices.

---

## **7. Performance Difference**

```go
package main

import (
	"strings"
	"time"
)

func concatNormal() string {
	s := ""
	for i := 0; i < 10000; i++ {
		s += "x"
	}
	return s
}

func concatBuilder() string {
	var b strings.Builder
	b.Grow(10000) // preallocate
	for i := 0; i < 10000; i++ {
		b.WriteByte('x')
	}
	return b.String()
}

func main() {
	start := time.Now()
	concatNormal()
	fmt.Println("Normal concat took:", time.Since(start))

	start = time.Now()
	concatBuilder()
	fmt.Println("Builder concat took:", time.Since(start))
}
```

On large iterations, `strings.Builder` can be **5x–10x faster**.

---

## **8. Gotchas & Best Practices**

1. **Don’t copy `strings.Builder` values.**
   Copying a builder will break things because it internally manages a slice of bytes. Always pass a pointer if needed.

2. **Not safe for concurrent use.**
   If multiple goroutines need to build a string together, use a lock or separate builders.

3. **Reuse carefully.**
   `Reset()` clears data but does not shrink capacity — useful for reusing memory in loops.

4. **Preallocate when possible.**
   Using `Grow()` helps avoid repeated memory growth.

---

Alright — let’s go deep into **Text Templates** in Go, because they’re one of the most powerful tools in the standard library for generating dynamic text output (HTML, configuration files, CLI output, emails, etc.) without resorting to messy string concatenation.

We’ll cover **how they work, syntax, data passing, functions, advanced features, and best practices**.

---

## **1. What is a Text Template in Go?**

A **text template** is like a blueprint for generating text output, where we can:

* **Insert values dynamically** from variables/structs/maps.
* **Use control structures** (loops, conditionals).
* **Call functions**.
* **Nest templates** for modularity.

They come from the **`text/template`** package:

```go
import "text/template"
```

Unlike `html/template`, it **does not escape HTML** — so it’s used for plain text, not HTML rendering (though syntax is almost identical).

---

## **2. Basic Workflow**

Using templates in Go generally follows these steps:

1. **Define** the template (string or file).
2. **Parse** it into a `*template.Template`.
3. **Execute** it with some data (the `Execute` method).

---

### **Basic Example**

```go
package main

import (
	"os"
	"text/template"
)

func main() {
	// Step 1: Create a template string
	tmpl := `Hello, {{.Name}}! You are {{.Age}} years old.`

	// Step 2: Parse it
	t, err := template.New("greeting").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Step 3: Execute with data
	data := struct {
		Name string
		Age  int
	}{"Skyy", 29}

	t.Execute(os.Stdout, data)
}
```

**Output:**

```
Hello, Skyy! You are 29 years old.
```

---

## **3. Template Syntax Basics**

All Go templates use `{{ ... }}` for expressions.

| Syntax                 | Meaning                     | Example                             |
| ---------------------- | --------------------------- | ----------------------------------- |
| `{{.}}`                | Current data item           | `{{.}}`                             |
| `{{.Field}}`           | Access struct field         | `{{.Name}}`                         |
| `{{.Key}}`             | Access map key              | `{{.City}}`                         |
| `{{.Method}}`          | Call method on data         | `{{.UpperCase}}`                    |
| `{{ $var := .Field }}` | Assign variable in template | `{{$name := .Name}}Hello {{$name}}` |

---

## **4. Control Structures**

### **If**

```go
{{if .IsAdmin}}
    Welcome, admin!
{{else}}
    Hello, user.
{{end}}
```

### **Range (loop)**

```go
{{range .Items}}
    - {{.}}
{{end}}
```

If `.Items` is a slice of strings, `{{.}}` will print each string.

### **With**

Narrow the scope to a specific value:

```go
{{with .Address}}
    City: {{.City}}
    Zip: {{.Zip}}
{{end}}
```

---

## **5. Functions in Templates**

Go templates have built-in functions (`and`, `or`, `not`, `len`, `index`, `printf`, `html`, `urlquery`, etc.).

Example:

```go
{{printf "My age is %d" .Age}}
```

---

## **6. Custom Functions**

We can add **our own functions** with `Funcs()`.

```go
package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	funcs := template.FuncMap{
		"upper": strings.ToUpper,
	}

	tmpl := `{{upper .Name}}`
	t := template.New("test").Funcs(funcs)
	t, _ = t.Parse(tmpl)

	data := struct{ Name string }{"skyy"}
	t.Execute(os.Stdout, data)
}
```

**Output:**

```
SKYY
```

---

## **7. Parsing Templates from Files**

```go
t, err := template.ParseFiles("file1.txt", "file2.txt")
t.Execute(os.Stdout, data)
```

Or from a directory:

```go
t, _ := template.ParseGlob("templates/*.txt")
```

---

## **8. Nested Templates**

We can define reusable chunks inside a file:

```gotemplate
{{define "header"}}== HEADER =={{end}}
{{define "footer"}}== FOOTER =={{end}}

{{template "header"}}
Body content here
{{template "footer"}}
```

---

## **9. Error Handling**

* **`Parse` errors** → bad template syntax.
* **`Execute` errors** → invalid data access (like `{{.NonExistent}}` when field doesn’t exist).

---

## **10. Common Gotchas**

❌ **Dot (`.`) confusion** – `.` changes inside `range` and `with`. If you need outer context, use `$` to store it:

```go
{{$root := .}}
{{range .Items}}
    {{.}} from {{$root.Category}}
{{end}}
```

❌ **No HTML escaping** – `text/template` doesn’t protect from XSS. Use `html/template` for web content.

❌ **Function order** – You must attach `Funcs()` **before** parsing.

---

## **11. Do’s & Don’ts**

✅ **Use `text/template` for non-HTML** (emails, logs, config).
✅ **Pre-parse templates** at startup for performance.
✅ **Attach functions before parsing**.
✅ **Use `template.Must()`** for compile-time checking.
❌ **Don’t store `template.Template` in global mutable vars** without sync in concurrent apps.
❌ **Don’t overcomplicate logic in templates** — keep them simple and readable.

---

Now, let’s unpack this Golang **Text Template** program in detail — section by section — so we fully understand what’s happening and why.

This program is essentially:

* Demonstrating **`text/template`** usage.
* Showing how to dynamically populate strings (like HTML, notifications, error messages).
* Providing an **interactive menu** to render templates with user-provided data.

---

## **1️⃣ Imports and Purpose**

```go
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)
```

We are importing:

* **`bufio`** → Buffered I/O to read input from the console.
* **`fmt`** → Printing and formatting output.
* **`os`** → For standard input/output streams (e.g., `os.Stdout`).
* **`strings`** → For trimming spaces/newlines from user input.
* **`text/template`** → Go’s template engine for text-based templates (safe for console, text files, and non-HTML).

---

## **2️⃣ Creating a Simple Template**

```go
tmpl, err := template.New("example").Parse("Welcome , {{.name}}! How are you doing?\n")
```

* **`template.New("example")`** → Creates a new template instance named `"example"`.
* **`.Parse()`** → Compiles a template string with placeholders.
* **`{{.name}}`** → Dot (`.`) refers to the data object passed into the template.

  * In this case, `.name` means “look for the field/key `name` in the provided data.”

**Error check:**

```go
if err != nil {
	panic(err.Error())
}
```

If parsing fails, the program stops.

---

## **3️⃣ Executing the Template**

```go
data := map[string]any{
	"name": "John",
}
err = tmpl.Execute(os.Stdout, data)
```

* **`data`** → A map with `"name": "John"`.
* **`tmpl.Execute(os.Stdout, data)`** → Replaces `{{.name}}` with `"John"` and prints to console.

---

## **4️⃣ Using `template.Must` for Simplicity**

```go
tmpl1 := template.Must(template.New("example").Parse("Welcome , {{.name}}! How are you doing?\n"))
```

* `template.Must()` → Panics immediately if parsing fails.
  This is handy for templates that should **always** be valid.

```go
data1 := map[string]any{
	"name": "Skyy",
}
tmpl1.Execute(os.Stdout, data1)
```

Prints `"Welcome, Skyy!"`.

---

## **5️⃣ Reading User Input**

```go
reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter your name: ")
name, err := reader.ReadString('\n')
name = strings.TrimSpace(name)
```

* Reads the user’s name.
* `strings.TrimSpace()` removes `\n` and extra spaces.

---

## **6️⃣ Defining Multiple Templates**

```go
templates := map[string]string{
	"welcome": "Herzlich Willkommen {{.name}}! Freut mich sehr 💖",
	"notification": "{{.nm}} , you have a new notification: {{.ntf}}",
	"error": "Oops! An ERROR occurred: {{.errorMsg}} ⚠️",
}
```

* **Why map?** → Easy to store many templates and retrieve them by key.

---

## **7️⃣ Parsing and Storing Templates**

```go
parsedTemps := make(map[string]*template.Template)
for name, tmp := range templates {
	parsedTemps[name] = template.Must(template.New(name).Parse(tmp))
}
```

* Loop over all template strings.
* Parse them into **`*template.Template`** objects and store for later use.

---

## **8️⃣ Menu Loop (While-like for { })**

```go
for {
	// Show menu
	fmt.Println("\n📃Menu:")
	fmt.Println("1. Join 👤")
	fmt.Println("2. Get Notification 🔔")
	fmt.Println("3. Get Error 🔴")
	fmt.Println("4. Exit 🚪")
	fmt.Println("Choose an option: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var data map[string]any
	var tmplt *template.Template
```

* This runs forever until the user chooses **Exit**.
* `data` and `tmplt` are set based on menu choice.

---

## **9️⃣ Handling User Choice**

```go
switch choice {
case "1":
	tmplt = parsedTemps["welcome"]
	data = map[string]any{"name": name}

case "2":
	fmt.Println("Enter your notification msg:")
	notification, _ := reader.ReadString('\n')
	notification = strings.TrimSpace(notification)
	tmplt = parsedTemps["notification"]
	data = map[string]any{"nm": name, "ntf": notification}

case "3":
	fmt.Println("Enter your errorMessage:")
	errorMessage, _ := reader.ReadString('\n')
	errorMessage = strings.TrimSpace(errorMessage)
	tmplt = parsedTemps["error"]
	data = map[string]any{"nm": name, "errorMsg": errorMessage}

case "4":
	fmt.Println("Exiting... ✔️")
	return

default:
	fmt.Println("Invalid Choice.. Please select a valid option!")
	continue
}
```

* Each case sets:

  * **`tmplt`** → Which template to use.
  * **`data`** → What data to insert.

---

## **🔟 Rendering the Template**

```go
err := tmplt.Execute(os.Stdout, data)
if err != nil {
	fmt.Println("Error executing template:", err)
}
```

* Executes the chosen template with its data.
* If there’s a missing key or syntax error, we’ll see an error.

---

## **🔍 Key Takeaways**

* **`text/template`** is for plain text (HTML uses `html/template` for auto-escaping).
* Templates use **`{{.FieldName}}`** syntax to insert values.
* `template.Must()` simplifies error handling at parse time.
* We can store multiple templates in **maps** for flexible retrieval.
* Templates are **type-safe** — missing fields cause execution errors.

---

## **✅ Do’s**

* ✅ Use `template.Must()` for static templates.
* ✅ Always check `Execute()` errors.
* ✅ Store templates in a map if they’re reused often.
* ✅ Trim spaces/newlines when taking user input.

## **❌ Don’ts**

* ❌ Don’t use `text/template` for HTML output on websites (use `html/template` instead for XSS safety).
* ❌ Don’t ignore execution errors — templates fail silently if we skip error checks.
* ❌ Don’t hardcode dynamic user data without sanitizing when it’s coming from untrusted input.

---

Alright — let’s deep dive into **Go’s `time` package** and **time handling** so we can confidently work with dates, durations, timers, and formatting.

We’ll go step by step, covering *everything* with examples.

---

## **1. Importing the Package**

```go
import "time"
```

The `time` package is part of Go’s standard library, so no extra installation is needed.

---

## **2. Getting the Current Time**

```go
now := time.Now()
fmt.Println(now) // e.g., 2025-08-09 14:23:45.123456 +0530 IST
```

* Returns a `time.Time` object.
* Includes date, time, timezone, and location.

---

## **3. Creating Specific Times**

```go
t := time.Date(2025, time.August, 30, 12, 0, 0, 0, time.UTC)
fmt.Println(t) // 2025-08-30 12:00:00 +0000 UTC
```

**Parameters**:
`year, month, day, hour, min, sec, nsec, location`

---

## **4. Parsing & Formatting Time**

Go uses **reference time**:

```
Mon Jan 2 15:04:05 MST 2006
```

* The numbers must be exactly this date & time to define a format.

### **Formatting**

```go
fmt.Println(now.Format("2006-01-02 15:04:05"))
fmt.Println(now.Format("02-Jan-2006"))
```

### **Parsing**

```go
layout := "2006-01-02 15:04:05"
t, err := time.Parse(layout, "2025-08-09 18:30:00")
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(t)
```

---

## **5. Time Zones**

```go
loc, _ := time.LoadLocation("America/New_York")
ny := now.In(loc)
fmt.Println(ny)
```

* Time zones come from the system’s `zoneinfo` database.

---

## **6. Getting Time Components**

```go
fmt.Println(now.Year())       // 2025
fmt.Println(now.Month())      // August
fmt.Println(now.Day())        // 9
fmt.Println(now.Weekday())    // Saturday
fmt.Println(now.Hour())       // 14
fmt.Println(now.Minute())     // 23
fmt.Println(now.Second())     // 45
```

---

## **7. Comparing Times**

```go
t1 := time.Now()
time.Sleep(2 * time.Second)
t2 := time.Now()

fmt.Println(t2.After(t1))  // true
fmt.Println(t1.Before(t2)) // true
fmt.Println(t1.Equal(t2))  // false
```

---

## **8. Durations**

`time.Duration` is a type representing nanoseconds.

```go
d := 2 * time.Hour
fmt.Println(d.Hours())       // 2
fmt.Println(d.Minutes())     // 120
fmt.Println(d.Seconds())     // 7200
```

---

## **9. Adding & Subtracting Time**

```go
fmt.Println(now.Add(48 * time.Hour)) // Add 2 days
fmt.Println(now.Add(-time.Hour))     // Subtract 1 hour
fmt.Println(now.Sub(t1))             // Difference
```

---

## **10. Timers**

A timer triggers once after a set duration.

```go
timer := time.NewTimer(3 * time.Second)
<-timer.C
fmt.Println("3 seconds passed")
```

---

## **11. Tickers**

A ticker triggers repeatedly at intervals.

```go
ticker := time.NewTicker(2 * time.Second)
go func() {
    for t := range ticker.C {
        fmt.Println("Tick at", t)
    }
}()
time.Sleep(7 * time.Second)
ticker.Stop()
```

---

## **12. Sleeping**

```go
time.Sleep(5 * time.Second) // Pause execution
```

---

## **13. Measuring Execution Time**

```go
start := time.Now()
time.Sleep(1 * time.Second)
elapsed := time.Since(start)
fmt.Println("Execution time:", elapsed)
```

---

## **14. UTC vs Local Time**

```go
fmt.Println(now.UTC())   // Convert to UTC
fmt.Println(now.Local()) // Convert to local timezone
```

---

## **15. Unix Timestamps**

```go
fmt.Println(now.Unix())      // Seconds since Jan 1 1970
fmt.Println(now.UnixMilli()) // Milliseconds
fmt.Println(now.UnixNano())  // Nanoseconds
```

Convert from Unix:

```go
t := time.Unix(1733790000, 0)
fmt.Println(t)
```

---

✅ **Summary Table — Common Functions in `time`**

| Function/Method          | Purpose               |
| ------------------------ | --------------------- |
| `time.Now()`             | Get current time      |
| `time.Date()`            | Create custom time    |
| `time.Parse()`           | Convert string → time |
| `t.Format()`             | Convert time → string |
| `t.Year()`, `.Month()`   | Extract components    |
| `t.After()`, `.Before()` | Compare times         |
| `t.Add()`, `.Sub()`      | Add/subtract time     |
| `time.Since()`           | Elapsed time          |
| `time.Sleep()`           | Pause execution       |
| `time.NewTimer()`        | Run once after delay  |
| `time.NewTicker()`       | Repeat at intervals   |
| `t.Unix()`               | Unix timestamp        |
| `t.In(location)`         | Convert timezone      |

---

Here’s a deep, practical guide to **Epoch time in Go** (aka “Unix time”). This is everything we’ll use when building APIs, storing timestamps in DBs, logging, TTLs, etc.

---

# What is “Epoch” (Unix time)?

* **Unix epoch** = **Jan 1, 1970 00:00:00 UTC**.
* A Unix timestamp is a **count of time since the epoch**:

  * **seconds** (classic): `int64`
  * **milliseconds**, **microseconds**, **nanoseconds**: finer units for higher precision

---

# Getting “now” as Unix time

```go
now := time.Now()

secs  := now.Unix()        // int64 seconds since epoch
millis:= now.UnixMilli()   // int64 milliseconds
micros:= now.UnixMicro()   // int64 microseconds
nanos := now.UnixNano()    // int64 nanoseconds
```

> Use `Unix()` for DBs that store seconds. Many JS/HTTP clients use **ms**; mobile SDKs often use **ms** or **s**.

---

# Converting Unix → time.Time (and timezone)

```go
// From seconds + optional nanoseconds
tUTC := time.Unix(1734295200, 0)      // 2025-12-16T12:00:00Z (example)

// From milliseconds
ms := int64(1734295200000)
tFromMs := time.Unix(0, ms*int64(time.Millisecond))

// Show in a specific zone (e.g., Asia/Kolkata)
loc, _ := time.LoadLocation("Asia/Kolkata")
tIST := tFromMs.In(loc)
fmt.Println(tIST.Format(time.RFC3339)) // 2025-12-16T17:30:00+05:30
```

**Notes**

* `time.Unix(sec, nsec)` treats arguments as **UTC-based** epoch values.
* Convert display with `t.In(location)`.

---

# Converting time.Time → Unix

```go
t := time.Date(2025, 8, 15, 10, 0, 0, 0, time.UTC)
fmt.Println(t.Unix())      // seconds
fmt.Println(t.UnixMilli()) // milliseconds
fmt.Println(t.UnixNano())  // nanoseconds
```

---

# Formatting vs. Unix

* **Unix** is numeric; **formatting** is text.
* Go’s layout reference is: `2006-01-02 15:04:05` (Mon Jan 2 15:04:05 MST 2006).

```go
fmt.Println(time.Now().Format(time.RFC3339))           // 2025-08-15T12:34:56+05:30
fmt.Println(time.Now().Format("2006-01-02 15:04:05"))  // custom format
```

There’s **no direct parser** for numeric epoch strings; parse to an integer, then `time.Unix`.

```go
s := "1734295200" // seconds as string
sec, _ := strconv.ParseInt(s, 10, 64)
t := time.Unix(sec, 0)
```

---

# Choosing the right unit (s/ms/µs/ns)

* **Seconds**: compact, common in DBs and APIs; good default.
* **Milliseconds**: common in JS/frontends; good for events/UI timelines.
* **Micro/Nano**: use when you need high precision (traces, profiling).

**DB tip:** store as **INTEGER** (seconds or ms). Document the unit!

---

# Epoch in APIs (typical patterns)

### 1) Return timestamps in JSON

```go
type Event struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    CreatedAt int64  `json:"createdAt"` // seconds since epoch
}
e := Event{ID: 1, Name: "Go Meetup", CreatedAt: time.Now().Unix()}
```

### 2) Accept ms from client → normalize to time.Time

```go
var payload struct {
    StartAtMs int64 `json:"startAtMs"`
}
_ = c.ShouldBindJSON(&payload)
start := time.Unix(0, payload.StartAtMs*int64(time.Millisecond))
```

### 3) TTL / expiration

```go
exp := time.Now().Unix() + 3600 // +1 hour (seconds)
```

---

# Monotonic time vs. Unix time (IMPORTANT)

* `time.Time` in Go can carry a **monotonic clock** reading for **duration measurements**.
* **Use** `time.Since(start)` / `time.Until(t)` or `t2.Sub(t1)` to measure elapsed time.
* **Don’t** compute elapsed time by subtracting Unix timestamps; wall clocks can jump (NTP, DST).

```go
start := time.Now()
doWork()
elapsed := time.Since(start) // correct; robust to clock changes
```

Calling `Unix()` / formatting **drops** the monotonic part — that’s expected.

---

# Edge cases & pitfalls

1. **Time zone confusion**
   Store epoch in UTC; convert to user’s zone at the edge (UI).

2. **Milliseconds vs. seconds mixups**
   If a timestamp seems “in 1970”, you probably treated **ms as s**.
   If it’s “way in the future”, you probably treated **s as ms**.

3. **Overflow concerns**

   * `int64` **nanoseconds** since 1970 will overflow around year **2262**.
   * `seconds` and `milliseconds` are safe for far longer with `int64`.

4. **SQLite helpers** (if you use SQLite)

   * Save as `INTEGER` seconds and use `datetime(created_at, 'unixepoch')` to render.
   * Or use Go to format.

---

# Handy utilities you’ll reuse

```go
// Now (UTC) as seconds
func NowSec() int64 { return time.Now().UTC().Unix() }

// Convert seconds → time in a zone
func SecToTimeIn(sec int64, tz string) (time.Time, error) {
    loc, err := time.LoadLocation(tz)
    if err != nil { return time.Time{}, err }
    return time.Unix(sec, 0).In(loc), nil
}

// Parse epoch (ms) string → time.Time (UTC)
func ParseMsStringToTimeUTC(s string) (time.Time, error) {
    ms, err := strconv.ParseInt(s, 10, 64)
    if err != nil { return time.Time{}, err }
    return time.Unix(0, ms*int64(time.Millisecond)).UTC(), nil
}
```

---

# Quick reference (cheat list)

* Now: `time.Now()`
* Seconds since epoch: `t.Unix()`
* Milliseconds since epoch: `t.UnixMilli()`
* From seconds: `time.Unix(sec, 0)`
* From ms: `time.Unix(0, ms*int64(time.Millisecond))`
* To specific TZ: `t.In(loc)`
* Format: `t.Format(time.RFC3339)`
* Parse layout/string: `time.Parse(layout, s)`
* Measure durations: `time.Since(start)` / `t2.Sub(t1)`

---

Random numbers in Go (`math/rand` and `crypto/rand`) can be a bit tricky for newcomers because they work differently depending on whether we want **fast pseudo-random numbers** or **secure cryptographic randomness**. Let’s go deep.

---

# 🔹 Random Numbers in Go

## 1️⃣ Two Main Packages

1. **`math/rand`**

   * Generates **pseudo-random numbers** (deterministic, reproducible).
   * Good for **games, simulations, shuffling, sampling**.
   * Not secure for cryptography.

2. **`crypto/rand`**

   * Generates **cryptographically secure random numbers**.
   * Uses OS’s secure entropy source.
   * Good for **passwords, keys, tokens, salts**.

---

## 2️⃣ `math/rand` Basics

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed with current time to avoid same sequence
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Random int:", rand.Int())        // random int
	fmt.Println("Random int < 100:", rand.Intn(100)) // 0 ≤ n < 100
	fmt.Println("Random float [0.0,1.0):", rand.Float64())
	fmt.Println("Random float32:", rand.Float32())

	// Random permutation
	fmt.Println("Random permutation:", rand.Perm(5)) // e.g. [3 1 4 0 2]
}
```

👉 Without `rand.Seed(...)`, Go always produces the **same sequence** (deterministic for reproducibility).
👉 With `Seed(time.Now().UnixNano())`, we get different values every run.

---

## 3️⃣ Generating Random Ranges

```go
// Random int in [min, max]
func randInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
```

Example:

```go
fmt.Println(randInt(10, 20)) // 10 ≤ n ≤ 20
```

---

## 4️⃣ Shuffling a Slice

```go
nums := []int{1, 2, 3, 4, 5}
rand.Shuffle(len(nums), func(i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
})
fmt.Println("Shuffled:", nums)
```

---

## 5️⃣ `crypto/rand` (Secure Randomness)

```go
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Secure random int in [0, 99]
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println("Secure random number:", n)

	// Secure random bytes
	b := make([]byte, 16) // 16 bytes = 128-bit
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Secure random bytes: %x\n", b)
}
```

👉 Use this for **tokens, salts, UUIDs, keys**.

---

## 6️⃣ Random Strings (Example: Password/Token)

```go
import "crypto/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[num.Int64()]
	}
	return string(b)
}
```

Example:

```go
fmt.Println(randomString(12)) // e.g. "aX3pQ8rTk1Zm"
```

---

## 7️⃣ Summary

* **`math/rand`**

  * Fast, reproducible
  * Needs `Seed()` for variety
  * Great for games, non-security apps
* **`crypto/rand`**

  * Truly random (uses OS entropy)
  * Safe for security-sensitive work
  * Slightly slower

---

✅ Rule of Thumb:

* Use **`math/rand`** → simulations, shuffling, random choices.
* Use **`crypto/rand`** → passwords, tokens, encryption, security.

---

Let’s build a **side-by-side cheat sheet** for random numbers in Go. This will help us quickly decide when to use `math/rand` vs `crypto/rand`.

---

# 🎲 Go Random Numbers Cheat Sheet

| Feature              | `math/rand`                             | `crypto/rand`                           |
| -------------------- | --------------------------------------- | --------------------------------------- |
| **Type**             | Pseudo-random (deterministic sequence)  | Cryptographically secure randomness     |
| **Source**           | PRNG seeded with `rand.Seed(...)`       | OS entropy (randomness from kernel/CPU) |
| **Performance**      | Very fast                               | Slower (secure entropy gathering)       |
| **Security**         | ❌ Not safe for cryptography             | ✅ Safe for cryptography                 |
| **When to use?**     | Games, simulations, sampling, shuffling | Passwords, tokens, keys, salts          |
| **Repeatable?**      | Yes, with same seed                     | No                                      |
| **Package**          | `math/rand`                             | `crypto/rand`                           |
| **Example Use Case** | Shuffle deck of cards                   | Generate secure API key                 |

---

## 📌 Code Examples

### 1️⃣ `math/rand` → (fast, non-secure)

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // important for varied results

	fmt.Println("Random Int:", rand.Int())
	fmt.Println("Random Int < 100:", rand.Intn(100))
	fmt.Println("Random Float:", rand.Float64())
	fmt.Println("Random Permutation:", rand.Perm(5))

	// Random int between [min, max]
	min, max := 10, 20
	fmt.Println("Random in range:", rand.Intn(max-min+1)+min)
}
```

👉 Output (changes every run):

```
Random Int: 557939843
Random Int < 100: 42
Random Float: 0.8429
Random Permutation: [3 1 4 0 2]
Random in range: 17
```

---

### 2️⃣ `crypto/rand` → (secure randomness)

```go
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Secure int in [0, 99]
	n, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println("Secure Random Int:", n)

	// Secure bytes
	b := make([]byte, 16) // 128-bit random
	rand.Read(b)
	fmt.Printf("Secure Bytes: %x\n", b)
}
```

👉 Output (different every run, not reproducible):

```
Secure Random Int: 73
Secure Bytes: 3f8b1a9c3ddfa4c84f6a23b7e2ab3c2d
```

---

### 3️⃣ Random Secure String Generator (useful for API keys, tokens)

```go
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[num.Int64()]
	}
	return string(b)
}
```

Example:

```go
fmt.Println(randomString(12)) // "aX9rTzQ1PmLd"
```

---

## 🔑 TL;DR Rule of Thumb

* **Need speed & reproducibility?** → `math/rand`
* **Need security (tokens, passwords)?** → `crypto/rand`

---

Perfect topic 🚀 — `strconv` is one of those core Go packages we keep using but often only know bits of. Let’s go **all in** and break it down in detail.

---

# 📦 `strconv` Package in Go

The **`strconv`** package in Go is used to **convert between strings and other basic data types** (integers, floats, booleans, etc.).
Think of it as the bridge between `"123"` (string) and `123` (int).

---

## ✨ Major Functions in `strconv`

### 🔹 1. String ↔ Integer

* **`strconv.Atoi(s string) (int, error)`**

  * Converts a string to an `int`.
  * Simple shortcut for `strconv.ParseInt(s, 10, 0)`.

```go
i, err := strconv.Atoi("42")
fmt.Println(i) // 42
```

---

* **`strconv.Itoa(i int) string`**

  * Converts an `int` to a string.

```go
s := strconv.Itoa(42)
fmt.Println(s) // "42"
```

---

* **`strconv.ParseInt(s string, base int, bitSize int)`**

  * More powerful, can handle bases and different integer sizes.
  * `bitSize` can be `0, 8, 16, 32, 64`.

```go
n, err := strconv.ParseInt("1010", 2, 64) // binary
fmt.Println(n) // 10
```

---

* **`strconv.FormatInt(i int64, base int)`**

  * Converts int64 → string in any base (2 to 36).

```go
s := strconv.FormatInt(255, 16)
fmt.Println(s) // "ff"
```

---

### 🔹 2. String ↔ Unsigned Integer

* **`strconv.ParseUint(s string, base int, bitSize int)`**
* **`strconv.FormatUint(i uint64, base int)`**

```go
u, _ := strconv.ParseUint("FF", 16, 64)
fmt.Println(u) // 255

s := strconv.FormatUint(255, 2)
fmt.Println(s) // "11111111"
```

---

### 🔹 3. String ↔ Float

* **`strconv.ParseFloat(s string, bitSize int)`**

  * Parses a string into float32/64 depending on `bitSize`.

```go
f, _ := strconv.ParseFloat("3.14159", 64)
fmt.Println(f) // 3.14159
```

---

* **`strconv.FormatFloat(f float64, fmt byte, prec, bitSize int)`**

  * Converts float → string with control over formatting.

Params:

* `fmt`: `'f'` (decimal), `'e'` (scientific), `'g'` (compact)
* `prec`: number of digits after decimal (`-1` means auto)
* `bitSize`: `32` or `64`

```go
s := strconv.FormatFloat(3.14159, 'f', 2, 64)
fmt.Println(s) // "3.14"
```

---

### 🔹 4. String ↔ Boolean

* **`strconv.ParseBool(s string)`**

  * Accepts: `"1"`, `"t"`, `"T"`, `"true"`, `"TRUE"`, `"True"` → true
    `"0"`, `"f"`, `"F"`, `"false"`, `"FALSE"`, `"False"` → false

```go
b, _ := strconv.ParseBool("true")
fmt.Println(b) // true
```

---

* **`strconv.FormatBool(b bool)`**

  * Converts bool → `"true"` / `"false"`

```go
s := strconv.FormatBool(false)
fmt.Println(s) // "false"
```

---

### 🔹 5. Quote & Unquote (escaping strings)

* **`strconv.Quote(s string)`**

  * Returns a **double-quoted Go string literal** with escapes.

```go
s := strconv.Quote("Hello\tWorld")
fmt.Println(s) // "\"Hello\tWorld\""
```

---

* **`strconv.QuoteToASCII(s string)`**

  * Escapes non-ASCII chars.

```go
s := strconv.QuoteToASCII("こんにちは")
fmt.Println(s) // "\"\\u3053\\u3093\\u306b\\u3061\\u306f\""
```

---

* **`strconv.Unquote(s string)`**

  * Removes surrounding quotes and interprets escapes.

```go
s, _ := strconv.Unquote("\"Hello\\tWorld\"")
fmt.Println(s) // Hello	World
```

---

### 🔹 6. Append Functions (efficient string building)

Instead of creating a new string, these append to an existing `[]byte`.

* **`strconv.AppendInt(dst []byte, i int64, base int) []byte`**
* **`strconv.AppendFloat(...)`**
* **`strconv.AppendBool(...)`**
* **`strconv.AppendQuote(...)`**

```go
b := []byte("Number: ")
b = strconv.AppendInt(b, 42, 10)
fmt.Println(string(b)) // "Number: 42"
```

---

## ⚡ Quick Overview Table

| Conversion           | String → X          | X → String                       |
| -------------------- | ------------------- | -------------------------------- |
| **Int**              | `Atoi` / `ParseInt` | `Itoa` / `FormatInt`             |
| **Uint**             | `ParseUint`         | `FormatUint`                     |
| **Float**            | `ParseFloat`        | `FormatFloat`                    |
| **Bool**             | `ParseBool`         | `FormatBool`                     |
| **Quote**            | `Unquote`           | `Quote`, `QuoteToASCII`          |
| **Efficient Append** | —                   | `AppendInt`, `AppendFloat`, etc. |

---

## 🔑 Key Points

* Always check for **errors** when parsing (`ParseX`).
* Use `Atoi` / `Itoa` for **simple int <-> string**.
* Use `Format`/`Parse` variants for **control (base, precision, bit size)**.
* Use `AppendX` when building strings efficiently in loops.
* `Quote` / `Unquote` are helpful for safe string literals and JSON-like escaping.

---
URL parsing in Go is an **essential skill** because almost every web-related program (HTTP clients, APIs, servers) requires dealing with URLs. Go gives us a **robust package** for this:

---

# 🌍 URL Parsing in Go (`net/url`)

The Go standard library has the **`net/url`** package for:

* Parsing URLs into components (`scheme`, `host`, `path`, `query`, `fragment`, etc.)
* Building and modifying URLs
* Encoding/decoding query parameters
* Ensuring URLs are properly escaped

---

## 🔹 Basic Structure of a URL

```
scheme://[userinfo@]host/path[?query][#fragment]
```

Example:

```
https://user:pass@www.example.com:8080/search?q=golang#section2
```

Breakdown:

* **scheme** → `https`
* **userinfo** → `user:pass`
* **host** → `www.example.com:8080`

  * host **name** → `www.example.com`
  * host **port** → `8080`
* **path** → `/search`
* **query** → `q=golang`
* **fragment** → `section2`

---

## 🔹 Parsing URLs

### `url.Parse(rawurl string) (*url.URL, error)`

Parses a string into a `*url.URL` struct.

```go
package main

import (
	"fmt"
	"net/url"
)

func main() {
	rawURL := "https://user:pass@www.example.com:8080/search?q=golang#section2"
	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", parsed.Scheme)       // https
	fmt.Println("User:", parsed.User)           // user:pass
	fmt.Println("Username:", parsed.User.Username()) // user
	pass, _ := parsed.User.Password()
	fmt.Println("Password:", pass)              // pass
	fmt.Println("Host:", parsed.Host)           // www.example.com:8080
	fmt.Println("Hostname:", parsed.Hostname()) // www.example.com
	fmt.Println("Port:", parsed.Port())         // 8080
	fmt.Println("Path:", parsed.Path)           // /search
	fmt.Println("RawQuery:", parsed.RawQuery)   // q=golang
	fmt.Println("Fragment:", parsed.Fragment)   // section2
}
```

---

## 🔹 Working with Query Parameters

### Parse queries from URL

```go
values, _ := url.ParseQuery("q=golang&lang=en&lang=fr")

fmt.Println(values.Get("q"))       // "golang"
fmt.Println(values["lang"])        // [en fr]
```

---

### Extract from URL object

```go
parsed, _ := url.Parse("https://example.com/search?q=golang&lang=en")
query := parsed.Query()
fmt.Println(query.Get("q")) // golang
```

---

### Modify Queries

```go
q := parsed.Query()
q.Set("q", "golang tutorials")
q.Add("page", "2")
parsed.RawQuery = q.Encode()

fmt.Println(parsed.String())
// https://example.com/search?q=golang+tutorials&page=2&lang=en
```

---

## 🔹 Building URLs

```go
u := &url.URL{
	Scheme:   "https",
	Host:     "www.example.com",
	Path:     "docs",
	RawQuery: "q=golang",
	Fragment: "intro",
}

fmt.Println(u.String())
// https://www.example.com/docs?q=golang#intro
```

---

## 🔹 URL Escaping & Unescaping

To safely embed values in URLs (spaces, special characters, etc.), we use:

* **`url.QueryEscape(s string) string`**
  Escapes a string so it’s safe in query parameters.

```go
escaped := url.QueryEscape("golang tutorials & examples")
fmt.Println(escaped) // golang+tutorials+%26+examples
```

---

* **`url.QueryUnescape(s string) (string, error)`**
  Decodes back to original.

```go
decoded, _ := url.QueryUnescape("golang+tutorials+%26+examples")
fmt.Println(decoded) // golang tutorials & examples
```

---

## 🔹 Relative vs Absolute URLs

* **Absolute URL** → has scheme + host
* **Relative URL** → path only, must resolve against base

### Resolving Relative URLs

```go
base, _ := url.Parse("https://example.com/docs/")
ref, _ := url.Parse("../about")

resolved := base.ResolveReference(ref)
fmt.Println(resolved.String()) // https://example.com/about
```

---

## 🔹 The `url.URL` Struct

The main struct has these fields:

```go
type URL struct {
    Scheme     string
    Opaque     string
    User       *Userinfo
    Host       string
    Path       string
    RawPath    string
    ForceQuery bool
    RawQuery   string
    Fragment   string
}
```

We usually deal with:

* `Scheme`
* `User`
* `Host`, `Hostname()`, `Port()`
* `Path`
* `RawQuery` / `Query()`
* `Fragment`

---

## ⚡ Key Points to Remember

1. Use **`url.Parse`** to safely break down URLs.
2. Use **`url.ParseQuery`** or `parsed.Query()` for query params.
3. Always reassign `parsed.RawQuery = q.Encode()` after modifying queries.
4. Escape untrusted input with **`url.QueryEscape`**.
5. Use `ResolveReference` for relative → absolute URL resolution.

---

`bufio` is one of the most **practical packages** in Go when working with input/output. It adds a **buffer layer** on top of readers and writers, making data handling faster and more efficient. Let’s go deep 👇

---

# 📦 `bufio` Package in Go

## 🔹 What is `bufio`?

* `bufio` stands for **buffered I/O**.
* It **wraps `io.Reader` and `io.Writer` objects** (like files, network connections, or stdin/stdout).
* Instead of reading/writing **one byte at a time**, it uses a **buffer in memory**, which:

  * **Reduces system calls** (expensive operations).
  * **Improves performance** when working with large data streams.
* Very useful for reading text files line by line, scanning input, or writing lots of small chunks efficiently.

---

## 🔹 Core Types in `bufio`

### 1. **Reader**

* Provides buffered reading.
* Common methods:

  * `Read(p []byte)` – reads into a byte slice.
  * `ReadString(delim byte)` – reads until a delimiter.
  * `ReadBytes(delim byte)` – like `ReadString` but returns `[]byte`.
  * `ReadLine()` – reads a single line (without newline).
  * `Peek(n int)` – looks ahead without consuming.
  * `UnreadByte()` – pushes one byte back.

---

### 2. **Writer**

* Provides buffered writing.
* Instead of writing to the underlying writer each time, it **stores data in a buffer** and writes when:

  * The buffer is full
  * `Flush()` is explicitly called
* Common methods:

  * `Write(p []byte)` – writes byte slice to buffer.
  * `WriteString(s string)` – writes string to buffer.
  * `Flush()` – pushes buffered data to underlying writer.

---

### 3. **Scanner**

* High-level, line-by-line or token-based reader.
* Wraps a `Reader` and **splits input by a scanning function** (default: lines).
* Common methods:

  * `Scan()` – advances to next token (returns `false` at EOF).
  * `Text()` – returns the token as a string.
  * `Bytes()` – returns the token as `[]byte`.
  * `Split(splitFunc)` – defines how to split input (lines, words, custom).

---

## 🔹 Examples

### ✅ Buffered Reader Example

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("sample.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("EOF reached or error:", err)
	}
	fmt.Println("Line read:", line)
}
```

👉 Reads one line at a time from a file.

---

### ✅ Buffered Writer Example

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("output.txt")
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString("Hello, World!\n")
	writer.WriteString("Buffered writing is fast!\n")

	writer.Flush() // must flush, otherwise data stays in buffer
	fmt.Println("Data written to file")
}
```

👉 Writes efficiently and flushes once.

---

### ✅ Scanner Example (Reading Line by Line)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("sample.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("Line:", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
```

👉 Clean and idiomatic way to read text files line by line.

---

### ✅ Scanner Example (Custom Split Function)

```go
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	text := "one,two,three,four"
	scanner := bufio.NewScanner(strings.NewReader(text))

	// Custom split by commas
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if atEOF && len(data) > 0 {
			return len(data), data, nil
		}
		return 0, nil, nil
	})

	for scanner.Scan() {
		fmt.Println("Token:", scanner.Text())
	}
}
```

👉 Shows how to tokenize input with custom rules.

---

## 🔹 When to Use What?

* **Reader** → when you want low-level control (peek, read until a delimiter, etc.).
* **Writer** → when writing lots of small chunks; remember to `Flush()`.
* **Scanner** → when reading line by line or word by word (most common for text files).

---

## ⚡ Important Notes (Do’s ✅ / Don’ts ❌)

### ✅ Do’s

* Use `Scanner` for reading **line by line** or **words**.
* Always **check for errors** (`scanner.Err()`, `reader.Read…()`).
* Always call `writer.Flush()` after writing.
* Prefer `bufio.Writer` for large writes with small chunks (performance boost).

### ❌ Don’ts

* Don’t forget to `Flush()` → data will stay in memory buffer.
* Don’t use `Scanner` for **huge tokens** (it has a default 64KB limit per token; can be increased via `scanner.Buffer()`).
* Don’t assume `ReadLine()` reads the whole line — it may split if too long.

---

## ⚡ Summary

* `bufio.Reader` → efficient buffered input (supports peeking, reading until delimiters).
* `bufio.Writer` → efficient buffered output (must `Flush()`).
* `bufio.Scanner` → easy line/word/token reading, customizable with `Split`.

It’s one of the most **practical packages** in Go — almost every Go project working with files or network streams uses it.

---

**Base64 encoding/decoding** is a very common need (storing/transmitting binary data like images, PDFs, or credentials in JSON, XML, or HTTP).
Go provides full support in the `encoding/base64` package.

---

# 📦 Base64 in Go

## 🔹 What is Base64?

* A way to represent **binary data as ASCII text**.
* Takes 3 bytes (24 bits) of input and encodes them into 4 characters (6 bits each).
* Safe for transmission in:

  * URLs
  * JSON, XML
  * Email (MIME encoding)
* Output always consists of **A–Z, a–z, 0–9, +, /** (or `-` and `_` for URL-safe).

---

## 🔹 Go’s `encoding/base64` Package

The package provides:

* **Standard encoding** → Uses `+` and `/`.
* **URL-safe encoding** → Uses `-` and `_`.
* **WithPadding / NoPadding** → Handles `=` padding.

---

## 🔹 Common Encoders

```go
var StdEncoding = NewEncoding(encodeStd)       // A-Z, a-z, 0-9, +, /
var URLEncoding = NewEncoding(encodeURL)       // A-Z, a-z, 0-9, -, _
var RawStdEncoding = StdEncoding.WithPadding(NoPadding)
var RawURLEncoding = URLEncoding.WithPadding(NoPadding)
```

---

## 🔹 Encoding

### ✅ Example: Encode to Base64 String

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hello, World!"
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", encoded)
}
```

👉 Output:

```
Encoded: SGVsbG8sIFdvcmxkIQ==
```

---

### ✅ URL-Safe Encoding

```go
encoded := base64.URLEncoding.EncodeToString([]byte("https://golang.org?name=go"))
fmt.Println(encoded)
```

👉 Output (no `+` or `/`):

```
aHR0cHM6Ly9nb2xhbmcub3JnP25hbWU9Z28=
```

---

### ✅ No Padding

```go
encoded := base64.RawURLEncoding.EncodeToString([]byte("Hello"))
fmt.Println(encoded)
```

👉 Output (notice no `=` padding):

```
SGVsbG8
```

---

## 🔹 Decoding

### ✅ Decode Base64 String

```go
decoded, err := base64.StdEncoding.DecodeString("SGVsbG8sIFdvcmxkIQ==")
if err != nil {
    panic(err)
}
fmt.Println("Decoded:", string(decoded))
```

👉 Output:

```
Decoded: Hello, World!
```

---

### ✅ Decode URL-Safe Base64

```go
decoded, _ := base64.URLEncoding.DecodeString("aHR0cHM6Ly9nb2xhbmcub3JnP25hbWU9Z28=")
fmt.Println("Decoded:", string(decoded))
```

👉 Output:

```
Decoded: https://golang.org?name=go
```

---

## 🔹 Streaming Encoders/Decoders

Sometimes we don’t want to load all data in memory (large files).
We can **wrap `io.Writer` or `io.Reader`** with Base64.

### ✅ Encode Stream

```go
package main

import (
	"encoding/base64"
	"os"
	"strings"
)

func main() {
	data := "Stream encoding example"
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	defer encoder.Close()

	encoder.Write([]byte(data))
}
```

👉 Output:

```
U3RyZWFtIGVuY29kaW5nIGV4YW1wbGU=
```

---

### ✅ Decode Stream

```go
package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := "U3RyZWFtIGVuY29kaW5nIGV4YW1wbGU="
	decoder := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input))

	buf := make([]byte, 1024)
	n, _ := decoder.Read(buf)

	fmt.Println("Decoded:", string(buf[:n]))
}
```

👉 Output:

```
Decoded: Stream encoding example
```

---

## 🔹 When to Use Which?

| Encoding Type    | Use Case                                  |
| ---------------- | ----------------------------------------- |
| `StdEncoding`    | General encoding/decoding (default).      |
| `URLEncoding`    | Safe for URLs and filenames (no `+` `/`). |
| `RawStdEncoding` | No padding version of standard encoding.  |
| `RawURLEncoding` | No padding, safe for URLs.                |

---

## ⚡ Important Notes

✅ Base64 **increases size** by \~33% (every 3 bytes → 4 bytes).
✅ Always handle **padding `=`** carefully if using `Raw` encodings.
✅ Use **stream encoders/decoders** for large files instead of `EncodeToString()`.
✅ Useful for binary-to-text conversions: images, JWT tokens, cryptographic keys.

---

## 🔹 Summary

* Use `base64.StdEncoding` for default safe encoding.
* Use `base64.URLEncoding` for URLs/filenames.
* Use `NewEncoder`/`NewDecoder` for streaming large data.
* Remember padding rules (`=` vs Raw).

---

Let’s walk through **hashing in Go** end-to-end, from concepts to idiomatic code.

# What hashing is (and why we care)

A **hash function** maps data of any size to a **fixed-size digest** (aka hash). Good cryptographic hashes have these properties:

* **Deterministic**: same input → same output.
* **Preimage resistance**: given a hash, it’s hard to find a message that produces it.
* **Second-preimage resistance**: given one message, it’s hard to find a *different* message with the same hash.
* **Collision resistance**: it’s hard to find *any* two different messages with the same hash.
* **Avalanche**: tiny input changes radically change the output.

We reach for hashing to:

* Verify **integrity** (file checksums, content addressing).
* Create **fingerprints** (deduplication, caching keys).
* Build **MACs** (HMAC) for message authenticity.
* **Never** store passwords as plain text—use **password hashing** (bcrypt/argon2/scrypt), not general-purpose hashing.

# Hashing families in Go

Standard library (safe defaults in bold):

* `crypto/sha256` → **SHA-256** (32-byte digest) – general secure default.
* `crypto/sha512` → SHA-512 (64-byte), plus variants like `Sum512_256` (32-byte).
* `crypto/sha1` → SHA-1 (collision-broken; avoid for security).
* `crypto/md5` → MD5 (collision-broken; only for non-security fingerprints).
* `hash/crc32`, `hash/crc64`, `hash/adler32` → fast **checksums** for accidental error detection (not secure).

Common extra (in `golang.org/x/crypto`, add as a module dep):

* `blake2b`, `blake2s` – fast, secure; keyed mode can replace HMAC.
* `sha3`/`keccak` – modern hash family.
* Password hashing: `bcrypt`, `scrypt`, `argon2`.

# Two APIs we use

Go gives us two idioms:

### 1) One-shot “Sum” helpers (simple)

```go
import (
	"crypto/sha256"
	"fmt"
)

func digestBytes(b []byte) string {
	h := sha256.Sum256(b)         // returns [32]byte
	return fmt.Sprintf("%x", h[:]) // hex encode
}
```

### 2) Streaming via `hash.Hash` (for large data / incremental)

Every `xxx.New()` in `crypto/*` returns a `hash.Hash` which is also an `io.Writer`.

```go
import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"os"
)

func digestFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil { return "", err }
	defer f.Close()

	var h hash.Hash = sha256.New()
	if _, err := io.Copy(h, f); err != nil { // stream the file
		return "", err
	}
	sum := h.Sum(nil)              // appends digest to provided slice (nil → new)
	return hex.EncodeToString(sum), nil
}
```

> Note: `h.Sum(dst)` **appends** the digest to `dst`. We pass `nil` to get just the digest.

# Encoding the digest (hex/base64)

We rarely keep raw bytes. We typically encode as hex:

```go
import "encoding/hex"

hexStr := hex.EncodeToString(sumBytes)
b, err := hex.DecodeString(hexStr)
```

Or base64:

```go
import "encoding/base64"

b64 := base64.StdEncoding.EncodeToString(sumBytes)
raw, _ := base64.StdEncoding.DecodeString(b64)
```

# HMAC (message authentication)

For authenticity + integrity with a shared secret, we use **HMAC**:

```go
import (
	"crypto/hmac"
	"crypto/sha256"
)

func sign(msg, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	return mac.Sum(nil)
}

func verify(msg, key, tag []byte) bool {
	m := hmac.New(sha256.New, key)
	m.Write(msg)
	expected := m.Sum(nil)
	return hmac.Equal(tag, expected) // constant-time compare
}
```

# Password hashing (different from hashing!)

We **must not** use SHA-256/MD5/etc. for passwords. We use **slow, memory-hard** algorithms that include salts and cost parameters.

### Bcrypt (Go’s most common)

```go
import "golang.org/x/crypto/bcrypt"

func hashPassword(pw string) (string, error) {
	// cost 10–14; higher = slower = stronger
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), 14)
	return string(hash), err
}

func checkPassword(pw string, stored string) bool {
	// IMPORTANT: (hashed, password) order
	return bcrypt.CompareHashAndPassword([]byte(stored), []byte(pw)) == nil
}
```

> Bcrypt **embeds a salt and cost** in the hash string; we just store that string.

### Argon2id (modern, memory-hard)

```go
import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

func hashArgon2id(pw string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil { return "", err }

	// Tunable params (example values):
	memKB := uint32(64 * 1024) // 64MB
	iterations := uint32(3)
	parallel := uint8(2)
	keyLen := uint32(32)

	key := argon2.IDKey([]byte(pw), salt, iterations, memKB, parallel, keyLen)
	return fmt.Sprintf("argon2id$%d$%d$%d$%s$%s",
		iterations, memKB, parallel,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(key),
	), nil
}
```

> With Argon2 we store: algorithm, params, salt, derived key. On verify, we parse them back and recompute.

# Choosing an algorithm

* **General integrity / fingerprints**: SHA-256 or SHA-512/256 (good default, widely supported).
* **Fast, modern alternative**: BLAKE2 (via `x/crypto/blake2b/blake2s`).
* **Non-security checks** (e.g., accidental corruption): CRC32/64 (very fast).
* **Passwords**: bcrypt (simple & standard) or Argon2id (preferred when we can manage params).
* **MACs**: HMAC-SHA-256 (interoperable, battle-tested).
* **Avoid** MD5 & SHA-1 for new security designs (collision attacks).

# Common pitfalls (and fixes)

* **Mixing up bcrypt args**: we must pass `(hashed, password)` to `CompareHashAndPassword`. (We hit this earlier!)
* **Reusing `hash.Hash`** across goroutines: instances aren’t safe for concurrent `Write`.
* **Using raw `==`** for MACs: use `hmac.Equal` (constant-time).
* **Hex vs bytes confusion**: printing a byte slice with `%x` is hex; storing that string requires decoding back if we need bytes again.
* **Assuming “hash = encryption”**: hashing is one-way; we can’t “decrypt” a hash.

# Full examples

### 1) Hash a string (one-shot)

```go
data := []byte("hello, hashing")
sum := sha256.Sum256(data)  // [32]byte
fmt.Printf("SHA-256: %x\n", sum)
```

### 2) Hash a large file (streaming)

```go
f, _ := os.Open("big.bin")
defer f.Close()
h := sha512.New()
io.Copy(h, f)
fmt.Printf("SHA-512: %x\n", h.Sum(nil))
```

### 3) HMAC over JSON payload

```go
payload := []byte(`{"ok":true}`)
key := []byte("our-shared-secret")
tag := sign(payload, key)
fmt.Printf("HMAC: %x\n", tag)
```

### 4) Store & verify a password (bcrypt)

```go
stored, _ := hashPassword("s3cret!")
ok := checkPassword("s3cret!", stored)   // true
bad := checkPassword("wrong", stored)    // false
```

# Performance tips

* For repeated small messages, reuse the hasher: `h.Reset()` between uses to reduce allocations.
* For very large inputs, let `io.Copy(h, r)` stream rather than reading whole files into memory.
* SHA-512 can be faster than SHA-256 on 64-bit CPUs; `sha512.Sum512_256` yields a 32-byte digest with SHA-512’s speed.

# Quick reference

* **Hex**: `fmt.Printf("%x", sum)`, `hex.EncodeToString(sum)`.
* **Base64**: `base64.StdEncoding.EncodeToString(sum)`.
* **One-shot**: `sha256.Sum256(b)`.
* **Stream**: `h := sha256.New(); h.Write(...); h.Sum(nil)`.
* **HMAC**: `mac := hmac.New(sha256.New, key)`.
* **Passwords**: `bcrypt.GenerateFromPassword`, `bcrypt.CompareHashAndPassword`.

---
# 🔐 **Hashing & Salting Explained in Go**

### ✅ 1. **Hashing**

* **What it is:**
  A **one-way function** that takes an input (password, data, file) and produces a **fixed-length output** (hash digest).

  * Same input → always same output.
  * Irreversible (cannot get the original back).
  * Used for **password storage, data integrity, digital signatures**.

* In your code:

  ```go
  hashed := sha256.Sum256([]byte(password))
  hashed512 := sha512.Sum512([]byte(password))
  ```

  * `sha256.Sum256` always gives a 32-byte array (256 bits).
  * `sha512.Sum512` always gives a 64-byte array (512 bits).

---

### ✅ 2. **Salting**

* **Why salt passwords?**

  * Without salt: attackers can use **rainbow tables** (precomputed hashes for common passwords) or **dictionary attacks**.
  * With salt: each password hash is unique, even if two users have the same password.

* **What your `GenSalt()` does:**

  ```go
  func GenSalt() ([]byte, error) {
      salt := make([]byte, 16) // 16 random bytes
      _, err := io.ReadFull(rand.Reader, salt)
      return salt, err
  }
  ```

  * Uses **`crypto/rand`** (cryptographically secure RNG, not predictable).
  * Generates a **random 16-byte salt**.
  * This salt is **unique per user** and stored **alongside the hash in the database** (not secret!).

---

### ✅ 3. **Hashing with Salt**

* Your function:

  ```go
  func HashPassword(password string, salt []byte) string {
      saltedPassword := append(salt, []byte(password)...)
      hash := sha256.Sum256(saltedPassword)
      return base64.StdEncoding.EncodeToString(hash[:])
  }
  ```

  Steps:

  1. Append salt to the password (`salt || password`).
  2. Compute hash.
  3. Encode it in Base64 (easier for storage/transport, DB safe).

* Example:

  ```
  password = "password123"
  salt = [16 random bytes]

  Hash("password123", salt) -> "f97a...4c8"
  ```

  Now even if two users both have "password123", their hashes differ because their salts differ.

---

### ✅ 4. **Verification**

* During login:

  * Retrieve stored `salt` + `hash`.
  * Re-hash the incoming password with the same salt.
  * Compare with stored hash.

In your code:

```go
if signupHash == loginHash {
    fmt.Println("Password is correct. ✅")
} else {
    fmt.Println("Login failed. ⚠️")
}
```

---

# 🧠 **Quirks & Best Practices (Junior Engineers Must Know)**

### ⚡ **1. Salts are Public, but Must Be Unique**

* **Salts are NOT secrets.**
  They should be stored in plaintext in the DB alongside the hash.
* What matters is that salts are **long and random** to prevent precomputed attacks.

✅ DO:

```sql
user_id | salt (base64)      | password_hash
-------------------------------------------------------
1       | F7hsH2lk9fU=       | sP94kUd+u9tK8s...
```

---

### ⚡ **2. Don’t Use Raw SHA256/512 for Passwords**

* SHA256 is **fast** → attackers can brute-force billions of guesses per second on GPUs.
* For passwords, we should use **slow hashing functions** designed for passwords:

  * `bcrypt` (Go: `golang.org/x/crypto/bcrypt`)
  * `scrypt`
  * `Argon2`

⚠️ Your current code is good for **learning**, but in production use **bcrypt or Argon2**.

---

### ⚡ **3. Difference Between `math/rand` and `crypto/rand`**

* `math/rand`: deterministic, needs seeding, predictable → ❌ insecure.
* `crypto/rand`: cryptographically secure randomness → ✅ always use for salts, keys.

---

### ⚡ **4. Base64 Encoding**

* You correctly store the salt and hash as Base64 strings.
* Never store raw bytes directly in DBs unless the column supports it (`BLOB`).

---

### ⚡ **5. Timing Attacks**

* Use `subtle.ConstantTimeCompare()` instead of `==` to avoid **timing attacks**:

  ```go
  import "crypto/subtle"

  if subtle.ConstantTimeCompare([]byte(signupHash), []byte(loginHash)) == 1 {
      fmt.Println("Password correct ✅")
  }
  ```

  ⚠️ `==` can leak timing differences to attackers.

---

### ⚡ **6. Length of Salt**

* Common: 16–32 bytes.
* More is unnecessary but not harmful.

---

### ⚡ **7. Peppering (Optional)**

* A "pepper" is a **secret key** stored outside the DB (like in environment variables).
* Workflow:

  ```
  hash(password + pepper + salt)
  ```

  * Even if DB leaks, without pepper, attacker can’t easily crack passwords.

---

# 📝 **Summary**

* Hashing = irreversible, fixed-size fingerprint.
* Salting = unique randomness added to password → prevents rainbow/dictionary attacks.
* Always use `crypto/rand` for secure salt generation.
* Store **salt + hash** together.
* **Never use SHA256/512 directly for real password storage** → use bcrypt/argon2.
* Use **constant-time comparison** to avoid timing attacks.
* Encode results (`hex` or `base64`) before storing.

---

## 🔹 1. What is Salting?

Salting is the process of **adding random data (salt)** to a password **before hashing** it.

* If two users both choose the same password (`"password123"`), **without salting** they’ll have the same hash — dangerous because attackers could spot duplicate passwords in your DB.
* With salting, each password gets a **unique hash**, because the salt is different each time.

👉 Example:

```
Password: "password123"
Salt: 4fj92k... (random 16 bytes)
Salted: "4fj92k...password123"
Hash: (SHA256 of above)
```

Two users with the same password will still produce different hashes because their salts are different.

---

## 🔹 2. Oour `GenSalt()` function

```go
func GenSalt() ([]byte, error) {
    salt := make([]byte, 16) // 16 bytes
    _, err := io.ReadFull(rand.Reader, salt)
    return salt, err
}
```

✅ This is solid because:

* Uses `crypto/rand` (cryptographically secure RNG).
* **16 bytes** of randomness → 128 bits of entropy (enough for password salts).
* Each call generates a new, unpredictable salt.

⚠️ **Quirk:** Some juniors mistakenly use `math/rand` instead of `crypto/rand`. That’s not secure because `math/rand` is deterministic if seed is known.

---

## 🔹 3. Our `HashPassword()` function

```go
func HashPassword(password string, salt []byte) string {
    saltedPassword := append(salt, []byte(password)...)
    hash := sha256.Sum256(saltedPassword)
    return base64.StdEncoding.EncodeToString(hash[:])
}
```

### Step-by-step:

1. **Salt + Password concatenation:**

   ```
   saltedPassword := append(salt, []byte(password)...)
   ```

   This prepends salt to the password before hashing.

2. **Hashing:**

   ```
   hash := sha256.Sum256(saltedPassword)
   ```

   Generates a **fixed 256-bit hash** irrespective of input size.

3. **Encoding to Base64:**

   ```
   return base64.StdEncoding.EncodeToString(hash[:])
   ```

   Encodes hash in Base64 → easier to store/compare in DB.

---

## 🔹 4. Login Verification

On signup:

* Store **`salt`** and **`hash(password+salt)`** in the DB.

On login:

* Retrieve the stored salt.
* Recompute `hash(inputPassword + storedSalt)`.
* Compare with stored hash.

👉 If they match → password is correct.

---

## 🔹 5. Important Quirks for Junior Engineers

1. **Don’t invent your own crypto.**
   Always use proven libraries like `bcrypt`, `scrypt`, or `Argon2`. SHA-256 is okay for learning but weak for production (fast → brute-forceable).

2. **Always store the salt with the hash.**
   Salt is not secret! It just prevents pre-computed attacks (rainbow tables). Storing salt in plain Base64 is fine.

3. **Comparison must be constant-time.**
   Don’t use `==` in production for comparing hashes (can lead to timing attacks). Instead:

   ```go
   import "crypto/subtle"
   subtle.ConstantTimeCompare(hash1, hash2)
   ```

4. **Iteration counts (key stretching).**
   SHA-256 once is too fast. Attackers can try billions of passwords per second.
   Use functions that **intentionally slow hashing down**:

   * `bcrypt`
   * `scrypt`
   * `Argon2id`

   Example with bcrypt:

   ```go
   import "golang.org/x/crypto/bcrypt"

   hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
   bcrypt.CompareHashAndPassword(hashed, []byte(password))
   ```

5. **Peppering (optional extra security).**
   A **pepper** is a secret key (stored in app, not DB). Example:

   ```
   hash = sha256(salt + password + pepper)
   ```

   If DB leaks, attacker still needs the pepper.

---

## 🔹 6. Common Mistakes Juniors Make

* ❌ Using `math/rand` for salt.
* ❌ Re-using the same salt for all users.
* ❌ Storing only the hash, forgetting the salt.
* ❌ Using plain SHA256 without key stretching (bcrypt/Argon2).
* ❌ Returning detailed error messages like *“password incorrect”* vs *“username incorrect”*. (Helps attackers).

---

✅ **Summary:**
Our implementation is **great for learning** but in **production** we should use `bcrypt`, `scrypt`, or `Argon2`.
Salting prevents rainbow-table attacks. Key stretching prevents brute-force attacks. Constant-time comparison prevents timing attacks.

---

Let’s **upgrade our code** into a **production-ready version** using `bcrypt` (from `golang.org/x/crypto/bcrypt`).

This will include:

* Signup flow (hash + store password securely).
* Login flow (compare password with stored hash).
* Explanations of why this is secure.

---

# 🔹 Production-Ready Password Hashing in Go (with `bcrypt`)

```go
package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ===============================
// SIGNUP: Hash the password
// ===============================
func HashPassword(password string) (string, error) {
	// bcrypt.DefaultCost = 10 (safe default).
	// Higher cost = more secure but slower (range: 4–31).
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ===============================
// LOGIN: Compare hash with input
// ===============================
func CheckPasswordHash(password, hashedPassword string) bool {
	// bcrypt.CompareHashAndPassword returns nil if match.
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func main() {
	// ---- SIGNUP ----
	password := "SuperSecret123!"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("🔐 Stored hash in DB:", hashedPassword)

	// ---- LOGIN ----
	loginPassword := "SuperSecret123!" // User input
	match := CheckPasswordHash(loginPassword, hashedPassword)

	if match {
		fmt.Println("✅ Password is correct, login success!")
	} else {
		fmt.Println("❌ Invalid password!")
	}
}
```

---

## 🔹 How it Works

1. **Hashing on Signup**

   ```go
   bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
   ```

   * `bcrypt` automatically generates a **unique salt for each password** (you don’t need to store it separately).
   * The salt is **embedded inside the hash string**.
   * `bcrypt.DefaultCost = 10` is secure. Increase to `12` or `14` for stronger hashing (but slower).

   Example stored hash (looks like gibberish but includes salt + cost):

   ```
   $2a$10$N9qo8uLOickgx2ZMRZoMyeIjZK0s8AjtKq6HgMHqYFQW1Ql9kZC8W
   ```

2. **Verification on Login**

   ```go
   bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
   ```

   * Extracts the salt from the stored hash.
   * Re-hashes the input password with that salt.
   * Constant-time comparison → prevents timing attacks.

3. **No need to manage salts manually** 👍

---

## 🔹 Advantages of `bcrypt`

✅ Automatically salts passwords.
✅ Adjustable cost factor → slows brute-force attacks.
✅ Resistant to rainbow-table attacks.
✅ Constant-time comparison for security.

---

## 🔹 Signup + Login Simulation

### Signup

* Input: `"SuperSecret123!"`
* Store in DB:

  ```
  $2a$10$3lK1O0t9WqL.jf9qV5sQvO9iM8f42o8mD8B6P6WzvHbR1b6OJh4tC
  ```

### Login

* User enters `"SuperSecret123!"`.
* App checks with `bcrypt.CompareHashAndPassword`.
* ✅ Match → Login successful.
* ❌ Wrong password → Login rejected.

---

## 🔹 Best Practices for Devs.

* Always hash passwords before storing. Never store plain-text passwords.
* Use **bcrypt/Argon2id** (not SHA256 directly).
* Store the whole hash string (bcrypt handles salt internally).
* Use `bcrypt.DefaultCost` or `bcrypt.MinCost` in dev only; raise cost in production.
* Never show “username OR password incorrect” separately — always show a generic “invalid login” error.

---

In Go, working with files (reading/writing) revolves around the `os`, `io`, `ioutil` (deprecated, but still seen), and `bufio` packages.

---

# 🔹 1. Basics: Files in Go

* Files are treated as a resource with a **file descriptor**.
* Before reading or writing, we must **open** a file using `os.Open` (read-only) or `os.OpenFile` (custom flags).
* Always `defer file.Close()` after opening — prevents leaks.

---

# 🔹 2. Opening a File

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt") // read-only
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close() // ✅ Always close

	fmt.Println("File opened successfully:", file.Name())
}
```

---

# 🔹 3. Reading from a File

## (a) Read Whole File at Once

```go
data, err := os.ReadFile("example.txt") // replaces ioutil.ReadFile in Go 1.16+
if err != nil {
	log.Fatal(err)
}
fmt.Println(string(data)) // Convert []byte → string
```

✅ Simple, but loads the entire file into memory (not good for very large files).

---

## (b) Read with `bufio.Scanner` (line by line)

```go
file, _ := os.Open("example.txt")
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
	fmt.Println(scanner.Text()) // each line
}

if err := scanner.Err(); err != nil {
	fmt.Println("Error:", err)
}
```

✅ Best for reading logs, config files, etc.

---

## (c) Read with `file.Read()`

```go
buffer := make([]byte, 64) // read 64 bytes
n, err := file.Read(buffer)
if err != nil && err != io.EOF {
	log.Fatal(err)
}
fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
```

✅ Good for streaming raw data (binary files, large files).

---

# 🔹 4. Writing to a File

## (a) Write Whole File

```go
data := []byte("Hello, Golang!\n")
err := os.WriteFile("output.txt", data, 0644) // perms: rw-r--r--
if err != nil {
	log.Fatal(err)
}
```

---

## (b) Append or Create with `os.OpenFile`

```go
file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Fatal(err)
}
defer file.Close()

if _, err := file.WriteString("Appending new line!\n"); err != nil {
	log.Fatal(err)
}
```

Flags:

* `os.O_CREATE` → create if not exists.
* `os.O_APPEND` → append instead of overwrite.
* `os.O_WRONLY` → write-only.
* `os.O_RDWR` → read + write.
* `os.O_TRUNC` → truncate existing file.

---

## (c) Buffered Writer

```go
file, _ := os.Create("buffered.txt")
defer file.Close()

writer := bufio.NewWriter(file)
writer.WriteString("First line\n")
writer.WriteString("Second line\n")
writer.Flush() // ✅ must flush buffer
```

✅ Faster when writing lots of small data.

---

# 🔹 5. File Metadata

```go
info, err := os.Stat("example.txt")
if err != nil {
	log.Fatal(err)
}
fmt.Println("Name:", info.Name())
fmt.Println("Size:", info.Size())
fmt.Println("IsDir:", info.IsDir())
fmt.Println("Permissions:", info.Mode())
fmt.Println("Modified:", info.ModTime())
```

---

# 🔹 6. File Paths

```go
import "path/filepath"

abs, _ := filepath.Abs("example.txt")
fmt.Println("Absolute path:", abs)
```

---

# 🔹 7. Example: Copy File

```go
src, _ := os.Open("example.txt")
defer src.Close()

dst, _ := os.Create("copy.txt")
defer dst.Close()

_, err := io.Copy(dst, src) // copies contents
if err != nil {
	log.Fatal(err)
}
```

---

# 🔹 8. Best Practices (Jr. Engineers 🚀)

1. Always `defer file.Close()`.
2. Use `os.ReadFile` / `os.WriteFile` for **small files**.
3. Use `bufio.Scanner` / `io.Reader` for **large files**.
4. When writing in loops → use `bufio.Writer` + `Flush()`.
5. Handle errors properly (don’t ignore them).
6. Remember file permissions (e.g., `0644`).

---

✅ TL;DR

* **Read small files** → `os.ReadFile`.
* **Read big files** → `bufio.Scanner` or `Read()`.
* **Write small files** → `os.WriteFile`.
* **Write big files / appending** → `os.OpenFile` + `bufio.Writer`.
* **Always close files**.

---

## 🔹 1. What is a Line Filter?

A **line filter** is a program that:

* Reads input **line by line** (from a file, stdin, or network).
* Processes/transforms each line.
* Writes the result **line by line** (to stdout, file, etc.).

They’re super common in Unix-like systems (e.g., `grep`, `sort`, `uniq`) and Go makes it easy to implement them.

---

## 🔹 2. Why Line Filters in Go?

* Go has strong support for **concurrent I/O** and **streaming text**.
* Perfect for building small tools that process data efficiently.
* Works well in **pipelines**: you can run your Go program and pipe input/output with other commands.

---

## 🔹 3. Building Blocks in Go

To make a line filter in Go, we typically use:

* **`os.Stdin`** → Read input from standard input (keyboard or piped data).
* **`bufio.Scanner`** → Efficient line-by-line reading.
* **`os.Stdout`** or `fmt.Println` → Write output.

---

## 🔹 4. Example: Uppercase Line Filter

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a scanner to read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Process line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Transform (make uppercase)
		upperLine := strings.ToUpper(line)

		// Write to stdout
		fmt.Println(upperLine)
	}

	// Handle errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
```

### ▶ Usage

```bash
echo "hello world" | go run main.go
```

Output:

```
HELLO WORLD
```

---

## 🔹 5. Example: Filter Specific Lines (like `grep`)

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		// Print only lines containing "go"
		if strings.Contains(line, "go") {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}
}
```

Usage:

```bash
cat file.txt | go run main.go
```

---

## 🔹 6. Quirks / Best Practices

* Always check `scanner.Err()` after looping.
* `bufio.Scanner` has a **64 KB default buffer**. For very large lines, you may need `scanner.Buffer()` to increase it.
* Use `os.Stdin` and `os.Stdout` to make the filter **composable in pipelines** (Unix philosophy).
* Keep transformations **stateless per line** if possible → makes them faster and memory efficient.

---

✅ So in short:

* A **line filter in Go** = read line → process line → output line.
* Use `bufio.Scanner` + `os.Stdin` + `os.Stdout`.
* They’re Go’s way of building Unix-style small programs.

---

Perfect 👍 let’s dig into **file paths in Go** in detail.

Go provides tools in its standard library (`path`, `path/filepath`, and `os`) for **working with file paths** across operating systems.

---

# 🔹 1. Paths in Go

* A **path** is just a string representing the location of a file or directory.
* Paths differ between OS:

  * **Unix/Linux/macOS** → `/home/skyy/docs/file.txt`
  * **Windows** → `C:\Users\Skyy\docs\file.txt`
* Go’s `path/filepath` package automatically handles these OS differences.

---

# 🔹 2. Packages for Paths

### ✅ `path` package

* Works only with **forward-slash (`/`) separated paths**.
* Mainly used for **URLs** and **virtual paths** (not actual filesystem paths).

### ✅ `path/filepath` package

* OS-aware → uses the correct separator (`/` or `\`).
* Should be used for **filesystem paths**.

---

# 🔹 3. Common Functions in `path/filepath`

### 🔸 `filepath.Join`

Safely concatenates path parts using the correct separator.

```go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("home", "skyy", "docs", "file.txt")
	fmt.Println("Path:", path)
}
```

* Linux/macOS → `home/skyy/docs/file.txt`
* Windows → `home\skyy\docs\file.txt`

---

### 🔸 `filepath.Dir` and `filepath.Base`

```go
p := "/home/skyy/docs/file.txt"

fmt.Println(filepath.Dir(p))   // "/home/skyy/docs"
fmt.Println(filepath.Base(p))  // "file.txt"
```

---

### 🔸 `filepath.Ext`

Get file extension.

```go
fmt.Println(filepath.Ext("report.pdf")) // ".pdf"
```

---

### 🔸 `filepath.Abs`

Get the absolute path (resolves relative paths).

```go
absPath, _ := filepath.Abs("file.txt")
fmt.Println("Absolute Path:", absPath)
```

---

### 🔸 `filepath.Clean`

Cleans up redundant `.` or `..` or slashes.

```go
p := filepath.Clean("/home/skyy/../docs//file.txt")
fmt.Println(p) // "/docs/file.txt"
```

---

### 🔸 `filepath.Split`

Splits directory + file.

```go
dir, file := filepath.Split("/home/skyy/docs/file.txt")
fmt.Println("Dir:", dir)   // "/home/skyy/docs/"
fmt.Println("File:", file) // "file.txt"
```

---

### 🔸 `filepath.WalkDir` (Traverse Directory Tree)

```go
import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	root := "./"

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println("Found:", path)
		return nil
	})
}
```

This walks through every file and folder from the given root.

---

# 🔹 4. OS-Specific Path Handling

### Path Separator

```go
fmt.Println("Separator:", string(filepath.Separator))
// Linux/macOS → "/"
// Windows → "\"
```

### List Split

```go
list := filepath.SplitList("/usr/bin:/bin:/usr/local/bin")
fmt.Println(list) // ["/usr/bin", "/bin", "/usr/local/bin"]
```

---

# 🔹 5. Checking Files/Directories with `os`

Often we combine `os` with `filepath`.

```go
import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Name:", info.Name())
	fmt.Println("IsDir:", info.IsDir())
}
```

---

# 🔹 6. Real-World Example

Suppose we want to **get all `.txt` files** in a directory:

```go
package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	matches, _ := filepath.Glob("*.txt")
	for _, file := range matches {
		fmt.Println("Found file:", file)
	}
}
```

---

✅ **Summary**:

* Use **`path/filepath`** for filesystem paths (OS-aware).
* Use **`path`** for slash-separated paths (URLs).
* Functions like `Join`, `Dir`, `Base`, `Ext`, `Abs`, `Clean`, `Split`, `WalkDir`, and `Glob` are key tools.

---

Let’s go **step by step** into **directories in Go (Golang)**.
Working with directories is very common in real-world apps — reading, creating, navigating, walking, deleting, etc. Go’s standard library (`os`, `io/ioutil` \[deprecated], `path/filepath`) gives us all the tools.

---

# 🔹 1. What is a Directory in Go?

* A **directory** is a container for files and subdirectories.
* Directories can be:

  * **Absolute path**: `/home/skyy/projects`
  * **Relative path**: `./projects` (relative to where the program runs).
* Go treats directories as **first-class filesystem objects** (like files).

---

# 🔹 2. Creating Directories

We use **`os.Mkdir`** and **`os.MkdirAll`**.

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Create single directory
	err := os.Mkdir("mydir", 0755) // rwxr-xr-x
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create nested directories
	err = os.MkdirAll("parent/child/grandchild", 0755)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Directories created ✅")
}
```

* **`Mkdir`** → creates a single directory. Fails if parent doesn’t exist.
* **`MkdirAll`** → creates nested directories, like `mkdir -p` in Linux.
* **Permissions**: `0755` = `rwxr-xr-x`.

---

# 🔹 3. Removing Directories

```go
// Remove empty dir
err := os.Remove("mydir")

// Remove dir with contents (recursive delete)
err := os.RemoveAll("parent")
```

* `Remove` → works only if directory is empty.
* `RemoveAll` → deletes everything inside (use carefully ⚠️).

---

# 🔹 4. Changing Directory

```go
err := os.Chdir("parent")
if err != nil {
	fmt.Println("Error changing dir:", err)
}
```

* Changes the **current working directory** of the running program.
* Useful if we want to work relative to another path.

---

# 🔹 5. Getting Current Directory

```go
cwd, _ := os.Getwd()
fmt.Println("Current working dir:", cwd)
```

---

# 🔹 6. Reading Directory Contents

We can **list files** inside a directory.

### Method 1: `os.ReadDir` (Go 1.16+ preferred)

```go
entries, _ := os.ReadDir(".")
for _, e := range entries {
	if e.IsDir() {
		fmt.Println("Dir:", e.Name())
	} else {
		fmt.Println("File:", e.Name())
	}
}
```

### Method 2: `os.Open` + `Readdir`

```go
f, _ := os.Open(".")
defer f.Close()

files, _ := f.Readdir(-1) // -1 → all entries
for _, file := range files {
	fmt.Println(file.Name(), file.IsDir())
}
```

---

# 🔹 7. Traversing Directories (Recursive Walk)

### `filepath.WalkDir`

```go
import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	root := "./"

	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			fmt.Println("Directory:", path)
		} else {
			fmt.Println("File:", path)
		}
		return nil
	})
}
```

* Walks **all files and subdirectories** starting from `root`.
* We can filter by extension, size, etc.

---

# 🔹 8. Checking if Path is Directory

```go
info, err := os.Stat("parent")
if err != nil {
	fmt.Println("Error:", err)
	return
}

if info.IsDir() {
	fmt.Println("It's a directory ✅")
} else {
	fmt.Println("It's a file 📄")
}
```

---

# 🔹 9. Directory Permissions

* Directories in Go follow Unix permissions:

  * **Read (r)** → list files inside.
  * **Write (w)** → create/remove files inside.
  * **Execute (x)** → enter the directory.
* Example: `0755` → `rwxr-xr-x`.

---

# 🔹 10. Example: Copying All Files from One Directory to Another

```go
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	src := "parent"
	dst := "backup"

	os.MkdirAll(dst, 0755)

	filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			// open source file
			in, _ := os.Open(path)
			defer in.Close()

			// create destination file
			outPath := filepath.Join(dst, filepath.Base(path))
			out, _ := os.Create(outPath)
			defer out.Close()

			io.Copy(out, in)
			fmt.Println("Copied:", path, "→", outPath)
		}
		return nil
	})
}
```

This copies all files from `parent` to `backup`.

---

# 🔹 11. Summary

✅ **Creating** → `os.Mkdir`, `os.MkdirAll`
✅ **Removing** → `os.Remove`, `os.RemoveAll`
✅ **Navigating** → `os.Chdir`, `os.Getwd`
✅ **Listing** → `os.ReadDir`, `os.Open` + `Readdir`
✅ **Walking** → `filepath.WalkDir` (recursive)
✅ **Checking** → `os.Stat().IsDir()`
✅ **Permissions** → Unix style (`0755`, `0777`, etc.)

---

Let’s build a **Directory Handling Cheat Sheet in Go**.
This will be a **quick reference table** for everything we covered — creation, navigation, deletion, listing, walking, etc.

---

# 📂 Golang Directory Handling Cheat Sheet

| **Action**                | **Function / Code**                                                                   | **Notes**                                             |
| ------------------------- | ------------------------------------------------------------------------------------- | ----------------------------------------------------- |
| **Create a directory**    | `os.Mkdir("mydir", 0755)`                                                             | Creates a single directory (fails if parent missing). |
| **Create nested dirs**    | `os.MkdirAll("a/b/c", 0755)`                                                          | Like `mkdir -p`. Creates parents automatically.       |
| **Remove empty dir**      | `os.Remove("mydir")`                                                                  | Works only if directory is empty.                     |
| **Remove dir + contents** | `os.RemoveAll("mydir")`                                                               | Recursive delete (⚠ dangerous).                       |
| **Change current dir**    | `os.Chdir("mydir")`                                                                   | Sets working dir of the program.                      |
| **Get current dir**       | `cwd, _ := os.Getwd()`                                                                | Returns absolute working dir.                         |
| **Check if path is dir**  | `info, _ := os.Stat("path"); info.IsDir()`                                            | Returns `true` if path is directory.                  |
| **List files in dir**     | `entries, _ := os.ReadDir("path")`                                                    | Returns slice of `DirEntry`.                          |
| **List (old way)**        | `f, _ := os.Open("path"); f.Readdir(-1)`                                              | Legacy, returns `FileInfo`.                           |
| **Read dir entry names**  | `for _, e := range entries { fmt.Println(e.Name()) }`                                 | Works with `os.ReadDir`.                              |
| **Walk directory tree**   | `filepath.WalkDir("root", func(path string, d fs.DirEntry, err error) error { ... })` | Recursively visits all dirs/files.                    |
| **Get absolute path**     | `abs, _ := filepath.Abs("mydir")`                                                     | Converts relative → absolute path.                    |
| **Join paths**            | `filepath.Join("a","b","c")`                                                          | Cross-platform safe path joining.                     |
| **Split path**            | `dir, file := filepath.Split("/a/b/c.txt")`                                           | Splits dir and filename.                              |
| **Get parent dir**        | `filepath.Dir("/a/b/c.txt")`                                                          | Returns `/a/b`.                                       |
| **Get base (last elem)**  | `filepath.Base("/a/b/c.txt")`                                                         | Returns `c.txt`.                                      |
| **Get extension**         | `filepath.Ext("file.txt")`                                                            | Returns `.txt`.                                       |

---

# 🔹 Example: Quick Directory Utility

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Create
	os.MkdirAll("parent/child", 0755)

	// Current working dir
	cwd, _ := os.Getwd()
	fmt.Println("CWD:", cwd)

	// List directory
	entries, _ := os.ReadDir(".")
	for _, e := range entries {
		if e.IsDir() {
			fmt.Println("Dir:", e.Name())
		} else {
			fmt.Println("File:", e.Name())
		}
	}

	// Walk tree
	filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if d.IsDir() {
			fmt.Println("[DIR]", path)
		} else {
			fmt.Println("[FILE]", path)
		}
		return nil
	})
}
```

---

Let’s dive into **Temporary Files & Directories in Go**.
These are very useful when we need scratch space for intermediate data, caching, or testing.

---

# 📂 Temporary Files and Directories in Go

Go provides support for **temporary files and dirs** via the `os` and `io/ioutil` (deprecated, use `os` or `os.CreateTemp`) packages.

---

## 🔹 1. Creating a Temporary File

We use `os.CreateTemp`:

```go
file, err := os.CreateTemp("", "example-*.txt")
if err != nil {
    panic(err)
}
defer os.Remove(file.Name()) // cleanup
defer file.Close()

fmt.Println("Temp file created:", file.Name())

file.WriteString("Hello, temp file!")
```

### Explanation:

* **`os.CreateTemp(dir, pattern)`**

  * `dir`: Directory to create file in. If `""`, defaults to system temp dir (`os.TempDir()`).
  * `pattern`: Name pattern. `*` is replaced with random string (ensures uniqueness).
* Returns a `*os.File`.
* File is created **opened for writing**.
* We should `defer file.Close()` after use.
* Usually also `defer os.Remove(file.Name())` for cleanup.

👉 Example result:
`/tmp/example-12345.txt` (Linux/Mac) or `%TEMP%\example-12345.txt` (Windows).

---

## 🔹 2. Creating a Temporary Directory

We use `os.MkdirTemp`:

```go
dir, err := os.MkdirTemp("", "mytempdir-*")
if err != nil {
    panic(err)
}
defer os.RemoveAll(dir) // cleanup entire dir

fmt.Println("Temp dir created:", dir)

// Create file inside temp dir
fpath := filepath.Join(dir, "file.txt")
os.WriteFile(fpath, []byte("data inside temp dir"), 0644)
```

### Explanation:

* **`os.MkdirTemp(dir, pattern)`**

  * `dir`: Parent dir for new temp dir. `""` = system temp.
  * `pattern`: Prefix + random suffix.
* Returns **absolute path** of new dir.
* We usually clean with `os.RemoveAll(dir)`.

---

## 🔹 3. Get System Temp Directory

```go
fmt.Println("System temp directory:", os.TempDir())
```

* Returns system-defined temp location:

  * Linux/Mac → `/tmp`
  * Windows → `%TEMP%`

---

## 🔹 4. Cleaning Up

* Always cleanup temp files/dirs after use.
* Use:

  * `os.Remove(file.Name())` → remove file
  * `os.RemoveAll(dir)` → remove dir recursively

⚠️ Best practice: Always `defer` cleanup immediately after creation.

---

## 🔹 5. Practical Example: Temp File + Dir Workflow

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// 1. Create temp dir
	dir, _ := os.MkdirTemp("", "session-*")
	defer os.RemoveAll(dir)
	fmt.Println("Temp dir:", dir)

	// 2. Create temp file inside dir
	file, _ := os.CreateTemp(dir, "data-*.txt")
	defer file.Close()
	fmt.Println("Temp file:", file.Name())

	// 3. Write into temp file
	file.WriteString("Temporary content here")

	// 4. Read back
	content, _ := os.ReadFile(file.Name())
	fmt.Println("File content:", string(content))
}
```

✅ Output:

```
Temp dir: /tmp/session-45672
Temp file: /tmp/session-45672/data-8374.txt
File content: Temporary content here
```

---

## 🔹 6. Common Use Cases

* **Testing**: Create isolated scratch files/directories.
* **Caching**: Store temporary downloads, processing data.
* **Intermediate results**: During file conversions, compression, etc.
* **Cross-platform safety**: `os.TempDir()` handles platform-specific paths.

---

# 📝 Summary

* `os.CreateTemp("", "prefix-*")` → Temp file.
* `os.MkdirTemp("", "prefix-*")` → Temp dir.
* `os.TempDir()` → Default system temp location.
* Always cleanup: `os.Remove()` or `os.RemoveAll()`.

---

The **`//go:embed` directive** in Go is one of the coolest features introduced in **Go 1.16**, making it possible to include static files **directly inside Go binaries** without shipping separate assets. Let’s break it down thoroughly. 🚀

---

# 📦 The `embed` Package in Go

## 🔹 What is `//go:embed`?

* `//go:embed` is a **compiler directive**.
* It allows us to embed **files or folders** into the compiled Go program.
* Managed via the **`embed` standard package**.
* Very useful for embedding **config files, templates, static assets, migrations, etc.**

---

## 🔹 1. Importing the Package

```go
import _ "embed"
```

⚠️ Notice the underscore `_`.
We don’t use functions from the package directly, but it enables the `//go:embed` directive.

---

## 🔹 2. Embedding a Single File

```go
package main

import (
	_ "embed"
	"fmt"
)

//go:embed hello.txt
var message string

func main() {
	fmt.Println(message)
}
```

### Explanation:

* `hello.txt` content will be **compiled into the binary**.
* Variable type can be:

  * `string` → file content as text
  * `[]byte` → raw binary data
* Here, if `hello.txt` contains `Hello, Go!`, running program prints:

  ```
  Hello, Go!
  ```

---

## 🔹 3. Embedding Binary Data

```go
//go:embed image.png
var imageBytes []byte

func main() {
	fmt.Println("Image size:", len(imageBytes))
}
```

✅ Useful for embedding logos, PDFs, images, etc.

---

## 🔹 4. Embedding Multiple Files

```go
//go:embed config.json schema.sql notes.txt
var files embed.FS
```

* `embed.FS` represents a **virtual filesystem** of embedded files.
* We can use `ReadFile`:

```go
data, _ := files.ReadFile("config.json")
fmt.Println(string(data))
```

---

## 🔹 5. Embedding Entire Directories

```go
//go:embed static/*
var staticFiles embed.FS
```

* Embeds all files inside `static/` directory.
* Access:

```go
data, _ := staticFiles.ReadFile("static/index.html")
fmt.Println(string(data))
```

---

## 🔹 6. Using `embed.FS` like `io/fs`

* `embed.FS` implements the `fs.FS` interface from `io/fs`.
* Which means:

  * Can be used with functions like `fs.ReadDir`, `fs.Glob`, etc.
  * Works with `http.FileServer` for serving static sites.

Example (serve website with embedded HTML/JS/CSS):

```go
package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed static/*
var content embed.FS

func main() {
	subFS, _ := fs.Sub(content, "static")
	http.Handle("/", http.FileServer(http.FS(subFS)))
	http.ListenAndServe(":8080", nil)
}
```

✅ This serves files from `static/` as a web server without shipping extra files.

---

## 🔹 7. Rules & Limitations

1. **Directive position**:

   * Must appear **immediately before** the variable declaration.
2. **Variable types allowed**:

   * `string`
   * `[]byte`
   * `embed.FS`
3. **Build-time resolution**:

   * Files must exist at **build time**.
   * If not, build fails.
4. **File names are relative** to the source file containing the directive.
5. **Globs allowed**:

   * `*` (matches files in directory, not subdirectories).
   * `static/*` works.
   * `static/**` (recursive glob) ❌ not supported.

---

## 🔹 8. Practical Use Cases

* 📜 **Configuration files**: ship default configs in binary.
* 🎨 **Templates**: embed HTML templates for web servers.
* 🌍 **Static websites**: serve JS/CSS/HTML from Go binary.
* 🗄️ **Database migrations**: embed `.sql` migration files.
* 📦 **Single-binary apps**: no external dependencies required.

---

## 🔹 9. Example: Templates with Embed

```go
package main

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*
var tmplFS embed.FS

func main() {
	tmpl := template.Must(template.ParseFS(tmplFS, "templates/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", map[string]string{"Name": "Skyy"})
	})

	http.ListenAndServe(":8080", nil)
}
```

✅ The `templates/` directory is embedded into the binary.

---

# 📝 Summary (Cheat Sheet)

| Directive                | Type Allowed | Purpose                      |
| ------------------------ | ------------ | ---------------------------- |
| `//go:embed file.txt`    | `string`     | Embed text file content      |
| `//go:embed file.bin`    | `[]byte`     | Embed binary file content    |
| `//go:embed file1 file2` | `embed.FS`   | Embed multiple files         |
| `//go:embed dir/*`       | `embed.FS`   | Embed all files in directory |

* Use `embed.FS.ReadFile("path")` to read embedded files.
* Clean solution for bundling assets directly inside Go binary.

---

**Command-Line Arguments and Flags in Go**. 🏳️
This is an essential part of building **CLI tools, utilities, and configurable apps**.

---

# 🖥 Command Line Arguments & Flags in Go

Go gives us **two levels** of handling command-line input:

1. **Arguments (`os.Args`)** → raw input as a slice of strings.
2. **Flags (`flag` package)** → structured parsing with types, defaults, and help text.

---

## 🔹 1. Command-Line Arguments with `os.Args`

### Example

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args gives us all arguments as []string
	args := os.Args
	fmt.Println("All arguments:", args)

	if len(args) > 1 {
		fmt.Println("First argument:", args[1])
	} else {
		fmt.Println("No arguments provided")
	}
}
```

### Run:

```bash
go run main.go hello world
```

### Output:

```
All arguments: [./main hello world]
First argument: hello
```

🔹 Notes:

* `os.Args[0]` is always the program name (path).
* Arguments start from index 1.

⚠️ Problem: All arguments are **strings**, no automatic type conversion, no help messages. That’s where **`flag` package** comes in.

---

## 🔹 2. Command-Line Flags with `flag` Package

The `flag` package is the standard way to define **options with names, default values, types, and help text**.

---

### Example: Basic Flags

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// define flags
	name := flag.String("name", "Guest", "Your name")
	age := flag.Int("age", 18, "Your age")
	isMember := flag.Bool("member", false, "Are you a member?")

	// parse flags
	flag.Parse()

	// use values
	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
	fmt.Println("Is Member:", *isMember)
}
```

### Run:

```bash
go run main.go -name=Skyy -age=29 -member=true
```

### Output:

```
Name: Skyy
Age: 29
Is Member: true
```

🔹 Notes:

* `flag.String` returns a **pointer**, so we dereference with `*name`.
* Defaults: if we don’t pass `-name`, it uses `"Guest"`.
* `flag.Parse()` is **required** before using the flags.

---

## 🔹 3. Positional Arguments + Flags

Anything **after flags** is stored in `flag.Args()`.

```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "dev", "Mode of operation")
	flag.Parse()

	fmt.Println("Mode:", *mode)
	fmt.Println("Remaining args:", flag.Args())
}
```

### Run:

```bash
go run main.go -mode=prod file1.txt file2.txt
```

### Output:

```
Mode: prod
Remaining args: [file1.txt file2.txt]
```

✅ Useful for commands like:

```
mytool -mode=prod input.txt output.txt
```

---

## 🔹 4. Custom Flag Types

We can define our own flag type by implementing `flag.Value`.

### Example: Duration

```go
package main

import (
	"flag"
	"fmt"
	"time"
)

type durationFlag struct {
	val time.Duration
}

func (d *durationFlag) String() string {
	return d.val.String()
}

func (d *durationFlag) Set(s string) error {
	v, err := time.ParseDuration(s)
	if err != nil {
		return err
	}
	d.val = v
	return nil
}

func main() {
	var d durationFlag
	flag.Var(&d, "timeout", "Timeout duration (e.g. 5s, 2m)")

	flag.Parse()
	fmt.Println("Timeout:", d.val)
}
```

### Run:

```bash
go run main.go -timeout=10s
```

### Output:

```
Timeout: 10s
```

---

## 🔹 5. Help & Usage

By default, `flag` auto-generates help text.

```go
package main

import (
	"flag"
)

func main() {
	name := flag.String("name", "Guest", "Your name")
	flag.Parse()
}
```

Run:

```bash
go run main.go -h
```

Output:

```
Usage of main:
  -name string
        Your name (default "Guest")
```

✅ `flag.Usage` can be overridden to customize help messages.

---

## 🔹 6. Advanced: Subcommands (Like `git commit`, `git push`)

For complex CLIs, we can simulate subcommands:

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	removeCmd := flag.NewFlagSet("remove", flag.ExitOnError)

	addName := addCmd.String("name", "", "Name to add")
	removeID := removeCmd.Int("id", 0, "ID to remove")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'remove'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Println("Adding:", *addName)
	case "remove":
		removeCmd.Parse(os.Args[2:])
		fmt.Println("Removing ID:", *removeID)
	default:
		fmt.Println("Unknown command")
	}
}
```

Run:

```bash
go run main.go add -name=Skyy
go run main.go remove -id=42
```

---

# 📝 Summary (Cheat Sheet)

| Method              | Use Case                                            |
| ------------------- | --------------------------------------------------- |
| `os.Args`           | Raw arguments (simple cases, positional arguments). |
| `flag.String` etc   | Typed flags (string, int, bool, etc.).              |
| `flag.Parse()`      | Must be called before reading flag values.          |
| `flag.Args()`       | Extra positional args after flags.                  |
| `flag.Usage`        | Custom help messages.                               |
| `flag.NewFlagSet`   | Subcommands (e.g. `git commit`).                    |
| Custom `flag.Value` | Support custom flag types (e.g. time.Duration).     |

---

👉 So in short:

* Use **`os.Args`** when we just need raw strings.
* Use **`flag` package** for real CLI tools (typed, defaults, help).
* For very advanced CLIs (like Docker, Git), we’d move to frameworks like **Cobra** or **urfave/cli**.

---

Let’s go step by step and deeply explain **command-line subcommands in Go**.

---

## 1. 📌 What are Subcommands?

* **Commands**: Arguments we pass when running a program.
* **Subcommands**: Think of them as *mini-programs* or *modes* inside the main program.
  Each subcommand usually has its **own set of flags and behavior**.

👉 Example from Git:

```sh
git commit -m "message"
git push origin main
git status
```

* `commit`, `push`, `status` → **subcommands** of `git`.

So in Go, we can build CLIs that behave the same way.

---

## 2. 📌 Subcommands in Go using `flag` package

Go’s standard library (`flag`) does not directly support subcommands,
but we can simulate them by creating **separate `FlagSet`s**.

### Example: Simple Subcommands

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define subcommands as FlagSets
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	subCmd := flag.NewFlagSet("sub", flag.ExitOnError)

	// Define flags for "add"
	addX := addCmd.Int("x", 0, "first number")
	addY := addCmd.Int("y", 0, "second number")

	// Define flags for "sub"
	subX := subCmd.Int("x", 0, "first number")
	subY := subCmd.Int("y", 0, "second number")

	if len(os.Args) < 2 {
		fmt.Println("expected 'add' or 'sub' subcommand")
		os.Exit(1)
	}

	// Switch based on subcommand
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Println("Result:", *addX+*addY)
	case "sub":
		subCmd.Parse(os.Args[2:])
		fmt.Println("Result:", *subX-*subY)
	default:
		fmt.Println("expected 'add' or 'sub' subcommand")
		os.Exit(1)
	}
}
```

### Run it:

```sh
go run main.go add -x 5 -y 3
# Output: Result: 8

go run main.go sub -x 10 -y 4
# Output: Result: 6
```

🔑 Key ideas:

* `flag.NewFlagSet` lets us define flags **specific** to a subcommand.
* We manually check `os.Args[1]` to see which subcommand the user typed.

---

## 3. 📌 Real-World Example: A CLI with multiple tools

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// greet subcommand
	greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
	name := greetCmd.String("name", "World", "Who to greet")

	// math subcommand
	mathCmd := flag.NewFlagSet("math", flag.ExitOnError)
	num1 := mathCmd.Int("a", 0, "First number")
	num2 := mathCmd.Int("b", 0, "Second number")

	if len(os.Args) < 2 {
		fmt.Println("expected 'greet' or 'math'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "greet":
		greetCmd.Parse(os.Args[2:])
		fmt.Printf("Hello, %s!\n", *name)
	case "math":
		mathCmd.Parse(os.Args[2:])
		fmt.Printf("%d + %d = %d\n", *num1, *num2, *num1+*num2)
	default:
		fmt.Println("expected 'greet' or 'math'")
		os.Exit(1)
	}
}
```

### Run:

```sh
go run main.go greet -name Skyy
# Output: Hello, Skyy!

go run main.go math -a 7 -b 5
# Output: 7 + 5 = 12
```

---

## 4. 📌 Limitations of `flag` for subcommands

* Manual parsing of `os.Args` is required.
* No built-in help menus for subcommands.
* Gets messy with many subcommands.

---

## 5. 📌 Using Third-Party CLI Libraries

For complex CLIs, Go developers usually use **Cobra** or **urfave/cli**.

### Cobra (used in `kubectl`, `git-lfs`, etc.)

* Auto-generates help menus.
* Easier subcommand hierarchy.
* Widely used for production-grade CLIs.

Example with Cobra:

```sh
go get -u github.com/spf13/cobra@latest
```

```go
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}

	var greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "Greets the user",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello from Cobra!")
		},
	}

	rootCmd.AddCommand(greetCmd)
	rootCmd.Execute()
}
```

Run:

```sh
go run main.go greet
# Output: Hello from Cobra!
```

---

## 6. 📌 Summary

* **Subcommands** = mini-programs inside a CLI tool.
* With standard lib → use `flag.NewFlagSet` + manual parsing.
* For advanced apps → use **Cobra** or **urfave/cli**.
* Subcommands let us build CLIs like `git`, `docker`, `kubectl`.

---

**Environment Variables in Go (Golang)** in detail.

---

# 🌍 1. What are Environment Variables?

* Environment variables are **key-value pairs** stored outside the program.
* They provide configuration (like database URL, API keys, ports) without hardcoding values in code.
* This makes apps more **flexible, secure, and portable**.

👉 Example in Linux/macOS:

```sh
export DB_USER=root
export DB_PASS=secret
```

👉 Example in Windows (PowerShell):

```ps
$env:DB_USER="root"
$env:DB_PASS="secret"
```

---

# 🛠 2. Working with Environment Variables in Go

Go provides the **`os` package** for handling env vars.

---

## 📌 (a) Reading Environment Variables

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	if dbUser == "" || dbPass == "" {
		fmt.Println("Environment variables not set!")
	} else {
		fmt.Println("DB_USER:", dbUser)
		fmt.Println("DB_PASS:", dbPass)
	}
}
```

Run:

```sh
export DB_USER=root
export DB_PASS=secret
go run main.go
```

Output:

```
DB_USER: root
DB_PASS: secret
```

👉 `os.Getenv("KEY")` returns:

* The value if the variable exists.
* An **empty string** if it does not exist.

---

## 📌 (b) MustGetEnv (Fail if Missing)

Sometimes we want to **fail fast** if a required variable is missing.

```go
func mustGetEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Missing required env var: %s", key))
	}
	return val
}
```

---

## 📌 (c) Setting Environment Variables (inside Go)

```go
os.Setenv("PORT", "8080")
port := os.Getenv("PORT")
fmt.Println("Server running on port:", port)
```

👉 Note: `os.Setenv` only changes env vars for the **current process**, not globally.

---

## 📌 (d) Listing All Environment Variables

```go
for _, e := range os.Environ() {
	fmt.Println(e)
}
```

This prints values like:

```
PATH=/usr/local/bin:/usr/bin
HOME=/home/skyy
SHELL=/bin/bash
```

---

## 📌 (e) Unsetting Environment Variables

```go
os.Unsetenv("DB_PASS")
```

---

# 🧑‍💻 3. Loading Environment Variables from `.env` File

By default, Go doesn’t load `.env` files like Node.js does.
We need a package such as [`github.com/joho/godotenv`](https://github.com/joho/godotenv).

### Install:

```sh
go get github.com/joho/godotenv
```

### Example:

**.env file**

```env
DB_USER=root
DB_PASS=secret
PORT=8080
```

**main.go**

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access variables
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))
	fmt.Println("DB_PASS:", os.Getenv("DB_PASS"))
	fmt.Println("PORT:", os.Getenv("PORT"))
}
```

Run:

```sh
go run main.go
```

Output:

```
DB_USER: root
DB_PASS: secret
PORT: 8080
```

---

# 🚀 4. Best Practices

✅ Don’t hardcode secrets in code.
✅ Use `.env` in **local development**, but in production use system env vars.
✅ Don’t commit `.env` to GitHub (add to `.gitignore`).
✅ Provide **default values** if env vars are optional.

---

# 📌 5. Example: Using Environment Variables in Web Server

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default fallback
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go! Running on port %s", port)
	})

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
```

Run:

```sh
PORT=8080 go run main.go
```

Output:

```
Server starting on port 8080...
```

---

# 🎯 Summary

* `os.Getenv(key)` → get value.
* `os.Setenv(key, val)` → set value.
* `os.Unsetenv(key)` → remove variable.
* `os.Environ()` → list all env vars.
* Use `godotenv` for `.env` files.
* Use env vars for **config, secrets, and portability**.

---

Let’s break it down clearly — how **environment variables work in Linux vs Windows**, and how **Go handles them**.

---

# 🌍 1. Environment Variables: Cross-Platform Concept

* Environment variables are just **key–value pairs attached to a process**.
* When we launch a program, the OS gives it a **copy of the current environment**.
* The Go runtime (`os` package) provides a **unified API**, so Go code works the same on **Linux, macOS, and Windows**.
* But the **way we define/set/use environment variables** differs between operating systems.

---

# 🐧 2. Linux/macOS (Unix-like systems)

### ✅ Setting Environment Variables

* **Temporary (only for this session/program):**

```sh
export DB_USER=root
export DB_PASS=secret
go run main.go
```

* The `export` command makes the variables available to child processes (like our Go app).

* **Inline (just for one command):**

```sh
DB_USER=root DB_PASS=secret go run main.go
```

### ✅ Checking Environment Variables

```sh
echo $DB_USER
printenv
env
```

### ✅ Removing an Environment Variable

```sh
unset DB_PASS
```

---

# 🪟 3. Windows

Windows has different syntax depending on **Command Prompt** or **PowerShell**.

### ✅ Setting Environment Variables

* **In Command Prompt (cmd.exe):**

```cmd
set DB_USER=root
set DB_PASS=secret
go run main.go
```

* **In PowerShell:**

```ps
$env:DB_USER="root"
$env:DB_PASS="secret"
go run main.go
```

### ✅ Checking Environment Variables

* **cmd.exe**

```cmd
echo %DB_USER%
```

* **PowerShell**

```ps
echo $env:DB_USER
```

### ✅ Removing an Environment Variable

* **cmd.exe**

```cmd
set DB_PASS=
```

* **PowerShell**

```ps
Remove-Item Env:DB_PASS
```

---

# 🏗 4. Permanent Environment Variables

* **Linux/macOS** → add to `~/.bashrc`, `~/.zshrc`, or `/etc/environment`
  Example:

```sh
export DB_USER=root
```

* **Windows** → use **System Properties → Environment Variables** GUI, or PowerShell:

```ps
setx DB_USER "root"
```

⚠️ `setx` makes it **permanent** but not available until a **new terminal** is opened.

---

# 🔄 5. Go’s Behavior Across Platforms

* **Reading**:

```go
os.Getenv("DB_USER")
```

✅ Works the same on Linux, macOS, and Windows.

* **Setting** (inside Go):

```go
os.Setenv("PORT", "8080")
```

✅ Affects **only the current process** and any child processes it spawns.
❌ It will **not persist** after the program exits, on either Linux or Windows.

* **Unsetting**:

```go
os.Unsetenv("DB_PASS")
```

✅ Works identically across OS.

* **Listing all environment variables**:

```go
for _, e := range os.Environ() {
    fmt.Println(e)
}
```

✅ Works the same on both OS.

---

# 🎯 Quick Comparison

| Operation       | Linux/macOS                     | Windows (cmd)    | Windows (PowerShell)  |
| --------------- | ------------------------------- | ---------------- | --------------------- |
| Set (temporary) | `export KEY=value`              | `set KEY=value`  | `$env:KEY="value"`    |
| Inline          | `KEY=value go run main.go`      | ❌ Not supported  | ❌ Not supported       |
| Get             | `echo $KEY`                     | `echo %KEY%`     | `echo $env:KEY`       |
| Unset           | `unset KEY`                     | `set KEY=`       | `Remove-Item Env:KEY` |
| Permanent       | `~/.bashrc`, `/etc/environment` | `setx KEY value` | `setx KEY value`      |

---

✅ **Bottom line:**
Go itself doesn’t care whether we’re on Linux or Windows — it always uses the same `os.Getenv`, `os.Setenv`, `os.Environ`, etc. The only difference is **how we define/export variables in the shell** before running the program.

---

Let’s go step by step and break down **logging in Go (Golang)** in depth — from the **standard library** to **production-grade logging setups**.

---

# 📜 1. What is Logging?

Logging means recording **events, errors, and debug information** while a program runs.
It helps us:

* Debug issues
* Monitor application behavior
* Collect audit trails
* Diagnose production issues

---

# 🛠 2. Logging with Go’s Standard Library (`log` package)

Go provides the built-in **`log`** package.

### Basic Usage

```go
package main

import (
	"log"
)

func main() {
	log.Println("This is a normal log")
	log.Printf("Hello %s", "World")
	log.Fatal("Something went really wrong") // logs and exits
	log.Panic("Unexpected situation")        // logs and panics
}
```

### Important Functions

* `log.Println(v ...any)` → Prints with newline.
* `log.Printf(format string, v ...any)` → Formatted print.
* `log.Fatal(v ...any)` → Logs + calls `os.Exit(1)`.
* `log.Panic(v ...any)` → Logs + panics (stack trace).
* `log.SetPrefix("INFO: ")` → Add prefix to logs.
* `log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)` → Control metadata.

---

# ⚙️ 3. Controlling Output

By default, logs go to **stderr**.

### Redirect logs to a file

```go
package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	log.Println("Logging to a file instead of console")
}
```

---

# 🧩 4. Log Flags (Metadata)

`log.SetFlags` controls what info is attached:

```go
log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
log.Println("Example log")
```

Common flags:

* `Ldate` → Date (2009/01/23)
* `Ltime` → Time (01:23:23)
* `Lmicroseconds` → Microsecond resolution
* `Llongfile` → Full file path + line number
* `Lshortfile` → File name + line number
* `LUTC` → Use UTC instead of local

---

# 🏷 5. Custom Loggers

We can create multiple loggers with different configurations:

```go
infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
errorLogger := log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

infoLogger.Println("This is an info message")
errorLogger.Println("This is an error message")
```

---

# 🏭 6. Logging Levels (Not Built-In)

Unlike Python’s `logging` or Java’s `Log4j`, Go’s `log` package has **no built-in log levels** (info, warn, error, debug).

👉 We can simulate levels using prefixes or custom loggers:

```go
debugLogger := log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
warnLogger  := log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
```

---

# 📦 7. Structured Logging (3rd Party Packages)

For **production systems**, plain logs aren’t enough. We use **structured logging** (JSON logs that are machine-readable).

### Popular Libraries

* **[logrus](https://github.com/sirupsen/logrus)**
* **[zap](https://github.com/uber-go/zap)** (fast, production-ready)
* **[zerolog](https://github.com/rs/zerolog)** (zero allocation, very fast)

---

## Example: Logrus

```go
import log "github.com/sirupsen/logrus"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	log.WithFields(log.Fields{
		"user": "sky",
		"age":  29,
	}).Info("User logged in")
}
```

Output (JSON):

```json
{"level":"info","msg":"User logged in","time":"2025-08-18T14:05:03Z","user":"sky","age":29}
```

---

## Example: Zap (Uber’s High-Performance Logger)

```go
import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("User login",
		zap.String("user", "sky"),
		zap.Int("age", 29),
	)
}
```

Output:

```json
{"level":"info","ts":1692378341.123,"msg":"User login","user":"sky","age":29}
```

---

# 🏗 8. Best Practices for Logging in Go

1. **Use structured logging** (JSON) in production → easier for log aggregators like ELK, Loki, Splunk.
2. **Separate log levels** (info, warn, error, debug).
3. **Log errors with context** (not just "error occurred").

   ```go
   log.Printf("Error saving user %s: %v", userID, err)
   ```
4. **Do not over-log** → too many logs = noise.
5. **Redirect logs to files or stdout/stderr** (let container/orchestrator handle collection).
6. **Use context** (with `context.Context`) for request IDs, trace IDs.

---

# 🎯 Summary

* ✅ Standard `log` package is simple and useful for small apps.
* ✅ Supports prefix, flags, multiple loggers, and file redirection.
* ❌ No log levels → need to implement manually or use third-party libraries.
* 🚀 For production → use **Logrus, Zap, or Zerolog** with JSON structured logging.

---

Now.. Let’s compare **`logrus`**, **`zap`**, and **`zerolog`** side by side so we can decide which logger is best for different use cases.

---

# ⚔️ Go Logging Libraries Comparison

| Feature                       | **Logrus** 🐹                  | **Zap** ⚡ (Uber)                                     | **Zerolog** 🪶                            |
| ----------------------------- | ------------------------------ | ---------------------------------------------------- | ----------------------------------------- |
| **Maturity**                  | Oldest, very popular           | Production-ready at Uber, very stable                | Newer but growing fast                    |
| **Performance** (Speed)       | ❌ Slowest (reflection, allocs) | ⚡ Very fast (zero reflection, precompiled log calls) | ⚡ Fastest (zero alloc, extremely lean)    |
| **API Style**                 | High-level, friendly           | Structured, somewhat verbose                         | Fluent, chained, concise                  |
| **Output Format**             | Text (default) + JSON          | JSON only (but configurable)                         | JSON only (very compact)                  |
| **Ease of Use**               | ✅ Beginner-friendly            | ⚠️ Verbose for beginners                             | ✅ Concise but may feel unusual            |
| **Features**                  | Hooks, Levels, JSON, Text      | Structured logging, Sampling, Levels                 | Structured logging, No reflection, Levels |
| **File Size (binary impact)** | \~600KB extra                  | \~1.5MB extra                                        | \~300KB extra                             |
| **When to Use**               | Small/medium apps, quick start | Large-scale production, performance critical         | Microservices, ultra-fast logging         |

---

# 📝 Example Code Comparison

### 1. **Logrus**

```go
import log "github.com/sirupsen/logrus"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	log.WithFields(log.Fields{
		"user": "sky",
		"age":  29,
	}).Info("User logged in")
}
```

✅ Simple API, supports both text & JSON.
❌ Slower in high-performance apps.

---

### 2. **Zap**

```go
import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("User login",
		zap.String("user", "sky"),
		zap.Int("age", 29),
	)
}
```

✅ Very fast, structured, production-ready.
❌ API is verbose (`zap.String`, `zap.Int` everywhere).

👉 You can also use `zap.SugaredLogger` for a friendlier API:

```go
sugar := logger.Sugar()
sugar.Infof("User %s logged in, age %d", "sky", 29)
```

---

### 3. **Zerolog**

```go
import (
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	logger.Info().
		Str("user", "sky").
		Int("age", 29).
		Msg("User logged in")
}
```

✅ Extremely fast, fluent API, tiny binary size.
✅ Perfect for microservices & JSON-first logs.
❌ JSON only, no pretty text logging out-of-the-box (need `zerolog.ConsoleWriter`).

Example pretty printing:

```go
logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
```

---

# 🏆 Verdict

* **Use Logrus** → If we’re building a small/medium project, want **easy setup**, and don’t care about performance much. (Good for beginners.)
* **Use Zap** → If we’re building a **serious production system** where performance matters (APIs, backend services, large traffic).
* **Use Zerolog** → If we’re building **microservices, cloud-native apps, or need ultra-fast JSON logging** with small binary footprint.

---

📌 Recommendation : (since we’re aiming for **professional backend development with Node/Go background**):

* Start with **Zap** → solid balance of performance + production-readiness.
* For hobby projects or learning → **Logrus** (easier).
* For high-performance microservices → **Zerolog**.

---

**deep dive into Logrus 🐹 and Zap ⚡** — two of the most popular logging libraries in Go.

We’ll cover:

1. What they are & why they exist
2. Features & internals
3. API design differences
4. Performance & benchmarks
5. Practical examples for each
6. When to use which

---

# 1. **Logrus** 🐹

### 📌 Overview

* Created by **Sirupsen**, one of the oldest & most widely used Go logging libraries.
* Provides **structured logging** with a simple API.
* Very **beginner-friendly** → feels similar to `log.Print*`.
* Supports **both text & JSON** output.
* Has **hooks** (run custom actions when logging, e.g., send to Sentry, Kafka).

### 📦 Installation

```bash
go get github.com/sirupsen/logrus
```

### ✨ Key Features

* **Structured logging** via `WithFields`
* **Log levels** → `Trace`, `Debug`, `Info`, `Warn`, `Error`, `Fatal`, `Panic`
* **Hooks** → Extend logging (e.g., send to file, monitoring system)
* **Formatters** → JSONFormatter, TextFormatter (default), or custom
* **Output control** → logs to `os.Stdout` by default, can redirect

### 🔑 Example

```go
import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	// Output as JSON instead of the default text
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of stderr
	log.SetOutput(os.Stdout)

	// Set log level
	log.SetLevel(log.InfoLevel)

	// Basic log
	log.Info("Application started")

	// Structured logging
	log.WithFields(log.Fields{
		"user": "Skyy",
		"age":  29,
	}).Warn("Invalid login attempt")
}
```

✅ Pros:

* Easy to start using.
* Rich ecosystem (hooks for Graylog, ELK, Sentry, etc).
* Flexible formatting.

❌ Cons:

* **Slower** compared to Zap/Zerolog (uses reflection, allocs).
* Somewhat outdated compared to Zap.

---

# 2. **Zap** ⚡ (Uber)

### 📌 Overview

* Created by **Uber** for their high-performance, production-scale systems.
* Focused on **speed and zero-allocation structured logging**.
* Used in **large-scale production systems**.
* Two APIs:

  * **Sugared Logger** → easier, allows printf-style logging.
  * **Logger** → strict, structured logging with typed fields.

### 📦 Installation

```bash
go get go.uber.org/zap
```

### ✨ Key Features

* **Zero-allocation JSON logging** (super fast 🚀)
* **Strongly typed structured fields** (e.g., `zap.String`, `zap.Int`)
* **Log levels** → `Debug`, `Info`, `Warn`, `Error`, `DPanic`, `Panic`, `Fatal`
* **Sampling** → prevents log spam
* **Predefined configs** → `NewDevelopment()`, `NewProduction()`
* **Sync()** → flushes buffered logs

### 🔑 Example

```go
import (
	"go.uber.org/zap"
)

func main() {
	// Production logger (JSON output)
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Structured logging
	logger.Info("User login",
		zap.String("user", "Skyy"),
		zap.Int("age", 29),
	)

	// Sugared Logger (friendlier API)
	sugar := logger.Sugar()
	sugar.Infof("User %s logged in, age %d", "Skyy", 29)
}
```

✅ Pros:

* **Extremely fast** (often 10x faster than Logrus).
* Perfect for structured JSON logging in **microservices / cloud-native apps**.
* Built-in sampling, production/development configs.

❌ Cons:

* API is **verbose** without Sugared Logger (`zap.String`, `zap.Int` everywhere).
* JSON-only by default (pretty console output requires dev mode or extra setup).

---

# 3. **Performance Benchmark**

(From Uber’s official benchmark and community tests)

| Logger  | Ops/sec (higher = better) | Allocations per log |
| ------- | ------------------------- | ------------------- |
| Stdlib  | \~300k                    | \~2 allocs/log      |
| Logrus  | \~40k                     | \~20 allocs/log     |
| Zap     | \~1.5M                    | \~0 allocs/log      |
| Zerolog | \~2M                      | \~0 allocs/log      |

👉 Clearly: **Zap & Zerolog are much faster** than Logrus.

---

# 4. **When to Use Which**

* **Logrus** 🐹
  ✅ Great for **small/medium apps, hobby projects, quick prototypes**
  ✅ Easy for beginners, flexible formatting
  ❌ Avoid for high-performance production systems

* **Zap** ⚡
  ✅ Best for **serious production apps, APIs, microservices**
  ✅ Handles **high throughput** logging efficiently
  ❌ API more verbose (use `SugaredLogger` for ease)

---

# 5. **Summary**

* **Logrus** = Friendly, flexible, slower → “good for getting started”
* **Zap** = High-performance, structured, production-grade → “use in real-world systems”

---

JSON is one of the most important things we’ll use in Go, especially for **APIs, configuration, and data exchange**.

---

# 📌 JSON in Golang

JSON (**JavaScript Object Notation**) is a lightweight format for structuring data.
Go has built-in support for JSON via the **`encoding/json`** package.

---

## 1. 🔹 Importing JSON package

```go
import "encoding/json"
```

---

## 2. 🔹 Basic Concepts

* **Marshalling** → Convert Go values (structs, maps, slices) into JSON (bytes/string).
* **Unmarshalling** → Convert JSON data into Go values.

---

## 3. 🔹 Marshalling (Go → JSON)

We use `json.Marshal` or `json.MarshalIndent`.

### Example:

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	user := User{"Skyy", 29, "skyy@example.com"}

	// Convert struct to JSON
	data, _ := json.Marshal(user)
	fmt.Println(string(data)) // {"name":"Skyy","age":29,"email":"skyy@example.com"}

	// Pretty print (indented)
	pretty, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(pretty))
}
```

✅ Output:

```json
{
  "name": "Skyy",
  "age": 29,
  "email": "skyy@example.com"
}
```

---

## 4. 🔹 Unmarshalling (JSON → Go)

We use `json.Unmarshal`.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	jsonData := `{"name":"Skyy","age":29,"email":"skyy@example.com"}`
	var user User

	// Parse JSON into struct
	err := json.Unmarshal([]byte(jsonData), &user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(user.Name, user.Age, user.Email) // Skyy 29 skyy@example.com
}
```

---

## 5. 🔹 Struct Tags

* Go uses **struct tags** to control JSON key names.
* Format: `` `json:"fieldName"` ``
* Options:

  * `json:"name"` → rename key.
  * `json:"name,omitempty"` → omit if empty.
  * `json:"-"` → skip field.

### Example:

```go
type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"` // ignored
}
```

---

## 6. 🔹 JSON with Maps & Slices

If we don’t know the structure, we can use `map[string]interface{}`.

```go
jsonData := `{"name":"Skyy","age":29,"skills":["Go","React"]}`

var result map[string]interface{}
json.Unmarshal([]byte(jsonData), &result)

fmt.Println(result["name"])        // Skyy
fmt.Println(result["skills"])      // [Go React]
```

⚠️ But `interface{}` means values are stored as generic types:

* Numbers → `float64`
* Strings → `string`
* Arrays → `[]interface{}`
* Objects → `map[string]interface{}`

---

## 7. 🔹 Streaming JSON

For **large data**, we can use:

* `json.NewEncoder(w).Encode(v)` → stream to `io.Writer`
* `json.NewDecoder(r).Decode(v)` → stream from `io.Reader`

### Example:

```go
import (
	"encoding/json"
	"os"
)

func main() {
	user := map[string]string{"name": "Skyy", "role": "developer"}
	json.NewEncoder(os.Stdout).Encode(user)
}
```

---

## 8. 🔹 Handling Unknown JSON

We can unmarshal into `interface{}` when structure is unknown.

```go
var data interface{}
json.Unmarshal([]byte(`{"foo":42,"bar":["a","b"]}`), &data)

m := data.(map[string]interface{})
fmt.Println(m["foo"]) // 42
```

---

## 9. 🔹 Custom Marshal/Unmarshal

We can control how structs are converted using methods:

* `MarshalJSON() ([]byte, error)`
* `UnmarshalJSON([]byte) error`

### Example:

```go
type Age int

func (a Age) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d years\"", a)), nil
}
```

---

## 10. 🔹 Common Pitfalls

1. JSON numbers → Go interprets as `float64` when using `interface{}`.
2. Field names must be **exported** (start with capital letter) to be marshalled.

   ```go
   type User struct {
       name string // not exported → ignored
   }
   ```
3. Always check for errors in Marshal/Unmarshal.

---

## 11. 🔹 Real-World Use Case (API)

```go
import (
	"encoding/json"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := User{"Skyy", 29}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
```

---

# 📌 Summary

* **`json.Marshal`** → Go → JSON
* **`json.Unmarshal`** → JSON → Go
* Use **struct tags** to control key names.
* Unknown JSON → `map[string]interface{}` or `interface{}`.
* Use **Encoder/Decoder** for streaming.
* JSON in Go is strongly tied to **exported fields**.

---

# 📌 Struct Tags in Go

Struct tags are **metadata annotations** attached to struct fields.
They help external packages (like `encoding/json`, ORMs, validation libraries, etc.) understand **how to process struct fields**.

---

## 1. 🔹 Basic Syntax

```go
type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age,omitempty"`
    Email string `json:"email" db:"email_address"`
}
```

👉 Syntax rules:

* Tags are **raw string literals** enclosed in backticks: `` `...` ``
* They are **key:"value"** pairs, separated by spaces if multiple.
* **Exported fields only** (fields starting with a capital letter) are affected by tags.

---

## 2. 🔹 Why Use Struct Tags?

* Control **serialization/deserialization** (JSON, XML, YAML).
* Work with **databases** (e.g., GORM, SQLx).
* Define **validation rules** (with `validator` package).
* Provide **metadata** for reflection-based libraries.

---

## 3. 🔹 JSON Struct Tags

The most common example is with `encoding/json`.

```go
type User struct {
    FullName string `json:"full_name"`
    Age      int    `json:"age,omitempty"`
    Password string `json:"-"` // ignored
}
```

Rules:

* `json:"full_name"` → use `full_name` instead of `FullName` in JSON.
* `json:"age,omitempty"` → omit if zero value (0, "", nil).
* `json:"-"` → completely ignore field.

### Example:

```go
u := User{"Skyy", 0, "secret"}
b, _ := json.Marshal(u)
fmt.Println(string(b)) 
// {"full_name":"Skyy"} → Age omitted, Password skipped
```

---

## 4. 🔹 Multiple Tags

We can define multiple tags on the same field.

```go
type User struct {
    Name  string `json:"name" xml:"name" db:"username"`
    Email string `json:"email" validate:"required,email"`
}
```

Here:

* JSON uses `"name"`.
* XML uses `"name"`.
* Database ORM maps to `username`.
* Validator ensures email is required and valid.

---

## 5. 🔹 Accessing Struct Tags with Reflection

Go’s **`reflect`** package lets us read tags at runtime.

```go
import (
    "fmt"
    "reflect"
)

type User struct {
    Name string `json:"name" db:"username"`
}

func main() {
    t := reflect.TypeOf(User{})
    field, _ := t.FieldByName("Name")
    fmt.Println(field.Tag.Get("json")) // name
    fmt.Println(field.Tag.Get("db"))   // username
}
```

👉 `.Tag.Get("json")` gives us the value of the `json` tag.

---

## 6. 🔹 Common Uses of Struct Tags

### ✅ JSON/YAML/XML

Control serialization.

```go
type Config struct {
    Host string `json:"host" yaml:"host"`
    Port int    `json:"port" yaml:"port"`
}
```

### ✅ Database ORMs

GORM, SQLx, etc. use struct tags for schema mapping.

```go
type Product struct {
    ID    int    `gorm:"primaryKey;autoIncrement"`
    Name  string `gorm:"size:255;not null"`
    Price int    `db:"price"`
}
```

### ✅ Validation

```go
type User struct {
    Email string `validate:"required,email"`
    Age   int    `validate:"gte=18"`
}
```

### ✅ Protobuf/GraphQL/Other Libraries

Struct tags define mappings for serialization/deserialization.

---

## 7. 🔹 Rules and Gotchas

1. **Only applies to exported fields**

   ```go
   type User struct {
       name string `json:"name"` // ignored (lowercase field)
   }
   ```

2. **OmitEmpty works only on zero values**

   * Numbers → 0
   * Strings → ""
   * Pointers/interfaces → nil
   * Arrays/slices/maps → len=0

3. **Multiple tags must be space-separated**

   ```go
   `json:"name" db:"username"`
   ```

4. **Invalid tags are ignored silently**
   Go won’t throw errors, libraries just won’t recognize them.

---

## 8. 🔹 Custom Struct Tags

We can create **our own tags** for internal use.

```go
type User struct {
    Name string `mytag:"important"`
}
```

Read via reflection:

```go
field, _ := reflect.TypeOf(User{}).FieldByName("Name")
fmt.Println(field.Tag.Get("mytag")) // important
```

---

## 9. 🔹 Real-World Example

API struct:

```go
type User struct {
    ID        int    `json:"id" db:"id"`
    FirstName string `json:"first_name" db:"first_name"`
    LastName  string `json:"last_name" db:"last_name"`
    Email     string `json:"email" db:"email" validate:"required,email"`
    Password  string `json:"-"` // not exposed
}
```

* JSON API → converts field names to snake\_case.
* Database ORM → maps struct fields to DB columns.
* Validator → ensures email is valid.
* Password → excluded from JSON response.

---

# 📌 Summary

* Struct tags = **metadata** for struct fields.
* Written inside backticks: `` `json:"name,omitempty" db:"username"` ``
* Used by packages like **encoding/json, database ORMs, validators**.
* Accessed at runtime using **reflection**.
* Rules:

  * Exported fields only.
  * Multiple tags allowed, space-separated.
  * `"omitempty"` and `"-"` modify behavior.

---

Go has the standard package **`encoding/xml`**, similar to `encoding/json`, that lets us encode (marshal) and decode (unmarshal) XML data.

---

# 📌 XML in Go (`encoding/xml`)

## 🔹 Import

```go
import "encoding/xml"
```

---

## 🔹 Core Functions

Just like JSON:

* `xml.Marshal(v interface{}) ([]byte, error)`
  → Go value → XML (`[]byte`).
* `xml.Unmarshal(data []byte, v interface{}) error`
  → XML (`[]byte`) → Go value.
* `xml.MarshalIndent(v interface{}, prefix, indent string)`
  → Pretty-print XML.

---

## 🔹 Struct Tags for XML

Struct tags control how XML is generated/parsed.

Syntax:

```go
`xml:"tagName,option"`
```

### Common Options:

* `xml:"name"` → element name in XML
* `xml:"-"` → ignore this field
* `xml:"attr"` → encode field as **XML attribute**
* `xml:",chardata"` → field stores the **text inside an element**
* `xml:",innerxml"` → raw inner XML (not parsed)
* `xml:",omitempty"` → omit empty fields
* `xml:",comment"` → encodes field as a comment

---

## 🔹 Example 1: Basic Marshalling

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email,omitempty"`
}

func main() {
	john := Person{Name: "John Doe", Age: 30, Email: "john@example.com"}

	// Marshal struct → XML
	xmlData, err := xml.MarshalIndent(john, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(xmlData))
}
```

### 🔹 Output:

```xml
<person>
  <name>John Doe</name>
  <age>30</age>
  <email>john@example.com</email>
</person>
```

---

## 🔹 Example 2: Unmarshalling

```go
func main() {
	data := []byte(`
	<person>
	  <name>John Doe</name>
	  <age>30</age>
	  <email>john@example.com</email>
	</person>`)

	var john Person
	err := xml.Unmarshal(data, &john)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Struct: %+v\n", john)
}
```

### 🔹 Output:

```
Struct: {XMLName:{Space: Local:person} Name:John Doe Age:30 Email:john@example.com}
```

---

## 🔹 Example 3: Attributes

```go
type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author"`
	ID      string   `xml:"id,attr"` // attribute instead of element
}
```

Instance:

```go
book := Book{Title: "Go in Action", Author: "John Doe", ID: "101"}
data, _ := xml.MarshalIndent(book, "", "  ")
fmt.Println(string(data))
```

Output:

```xml
<book id="101">
  <title>Go in Action</title>
  <author>John Doe</author>
</book>
```

---

## 🔹 Example 4: Nested Structs & Slices

```go
type Library struct {
	XMLName xml.Name `xml:"library"`
	Books   []Book   `xml:"book"`
}

type Book struct {
	Title  string `xml:"title"`
	Author string `xml:"author"`
}
```

Input XML:

```xml
<library>
  <book>
    <title>Go Programming</title>
    <author>Alan Donovan</author>
  </book>
  <book>
    <title>Go in Action</title>
    <author>John Doe</author>
  </book>
</library>
```

Unmarshal:

```go
var lib Library
_ = xml.Unmarshal([]byte(xmlStr), &lib)
fmt.Printf("%+v\n", lib)
```

---

## 🔹 Special Struct Tags

* **Inner XML**

```go
type Note struct {
	To   string `xml:"to"`
	Body string `xml:",innerxml"`
}
```

If XML:

```xml
<note>
  <to>Alice</to>
  <body><b>Hello</b>, World!</body>
</note>
```

The `Body` field will contain `<b>Hello</b>, World!` as raw string.

* **Char Data**

```go
type Message struct {
	Body string `xml:",chardata"`
}
```

If XML:

```xml
<message>Hello World</message>
```

Then `Body = "Hello World"`.

* **Comments**

```go
type Comment struct {
	Text string `xml:",comment"`
}
```

---

## 🔹 Streaming with Encoder/Decoder

For **large XML files**, instead of `Marshal/Unmarshal`, we can use **`xml.Encoder`** and **`xml.Decoder`** to stream.

Example:

```go
import (
	"encoding/xml"
	"os"
)

type Item struct {
	Name string `xml:"name"`
}

func main() {
	// Encode
	encoder := xml.NewEncoder(os.Stdout)
	item := Item{Name: "Laptop"}
	encoder.Encode(item)

	// Decode
	xmlData := `<Item><name>Laptop</name></Item>`
	decoder := xml.NewDecoder(strings.NewReader(xmlData))
	var decoded Item
	decoder.Decode(&decoded)
}
```

---

# ✅ Summary

* Use `encoding/xml` (similar to `encoding/json`).
* `xml.Marshal`, `xml.Unmarshal` for basic usage.
* Tags (`xml:"tagName,attr,omitempty"`) control how struct maps to XML.
* Supports attributes, nested structs, slices, inner XML, comments, etc.
* Use **Encoder/Decoder** for large files (streaming).
* Great for **APIs, configs, or legacy systems** that use XML.

---

Now.. Let’s compare **JSON vs XML handling in Go** in detail.

---

# 📌 JSON vs XML in Go

Both are supported in Go via standard packages:

* **`encoding/json`**
* **`encoding/xml`**

They share a lot of concepts but have key differences in **representation, struct tags, and use cases**.

---

## 🔹 1. Import

```go
import "encoding/json"
import "encoding/xml"
```

---

## 🔹 2. Marshalling (Go struct → JSON/XML)

### JSON

```go
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}
```

Output:

```json
{
  "name": "Alice",
  "age": 25,
  "email": "alice@example.com"
}
```

---

### XML

```go
type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email,omitempty"`
}
```

Output:

```xml
<person>
  <name>Alice</name>
  <age>25</age>
  <email>alice@example.com</email>
</person>
```

---

## 🔹 3. Struct Tags Differences

| Feature                 | JSON                     | XML                     |
| ----------------------- | ------------------------ | ----------------------- |
| Rename field            | `json:"username"`        | `xml:"username"`        |
| Ignore field            | `json:"-"`               | `xml:"-"`               |
| Omit empty fields       | `json:"field,omitempty"` | `xml:"field,omitempty"` |
| Use attributes          | ❌ Not supported          | `xml:"id,attr"`         |
| Raw inner data          | ❌ Not supported          | `xml:",innerxml"`       |
| Raw text inside element | ❌ Not supported          | `xml:",chardata"`       |
| Comments                | ❌ Not supported          | `xml:",comment"`        |

👉 JSON is simpler: key-value only.
👉 XML is richer: can do attributes, nested elements, inner text, etc.

---

## 🔹 4. Unmarshalling (JSON/XML → Go struct)

Both work similarly:

```go
json.Unmarshal(jsonData, &obj)
xml.Unmarshal(xmlData, &obj)
```

---

## 🔹 5. Encoding/Decoding Streams

Both support streaming:

* `json.NewEncoder(writer).Encode(obj)`
* `xml.NewEncoder(writer).Encode(obj)`

and

* `json.NewDecoder(reader).Decode(&obj)`
* `xml.NewDecoder(reader).Decode(&obj)`

Useful for **large files or network streams**.

---

## 🔹 6. Readability & Size

* **JSON**

  * More compact
  * Easier to read for humans
  * Smaller payloads
  * Popular in modern APIs (REST, GraphQL, etc.)

* **XML**

  * Verbose, larger in size
  * Harder for humans to scan quickly
  * Supports attributes, comments, nested text → more descriptive
  * Still common in **legacy systems, SOAP APIs, configuration files**

---

## 🔹 7. Example Comparison

### JSON

```json
{
  "id": "101",
  "title": "Go in Action",
  "author": "John Doe"
}
```

### XML

```xml
<book id="101">
  <title>Go in Action</title>
  <author>John Doe</author>
</book>
```

👉 Notice XML supports **attributes (`id="101"`)**, while JSON does not.

---

## 🔹 8. Performance

* **JSON**: Faster, smaller payload, better for web APIs.
* **XML**: Slower, larger, but more flexible with complex data.

Benchmarks in Go typically show **`encoding/json` > `encoding/xml`** in speed.
That’s why **most modern Go APIs use JSON**.

---

# ✅ Summary

* **JSON**:

  * Lightweight, fast, compact.
  * Struct tags: `json:"field,omitempty"`.
  * Best for web APIs, microservices, mobile apps.

* **XML**:

  * Verbose, rich metadata, supports attributes & inner text.
  * Struct tags: `xml:"field,attr,omitempty"`.
  * Best for configs, legacy systems, SOAP, scientific data.

---

👉 Rule of thumb:

* Use **JSON** if we control the API or working with modern services.
* Use **XML** if we’re **interfacing with legacy APIs, enterprise systems, or config formats**.

---

Let’s go step by step and cover **everything about type conversions in Go (Golang)** in detail.

---

# 🔹 1. What is Type Conversion in Go?

In Go, **type conversion** means explicitly changing a value from one type to another.
👉 Go does **not allow implicit type conversion** (unlike JavaScript, Python, etc.).
We **must** tell the compiler exactly what type we want.

Example:

```go
var x int = 10
var y float64 = float64(x) // explicit conversion
```

✅ `int` → `float64`
❌ No automatic conversion in Go.

---

# 🔹 2. Basic Syntax

The general syntax is:

```go
T(v)
```

* `T` → the target type we want
* `v` → the value or variable we’re converting

Example:

```go
var a int = 5
var b float64 = float64(a)  // int → float64
var c int = int(b)          // float64 → int
```

---

# 🔹 3. Numeric Type Conversions

Go supports conversions between numeric types (`int`, `float32`, `float64`, `uint`, etc.).

### Example:

```go
package main

import "fmt"

func main() {
    var i int = 42
    var f float64 = float64(i)   // int → float64
    var u uint = uint(f)         // float64 → uint

    fmt.Println(i, f, u) // 42 42 42
}
```

⚠️ Important notes:

* Converting `float → int` **truncates** (drops the decimal, doesn’t round).

  ```go
  var x float64 = 9.99
  var y int = int(x) // y = 9 (not 10)
  ```
* Negative values converted to `uint` wrap around:

  ```go
  var n int = -5
  var u uint = uint(n)
  fmt.Println(u) // big number, since it wraps
  ```

---

# 🔹 4. String and Byte/Rune Conversions

Go treats **strings, runes, and byte slices** differently, so conversions are explicit.

### (a) String ↔ Byte Slice

```go
s := "hello"
b := []byte(s)       // string → []byte
fmt.Println(b)       // [104 101 108 108 111]

s2 := string(b)      // []byte → string
fmt.Println(s2)      // "hello"
```

### (b) String ↔ Rune Slice

```go
s := "你好"
r := []rune(s)       // string → []rune (Unicode code points)
fmt.Println(r)       // [20320 22909]

s2 := string(r)      // []rune → string
fmt.Println(s2)      // "你好"
```

---

# 🔹 5. String and Number Conversions

Go does **not** allow direct `string ↔ int/float` conversion.
We use the **`strconv` package**.

### Example:

```go
import "strconv"

// int → string
s := strconv.Itoa(123)    // "123"

// string → int
i, _ := strconv.Atoi("456") // 456

// string → float
f, _ := strconv.ParseFloat("3.14", 64) // 3.14

// float → string
s2 := strconv.FormatFloat(3.14, 'f', 2, 64) // "3.14"
```

---

# 🔹 6. Interface and Type Assertion (Special Case)

When we store values in an `interface{}`, we might need to **extract the original type**.

Example:

```go
var i interface{} = "hello"
str, ok := i.(string)  // type assertion
if ok {
    fmt.Println("string:", str)
}
```

This is **not conversion**, but related to runtime type handling.

---

# 🔹 7. Custom Type Conversions

If we define our own type, we must explicitly convert it to/from its base type.

```go
type Celsius float64
type Fahrenheit float64

func main() {
    var c Celsius = 30
    f := Fahrenheit(c) // explicit conversion
    fmt.Println(f)     // 30 (but type = Fahrenheit)
}
```

---

# 🔹 8. Conversions vs Casting (Important Difference)

In C/Java, **casting** can reinterpret memory bits (unsafe).
In Go, **conversion is safe** and always changes the value logically.

If we want low-level memory reinterpretation, we use `unsafe` package, but that’s discouraged.

---

# 🔹 9. Buffer Flush Example (Relating Back to Zap/Logs)

Sometimes people say “flush before converting/outputting” when buffered writers (like `bufio.Writer`, `os.File`, loggers like Zap) hold data in memory.
We must call `.Flush()` or `.Sync()` so data is **converted & written to disk/console immediately**.

---

# 🔹 10. Quick Conversion Examples Cheat Sheet

```go
// int ↔ float
float64(i), int(f)

// string ↔ byte slice
[]byte(s), string(b)

// string ↔ rune slice
[]rune(s), string(r)

// string ↔ int/float (strconv)
strconv.Atoi("123"), strconv.Itoa(456)
strconv.ParseFloat("3.14", 64), strconv.FormatFloat(3.14, 'f', 2, 64)

// custom types
type MyInt int
var x int = 5
y := MyInt(x)
```

---

Practical tour of Go’s **io** package—the tiny set of interfaces and helpers that power almost all input/output in our programs.
---

# What the io package is

`io` defines **small, composable interfaces** and **utility functions** for streaming data. Files, network sockets, in-memory buffers, HTTP bodies—nearly all of them implement `io.Reader`, `io.Writer`, or friends. Because of this, we can plug components together without caring about the concrete types.

---

# Core interfaces (the heart of io)

### Streams

* **`io.Reader`**

  ```go
  type Reader interface {
      Read(p []byte) (n int, err error)
  }
  ```

  Reads up to `len(p)` bytes into `p`. Returns `n` and possibly an `err`. Reaching end of stream returns `err == io.EOF` (sometimes with `n > 0`).

* **`io.Writer`**

  ```go
  type Writer interface {
      Write(p []byte) (n int, err error)
  }
  ```

  Writes as many bytes as it can; may write fewer than `len(p)`.

* **`io.Closer`** — has `Close() error`. We should `defer f.Close()` when we open a resource.

* **Combinations**: `io.ReadCloser`, `io.WriteCloser`, `io.ReadWriter`, `io.ReadWriteCloser`, etc.

### Random access & seeking

* **`io.Seeker`** — `Seek(offset, whence)` moves the read/write cursor.

  * `io.SeekStart` (0): from start
  * `io.SeekCurrent` (1): from current
  * `io.SeekEnd` (2): from end
* **`io.ReaderAt` / `io.WriterAt`** — read/write **at an absolute offset** (do not change the current offset). Useful for parallel reads. (We must not mix `Read` with `ReadAt` on the same object unless docs say it’s safe.)

### Optimized streaming hooks

* **`io.ReaderFrom`** — `ReadFrom(r Reader) (n int64, err error)`
* **`io.WriterTo`** — `WriteTo(w Writer) (n int64, err error)`

  `io.Copy` will use these when available for zero-copy/fast paths (e.g., `sendfile`).

### Byte/rune helpers

* **`io.ByteReader` / `io.ByteWriter`** — `ReadByte`, `WriteByte`.
* **`io.RuneReader` / `io.RuneScanner`** — Unicode-aware `ReadRune`, `UnreadRune`.
* **`io.StringWriter`** — optimized `WriteString`.

---

# “Batteries included”: key helpers and adapters

* **`io.Copy(dst, src)`**
  Streams from `src` to `dst` until EOF. Uses fast paths when possible.

* **`io.CopyN(dst, src, n)`**
  Copies exactly `n` bytes or returns an error (often `EOF`).

* **`io.CopyBuffer(dst, src, buf)`**
  Like `Copy`, but we supply the reusable buffer (handy for pooling).

* **`io.ReadAll(r)`**
  Reads everything into memory and returns a `[]byte`. Great for small payloads; dangerous for unbounded data.

* **`io.ReadFull(r, buf)`** / **`io.ReadAtLeast(r, buf, min)`**
  Ensure we fill the buffer (or at least `min`) or error.

* **`io.WriteString(w, "…")`**
  Efficient way to write strings to a writer.

* **`io.MultiWriter(w1, w2, …)`**
  Fan-out writes to multiple writers (e.g., stdout + a file).

* **`io.MultiReader(r1, r2, …)`**
  Concatenate readers into one long stream.

* **`io.LimitReader(r, n)`**
  Wraps `r` so it produces at most `n` bytes (defense against giant inputs).

* **`io.TeeReader(r, w)`**
  Reads from `r` while also writing what it reads to `w`. Great for hashing/logging streams while forwarding them.

* **`io.Pipe()`**
  In-memory pipe with a `Reader` and `Writer`. We connect two goroutines: one writes, the other reads.

* **`io.NopCloser(r)`**
  Adds a no-op `Close()` to a `Reader` when an API requires a `ReadCloser`.

* **`io.Discard`**
  A writer that throws data away (like `/dev/null`). Useful to drain streams quickly.

* **`io.NewSectionReader(r io.ReaderAt, off, n int64)`**
  A `Reader`/`ReaderAt`/`Seeker` that views a slice of a larger `ReaderAt` (handy for range requests or partial file reads).

* **`io.NewOffsetWriter(w io.WriterAt, off int64)`** (Go 1.20+)
  A `Writer` that writes to `w` starting at an offset.

---

# Practical patterns we use all the time

### 1) Copy a file to another file (streaming)

```go
src, _ := os.Open("in.bin")
defer src.Close()

dst, _ := os.Create("out.bin")
defer dst.Close()

if _, err := io.Copy(dst, src); err != nil { /* handle */ }
```

### 2) Limit incoming request body (defensive)

```go
r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10 MB cap (http helper)
limited := io.LimitReader(r.Body, 10<<20)       // io helper
data, err := io.ReadAll(limited)
```

### 3) Hash data while forwarding it (tee)

```go
h := sha256.New()
tee := io.TeeReader(src, h)
if _, err := io.Copy(dst, tee); err != nil { /* handle */ }
sum := h.Sum(nil)
```

### 4) Write to multiple sinks (console + file)

```go
logFile, _ := os.Create("app.log")
defer logFile.Close()

mw := io.MultiWriter(os.Stdout, logFile)
fmt.Fprintln(mw, "hello world") // goes to both
```

### 5) Producer/consumer with a pipe

```go
pr, pw := io.Pipe()

go func() {
    defer pw.Close()
    // produce
    io.Copy(pw, someSource)
}()

// consume
if _, err := io.Copy(someDest, pr); err != nil { /* handle */ }
```

### 6) Implement our own Reader wrapper

```go
type upperReader struct{ r io.Reader }

func (u upperReader) Read(p []byte) (int, error) {
    n, err := u.r.Read(p)
    for i := 0; i < n; i++ {
        if 'a' <= p[i] && p[i] <= 'z' { p[i] -= 'a' - 'A' }
    }
    return n, err
}

// usage: io.Copy(dst, upperReader{r: src})
```

### 7) Random access reads in parallel

```go
f, _ := os.Open("big.dat")
defer f.Close()

bufA := make([]byte, 1<<20)
bufB := make([]byte, 1<<20)

var wg sync.WaitGroup
wg.Add(2)
go func() { defer wg.Done(); f.ReadAt(bufA, 0) }()
go func() { defer wg.Done(); f.ReadAt(bufB, 1<<20) }()
wg.Wait()
```

*(Do not mix `ReadAt` with `Read` concurrently on the same file.)*

---

# The tricky bits we must remember

* **EOF semantics**
  `Read` may return `(n > 0, err == io.EOF)`. We must process the `n` bytes; EOF just means “no more after this”.

* **Partial reads/writes**
  Both `Read` and `Write` can return partial counts; loops are our friend:

  ```go
  for n < len(buf) {
      m, err := r.Read(buf[n:])
      n += m
      if err != nil { /* handle (maybe EOF) */ break }
  }
  ```

* **Closing matters**
  Always close what we open. For HTTP responses: `defer resp.Body.Close()`, even if we don’t read it fully (or drain it to let keep-alive work).

* **Streaming beats `ReadAll`**
  `io.ReadAll` is fine for small known-size data. For unknown/untrusted size, prefer streaming (`io.Copy`) plus limits (`LimitReader`/`MaxBytesReader`).

* **Buffer reuse**
  For high-throughput paths, reuse a `[]byte` buffer or supply our own to `io.CopyBuffer` to reduce allocations.

* **Concurrency**
  Most `io.Reader`/`Writer` implementations are **not** safe for concurrent use (unless documented). `ReaderAt` implementations must be safe for concurrent `ReadAt` calls.

* **`io.Copy` is smarter than it looks**
  It will call `WriterTo`/`ReaderFrom` if available for fewer copies and syscalls.

* **Adapting types**

  * We can wrap any `Reader` to become a `ReadCloser` using `io.NopCloser`.
  * Use `io.Discard` to sink data we must read but don’t care about.

---

# When to use what (quick guide)

* We need to just move bytes from A → B: **`io.Copy`**.
* We need to cap data size: **`io.LimitReader`** (or `http.MaxBytesReader`).
* We want to write to several places: **`io.MultiWriter`**.
* We want to concatenate sources: **`io.MultiReader`**.
* We need a “tee” for hashing/logging: **`io.TeeReader`**.
* We want to connect goroutines via a stream: **`io.Pipe`**.
* We need random access reads: **`io.ReaderAt`** (and maybe `NewSectionReader`).
* We need to parse everything into memory (small data): **`io.ReadAll`**.
* We must satisfy `ReadCloser` from a plain `Reader`: **`io.NopCloser`**.
* We need an efficient `WriteString`: **`io.StringWriter`** or `io.WriteString`.

---

This file is an **excellent practical deep dive into the `io` package in Go**, covering the three key abstractions:

* `io.Reader` (read streams of data)
* `io.Writer` (write streams of data)
* `io.Closer` (close resources)

Let’s break it **line by line and concept by concept** 👇

---

## 📌 Imports

```go
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
```

* `bytes` → Used to create buffers (`bytes.Buffer`) that can act as both readers and writers.
* `fmt` → Printing and formatting output.
* `io` → Provides core interfaces (`Reader`, `Writer`, `Closer`) and utilities (`MultiReader`, `Pipe`, etc.).
* `log` → Logging errors with fatal exits.
* `os` → For interacting with files (open, write, create).
* `strings` → Allows creating `io.Reader` from strings (`strings.NewReader`).

---

## 📌 `readFromReader(r io.Reader)`

```go
func readFromReader(r io.Reader) {
	buf:= make([]byte,1024) // buffer of 1024 bytes
	n,err:= r.Read(buf)     // read into buffer
	if err!=nil{
		log.Fatal("🔴ERROR reading from READER:",err)
	}
	fmt.Println("Data:",string(buf[:n]))
}
```

* Accepts **any type that implements `io.Reader`** (e.g., files, strings, buffers).
* Creates a `buf` (temporary storage).
* Calls `Read(buf)` → Fills `buf` with data and returns `n` (number of bytes read).
* `string(buf[:n])` ensures we only convert actual data, not empty parts of buffer.

⚡ Usage: `readFromReader(strings.NewReader("random str."))`

---

## 📌 `writeToWriter(w io.Writer, data string)`

```go
func writeToWriter(w io.Writer, data string){
	_,err:= w.Write([]byte(data))
	if err!=nil{
		log.Fatal("🔴ERROR writing to WRITER:",err)
	}
}
```

* Accepts **any type implementing `io.Writer`** (like files, buffers, network connections).
* Converts string to `[]byte` and writes it.
* Errors handled with logging.

⚡ Example: `writeToWriter(&bytes.Buffer{}, "Hello")`

---

## 📌 `closeResource(c io.Closer)`

```go
func closeResource(c io.Closer){
	err:=c.Close()
	if err!=nil{
		log.Fatal("🔴ERROR closing the resource:",err)
	}
}
```

* Accepts anything implementing `io.Closer`.
* Calls `.Close()` → Used for releasing resources (files, DB connections, etc.).
* Common in `defer file.Close()` patterns.

---

## 📌 `bufferExample()`

```go
func bufferExample(){
	var buf bytes.Buffer // stack allocated
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}
```

* `bytes.Buffer` is an in-memory buffer (both `Reader` + `Writer`).
* `WriteString` adds data, `String()` retrieves it.
* Very efficient for temporary string manipulation.

---

## 📌 `multiReaderExample()`

```go
func multiReaderExample(){
	r1:= strings.NewReader("Hey ")
	r2:= strings.NewReader("World")
	mr:= io.MultiReader(r1,r2)

	buf := new(bytes.Buffer)
	_,err:=buf.ReadFrom(mr)
	if err!=nil{
		log.Fatal("🔴ERROR reading from READER:",err)
	}
	fmt.Println("Data/Bufs:",buf.String())
}
```

* `io.MultiReader(r1, r2)` → Combines multiple readers into one.
* Reads sequentially: first `r1` then `r2`.
* `buf.ReadFrom(mr)` copies data from combined reader into buffer.
* Output: `"Hey World"`.

---

## 📌 `pipeExample()`

```go
func pipeExample(){
	pr,pw:=io.Pipe()

	go func(){
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}
```

* `io.Pipe()` creates a **reader (`pr`) and writer (`pw`)** that are connected.
* Data written to `pw` can be read from `pr`.
* Example of producer-consumer pattern.
* Used for streaming, like passing log data from one goroutine to another.

---

## 📌 `writeToFile(filepath string, data string)`

```go
func writeToFile(filepath string, data string){
	file,err:= os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		log.Fatalln("🔴ERROR opening/creating file:",err)
	}
	defer closeResource(file)

	_,err = file.Write([]byte(data))
	if err != nil{
		log.Fatalln("🔴ERROR writing file:",err)
	}
}
```

* Opens/creates file at `filepath`.

  * `os.O_APPEND` → Append to file.
  * `os.O_CREATE` → Create if doesn’t exist.
  * `os.O_WRONLY` → Write-only mode.
* `0644` → File permissions.
* Writes data to file, then ensures it’s closed.

⚡ Example: Appends `"Hello File\n"` to `io.txt`.

---

## 📌 `MyResource` with `io.Closer`

```go
type MyResource struct{
	name string
}
func (m MyResource) Close()error{
	fmt.Println("Closing resource!",m.name)
	return nil
}
```

* Implements `io.Closer` by defining `Close() error`.
* Allows `MyResource` to be closed like a file or DB connection.

⚡ Example: `closeResource(&MyResource{name:"testingResource"})`.

---

## 📌 `main()`

```go
func main() {
	fmt.Println("====== READ FROM READER =======")
	readFromReader(strings.NewReader("random str. to be read"))

	fmt.Println("\n====== WRITE TO WRITER =======")
	var writer bytes.Buffer
	writeToWriter(&writer, "random str. to be written to WRITER")
	fmt.Println(writer.String())

	fmt.Println("\n====== BUFFER EXAMPLE =======")
	bufferExample()

	fmt.Println("\n====== MULTI-READER EXAMPLE =======")
	multiReaderExample()

	fmt.Println("\n====== READING/WRITING PIPE =======")
	pipeExample()

	fmt.Println("\n====== WRITE TO FILE =======")
	writeToFile("io.txt","Hello File\n")

	fmt.Println("\n====== IMPLEMENTING IO.CLOSER =======")
	resource:= &MyResource{name:"testingResource"}
	closeResource(resource)
}
```

---

## 📌 Output Recap

```
====== READ FROM READER =======
Data: random str. to be read

====== WRITE TO WRITER =======
random str. to be written to WRITER

====== BUFFER EXAMPLE =======
Hello Buffer!

====== MULTI-READER EXAMPLE =======
Data/Bufs: Hey World

====== READING/WRITING PIPE =======
Hello Pipe

====== WRITE TO FILE =======

====== IMPLEMENTING IO.CLOSER =======
Closing resource! testingResource
```

---

## 🚀 Big Picture

This program demonstrates **the essence of Go’s I/O model**:

* `Reader`, `Writer`, and `Closer` are **interfaces**, not concrete implementations.
* We can plug in different sources/sinks: strings, files, buffers, pipes, network connections.
* Because they all share these interfaces, your code becomes **highly reusable and modular**.

👉 That’s why Go’s I/O is considered **simple but extremely powerful**.

---

This comes up a lot when we first dive into Go’s standard library.
Let’s break it down clearly 👇

---

## 📦 `io` Package

* **Purpose**: Provides **basic interfaces** for I/O.
* Think of it as the *foundation layer* in Go’s I/O model.

### Key Interfaces

* **`io.Reader`** → anything that can read bytes (`Read(p []byte) (n int, err error)`).
* **`io.Writer`** → anything that can write bytes (`Write(p []byte) (n int, err error)`).
* **`io.Closer`** → something that can be closed (`Close() error`).
* **`io.ReaderFrom`, `io.WriterTo`, `io.Seeker`, etc.** → small, composable behaviors.

🔑 **Important**:

* `io` itself doesn’t buffer, it just defines *how reading and writing works*.
* Examples: `os.File` implements both `io.Reader` and `io.Writer`.

---

## 📦 `bufio` Package

* **Purpose**: Provides **buffered wrappers** around `io.Reader` and `io.Writer`.
* Think of it as the *performance enhancer* for I/O operations.

### Why buffering?

* If we read/write one byte at a time directly with `io.Reader/Writer`, it’s **slow** (every call hits the underlying source, like disk or network).
* `bufio` improves efficiency by reading/writing data in **chunks** (buffer), reducing system calls.

### Key Types

* **`bufio.Reader`** → wraps an `io.Reader`, provides efficient buffered reading + helper methods like `ReadLine()`, `ReadString(delim)`, `Peek(n)`.
* **`bufio.Writer`** → wraps an `io.Writer`, buffers writes until flushed (with `Flush()`).
* **`bufio.Scanner`** → super handy for reading text line by line or token by token.

---

## ⚡ Example Difference

### Using `io.Reader` directly:

```go
file, _ := os.Open("data.txt")
defer file.Close()

buf := make([]byte, 10)
n, _ := file.Read(buf)
fmt.Println("Read:", string(buf[:n]))
```

* Reads raw bytes directly from file into `buf`.
* Each call hits the file system.

---

### Using `bufio.Reader`:

```go
file, _ := os.Open("data.txt")
defer file.Close()

reader := bufio.NewReader(file)
line, _ := reader.ReadString('\n')
fmt.Println("Line:", line)
```

* Reads efficiently with buffering.
* Gives higher-level helpers (like reading until newline).

---

## ✅ TL;DR

* **`io`** → defines the *basic building blocks* (interfaces for reading/writing).
* **`bufio`** → provides *buffered & convenient implementations* on top of `io` for performance and ease of use.

---

Lastly.. Let’s go deep into the **`math` package** in Go.
This is one of the most important standard library packages, especially if we work with **numerical computations, algorithms, or graphics**.

---

# 📦 `math` Package in Go

* Path: **`import "math"`**
* Purpose: Provides **basic constants and mathematical functions** for floating-point arithmetic.
* Works mainly with `float64` (most functions take/return `float64`).

---

## 1. 🔢 Mathematical Constants

Go gives us many **predefined constants** in `math` (all `float64`):

| Constant       | Value     | Meaning                            |
| -------------- | --------- | ---------------------------------- |
| `math.Pi`      | 3.14159…  | Ratio of circumference to diameter |
| `math.E`       | 2.71828…  | Base of natural logarithms         |
| `math.Phi`     | 1.61803…  | Golden ratio                       |
| `math.Sqrt2`   | 1.41421…  | √2                                 |
| `math.SqrtE`   | 1.64872…  | √e                                 |
| `math.SqrtPi`  | 1.77245…  | √π                                 |
| `math.SqrtPhi` | 1.27201…  | √φ                                 |
| `math.Ln2`     | 0.693147… | Natural log of 2                   |
| `math.Ln10`    | 2.30258…  | Natural log of 10                  |
| `math.Log2E`   | 1.44269…  | log₂(e)                            |
| `math.Log10E`  | 0.434294… | log₁₀(e)                           |

---

## 2. ➗ Basic Arithmetic Helpers

* `math.Abs(x)` → absolute value of x
* `math.Mod(x, y)` → remainder of x / y (like `%` but works with floats)
* `math.Modf(x)` → splits x into integer and fractional parts

---

## 3. 📈 Powers, Roots, and Exponentials

* `math.Pow(x, y)` → xʸ
* `math.Pow10(n)` → 10ⁿ
* `math.Sqrt(x)` → √x
* `math.Cbrt(x)` → ∛x
* `math.Exp(x)` → eˣ
* `math.Exp2(x)` → 2ˣ
* `math.Expm1(x)` → eˣ − 1 (more accurate for small x)

---

## 4. 📉 Logarithms

* `math.Log(x)` → natural log (ln)
* `math.Log2(x)` → base-2 log
* `math.Log10(x)` → base-10 log
* `math.Log1p(x)` → ln(1 + x) (accurate for small x)

---

## 5. 📐 Trigonometry

(All take radians)

* **Basic**:

  * `math.Sin(x)`, `math.Cos(x)`, `math.Tan(x)`
* **Inverse (arcs)**:

  * `math.Asin(x)`, `math.Acos(x)`, `math.Atan(x)`
  * `math.Atan2(y, x)` → angle (like `atan(y/x)` but quadrant-aware)
* **Hyperbolic**:

  * `math.Sinh(x)`, `math.Cosh(x)`, `math.Tanh(x)`
* **Inverse hyperbolic**:

  * `math.Asinh(x)`, `math.Acosh(x)`, `math.Atanh(x)`

---

## 6. 🎲 Rounding and Comparison

* `math.Ceil(x)` → round up
* `math.Floor(x)` → round down
* `math.Trunc(x)` → cut off fractional part
* `math.Round(x)` → round to nearest integer
* `math.RoundToEven(x)` → round to nearest even number if halfway
* `math.Max(x, y)` → larger of two
* `math.Min(x, y)` → smaller of two
* `math.Copysign(x, y)` → x with sign of y

---

## 7. 🧮 Special Functions

* `math.Hypot(x, y)` → √(x² + y²) (distance in 2D plane, safe from overflow)
* `math.Dim(x, y)` → max(x - y, 0)
* `math.Frexp(f)` → splits float into mantissa + exponent
* `math.Ldexp(frac, exp)` → frac × 2ᵉˣᵖ
* `math.IsNaN(x)` → check if NaN
* `math.IsInf(x, sign)` → check if +∞ or -∞
* `math.NaN()` → returns NaN
* `math.Inf(sign)` → returns +∞ (sign = 1) or -∞ (sign = -1)

---

## 8. 🎯 Example Usage

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Pi:", math.Pi)
	fmt.Println("Abs(-42):", math.Abs(-42))
	fmt.Println("Sqrt(16):", math.Sqrt(16))
	fmt.Println("Pow(2, 3):", math.Pow(2, 3))
	fmt.Println("Log10(1000):", math.Log10(1000))
	fmt.Println("Ceil(3.2):", math.Ceil(3.2))
	fmt.Println("Floor(3.8):", math.Floor(3.8))
	fmt.Println("Round(2.5):", math.Round(2.5))
	fmt.Println("Hypot(3,4):", math.Hypot(3,4)) // 5
}
```

**Output:**

```
Pi: 3.141592653589793
Abs(-42): 42
Sqrt(16): 4
Pow(2, 3): 8
Log10(1000): 3
Ceil(3.2): 4
Floor(3.8): 3
Round(2.5): 2
Hypot(3,4): 5
```

---

## ✅ TL;DR

* `math` is the **go-to package** for constants, powers, roots, logs, trig, rounding, and special functions.
* Works with **float64** mostly.
* For randomness, use **`math/rand`**, not `math`.
* For big integers, use **`math/big`** package.

---








