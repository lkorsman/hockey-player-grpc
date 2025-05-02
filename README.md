# Hockey Player gRPC Server

A simple Go-based gRPC web server for managing hockey player data. This project serves as a learning exercise to explore the fundamentals of Go, Protocol Buffers, and gRPC-based communication in a backend service context.

## 📌 Features

* gRPC service defined with Protocol Buffers
* Server implementation in Go
* Basic CRUD operations for player records (in-memory)
* Organized structure for future testing and extension

## 🛠 Technologies

* **Go** – Core server logic and concurrency
* **gRPC** – Remote procedure call framework
* **Protocol Buffers** – Data serialization

## 🚀 Getting Started

1. **Clone the repository:**

```bash
git clone https://github.com/lkorsman/hockey-player-grpc.git
cd hockey-player-grpc
```

2. **Generate the gRPC code:**

Make sure you have `protoc` and the Go plugins installed.

```bash
protoc --go_out=. --go-grpc_out=. player.proto
```

3. **Run the server:**

```bash
go run server/main.go
```

## 📚 Learning Goals

* Understand gRPC and protobuf basics
* Practice idiomatic Go development
* Explore service-oriented architecture concepts

## 🧱 Future Improvements

* Add persistent storage (e.g., PostgreSQL or CockroachDB)
* Implement client in Go or another language
* Add test coverage using Go’s `testing` package
