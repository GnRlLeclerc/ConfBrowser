# Build the project
build:
  bun run build
  go build .

# Check everything
check:
  golangci-lint run
  bun run prettier --check .
  bun run eslint .
