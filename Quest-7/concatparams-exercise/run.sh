#!/usr/bin/env bash
set -euo pipefail

echo "Running concatparams..."
echo ""

cat > temp_run.go << 'EOF'
package main

import "fmt"
EOF

sed '/^package student/d' student/concatparams.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
EOF

go run temp_run.go
rm -f temp_run.go