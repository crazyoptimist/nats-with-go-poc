# NATS Server POC

Hello, NATS

### Producer

```
go run ./cmd/producer/main.go <subject> <message>
```

Example:

```
go run ./cmd/producer/main.go "greet.joe" "hello from joe"
```

### Consumer

```
go run ./cmd/consumer/main.go <subject>
```

Example:

```
go run ./cmd/consumer/main.go "greet.*"
```
