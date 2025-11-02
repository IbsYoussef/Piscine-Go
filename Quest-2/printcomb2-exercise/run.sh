#!/usr/bin/env bash

echo "Running your solution:"
echo ""

cat > temp_run.go << 'EOF'
package main
EOF

sed 's/package student//' student/printcomb2.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	PrintComb2()
}
EOF

go run temp_run.go
rm -f temp_run.go