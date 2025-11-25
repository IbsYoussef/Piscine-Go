#!/usr/bin/env bash
set -euo pipefail

echo "Running makerange..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/makerange.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
EOF

go run temp_run.go
rm -f temp_run.go