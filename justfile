# Build the project
build:
  bun run build
  go build .

# Check everything
check:
  bun run prettier --check .
  bun run eslint .
