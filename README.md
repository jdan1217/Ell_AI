# Ell_AI – AI-Powered Penetration Testing Proxy

Ell_AI is an open-source, AI-powered web application security testing proxy being developed from scratch in Go.

Inspired by tools like Burp Suite, Ell_AI aims to combine HTTP/HTTPS traffic inspection and manipulation with AI-driven vulnerability discovery, request analysis, and automated penetration testing workflows.

The long-term goal is to build a modular security testing platform that combines traditional application security testing capabilities with modern AI models and agent-based workflows.

**Project Status:** Active Development – HTTP Proxy Engine, Traffic Capture, and Management API

---

## Project Goals

The primary objective is to develop a functional penetration testing proxy that provides visibility and control over HTTP/HTTPS traffic while leveraging AI to assist with security assessments.

Key objectives include:

- **Traffic Interception:** Capture, inspect, modify, and forward HTTP/HTTPS requests and responses.
- **Request Analysis:** Analyze HTTP methods, URLs, headers, parameters, cookies, and request bodies.
- **HTTPS Interception:** Implement certificate authority functionality to inspect encrypted HTTPS traffic.
- **Traffic Management:** Maintain request history and provide functionality to replay, modify, and manipulate requests.
- **AI Integration:** Incorporate LLMs and AI agents for vulnerability identification, request analysis, and security testing.
- **Automated Security Testing:** Develop workflows for identifying common web application vulnerabilities.
- **Extensibility:** Maintain a modular architecture that supports additional testing capabilities and integrations.

## Current Features

The following functionality has been implemented:

| Component | Functionality | Status |
|---|---|---|
| Proxy Engine | TCP listener and concurrent client connections | Implemented |
| HTTP Forwarding | Parse, forward, and return HTTP requests/responses | Implemented |
| Persistent Connections | Support multiple HTTP requests over a client connection | Implemented |
| Traffic Capture | Capture structured HTTP request/response data | Implemented |
| Traffic History | Store captured transactions in memory | Implemented |
| Management API | Retrieve traffic through HTTP endpoints | Implemented |
| Traffic Lookup | Retrieve individual transactions by unique ID | Implemented |
| Pagination | Limit and offset support for traffic history | Implemented |
| HTTPS Interception | CONNECT tunneling and TLS interception | Planned |
| Request Interception | Pause, modify, forward, or drop requests | Planned |
| Repeater / Intruder | Request replay and automated parameter testing | Planned |
| AI Analysis | AI-assisted vulnerability discovery and analysis | Planned |
| Frontend | Web-based traffic inspection dashboard | Planned |

## Architecture

Ell_AI follows a modular monolith architecture, separating proxy functionality, traffic management, configuration, and API functionality into individual packages.

```text
Ell_AI/
├── main.go
├── dial/
│   └── main.go
├── internal/
│   ├── api/
│   │   ├── server.go
│   │   └── server_test.go
│   ├── config/
│   │   ├── config.go
│   │   └── config_test.go
│   ├── proxy/
│   │   ├── server.go
│   │   ├── request.go
│   │   ├── request_test.go
│   │   ├── server_test.go
│   │   └── transport.go
│   └── traffic/
│       ├── entry.go
│       ├── entry_test.go
│       ├── history.go
│       ├── history_test.go
│       ├── request.go
│       ├── request_test.go
│       ├── response.go
│       └── response_test.go
├── go.mod
└── README.md
```

### Component Responsibilities

**Proxy (`internal/proxy`)**
- Accepts incoming TCP connections.
- Parses HTTP/1.x requests.
- Forwards requests to upstream servers using Go's `http.Transport`.
- Receives upstream responses and returns them to clients.
- Handles concurrent client connections and persistent HTTP connections.
- Integrates traffic capture into the request lifecycle.

**Traffic (`internal/traffic`)**
- Defines structured request, response, and traffic entry models.
- Generates cryptographically random transaction IDs.
- Records timestamps and request processing durations.
- Captures request and response headers and bodies.
- Maintains thread-safe in-memory traffic history.
- Supports retrieving all entries, paginated entries, and individual entries by ID.

**API (`internal/api`)**
- Exposes HTTP endpoints for accessing captured traffic.
- Provides individual traffic entry retrieval.
- Supports pagination through query parameters.
- Uses JSON serialization for communication with future frontend components.

**Configuration (`internal/config`)**
- Centralizes proxy configuration and default settings.

## Technology Stack

### Backend
- Go 1.26.4
- Go standard library:
  - `net`
  - `net/http`
  - `bufio`
  - `io`
  - `crypto/tls`
  - `crypto/x509`
  - `sync`
  - `encoding/json`

### Frontend (Planned)
- React
- TypeScript

### Data Storage
- In-memory traffic history (implemented)
- PostgreSQL (planned)

### AI Integration (Planned)
- Large Language Models
- Open-weight models
- Agent-based workflows
- API integrations

### Deployment (Planned)
- Docker
- Docker Compose

---

## Getting Started

### Prerequisites

- Go 1.26.4 or later
- Git
- PowerShell, Command Prompt, or another terminal

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

### Running Ell_AI

Start the application:

```bash
go run .
```

By default, Ell_AI starts two local services:

| Service | Address | Purpose |
|---|---|---|
| Proxy | `127.0.0.1:8080` | Accepts and forwards HTTP traffic |
| Management API | `127.0.0.1:8081` | Provides access to captured traffic |

The default proxy configuration:

| Setting | Default |
|---|---|
| Host | `127.0.0.1` |
| Port | `8080` |
| Log Level | `INFO` |
| Interception Enabled | `false` |

### Configuring a Client

Configure an HTTP client to use Ell_AI as its proxy.

Example using curl on Windows:

```powershell
curl.exe --proxy http://127.0.0.1:8080 http://example.com/
```

The proxy forwards the request to the upstream server and returns the response to the original client.

Captured traffic is automatically added to the in-memory traffic history.

---

## Management API

Ell_AI includes a separate HTTP management API for accessing captured traffic.

The API listens on:

```text
http://127.0.0.1:8081
```

### Available Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/api/traffic` | Retrieve captured traffic |
| GET | `/api/traffic/{id}` | Retrieve an individual traffic entry |

### Retrieve Traffic History

```powershell
curl.exe http://127.0.0.1:8081/api/traffic
```

Returns captured traffic entries in JSON format.

### Retrieve an Individual Traffic Entry

```powershell
curl.exe http://127.0.0.1:8081/api/traffic/{id}
```

Replace `{id}` with the unique ID of a captured transaction.

### Pagination

Traffic history supports pagination through `limit` and `offset` query parameters.

Example:

```powershell
curl.exe "http://127.0.0.1:8081/api/traffic?limit=20&offset=0"
```

| Parameter | Description | Default |
|---|---|---|
| `limit` | Maximum number of entries returned | `50` |
| `offset` | Number of entries to skip | `0` |

Pagination constraints:

- Minimum limit: `1`
- Maximum limit: `500`
- Offset must be nonnegative.
- Invalid pagination parameters return HTTP `400 Bad Request`.

### Traffic Entry Structure

Each captured transaction contains:

**Request Data**
- HTTP method
- URL
- Headers
- Body
- Protocol

**Response Data**
- Status code
- Headers
- Body
- Protocol

**Transaction Metadata**
- Unique transaction ID
- Timestamp
- Processing duration

Request and response bodies are represented as Base64-encoded values in JSON because the underlying Go fields use `[]byte`.

Traffic history currently resides in memory and is lost when the application stops.

---

## Running Tests

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

Run individual package tests:

```bash
go test ./internal/proxy/... -v
go test ./internal/traffic/... -v
go test ./internal/api/... -v
go test ./internal/config/... -v
```

Testing currently covers:
- Configuration defaults.
- HTTP request parsing.
- HTTP forwarding and connection handling.
- Traffic entry construction.
- Request and response data conversion.
- Thread-safe traffic history operations.
- Traffic pagination.
- Management API endpoints.
- API pagination and invalid parameter handling.

---

## Development Roadmap

Development is being approached incrementally, prioritizing a reliable proxy engine and traffic management foundation before introducing advanced security testing and AI functionality.

### 1. Project Foundation

- [x] Initialize Go project and repository.
- [x] Implement initial configuration.
- [x] Create TCP listener.
- [x] Accept concurrent client connections using goroutines.
- [x] Implement shared HTTP transport.
- [x] Separate request parsing functionality.
- [x] Implement request parsing unit tests.
- [x] Configure and run Go race detector tests.

### 2. HTTP Proxy Engine

- [x] Parse incoming HTTP requests.
- [x] Forward requests using `http.Transport.RoundTrip()`.
- [x] Receive upstream HTTP responses.
- [x] Return upstream responses to clients.
- [x] Verify end-to-end HTTP forwarding using curl.
- [x] Support persistent HTTP connections.
- [x] Implement basic request inspection and logging.
- [ ] Implement request modification before forwarding.
- [ ] Improve transport configuration, timeouts, and connection management.

### 3. Traffic Data Model and History

- [x] Create structured traffic entry models.
- [x] Generate cryptographically random traffic IDs.
- [x] Capture timestamps and processing durations.
- [x] Capture request methods, URLs, headers, bodies, and protocols.
- [x] Capture response status codes, headers, bodies, and protocols.
- [x] Deep-copy HTTP header values.
- [x] Restore request and response bodies after reading.
- [x] Implement traffic entry unit tests.
- [x] Integrate traffic capture into the proxy lifecycle.
- [x] Implement thread-safe in-memory traffic history.
- [x] Implement traffic retrieval by ID.
- [x] Implement paginated traffic retrieval.
- [x] Test history operations and concurrent access.

### 4. Management API

- [x] Implement separate management API server.
- [x] Implement `GET /api/traffic`.
- [x] Implement `GET /api/traffic/{id}`.
- [x] Implement JSON serialization.
- [x] Implement pagination using limit and offset.
- [x] Validate pagination parameters.
- [x] Add API endpoint and pagination tests.
- [ ] Implement traffic filtering by method, status code, host, and other attributes.
- [ ] Add traffic deletion and history management functionality.

### 5. HTTPS Support

- [ ] Implement HTTPS CONNECT tunneling.
- [ ] Implement local certificate authority.
- [ ] Generate and manage certificates for HTTPS interception.
- [ ] Implement TLS interception and upstream TLS connections.

### 6. Traffic Interception and Security Testing

- [ ] Build request/response interception engine.
- [ ] Implement pause, forward, modify, and drop functionality.
- [ ] Develop Repeater functionality.
- [ ] Develop Target functionality.
- [ ] Develop Intruder functionality.
- [ ] Implement scope restrictions and egress policies.

### 7. Frontend and AI Integration

- [ ] Build web-based frontend dashboard.
- [ ] Integrate frontend with management API.
- [ ] Implement traffic inspection interface.
- [ ] Integrate AI models and agent workflows.
- [ ] Implement AI-assisted request and response analysis.
- [ ] Develop automated vulnerability testing workflows.

### 8. Storage and Deployment

- [ ] Implement persistent traffic storage.
- [ ] Integrate PostgreSQL.
- [ ] Containerize with Docker.
- [ ] Implement Docker Compose deployment.
- [ ] Develop setup and usage documentation.

---

## Current Status and Limitations

Ell_AI currently has a functional HTTP forwarding proxy with automatic traffic capture and a separate management API.

Implemented functionality has been tested end-to-end using curl, and package tests include execution with Go's race detector.

Current limitations:

- Only basic HTTP forwarding is supported.
- HTTPS CONNECT tunneling and TLS interception are not implemented.
- Requests cannot yet be paused, modified, or dropped.
- Traffic history is stored in memory and is not persistent.
- Traffic filtering has not yet been implemented.
- There is no frontend dashboard.
- AI-driven analysis and automated security testing remain planned features.

### Next Development Milestone

Expand traffic management functionality by implementing filtering and additional retrieval capabilities.

This will establish a stronger foundation for the future traffic inspection interface, Repeater, and AI-powered analysis workflows.

---

## Learning Objectives

Ell_AI is being built from scratch to develop practical engineering knowledge across several areas:

- Go programming, concurrency, goroutines, and error handling.
- TCP/IP networking and HTTP protocol internals.
- TLS, certificate management, and HTTPS interception.
- Secure software architecture and modular application development.
- Database design and API development.
- AI integration, LLM orchestration, and agentic workflows.
- Automated application security testing.

The project is intended to serve as both a functional security engineering tool and a practical demonstration of backend development, application security, networking, and AI engineering.

## Security and Responsible Use

Ell_AI is intended for authorized security assessments, application testing, and educational research.

Future testing functionality will incorporate scope restrictions, authorization controls, and configurable egress policies to help prevent testing outside approved environments.

Users are responsible for ensuring they have permission to test the systems they target.