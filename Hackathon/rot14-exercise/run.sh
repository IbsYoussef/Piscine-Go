#!/usr/bin/env bash
set -euo pipefail

echo "Running rot14..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "github.com/01-edu/z01"
EOF

sed '/^package student/d' student/rot14.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test1 := Rot14("Hello! How are You?")
	test2 := Rot14("abc")
	test3 := Rot14("xyz")

	for _, r := range test1 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	for _, r := range test2 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')

	for _, r := range test3 {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
EOF

go run temp_run.go
rm -f temp_run.go