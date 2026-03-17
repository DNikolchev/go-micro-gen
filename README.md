# ⚙️ go-micro-gen - Generate Clean Microservice Templates Fast

[![Download go-micro-gen](https://img.shields.io/badge/Download-go--micro--gen-brightgreen?style=for-the-badge)](https://github.com/DNikolchev/go-micro-gen/releases)

## What is go-micro-gen?

go-micro-gen is a tool that helps you create ready-to-use project templates for microservices in the Go language. It builds clean templates that include important parts like databases, Docker support, message brokers, and cloud setup. You don’t need to write everything from scratch. The templates use common patterns to keep your projects organized and easy to manage. The tool also adds tracing support with OpenTelemetry. 

This tool works well if you want to start a new Go microservice quickly, using good practices and popular tools like PostgreSQL, Kafka, MongoDB, and NATS.

## 🖥 System Requirements

Before installing go-micro-gen, make sure your system has:

- Windows 10 or later (64-bit recommended)
- At least 4 GB of free RAM
- At least 500 MB of free disk space
- Internet connection to download files
- PowerShell 5.1 or later (comes with Windows by default)
- Docker Desktop installed (optional but required to use Docker templates)
- A terminal or command prompt to run commands

## 🔧 Key Features

- Generates microservice project templates using clean architecture principles.
- Supports popular databases: PostgreSQL and MongoDB.
- Includes ready-to-use Docker configurations.
- Preconfigured for message brokers: Kafka and NATS.
- Adds OpenTelemetry tracing integration.
- Cloud deployment setup for common providers.
- Easy-to-run templates to speed up development.

## 🚀 Getting Started with go-micro-gen on Windows

Follow these steps to download and run go-micro-gen on your Windows computer.

### 1. Visit the Download Page

Go to the release page where the program files are stored:

[Download go-micro-gen Releases](https://github.com/DNikolchev/go-micro-gen/releases)

Here, you will find the latest versions of the tool. Choose the latest stable release.

### 2. Download the Windows Version

Look for the Windows executable file. It usually ends with `.exe`. For example:

- `go-micro-gen-windows.exe`  
- Or a similar name containing "windows"

Click the file name to download it to your computer.

### 3. Run the Installer or Executable

Once downloaded:

- Open the folder where you saved the file.
- Double-click the `.exe` file to start the program.
- If a security prompt appears, confirm that you want to run the program.

go-micro-gen may run directly without installation. If it has an installer, follow the on-screen instructions.

### 4. Using go-micro-gen on Your PC

After launching:

- Open Command Prompt or PowerShell.
- Navigate to the folder where go-micro-gen is located using `cd` command.
  
  ```
  cd C:\path\to\go-micro-gen
  ```

- Run the command to generate a new project template:

  ```
  go-micro-gen generate
  ```

You will be asked a few simple questions. Based on your answers, the tool will create a ready project folder with all necessary files. You can then open this in your favorite editor or IDE.

### 5. Run the Generated Microservice

Most project templates will include instructions or scripts to start your service.

- Open the project folder.
- Look for a `README.md` or instructions file.
- Use the included commands to build and run your microservice.

If your template uses Docker, make sure Docker Desktop runs on your machine.

## 📥 Download Again Anytime

Access the releases page to download the latest updates or previous versions:

[https://github.com/DNikolchev/go-micro-gen/releases](https://github.com/DNikolchev/go-micro-gen/releases)

## ⚙️ How go-micro-gen Works

go-micro-gen uses templates to create the structure of your microservice. It sets up folders, config files, and code files that follow a pattern called Clean Architecture. This pattern helps separate different parts of your service so they are easier to manage.

The generated service will include:

- Models and entities for your data.
- Database connections and migrations.
- Message broker setup to exchange information.
- Docker files for containerization.
- Telemetry and tracing integrated for monitoring.
- Base configuration for cloud deployment.

## 🛠 Installation Details for Non-Developers

You don’t need programming skills to install go-micro-gen. The main task is to download the executable and run simple commands. 

If this is your first time using command-line tools:

- Open PowerShell by clicking Start and typing `PowerShell`.
- Navigate to the folder where you downloaded go-micro-gen using `cd`.
- Enter the command as shown above to generate a project.
  
Commands to try:

- `go-micro-gen help` — see available commands.
- `go-micro-gen generate` — create a new microservice template.

## 🎯 What You Can Expect

The generated projects provide a base that developers use to build real microservices. Even if you do not know Go programming, you can share this base with your team or use it to learn more about microservices in Go.

The templates include standard tools and good design practices. This saves weeks of setup time and reduces errors.

## 🔗 Useful Links and Resources

- Release downloads: https://github.com/DNikolchev/go-micro-gen/releases
- Repository topics cover tools and integration used in go-micro-gen:
  - clean-architecture
  - cloud setup
  - Docker container support
  - Kafka and NATS message brokers
  - PostgreSQL and MongoDB databases

## 📂 File Structure Example

When you generate a project, you can expect folders like:

- `/cmd` — entry points for your service
- `/internal` — main code base separated into layers
- `/configs` — configuration files
- `/deployments` — Docker and cloud files
- `/scripts` — helper scripts for building and running
- `/docs` — documentation for your project

These folders use clear names to simplify navigation.

## 🔄 Updating go-micro-gen

To get a newer version of go-micro-gen:

- Visit the releases page again.
- Download the latest Windows executable.
- Replace the old file with the new one.
- Run the new file to generate fresh templates.

No complex uninstallation is needed.

---

[![Download go-micro-gen](https://img.shields.io/badge/Download-go--micro--gen-blue?style=for-the-badge)](https://github.com/DNikolchev/go-micro-gen/releases)