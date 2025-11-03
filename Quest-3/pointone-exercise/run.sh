#!/usr/bin/env bash

echo "Running your solution:"
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/pointone.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	n := 0
	PointOne(&n)
	fmt.Println(n)
}
EOF

go run temp_run.go
rm -f temp_run.go