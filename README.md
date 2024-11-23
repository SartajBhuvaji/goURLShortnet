# goURLShortner 🔗

A lightning-fast URL shortening service built with Go and Redis, designed for high performance and reliability.

[![Go](https://github.com/SartajBhuvaji/goURLShortnet/actions/workflows/go.yml/badge.svg)](https://github.com/SartajBhuvaji/goURLShortnet/actions/workflows/go.yml)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## 🚀 Features

- **Fast URL Shortening**: Efficient Base62 encoding for short, memorable URLs
- **Redis Backend**: High-performance storage and retrieval
- **RESTful API**: Clean and intuitive API endpoints
- **Analytics**: Track creation time, last access, and usage count
- **Concurrent Processing**: Handle multiple requests efficiently
- **Production Ready**: Includes tests, CI/CD, and comprehensive error handling

## 🛠️ Tech Stack

- **Go** - Core programming language
- **Redis** - Primary database
- **godotenv** - Environment configuration
- **go-redis** - Redis client for Go

## 📋 Prerequisites

- Go 1.23.2 or higher
- Redis server
- Git

## 🔧 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/SartajBhuvaji/goURLShortnet.git
   cd goURLShortner
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment**
   Create a `.env` file in the root directory:
   ```env
   REDIS_HOST="REDIS_HOST"
   REDIS_PASSWORD="REDIS_PASSWORD"
   REDIS_DB=0
   ```

## 🚦 Usage

1. **Start the server**
   ```bash
   go run main.go
   ```
   Server starts on `http://localhost:8080`

2. **API Endpoints**

   ### Shorten URL
   ```http
   POST /shorten
   Content-Type: application/json

   {
       "url": "https://www.example.com/very/long/url/that/needs/shortening"
   }
   ```
   Response:
   ```json
   {
       "short_url": "www.goURLShortner/abc123"
   }
   ```

   ### Redirect to Original URL
   ```http
   GET /redirect?url=abc123
   ```
   Response:
   ```json
   {
       "long_url": "https://www.example.com/very/long/url/that/needs/shortening"
   }
   ```

## 🧪 Testing

Run the test suite:

```bash
go test ./tests -v
```
