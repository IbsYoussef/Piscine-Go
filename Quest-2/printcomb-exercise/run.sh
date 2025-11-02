#!/usr/bin/env bash

echo "Running your solution:"
echo ""

# Create test program
cat > temp_run.go << 'EOF'
package main
EOF

sed 's/package student//' student/printcomb.go >> temp_run.go

cat >> temp_run.go << 'EOF'

func main() {
	PrintComb()
}
EOF

go run temp_run.go
rm -f temp_run.go