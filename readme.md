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


