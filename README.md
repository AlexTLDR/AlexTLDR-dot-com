# AlexTLDR.com

[![CI](https://github.com/AlexTLDR/AlexTLDR-dot-com/actions/workflows/ci.yml/badge.svg)](https://github.com/AlexTLDR/AlexTLDR-dot-com/actions/workflows/ci.yml)

A modern, fast, and responsive personal website built with Go, templ, and DaisyUI.

## 🚀 Overview

This is the complete rewrite of my personal website [alextldr.com](https://alextldr.com). The previous version (v1) was built with Bootstrap and has been moved to the `v1` branch for historical reference. This new version leverages modern Go web technologies for better performance and developer experience.

## 🛠️ Technology Stack

- **Backend**: Go
- **Templating**: [templ](https://templ.guide/) - Type-safe HTML templates for Go
- **CSS Framework**: [DaisyUI](https://daisyui.com/) - Semantic component classes for Tailwind CSS
- **CSS Build**: [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
- **Frontend Build**: Node.js and npm for CSS compilation
- **Build Tool**: [Task](https://taskfile.dev/) - Task runner and build tool
- **Containerization**: Docker with multi-stage builds
- **Deployment**: Docker Compose

## 📋 Prerequisites

- Go 1.21 or later
- [Task](https://taskfile.dev/installation/) (recommended) or standard Go toolchain
- Node.js 16+ and npm (for Tailwind CSS and DaisyUI compilation)
- Docker and Docker Compose (for containerized deployment)
- [templ CLI](https://templ.guide/quick-start/installation) for template generation

## 🚀 Quick Start

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/AlexTLDR/AlexTLDR-dot-com.git
   cd AlexTLDR-dot-com
   ```

2. **Setup project and install dependencies**
   ```bash
   # Using Task (recommended)
   task setup

   # Or manually
   go mod download
   npm install
   ```

3. **Install templ CLI** (if not already installed)
   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

4. **Build CSS and generate templates**
   ```bash
   # Using Task (recommended)
   task css
   task templ

   # Or manually
   npm run build-css-prod

   templ generate
   ```

5. **Build and run**
   ```bash
   # Using Task (recommended)
   task run

   # Or using Go directly
   go build -o bin/AlexTLDR-dot-com ./cmd/server
   ./bin/AlexTLDR-dot-com
   ```

6. **Visit the site**
   Open your browser to `http://localhost:8080`

### Docker Development

1. **Build and run with Docker Compose**
   ```bash
   # Using Task (recommended)
   task docker-compose-up

   # Or directly
   docker-compose up --build
   ```

2. **Development with hot reload**
   ```bash
   # Using Task (recommended)
   task docker-compose-dev

   # Or directly
   docker-compose --profile dev up --build
   ```

3. **Access the application**
   The site will be available at `http://localhost:8080`

## 📁 Project Structure

```
├── bin/                    # Compiled binaries (gitignored)
├── templates/             # templ template files
├── static/               # Static assets (CSS, JS, images)
│   └── css/              # Generated CSS files
├── src/                  # Source CSS files
│   └── input.css         # Tailwind CSS input file

├── handlers/             # HTTP handlers
├── models/               # Data models
├── services/             # Business logic
├── cmd/                  # Application entry points
│   └── server/           # Main server application
├── internal/             # Private application code
├── node_modules/         # Node.js dependencies (gitignored)
├── package.json          # Node.js dependencies
├── tailwind.config.js    # Tailwind CSS configuration
├── docker-compose.yml    # Docker compose configuration
├── Dockerfile           # Multi-stage Docker build
├── Taskfile.yml         # Task automation
├── go.mod               # Go module definition
└── README.md            # This file
```

## 🔧 Available Tasks

This project uses [Task](https://taskfile.dev/) for build automation:

```bash
# Setup project and install dependencies
task setup

# Install dependencies
task install

# Build CSS with Tailwind and DaisyUI
task css

# Build CSS in watch mode
task css-watch

# Generate templ templates
task templ

# Generate templ templates in watch mode
task templ-watch

# Build the application
task build

# Build and run the application
task run

# Run in development mode with hot reload
task dev

# Docker commands
task docker-build
task docker-run
task docker-compose-up
task docker-compose-dev
task docker-compose-down

# Generate static files for deployment
task generate-static

# Build for production
task prod

# Clean build artifacts
task clean

# Run linters
task lint

# Run tests
task test

# Show all available tasks
task help
```

## 🌐 Deployment

### Docker Deployment

The application is containerized using Docker with a multi-stage build:

1. **CSS Builder stage**: Compiles Tailwind CSS and DaisyUI
2. **Go Builder stage**: Compiles the Go application and generates templ files
3. **Runtime stage**: Minimal Alpine Linux container with the compiled binary and assets

```bash
# Build the Docker image
docker build -t alextldr-dot-com .

# Run the container
docker run -p 8080:8080 alextldr-dot-com
```

### Production Deployment

The `docker-compose.yml` includes configuration for production deployment with:
- Automatic restart policy
- Health checks and monitoring
- Optimized CSS builds
- Static asset serving

## 🎨 Styling with DaisyUI and Tailwind CSS

This project uses DaisyUI and Tailwind CSS for styling:

### DaisyUI Features:
- Semantic CSS classes
- Pre-built components
- Dark/light theme support (dracula/acid themes configured)
- Responsive design utilities

### CSS Build Process:
- Source CSS in `src/input.css`
- Tailwind CSS processes templates and generates utility classes
- DaisyUI adds component classes on top of Tailwind
- Output CSS is generated to `static/css/styles.css`
- Custom animations and components defined in source CSS

### Development:
```bash
# Build CSS once
task css

# Build CSS in watch mode (rebuilds on changes)
task css-watch
```

DaisyUI is built on top of Tailwind CSS, so you can use both DaisyUI components and Tailwind utilities.

## 📝 templ Templates

Templates are written using [templ](https://templ.guide/), which provides:
- Type-safe HTML templates
- Go syntax for logic
- Compile-time template validation
- Hot reload during development

## 🔍 Development Workflow

1. **Make changes** to your Go code, templ templates, or CSS
2. **Rebuild CSS** if you modified CSS or templates:
   ```bash
   task css
   ```
3. **Generate templates** if you modified .templ files:
   ```bash
   task templ
   ```
4. **Build and run** the application:
   ```bash
   task run
   ```
5. **Test your changes** at `http://localhost:8080`

### Hot Reload Development:
```bash
# Run with auto-rebuild (CSS + Templates + Go)
task dev

# Or run individual watchers in separate terminals
task css-watch
task templ-watch
task dev-go
```

## 🐛 Debugging

- Application logs are sent to Highlight.io in production
- Use `fmt.Printf` or proper logging for local debugging
- Check Docker logs: `docker-compose logs -f`



## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- **Website**: [alextldr.com](https://alextldr.com)
- **Previous Version**: See the `v1` branch for the Bootstrap-based version
- **templ Documentation**: [templ.guide](https://templ.guide/)
- **DaisyUI Documentation**: [daisyui.com](https://daisyui.com/)

## 📧 Contact

Alex - alex@alextldr.com - [GitHub](https://github.com/AlexTLDR)

---

**Note**: This is a complete rewrite of my personal website. The previous version using Bootstrap can be found in the `v1` branch.
