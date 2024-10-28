# Distributed Chat System

## 🚀 Overview
A high-performance, distributed chat system implemented in Go, featuring a microservices architecture with HTTP and RPC servers. The system is designed for high scalability and reliable message delivery, utilizing Redis for message storage and Kitex for efficient RPC communication.

## 🏗️ Architecture
```
┌─────────────────┐         ┌─────────────────┐
│                 │   RPC   │                 │
│   HTTP Server   ├────────►│   RPC Server    │
│   (API Layer)   │         │ (Business Logic)│
└────────┬────────┘         └────────┬────────┘
         │                           │
         │                           │
         │         ┌────────────────►│
         │         │                 │
    ┌────▼─────────▼────┐            │
    │                   │            │
    │  Redis Storage    │◄───────────┘
    │                   │
    └───────────────────┘
```

### Key Components
1. **HTTP Server**
   - Handles external API requests
   - Implements RESTful endpoints
   - Manages request validation
   - Routes requests to RPC server

2. **RPC Server**
   - Processes business logic
   - Manages message operations
   - Handles data persistence
   - Implements message delivery logic

3. **Redis Storage**
   - Stores chat messages
   - Manages message cursors
   - Enables efficient message retrieval
   - Supports scalable data access

## ⚡ Features
- **High Performance**: Supports 20+ concurrent users
- **Message Persistence**: Reliable message storage in Redis
- **Pull-based Messaging**: Efficient message retrieval system
- **Scalable Architecture**: Microservices-based design
- **Docker Support**: Easy deployment with containers
- **Configurable Limits**: Adjustable message fetch limits

## 🛠️ Technical Stack
- **Language**: Go 1.19+
- **RPC Framework**: Kitex
- **Storage**: Redis
- **Containerization**: Docker
- **CI/CD**: GitHub Actions

## 📝 API Documentation

### 1. Health Check
```bash
# Check server status
GET /ping

# Response (200 OK)
{
    "message": "pong"
}
```

### 2. Send Message
```bash
# Send a message
POST /api/send
Content-Type: application/json

{
    "Chat": "a:b",
    "Text": "Hello World",
    "Sender": "a"
}

# Response (200 OK)
{
    "status": "success"
}
```

### 3. Pull Messages
```bash
# Retrieve messages
GET /api/pull
Content-Type: application/json

{
    "Chat": "a:b",
    "Cursor": 0,
    "Limit": 20,
    "Reverse": false
}

# Response (200 OK)
{
    "messages": [
        {
            "chat": "a:b",
            "text": "Hello World",
            "sender": "a",
            "send_time": 1684744610
        }
    ]
}
```

## 🚀 Getting Started

### Prerequisites
- Docker and Docker Compose
- Go 1.19 or higher (for local development)
- Redis (handled by Docker Compose)

### Quick Start
```bash
# Clone the repository
git clone https://github.com/yourusername/chat-system.git

# Start the application
docker-compose up -d

# Verify the setup
curl http://localhost:8080/ping
```

### Project Structure
```
.
├── .github/
│   └── workflows/
│       └── test.yml        # CI configuration
├── http-server/
│   ├── handler/           # HTTP request handlers
│   ├── middleware/        # HTTP middlewares
│   └── main.go           # HTTP server entry point
├── rpc-server/
│   ├── handler/          # RPC request handlers
│   ├── storage/          # Redis operations
│   └── main.go          # RPC server entry point
├── docker-compose.yml    # Container orchestration
├── idl_http.proto       # HTTP API definitions
└── idl_rpc.thrift      # RPC service definitions
```

## 🔧 Development

### Local Setup
```bash
# Start dependencies
docker-compose up redis -d

# Start RPC server
cd rpc-server
go run main.go

# Start HTTP server
cd http-server
go run main.go
```

### Running Tests
```bash
# Run unit tests
go test ./...

# Run integration tests
docker-compose -f docker-compose.test.yml up --abort-on-container-exit
```

## 📈 Performance

### Benchmarks
- Supports 20+ concurrent users
- Message delivery latency < 100ms
- Can handle 1000+ messages per second

### Scalability Features
1. Microservices architecture enables horizontal scaling
2. Redis clustering support for data distribution
3. Stateless servers for easy replication
