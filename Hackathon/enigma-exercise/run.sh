#!/usr/bin/env bash
set -euo pipefail

echo "Running enigma..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/enigma.go >> temp_run.go

cat >> temp_run.go << 'EOF'

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

	fmt.Println("=== Before Enigma ===")
	fmt.Printf("a: %d\n", ***a)
	fmt.Printf("b: %d\n", *b)
	fmt.Printf("c: %d\n", *******c)
	fmt.Printf("d: %d\n", ****d)
	fmt.Println()

	Enigma(a, b, c, d)

	fmt.Println("=== After Enigma ===")
	fmt.Printf("a: %d\n", ***a)
	fmt.Printf("b: %d\n", *b)
	fmt.Printf("c: %d\n", *******c)
	fmt.Printf("d: %d\n", ****d)
}
EOF

go run temp_run.go
rm -f temp_run.go