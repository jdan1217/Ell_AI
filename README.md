# Ell_AI – AI-Powered Penetration Testing Proxy

## Overview

Ell_AI is an open-source, AI-powered web application security testing proxy being developed from scratch in Go.

Inspired by tools like Burp Suite, the goal is to build a security testing platform that combines HTTP/HTTPS traffic inspection and manipulation with AI-driven vulnerability discovery, request analysis, and penetration testing automation.

Rather than building another traditional intercepting proxy, Ell_AI aims to combine the capabilities of a web security testing platform with modern AI agents and automated security analysis.

## Project Goals

The primary objective is to develop a functional penetration testing proxy that provides visibility and control over HTTP/HTTPS traffic while leveraging AI to assist with security assessments.

Key goals include:

- **Traffic Interception:** Capture, inspect, modify, and forward HTTP/HTTPS requests and responses.
- **Request Analysis:** Analyze HTTP methods, URLs, headers, parameters, cookies, and request bodies.
- **HTTPS Interception:** Implement certificate authority functionality to inspect encrypted HTTPS traffic.
- **Traffic Management:** Maintain request history and provide the ability to replay, modify, and manipulate requests.
- **AI Integration:** Incorporate LLMs and AI agents to assist with vulnerability identification, request analysis, and security testing.
- **Automated Security Testing:** Develop testing workflows for identifying common web application vulnerabilities.
- **Extensibility:** Build a modular architecture that supports additional testing capabilities and integrations over time.

## Planned Features

| Component | Description |
|---|---|
| Proxy Engine | TCP listener and HTTP/HTTPS traffic handling |
| Intercept | Pause, inspect, modify, forward, or drop requests |
| Target | Organize traffic by destination and application |
| Repeater | Modify and replay HTTP requests |
| Intruder | Automate parameter manipulation and testing |
| AI Analysis | Analyze traffic and identify potential vulnerabilities |
| AI Agents | Coordinate specialized security testing workflows |
| Traffic History | Store and review captured HTTP transactions |
| Dashboard | Web-based interface for interacting with the proxy |

## Technology Stack

### Backend
- Go
- Go standard library (`net`, `net/http`, `bufio`, `io`, `crypto/tls`, `crypto/x509`)
- PostgreSQL (planned)

### Frontend (Planned)
- React
- TypeScript

### AI Integration (Planned)
- Large Language Models (LLMs)
- Open-weight models
- Agent-based workflows
- API integrations

### Deployment (Planned)
- Docker
- Docker Compose

## Architecture

The project is being developed as a modular monolith, with individual packages responsible for specific functionality.

```text
Ell_AI/
├── main.go
├── dial/
│   └── main.go
├── internal/
│   ├── config/
│   │   ├── config.go
│   │   └── config_test.go
│   ├── proxy/
│   │   ├── server.go
│   │   ├── request.go
│   │   ├── request_test.go
│   │   └── transport.go
│   └── traffic/
│       ├── entry.go
│       ├── entry_test.go
│       ├── request.go
│       ├── request_test.go
│       ├── response.go
│       └── response_test.go
├── go.mod
└── README.md
```

Additional packages and directories will be introduced as the project grows.

## Getting Started

### Prerequisites

- Go 1.26.4 or later
- Git
- A terminal (PowerShell, Command Prompt, or equivalent)

### Installation

Clone the repository:

```bash
git clone https://github.com/jdan1217/Ell_AI.git
cd Ell_AI
```

Download dependencies:

```bash
go mod download
```

### Running the Proxy

Start the proxy using:

```bash
go run .
```

By default, the proxy listens on:

```text
127.0.0.1:8080
```

The default configuration currently includes:

| Setting | Default |
|---|---|
| Host | `127.0.0.1` |
| Port | `8080` |
| Log Level | `INFO` |
| Interception Enabled | `false` |

### Configuring a Client

Configure your HTTP client to use the local proxy.

For example, using curl on Windows:

```powershell
curl.exe --proxy http://127.0.0.1:8080 http://example.com/
```

To test multiple HTTP requests over the same client connection:

```powershell
curl.exe --proxy http://127.0.0.1:8080 http://example.com/ http://example.org/
```

The proxy will log incoming connections and HTTP request information, including methods, URLs, schemes, hosts, paths, and headers.

**Note:** The current implementation supports basic HTTP forwarding. HTTPS tunneling and TLS interception have not yet been implemented.

### Running Tests

Run all project tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test ./... -v
```

Run tests with the Go race detector:

```bash
go test -race ./...
```

Run tests for individual packages:

```bash
go test ./internal/proxy/... -v
go test ./internal/traffic/... -v
go test ./internal/config/... -v
```

## Development Roadmap

Development is being approached incrementally, prioritizing a reliable proxy engine before introducing more advanced functionality.

### 1. Project Foundation

- [x] Initialize Go project and repository.
- [x] Implement initial configuration.
- [x] Create TCP listener.
- [x] Accept concurrent client connections using goroutines.
- [x] Implement shared HTTP transport.
- [x] Separate request parsing functionality.
- [x] Implement HTTP request parsing unit tests.
- [x] Configure and run Go race detector tests.

### 2. HTTP Proxy Engine

- [x] Parse incoming HTTP requests.
- [x] Forward HTTP requests to upstream servers using `http.Transport.RoundTrip()`.
- [x] Receive upstream HTTP responses.
- [x] Return upstream responses to the original client.
- [x] Verify end-to-end HTTP forwarding using curl.
- [x] Support multiple HTTP requests over persistent client connections.
- [x] Implement request and response error handling.
- [x] Implement basic request inspection and logging.
- [ ] Implement request modification before forwarding.
- [ ] Improve transport configuration, timeouts, and connection management.

### 3. Traffic Data Model

- [x] Create a structured traffic entry model.
- [x] Implement cryptographically random traffic IDs.
- [x] Capture transaction timestamps and processing durations.
- [x] Capture HTTP request methods, URLs, headers, bodies, and protocols.
- [x] Capture HTTP response status codes, headers, bodies, and protocols.
- [x] Deep-copy HTTP header values to preserve captured data.
- [x] Restore request and response bodies after reading them.
- [x] Implement unit tests for request data conversion.
- [x] Implement unit tests for response data conversion.
- [x] Implement unit tests for traffic entry construction.
- [ ] Integrate traffic capture into the proxy request lifecycle.
- [ ] Implement in-memory traffic history.
- [ ] Implement traffic filtering and retrieval.

### 4. HTTPS Support

- [ ] Implement HTTPS CONNECT tunneling.
- [ ] Implement local certificate authority.
- [ ] Generate and manage certificates for HTTPS interception.
- [ ] Implement TLS interception and upstream TLS connections.

### 5. Traffic Management and Security Testing

- [ ] Build request/response interception engine.
- [ ] Implement pause, forward, modify, and drop functionality.
- [ ] Implement traffic history and persistent storage.
- [ ] Develop Repeater and Target functionality.
- [ ] Develop Intruder functionality.
- [ ] Implement scope restrictions and egress policies.

### 6. Frontend and AI Integration

- [ ] Build frontend dashboard.
- [ ] Develop backend API for frontend communication.
- [ ] Integrate AI models and agent workflows.
- [ ] Implement AI-assisted request and response analysis.
- [ ] Implement automated vulnerability testing.

### 7. Deployment

- [ ] Containerize with Docker.
- [ ] Implement Docker Compose deployment.
- [ ] Develop setup and usage documentation.

## Current Status

**Active Development – HTTP Proxy Engine and Traffic Data Model**

The initial TCP listener, HTTP forwarding functionality, persistent connection handling, and structured traffic data models have been implemented.

### Implemented Functionality

- TCP listener and concurrent client connection handling.
- HTTP/1.x request parsing using Go's `net/http` package.
- Upstream HTTP request forwarding using a shared `http.Transport`.
- Receiving and returning upstream HTTP responses.
- Persistent HTTP connections supporting multiple requests over the same client connection.
- Basic request inspection and logging.
- Structured request and response data models.
- Traffic entry construction with unique identifiers, timestamps, and request durations.
- Request and response body capture with stream restoration.
- Deep-copying HTTP headers to preserve captured transaction data.
- Unit tests covering configuration, request parsing, forwarding, connection handling, and traffic data models.

End-to-end HTTP forwarding has been successfully tested using curl against external HTTP endpoints. Persistent connections have also been verified by sending multiple requests over the same client TCP connection.

### Current Limitations

- Only basic HTTP forwarding is implemented.
- HTTPS CONNECT tunneling and TLS interception are not yet supported.
- Request modification and interception functionality have not yet been implemented.
- Traffic entries can be constructed, but automatic traffic history capture is not yet integrated into the proxy.
- Traffic history is not yet persisted or exposed through an API.
- Frontend and AI functionality remain planned features.

### Next Development Milestone

Integrate the traffic data model into the proxy request lifecycle.

The next objective is to automatically capture every forwarded HTTP transaction, measure request processing duration, and maintain an in-memory collection of captured traffic entries. This will establish the foundation for traffic history, request inspection, and future interception functionality.

## Learning Objectives

An important part of this project is developing the underlying engineering knowledge by building the application rather than relying on existing proxy frameworks.

Areas of focus include:

- Go programming, concurrency, goroutines, and error handling.
- TCP/IP networking and HTTP protocol internals.
- TLS, certificate management, and HTTPS interception.
- Secure software architecture and modular application development.
- Database design and API development.
- AI integration, LLM orchestration, and agentic workflows.
- Automated application security testing.

The project is intended to serve as both a functional security engineering tool and a practical demonstration of these skills.

## Security and Responsible Use

Ell_AI is intended for authorized security assessments, application testing, and educational research.

Future testing functionality will incorporate scope restrictions, authorization controls, and configurable egress policies to help prevent testing outside approved environments.

Users are responsible for ensuring they have permission to test the systems they target.