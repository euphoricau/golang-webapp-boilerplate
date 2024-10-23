# Go WebApp Boilerplate

This project serves as both a learning activity and an attempt to implement a quick and efficient business-oriented Server-side Rendered dashboard/systems using Golang in a OOP-like fashion. It provides a basic structure and common functionalities such as authentication and authorization to help jumpstart projects.
SPA like functionality will be available with the implementation of htmx.

Feed back and constructive critisism welcome 😊

### Tech Stack
- go-1.23
- gin
- sqlite
- htmx
- bootstrap

## Features


- ✅ Basic project structure
- ✅ Routing with Gin
- ✅ Basic Authentication
- ✅ User Registration
- ✅ Bootstrap 
- ❌ Email Verification
- ❌ Environment variable support 
- ❌ Implement OAuth authentication
- ❌ Implement LDAP
- ❌ Add Multi-Factor Authentication (MFA)
- ❌ Set up support for other SQL 
- ❌ Add Docker support for containerization
- ❌ Enhance documentation with examples
- ❌ Optimize performance and security
- ❌ Logging 
- ❌ HTTPS/TLS 

> **Disclaimer:** The features listed above are not in any particular order of priority or implementation.

## Getting Started

### Prerequisites

- Go 1.23 or higher

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/euphoricau/go-webapp-boilerplate.git
    cd go-webapp-boilerplate
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the Application

1. Run the application:
    ```sh
    go run main.go
    ```

2. Open your browser and navigate to `http://localhost:8080`.

### Development

Personally, I would suggest using a hot reload tool such as [air](https://github.com/cosmtrek/air) to improve your development workflow. Air watches for changes in your source code and automatically reloads the application, saving you the hassle of restarting the server manually.

## Project Structure

```
.
├── controllers     // 
├── db              // Database config
├── middleware      // Middlewares -> 
├── models          // Data Objects
├── main.go         // Entry point
├── routes          // for all routing 
└── templates       // views
    └── partials    // singular components

```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

