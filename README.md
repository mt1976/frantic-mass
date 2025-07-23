# Frantic Mass

Frantic Mass is a health and fitness tracking application written in Go. It provides tools for tracking weight, BMI, goals, and other health-related metrics, with a modular architecture and support for web-based interfaces.

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
