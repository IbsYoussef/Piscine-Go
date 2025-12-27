# Piscine-Go

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Shell Script](https://img.shields.io/badge/Shell_Script-121011?style=for-the-badge&logo=gnu-bash&logoColor=white)

This repo documents my 3-week 01 Founders Piscine-Go journey—an intensive selection pool learning Go from scratch through structured quests. Each quest targets a specific Go concept, progressing from fundamentals to advanced problem-solving. Includes all solutions and explanations from my Nov 2022 intake.

---

## 📋 Table of Contents

- [🏊 About Piscine-Go](#-about-piscine-go)
- [📚 Learning Journey](#-learning-journey)
  - [Week 1: Foundations 🌱](#week-1-foundations-)
  - [Week 2: Core Concepts 🔧](#week-2-core-concepts-)
  - [Week 3: Advanced Topics 🚀](#week-3-advanced-topics-)
  - [Hackathon: Challenge Mode 🏆](#hackathon-challenge-mode-)
- [🎯 Quest Overview](#-quest-overview)
- [📁 Repository Structure](#-repository-structure)
- [🚀 How to Use This Repository](#-how-to-use-this-repository)
  - [Exploring Exercises 🔍](#exploring-exercises-)
  - [Writing Your Solutions ✍️](#writing-your-solutions-️)
  - [Testing with Scripts 🧪](#testing-with-scripts-)
- [💡 Key Learnings](#-key-learnings)

---

## 🏊 About Piscine-Go

The **Piscine-Go** is a 3-week intensive coding bootcamp that serves as the selection process for the 01 Founders fellowship program. The term "Piscine" (French for "swimming pool") reflects the immersive, sink-or-swim nature of the experience.

**What makes it unique:**

- 🤝 **Peer-to-peer learning:** No traditional teachers—students learn collaboratively
- 🎮 **Quest-based progression:** Each quest focuses on a specific programming concept
- 💻 **Practical application:** Every exercise is hands-on, building real programming skills
- ⏱️ **Time-boxed challenges:** 3 weeks to progress from zero Go knowledge to advanced problem-solving

The Piscine tests not just coding ability, but also:

- 🧩 Problem-solving under pressure
- 📖 Self-directed learning
- 💬 Collaboration and communication
- 💪 Perseverance and time management

---

## 📚 Learning Journey

### Week 1: Foundations 🌱

**Quests 1-3** | Shell Scripting, Go Basics, String Manipulation

The first week establishes fundamental programming concepts:

- Command-line proficiency through shell scripting
- Go syntax, control structures, and basic I/O
- String manipulation and basic algorithms

**Key Milestone:** Transition from shell to Go, understanding compiled languages

### Week 2: Core Concepts 🔧

**Quests 4-7** | Recursion, Advanced Strings, Arrays & Slices

Week two deepens understanding of Go's type system and data structures:

- Iterative vs recursive problem-solving approaches
- Memory management with pointers
- Working with collections and dynamic arrays
- ASCII encoding and character manipulation

**Key Milestone:** Building complex algorithms and understanding memory references

### Week 3: Advanced Topics 🚀

**Quests 8-9** | Command-Line Programs, Higher-Order Functions

The final week introduces professional development practices:

- Building complete command-line applications
- Debugging techniques and error handling
- Functional programming concepts
- Function composition and callbacks

**Key Milestone:** Creating production-ready programs with proper file I/O

### Hackathon: Challenge Mode 🏆

**Optional Advanced Challenges** | Algorithm Design & Optimization

The Hackathon quest tests cumulative knowledge with unique challenges covering:

- 🔐 **Cryptography & Encoding:** ROT14 cipher, character transformations
- 🧮 **Mathematical Algorithms:** Median finding, Collatz conjecture, prime detection
- ⚡ **Bitwise Operations:** Power of 2 detection, active bit counting
- 📊 **Data Structures:** Hash maps, frequency analysis, unmatched element detection
- 🔗 **Pointer Manipulation:** Multi-level indirection, memory puzzles
- 📝 **String Processing:** Word frequency analysis, complex parsing
- 🔄 **Type Conversion:** Integer-to-string and reverse conversions
- 🐛 **Debugging:** Fixing broken code with multiple error types

These exercises emphasize algorithmic thinking, optimization, and creative problem-solving beyond standard curriculum.

---

## 🎯 Quest Overview

### Quest 1: Shell Scripting 🐚

**Focus:** Command-line basics, file operations, pipes, scripting fundamentals

```bash
# Finding and counting files
find . -type f -name "*.sh" | wc -l
```

---

### Quest 2: Go Fundamentals 🔤

**Focus:** Syntax, loops, conditionals, `z01.PrintRune`, basic I/O

```go
func PrintDigits() {
    for digit := '0'; digit <= '9'; digit++ {
        z01.PrintRune(digit)
    }
    z01.PrintRune('\n')
}
```

---

### Quest 3: String Manipulation ✂️

**Focus:** String operations, type conversions, basic algorithms

```go
func StrLen(s string) int {
    count := 0
    for range s {
        count++
    }
    return count
}
```

---

### Quest 4: Iteration & Recursion 🔁

**Focus:** Iterative vs recursive approaches, factorial, Fibonacci, prime numbers

```go
// Recursive approach
func Fibonacci(n int) int {
    if n < 2 {
        return n
    }
    return Fibonacci(n-1) + Fibonacci(n-2)
}

// Iterative approach
func IterativePower(base, power int) int {
    result := 1
    for i := 0; i < power; i++ {
        result *= base
    }
    return result
}
```

---

### Quest 5: Strings, Bytes & Runes 🔠

**Focus:** ASCII encoding, character manipulation, deep string processing

```go
func ToUpper(s string) string {
    result := ""
    for _, r := range s {
        if r >= 'a' && r <= 'z' {
            result += string(r - 32)  // ASCII magic!
        } else {
            result += string(r)
        }
    }
    return result
}
```

---

### Quest 6: Command-Line Arguments 💻

**Focus:** `os.Args`, argument parsing, program parameters

```go
func main() {
    args := os.Args[1:]  // Skip program name

    for i, arg := range args {
        fmt.Printf("Argument %d: %s\n", i, arg)
    }
}
```

---

### Quest 7: Arrays & Slices 📦

**Focus:** `append`, `make`, slice manipulation, dynamic arrays

```go
func MakeRange(min, max int) []int {
    if min >= max {
        return nil
    }

    size := max - min
    result := make([]int, size)

    for i := 0; i < size; i++ {
        result[i] = min + i
    }
    return result
}
```

---

### Quest 8: Program Debugging 🐛

**Focus:** File I/O, stdin/stdout, error handling, practical debugging

```go
func DisplayFile(filename string) {
    content, err := os.ReadFile(filename)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        return
    }
    fmt.Print(string(content))
}
```

---

### Quest 9: Higher-Order Functions 🎭

**Focus:** Functions as parameters, callbacks, functional programming

```go
func Map(f func(int) bool, arr []int) []bool {
    result := []bool{}
    for _, val := range arr {
        result = append(result, f(val))
    }
    return result
}

// Usage
func IsPrime(n int) bool {
    // ... prime checking logic
}

mapped := Map(IsPrime, []int{1, 2, 3, 4, 5})
// Result: [false, true, true, false, true]
```

---

### Hackathon: Advanced Challenges 🏆

**Focus:** Bitwise ops, pointers, algorithms, optimization, problem-solving

```go
// Bitwise: Check if power of 2
func IsPowerOfTwo(n int) bool {
    return n > 0 && (n&(n-1)) == 0
}

// Pointers: Multi-level indirection
func Enigma(a ***int, b *int, c *******int, d ****int) {
    tempA := ***a
    ***a = *b
    *b = ****d
    *******c = tempA
    ****d = *******c
}

// Hash maps: Find unpaired element
func Unmatch(arr []int) int {
    freq := make(map[int]int)
    for _, num := range arr {
        freq[num]++
    }
    for num, count := range freq {
        if count%2 != 0 {
            return num
        }
    }
    return -1
}
```

---

## 📁 Repository Structure

```
Piscine-Go/
├── Quest-1/              # 🐚 Shell scripting fundamentals
├── Quest-2/              # 🔤 Go basics and PrintRune
├── Quest-3/              # ✂️ String manipulation
├── Quest-4/              # 🔁 Iteration and recursion
├── Quest-5/              # 🔠 Advanced string/byte/rune operations
├── Quest-6/              # 💻 Command-line arguments
├── Quest-7/              # 📦 Arrays and slices
├── Quest-8/              # 🐛 File I/O and debugging
├── Quest-9/              # 🎭 Higher-order functions
└── Hackathon/            # 🏆 Advanced algorithmic challenges
```

### Exercise Structure 📂

**Each exercise follows a consistent format when navigating quest exercises:**

```
exercise-name-exercise/
├── README.md             # 📖 Problem description and requirements
├── solutions/            # ✅ Reference solutions
│   └── filename.go       # Complete working solution
├── student/              # ✏️ Your workspace
│   └── filename.go       # Where you write your code
├── run.sh                # ▶️ Quick execution script
└── test.sh               # 🧪 Automated testing script
```

**File Types:**

- **Single file exercises:** `solutions/filename.go` and `student/filename.go`
- **Program exercises:** `solutions/programname/main.go` and `student/programname/main.go`

**Example navigation:**

```bash
cd Quest-5/alphacount-exercise
ls
# Output: README.md  run.sh  solutions/  student/  test.sh
```

---

## 🚀 How to Use This Repository

### Exploring Exercises 🔍

1. **Navigate to a quest:**

```bash
   cd Quest-5
```

2. **List available exercises:**

```bash
   ls
```

3. **Read the problem:**

```bash
   cd exercise-name-exercise
   cat README.md
```

### Writing Your Solutions ✍️

1. **Open the student file:**

```bash
   # For function exercises
   nano student/filename.go

   # For program exercises
   nano student/programname/main.go
```

2. **Write your solution** following the problem requirements

3. **Save and exit**

### Testing with Scripts 🧪

#### Quick Run (see your output)

```bash
./run.sh
```

This executes your solution and displays the output without comparing to expected results.

#### Automated Testing (verify correctness)

```bash
./test.sh
```

This compares your output against the reference solution:

- ✅ **Green checkmark:** Your solution matches!
- ❌ **Red X:** Output differs—shows expected vs actual

**Example output:**

```bash
Testing alphacount...

✓ Test passed!

Output:
  Test 1: 26
  Test 2: 0
  Test 3: 13
```

#### Tips for Testing 💡

- Start with `./run.sh` to see basic output
- Use `./test.sh` to verify correctness
- Read error messages carefully—they show what's expected vs what you produced
- Test scripts automatically clean up temporary files

### Revision Workflow 🔄

**Recommended approach:**

1. 📖 Read the `README.md` to understand the problem
2. 💭 Try solving it yourself in the `student/` directory
3. ▶️ Run `./run.sh` to check basic functionality
4. 🧪 Run `./test.sh` to verify correctness
5. 👀 If stuck, peek at `solutions/` for hints (not the full solution!)
6. 🔍 Compare approaches after solving successfully

**Active recall strategy:**

- ⏰ Wait a few days after completing an exercise
- 🔄 Return and solve it from scratch without looking
- ✅ If you can rebuild it, you've internalized the concept

---

## 💡 Key Learnings

### Technical Skills 🛠️

- ✅ Go syntax and idioms
- ✅ Memory management with pointers
- ✅ Algorithm design (sorting, searching, recursion)
- ✅ String/byte/rune manipulation
- ✅ File I/O and error handling
- ✅ Functional programming concepts
- ✅ Command-line program development
- ✅ Bitwise operations and optimization
- ✅ Test-driven development mindset

### Problem-Solving 🧩

- ✅ Breaking down complex problems
- ✅ Edge case identification
- ✅ Debugging systematically
- ✅ Pattern recognition across exercises
- ✅ Optimization techniques
- ✅ Reading and understanding documentation

### Professional Development 📈

- ✅ Self-directed learning
- ✅ Code organization and structure
- ✅ Writing maintainable solutions
- ✅ Persistence through difficult challenges
- ✅ Learning from mistakes
