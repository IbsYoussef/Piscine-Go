#!/usr/bin/env bash

echo "Running your solution:"
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/divmod.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
EOF

go run temp_run.go
rm -f temp_run.go