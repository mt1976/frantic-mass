# Frantic Mass

Frantic Mass is a modular, cross-platform health and fitness tracking application written in Go. It allows users to track weight, BMI, and fitness goals via a web interface. The backend uses StormDB for data storage and the chi router for HTTP endpoints. The frontend leverages Pico.css and Bootstrap Icons for a clean, responsive UI. The project supports Docker for easy deployment, includes utility scripts, and is structured for extensibility with clear separation of data access, business logic, background jobs, and web handlers. Configuration is managed via TOML files, and the app provides endpoints for BMI calculation, enrichment, and weight projection.

[![Go](https://github.com/mt1976/frantic-mass/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/mt1976/frantic-mass/actions/workflows/go.yml)
[![Dependabot Updates](https://github.com/mt1976/frantic-mass/actions/workflows/dependabot/dependabot-updates/badge.svg)](https://github.com/mt1976/frantic-mass/actions/workflows/dependabot/dependabot-updates)

## Tech Stack
- **Language:** Go (Golang)
- **Database:** StormDB (see `data/database/`), file-based storage for backups/dumps
- **Web:** Native Go HTTP server, HTML templates, minimal JavaScript
- **Router:** [chi](https://github.com/go-chi/chi) — Lightweight, idiomatic HTTP router for Go, used for defining routes and middleware in the web server
- **Node Packages:**
  - [Pico.css](https://picocss.com/) — Minimal CSS framework for clean, responsive design (`node_modules/@picocss/pico`)
  - [Bootstrap Icons](https://icons.getbootstrap.com/) — Icon library for UI elements (`node_modules/bootstrap-icons`)
- **Containerization:** Docker, Docker Compose
- **Scripts:** Bash (`build.sh`, `clear.sh`, `docs.sh`)
- **Configuration:** TOML files in `data/config/`
- **Platform Support:** Cross-platform builds (macOS, Linux, Windows)

## Features
- Track weight, BMI, and other health metrics
- Set and monitor fitness goals
- Web interface for data entry and visualization
- Modular DAO structure for extensibility
- Data export and backup functionality
- Docker support for easy deployment

## Project Structure
- `main.go` — Application entry point
- `app/dao/` — Data access objects for various entities (weight, user, goal, etc.)
- `app/functions/` — Core business logic and utility functions
- `app/jobs/` — Background jobs and scheduled tasks
- `app/types/` — Type definitions for core entities
- `app/web/` — Web handlers, viewProviders, contentProviders, and static assets
- `data/` — Backups, configs, database, logs, and dumps
- `exec/` — Platform-specific executables
- `res/` — HTML templates, images, and JavaScript resources
- `docker-compose.yml` — Docker orchestration file
- `build.sh`, `clear.sh`, `docs.sh` — Utility scripts

## Getting Started

### Prerequisites
- Go 1.20+
- Docker (optional, for containerized deployment)

### Build and Run

#### Using Go
```bash
git clone https://github.com/mt1976/frantic-mass.git
cd frantic-mass
go build -o frantic-mass main.go
./frantic-mass
```

#### Using Docker
```bash
docker-compose up --build
```

### Configuration
- Edit configuration files in `data/config/` as needed (e.g., `common.toml`).

### Database
- The default StormDB database is located at `data/database/`.
- Backups and dumps are stored in `data/backups/` and `data/dumps/`.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

## Author
- [mt1976](https://github.com/mt1976)
