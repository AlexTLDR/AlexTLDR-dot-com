# AlexTLDR.com

A modern, fast, and responsive personal website built with Go, templ, and DaisyUI.

## 🚀 Overview

This is the complete rewrite of my personal website [alextldr.com](https://alextldr.com). The previous version (v1) was built with Bootstrap and has been moved to the `v1` branch for historical reference. This new version leverages modern Go web technologies for better performance and developer experience.

## 🛠️ Technology Stack

- **Backend**: Go 1.24.4
- **Templating**: [templ](https://templ.guide/) - Type-safe HTML templates for Go
- **CSS Framework**: [DaisyUI](https://daisyui.com/) - Semantic component classes for Tailwind CSS
- **Build Tool**: [Task](https://taskfile.dev/) - Task runner and build tool
- **Containerization**: Docker with multi-stage builds
- **Logging**: Fluentd with Highlight.io integration
- **Deployment**: Docker Compose

## 📋 Prerequisites

- Go 1.24.4 or later
- [Task](https://taskfile.dev/installation/) (recommended) or standard Go toolchain
- Docker and Docker Compose (for containerized deployment)
- Node.js and npm (for DaisyUI/Tailwind CSS compilation)

## 🚀 Quick Start

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/AlexTLDR/AlexTLDR-dot-com.git
   cd AlexTLDR-dot-com
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Install templ CLI** (if not already installed)
   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

4. **Generate templates** (when you have .templ files)
   ```bash
   templ generate
   ```

5. **Build and run**
   ```bash
   # Using Task (recommended)
   task run

   # Or using Go directly
   go build -o bin/AlexTLDR-dot-com
   ./bin/AlexTLDR-dot-com
   ```

6. **Visit the site**
   Open your browser to `http://localhost:3000`

### Docker Development

1. **Build and run with Docker Compose**
   ```bash
   docker-compose up --build
   ```

2. **Access the application**
   The site will be available at `http://localhost:3000`

## 📁 Project Structure

```
├── bin/                    # Compiled binaries (gitignored)
├── templates/             # templ template files
├── static/               # Static assets (CSS, JS, images)
├── handlers/             # HTTP handlers
├── models/               # Data models
├── services/             # Business logic
├── cmd/                  # Application entry points
├── internal/             # Private application code
├── docker-compose.yml    # Docker compose configuration
├── Dockerfile           # Multi-stage Docker build
├── Taskfile.yml         # Task automation
├── go.mod               # Go module definition
└── README.md            # This file
```

## 🔧 Available Tasks

This project uses [Task](https://taskfile.dev/) for build automation:

```bash
# Build the application
task build

# Build and run the application
task run

# Run tests (when implemented)
# task test
```

## 🌐 Deployment

### Docker Deployment

The application is containerized using Docker with a multi-stage build:

1. **Builder stage**: Compiles the Go application
2. **Runtime stage**: Minimal Alpine Linux container with the compiled binary

```bash
# Build the Docker image
docker build -t alextldr-dot-com .

# Run the container
docker run -p 3000:3000 alextldr-dot-com
```

### Production Deployment

The `docker-compose.yml` includes configuration for production deployment with:
- Automatic restart policy
- Fluentd logging integration with Highlight.io
- Health checks and monitoring

## 🎨 Styling with DaisyUI

This project uses DaisyUI for styling, which provides:
- Semantic CSS classes
- Pre-built components
- Dark/light theme support
- Responsive design utilities

DaisyUI is built on top of Tailwind CSS, so you can use both DaisyUI components and Tailwind utilities.

## 📝 templ Templates

Templates are written using [templ](https://templ.guide/), which provides:
- Type-safe HTML templates
- Go syntax for logic
- Compile-time template validation
- Hot reload during development

## 🔍 Development Workflow

1. **Make changes** to your Go code or templ templates
2. **Generate templates** if you modified .templ files:
   ```bash
   templ generate
   ```
3. **Build and run** the application:
   ```bash
   task run
   ```
4. **Test your changes** at `http://localhost:3000`

## 🐛 Debugging

- Application logs are sent to Highlight.io in production
- Use `fmt.Printf` or proper logging for local debugging
- Check Docker logs: `docker-compose logs -f`

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Make your changes
4. Commit your changes: `git commit -m 'Add some amazing feature'`
5. Push to the branch: `git push origin feature/amazing-feature`
6. Open a Pull Request

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Links

- **Website**: [alextldr.com](https://alextldr.com)
- **Previous Version**: See the `v1` branch for the Bootstrap-based version
- **templ Documentation**: [templ.guide](https://templ.guide/)
- **DaisyUI Documentation**: [daisyui.com](https://daisyui.com/)

## 📧 Contact

Alex - [GitHub](https://github.com/AlexTLDR)

---

**Note**: This is a complete rewrite of my personal website. The previous version using Bootstrap can be found in the `v1` branch.