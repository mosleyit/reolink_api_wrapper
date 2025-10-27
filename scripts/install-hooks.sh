#!/bin/bash
# Install git hooks for the Reolink API Wrapper project

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"
HOOKS_DIR="$PROJECT_ROOT/.git/hooks"

echo "Installing git hooks..."

# Create pre-commit hook
cat > "$HOOKS_DIR/pre-commit" << 'EOF'
#!/bin/bash
# Pre-commit hook for Reolink API Wrapper
# This hook runs formatting and linting checks before allowing a commit

set -e

echo "Running pre-commit checks..."

# Change to the Go SDK directory
cd sdk/go/reolink

# Run formatting
echo "1. Formatting code..."
gofmt -s -w .
if command -v goimports >/dev/null 2>&1; then
    goimports -w .
elif [ -f "$HOME/go/bin/goimports" ]; then
    "$HOME/go/bin/goimports" -w .
fi

# Add any formatting changes
git add -u

# Run linter
echo "2. Running linter..."
if command -v golangci-lint >/dev/null 2>&1; then
    golangci-lint run ./...
elif [ -f "$HOME/go/bin/golangci-lint" ]; then
    "$HOME/go/bin/golangci-lint" run ./...
else
    echo "Warning: golangci-lint not found. Skipping linter checks."
fi

# Run tests
echo "3. Running tests..."
go test -short ./...

echo "✓ All pre-commit checks passed!"
EOF

chmod +x "$HOOKS_DIR/pre-commit"

echo "✓ Git hooks installed successfully!"
echo ""
echo "To skip hooks for a specific commit, use: git commit --no-verify"

