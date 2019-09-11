# Rebel Services

Metrics.

Tasks:
- Import expvar to expose custom variables to external clients.
- Add middleware to track number of requests, number of errors, and current goroutine count.
- http://localhost:6060/debug/vars

## File Changes :
- main.go
- routing/route.go

## New File :
- middleware/metrics.go

## Adding Dependency :