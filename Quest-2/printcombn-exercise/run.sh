#!/usr/bin/env bash

echo "Running your solution (testing with n=1, n=2, n=3):"
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

# Remove package line AND any import statements from student file
sed '/^package student/d; /^import/d' student/printcombn.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println("n=1:")
	PrintCombN(1)
	fmt.Println("\nn=2:")
	PrintCombN(2)
	fmt.Println("\nn=3:")
	PrintCombN(3)
}
EOF

go run temp_run.go
rm -f temp_run.go