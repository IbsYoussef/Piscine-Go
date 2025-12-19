#!/usr/bin/env bash
set -euo pipefail

echo "Running map..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

# Include IsPrime helper
cat >> temp_run.go << 'EOF'

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
EOF

sed '/^package student/d' student/map.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}
EOF

go run temp_run.go
rm -f temp_run.go