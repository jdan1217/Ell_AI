# Ell_AI – AI-Powered Penetration Testing Proxy

## Overview

Ell_AI is an open-source, AI-powered web application security testing proxy being developed from scratch in Go.

The goal is to build a security testing platform inspired by tools like Burp Suite, with an emphasis on integrating AI-driven vulnerability discovery, request analysis, and penetration testing automation.

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

| Component       | Description                                            |
| --------------- | ------------------------------------------------------ |
| Proxy Engine    | TCP listener and HTTP/HTTPS traffic handling           |
| Intercept       | Pause, inspect, modify, forward, or drop requests      |
| Target          | Organize traffic by destination and application        |
| Repeater        | Modify and replay HTTP requests                        |
| Intruder        | Automate parameter manipulation and testing            |
| AI Analysis     | Analyze traffic and identify potential vulnerabilities |
| AI Agents       | Coordinate specialized security testing workflows      |
| Traffic History | Store and review captured HTTP transactions            |
| Dashboard       | Web-based interface for interacting with the proxy     |

## Technology Stack

**Backend**

- Go
- Go standard library (`net`, `net/http`, `bufio`, `crypto/tls`, `crypto/x509`)
- PostgreSQL (planned)

**Frontend (Planned)**

- React
- TypeScript

**AI Integration (Planned)**

- Large Language Models (LLMs)
- Open-weight models
- Agent-based workflows
- API integrations

**Deployment (Planned)**

- Docker
- Docker Compose

## Architecture

The project is being developed as a modular monolith, with individual components responsible for specific functionality.

```
Ell_AI/
├── main.go
├── dial/
│   └── main.go
├── internal/
│   ├── config/
│   │   ├── config.go
│   │   └── config_test.go
│   └── proxy/
│       ├── server.go
│       ├── request.go
│       ├── request_test.go
│       └── transport.go
├── go.mod
└── README.md
```

Additional packages and directories will be introduced as the project grows.

The architecture will evolve as additional functionality is implemented.

## Development Roadmap

Development is being approached incrementally, prioritizing a reliable proxy engine before introducing more advanced functionality.

### Project Foundation

- Initialize Go project and repository
- Implement initial configuration
- Create TCP listener
- Accept concurrent client connections using goroutines
- Implement shared HTTP transport
- Create separate request parsing functionality
- Implement HTTP request parsing unit tests
- Configure and run Go race detector tests

### HTTP Proxy Engine

- Parse incoming HTTP requests
- Forward HTTP requests to upstream servers using `http.Transport.RoundTrip()`
- Receive upstream HTTP responses
- Return upstream responses to the original client
- Verify end-to-end HTTP forwarding using curl
- Support multiple HTTP requests over persistent client connections
- Improve request and response handling, including error cases
- Implement request and response inspection
- Implement request modification before forwarding

### HTTPS Support

- Implement HTTPS CONNECT tunneling
- Implement local certificate authority
- Generate and manage certificates for HTTPS interception
- Implement TLS interception and upstream TLS connections

### Traffic Management and Security Testing

- Build request/response interception engine
- Implement pause, forward, modify, and drop functionality
- Implement traffic history and storage
- Develop Repeater and Target functionality
- Develop Intruder functionality
- Implement scope restrictions and egress policies

### Frontend and AI Integration

- Build frontend dashboard
- Develop backend API for frontend communication
- Integrate AI models and agent workflows
- Implement AI-assisted request and response analysis
- Implement automated vulnerability testing

### Deployment

- Containerize with Docker
- Implement Docker Compose deployment
- Develop setup and usage documentation

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

## Current Status

**Active Development – HTTP Proxy Engine**

The initial TCP and HTTP proxy functionality is now implemented.

The proxy currently supports:

- TCP listener and concurrent client connection handling.
- HTTP/1.x request parsing using Go's `net/http` package.
- Upstream HTTP request forwarding using a shared `http.Transport`.
- Receiving and returning upstream HTTP responses.
- Basic request logging, including methods, URLs, and headers.
- Unit testing for HTTP request parsing and configuration.
- Race detector testing.

End-to-end HTTP forwarding has been successfully tested using curl against `example.com`.

**Current Limitations:**

- Only basic HTTP forwarding is implemented.
- Each client connection currently handles one request before closing.
- Persistent HTTP connections are not yet supported.
- HTTPS tunneling and TLS interception have not yet been implemented.
- Request interception, modification, and traffic history remain planned features.

The next development milestone is implementing persistent HTTP connections before moving toward HTTPS support and the interception engine.



Add clear setup and usage instructions