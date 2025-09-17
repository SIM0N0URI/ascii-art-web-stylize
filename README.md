# ASCII Art Web Stylize

A Go web application that generates stylized ASCII art from text input with multiple banner styles, now fully containerized with Docker for easy deployment and portability.

## Description

This project is a web-based ASCII art generator built with **Go (Golang)** and containerized using Docker. The application provides a simple web interface where users can input custom text, choose from multiple banner styles (`standard`, `shadow`, or `thinkertoy`), and generate stylized ASCII art directly in their browser. The project demonstrates modern web development practices combined with containerization technologies.

## Features

- Web-based ASCII art generation with custom text input
- Multiple ASCII art banner styles: `standard`, `shadow`, and `thinkertoy`
- Clean, responsive HTML interface
- HTTP server built with Go standard library only
- Multi-line text support with proper newline handling
- Comprehensive error handling and validation
- Fully containerized with Docker for easy deployment
- Optimized Docker image with proper OCI metadata

## Prerequisites

- **For Docker deployment**: Docker installed on your system
- **For local development**: Go 1.22.2 or later
- Git for cloning the repository

## Installation & Usage

### Option 1: Docker Deployment (Recommended)

#### 1. Clone the Repository
```bash
git clone https://github.com/SIM0N0URI/ascii-art-web-stylize.git
cd ascii-art-web-stylize
```

#### 2. Build the Docker Image
```bash
docker build -t ascii-art-web-stylize .
```

#### 3. Run the Container
```bash
docker run -p 8080:8080 ascii-art-web-stylize
```

### Option 2: Local Development

#### 1. Clone the Repository
```bash
git clone https://github.com/SIM0N0URI/ascii-art-web-stylize.git
cd ascii-art-web-stylize
```

#### 2. Run the Server
```bash
go run main.go
```

### Access the Application

Open your web browser and navigate to:
```
http://localhost:8080
```

### How to Use

1. Type your text in the input field
2. Select a banner style: `standard`, `shadow`, or `thinkertoy`
3. Click the **Generate** button to view the ASCII-styled text output
4. The application handles multi-line input with `\n` for line breaks

## Implementation Details

### Route Handling

| Route | Method | Description |
|-------|--------|-------------|
| `/` | GET | Serves the HTML form (index.html) |
| `/ascii-art` | POST | Handles ASCII art generation |
| `/styles/*` | GET | Serves static CSS or asset files |

### ASCII Art Generation Algorithm

1. **Input Validation**: Ensures only printable ASCII characters are allowed
2. **Font Loading**: Reads the appropriate font file from the `/banners` directory
3. **Character Conversion**: Each character of the input text is converted into ASCII art lines
4. **Multi-line Support**: Handles newlines in input to preserve formatting
5. **Output Rendering**: Combines all character art into final ASCII output

### Error Handling

| Error Type | Description | HTTP Status |
|------------|-------------|-------------|
| ✅ Success | ASCII Art successfully rendered | 200 OK |
| ⚠️ Invalid Input | Empty input, unsupported characters, bad request | 400 Bad Request |
| ❌ Not Found | Invalid banner name or missing route | 404 Not Found |
| 💥 Server Error | Template load error, file read error, etc. | 500 Internal Server Error |

## Docker Implementation

### Dockerfile Features

- **Base Image**: `golang:1.22.2-alpine` for consistency and security
- **OCI Metadata**: Proper labels for maintainability and identification
- **Multi-stage Build**: Optimized image size and security
- **Port Exposure**: Configured for port 8080
- **Working Directory**: Clean project structure setup

### Docker Objects

- **Image**: `ascii-art-web-stylize` - Contains the compiled Go application
- **Container**: Runs the web server on port 8080
- **Volumes**: Static assets and banner files properly included

## Project Structure

```
ascii-art-web-stylize/
├── main.go                 # Main application entry point
├── banners/               # ASCII art font files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
├── templates/             # HTML templates
├── static/               # CSS and static assets
├── Dockerfile            # Docker configuration
└── README.md            # This file
```

## Technical Stack

- **Language**: Go (using standard library only)
- **Containerization**: Docker with Alpine Linux
- **Web Technologies**: HTML, CSS, HTTP
- **Server**: Go HTTP server with proper routing
- **Architecture**: Clean separation of concerns

## Learning Objectives

This project demonstrates:

- **Go Web Development**: HTTP servers, routing, templating, and data handling
- **Docker Fundamentals**: Containerization, image creation, and container management
- **ASCII Art Processing**: Character mapping, font handling, and text transformation
- **Error Handling**: Comprehensive validation and user feedback
- **DevOps Practices**: Container orchestration and deployment strategies

## Authors

<table>
  <tr>
    <td align="center" width="33.33%">
      <img src="https://learn.zone01oujda.ma/git/avatars/345e97055f0fdee7afd6db5b6c009f48e9122b3febe4e400c7c6273445aa31d8" width="100px;" alt="Mohamed NOURI"/><br />
      <b>Mohamed NOURI</b>
    </td>
    <td align="center" width="33.33%">
      <img src="https://learn.zone01oujda.ma/git/avatars/2e6f4b74e2e932c1ee7e924565216ec70c64175b58a543dab9a3376a2175c602" width="100px;" alt="Othmane BENMBAREK"/><br />
      <b>Othmane BENMBAREK</b>
    </td>
    <td align="center" width="33.33%">
      <img src="https://learn.zone01oujda.ma/git/avatars/011625ea603ee5fea0a6f073f068df03f888201f810ea9bd24b7a93aa3c1e96c" width="100px;" alt="BEMAMORY Nomenjanahary Luciano Loic"/><br />
      <b>BEMAMORY Nomenjanahary Luciano Loic</b>
    </td>
  </tr>
</table>

## License

This project is part of the Zone01 curriculum and follows their guidelines and policies.

---

## 🌐 Technologies Used

- Go (Golang) - Backend development
- Docker - Containerization
- HTML/CSS - Frontend interface
- Alpine Linux - Base container image
- Standard Library only (no third-party Go packages)