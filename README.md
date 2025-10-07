# openmrsctl 🩺

A powerful DevOps CLI tool written in Go for managing **OpenMRS deployments**, whether on Docker or bare-metal environments.

## 🚀 Features

- **Initialize Environment** — Setup OpenMRS environment configuration (`openmrsctl init`).
- **Start & Stop Services** — Start or stop OpenMRS (supports both Docker Compose and system services).
- **View Logs** — Stream or analyze OpenMRS logs with syntax highlighting and filtering.
- **Backup & Restore** — Backup and restore MySQL databases securely.
- **Health Checks** — Check the health of the database, API, and backend services.
- **Build & Deploy Automation** — Build and deploy OpenMRS in one step.
- **Config Management** — Central configuration file (`~/.openmrsctl/config.yaml`).
- **Cross-platform Support** — Works on Linux, macOS, and Windows.

## 🛠️ Installation

```bash
git clone https://github.com/yourusername/openmrsctl.git
cd openmrsctl
make build
sudo mv bin/openmrsctl /usr/local/bin/
```

## 🧩 Usage

```bash
openmrsctl init
openmrsctl start
openmrsctl logs
openmrsctl backup
openmrsctl deploy-module mymodule.omod
openmrsctl status
openmrsctl version
```

## ⚙️ Configuration

`openmrsctl init` creates a config file at:

```
~/.openmrsctl/config.yaml
```

Example:

```yaml
server_type: docker
mysql_host: localhost
mysql_user: openmrs
mysql_password: openmrs
openmrs_home: /var/lib/OpenMRS
```

## 🧱 Build Info

When built using the `Makefile`, version info is embedded automatically:

```bash
make build
openmrsctl version
# Output:
# openmrsctl v0.1.0 (commit 9b12d3a, built 2025-10-06)
```

## 🤝 Contributing

Contributions are welcome! To get started:

```bash
make dev
```

## 📄 License

MIT License © 2025 — OpenMRS DevOps Toolkit
