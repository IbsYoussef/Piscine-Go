#!/usr/bin/env bash
set -euo pipefail

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Testing enigma..."
echo ""

{
cat > temp_student.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/enigma.go >> temp_student.go

cat >> temp_student.go << 'EOF'

func main() {
	x := 5
	y := &x
	z := &y
	a := &z
	
	w := 2
	b := &w
	
	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j
	
	k := 6
	l := &k
	m := &l
	n := &m
	d := &n

	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
	Enigma(a, b, c, d)
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}
EOF

cat > temp_solution.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package solutions/d' solutions/enigma.go >> temp_solution.go

cat >> temp_solution.go << 'EOF'

func main() {
	x := 5
	y := &x
	z := &y
	a := &z
	
	w := 2
	b := &w
	
	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j
	
	k := 6
	l := &k
	m := &l
	n := &m
	d := &n

	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
	Enigma(a, b, c, d)
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}
EOF

student_output=$(go run temp_student.go)
solution_output=$(go run temp_solution.go)

rm -f temp_student.go temp_solution.go

} 2>/dev/null

if [ "$student_output" = "$solution_output" ]; then
    echo -e "${GREEN}✓ Test passed!${NC}"
    echo ""
    echo "Before Enigma:"
    echo "  a: $(echo "$student_output" | sed -n '1p')"
    echo "  b: $(echo "$student_output" | sed -n '2p')"
    echo "  c: $(echo "$student_output" | sed -n '3p')"
    echo "  d: $(echo "$student_output" | sed -n '4p')"
    echo ""
    echo "After Enigma:"
    echo "  a: $(echo "$student_output" | sed -n '5p')"
    echo "  b: $(echo "$student_output" | sed -n '6p')"
    echo "  c: $(echo "$student_output" | sed -n '7p')"
    echo "  d: $(echo "$student_output" | sed -n '8p')"
    exit 0
else
    echo -e "${RED}✗ Test failed${NC}"
    echo ""
    echo "Expected:"
    echo "  Before: $(echo "$solution_output" | sed -n '1p') $(echo "$solution_output" | sed -n '2p') $(echo "$solution_output" | sed -n '3p') $(echo "$solution_output" | sed -n '4p')"
    echo "  After:  $(echo "$solution_output" | sed -n '5p') $(echo "$solution_output" | sed -n '6p') $(echo "$solution_output" | sed -n '7p') $(echo "$solution_output" | sed -n '8p')"
    echo ""
    echo "Got:"
    echo "  Before: $(echo "$student_output" | sed -n '1p') $(echo "$student_output" | sed -n '2p') $(echo "$student_output" | sed -n '3p') $(echo "$student_output" | sed -n '4p')"
    echo "  After:  $(echo "$student_output" | sed -n '5p') $(echo "$student_output" | sed -n '6p') $(echo "$student_output" | sed -n '7p') $(echo "$student_output" | sed -n '8p')"
    exit 1
fi
```

---

**Now the output will be much cleaner:**

**run.sh output:**
```
Running enigma...

=== Before Enigma ===
a: 5
b: 2
c: 7
d: 6

=== After Enigma ===
a: 2
b: 6
c: 5
d: 7
```

**test.sh output (success):**
```
Testing enigma...

✓ Test passed!

Before Enigma:
  a: 5
  b: 2
  c: 7
  d: 6

After Enigma:
  a: 2
  b: 6
  c: 5
  d: 7