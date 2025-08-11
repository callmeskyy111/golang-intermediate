Here’s a **complete Go verb formatter cheat sheet** for `fmt.Printf`, `fmt.Sprintf`, `fmt.Fprintf`, etc., neatly organized so we can quickly find the right one when needed.

---

## **📜 Go Format Verbs Cheat Sheet**

### 1️⃣ **General**

| Verb  | Description              | Example                       |
| ----- | ------------------------ | ----------------------------- |
| `%v`  | Default format for value | `fmt.Printf("%v", myVar)`     |
| `%+v` | Struct fields with names | `fmt.Printf("%+v", myStruct)` |
| `%#v` | Go-syntax representation | `fmt.Printf("%#v", myStruct)` |
| `%T`  | Type of value            | `fmt.Printf("%T", myVar)`     |
| `%%`  | Literal `%` sign         | `fmt.Printf("%%")`            |

---

### 2️⃣ **Booleans**

| Verb | Description       | Example                           |
| ---- | ----------------- | --------------------------------- |
| `%t` | `true` or `false` | `fmt.Printf("%t", true)` → `true` |

---

### 3️⃣ **Integers**

| Verb | Description                       | Example                           |
| ---- | --------------------------------- | --------------------------------- |
| `%b` | Binary                            | `fmt.Printf("%b", 5)` → `101`     |
| `%c` | Character from Unicode code point | `fmt.Printf("%c", 65)` → `A`      |
| `%d` | Decimal (base 10)                 | `fmt.Printf("%d", 123)` → `123`   |
| `%o` | Octal                             | `fmt.Printf("%o", 9)` → `11`      |
| `%O` | Octal with `0o` prefix (Go 1.13+) | `fmt.Printf("%O", 9)` → `0o11`    |
| `%q` | Quoted character                  | `fmt.Printf("%q", 65)` → `'A'`    |
| `%x` | Hexadecimal (lowercase)           | `fmt.Printf("%x", 255)` → `ff`    |
| `%X` | Hexadecimal (uppercase)           | `fmt.Printf("%X", 255)` → `FF`    |
| `%U` | Unicode format                    | `fmt.Printf("%U", 65)` → `U+0041` |

---

### 4️⃣ **Floating-Point & Complex Numbers**

| Verb | Description                               | Example                                        |
| ---- | ----------------------------------------- | ---------------------------------------------- |
| `%b` | Binary exponent (for float)               | `fmt.Printf("%b", 1.5)`                        |
| `%e` | Scientific notation (lowercase)           | `fmt.Printf("%e", 1234.5678)` → `1.234568e+03` |
| `%E` | Scientific notation (uppercase)           | `fmt.Printf("%E", 1234.5678)` → `1.234568E+03` |
| `%f` | Decimal (no exponent)                     | `fmt.Printf("%f", 1234.5678)` → `1234.567800`  |
| `%F` | Same as `%f`                              |                                                |
| `%g` | Compact: `%f` or `%e` (whichever shorter) | `fmt.Printf("%g", 1234.5678)`                  |
| `%G` | Compact: `%F` or `%E`                     | `fmt.Printf("%G", 1234.5678)`                  |

---

### 5️⃣ **Strings & Slices of Bytes**

| Verb | Description                        | Example                                 |
| ---- | ---------------------------------- | --------------------------------------- |
| `%s` | String                             | `fmt.Printf("%s", "GoLang")` → `GoLang` |
| `%q` | Quoted string (Go syntax)          | `fmt.Printf("%q", "Go")` → `"Go"`       |
| `%x` | Hex (lowercase, 2 digits per byte) | `fmt.Printf("%x", "Go")` → `476f`       |
| `%X` | Hex (uppercase)                    | `fmt.Printf("%X", "Go")` → `476F`       |

---

### 6️⃣ **Pointers**

| Verb  | Description                  | Example                     |
| ----- | ---------------------------- | --------------------------- |
| `%p`  | Pointer address in hex       | `fmt.Printf("%p", &myVar)`  |
| `%#p` | Go-syntax pointer (Go 1.20+) | `fmt.Printf("%#p", &myVar)` |

---

### 7️⃣ **Width & Precision**

| Format  | Meaning                           | Example                                     |
| ------- | --------------------------------- | ------------------------------------------- |
| `%6d`   | Minimum width 6 (right-justified) | `fmt.Printf("%6d", 42)` → `"    42"`        |
| `%-6d`  | Minimum width 6 (left-justified)  | `fmt.Printf("%-6d", 42)` → `"42    "`       |
| `%06d`  | Zero-padded to width 6            | `fmt.Printf("%06d", 42)` → `"000042"`       |
| `%.2f`  | 2 decimal places                  | `fmt.Printf("%.2f", 3.14159)` → `3.14`      |
| `%6.2f` | Width 6, 2 decimals               | `fmt.Printf("%6.2f", 3.14159)` → `"  3.14"` |

---

### 8️⃣ **Extra Tricks**

| Usage                 | Meaning                                    |
| --------------------- | ------------------------------------------ |
| `fmt.Sprintf()`       | Store formatted string instead of printing |
| `fmt.Fprintf(w, ...)` | Write to a custom writer                   |
| `fmt.Errorf()`        | Create formatted errors                    |

---

✅ **Pro Tip:** If we ever forget a verb, run:

```go
fmt.Printf("%#v", myVar)
```

This will give a Go-syntax dump, which is super useful for debugging.

---

