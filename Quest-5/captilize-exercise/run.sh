#!/usr/bin/env bash
set -euo pipefail

echo "Running capitalize..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/capitalize.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}
EOF

go run temp_run.go
rm -f temp_run.go