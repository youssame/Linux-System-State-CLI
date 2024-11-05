
# Linux System State CLI - Golang Learning Project

## Project Overview

This project is a command-line interface (CLI) tool built in Golang that provides real-time information about the state of your Linux system. It fetches and displays system metrics such as CPU usage, memory consumption, storage usage, and more. The goal is to practice and improve Go programming skills while developing a useful tool for monitoring system performance.

## Features

- **CPU Usage**: Displays current CPU usage for all cores.
- **Memory Usage**: Shows current memory consumption, including free and used memory.
- **Storage Information**: Displays available and used disk space for all mounted file systems.
- **System Uptime**: Shows how long the system has been running.
- **Network Usage**: Displays current network traffic (optional).
- **Real-Time Monitoring**: Option to display updates continuously with a defined refresh interval.

## Requirements

- Golang (version 1.18 or higher)
- Linux-based operating system
- Libraries for system information retrieval, such as `gopsutil`

## Usage

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/linux-system-state-cli.git
   cd linux-system-state-cli
   ```

2. Install dependencies:

   ```bash
   go get -u github.com/shirou/gopsutil
   ```

3. Build the CLI:

   ```bash
   go build -o sysstate
   ```

4. Run the CLI:

   ```bash
   ./sysstate
   ```

5. Optional: Use flags for specific metrics:

   ```bash
   ./sysstate --cpu --memory --storage
   ```

## Future Enhancements

- Add network traffic monitoring.
- Allow output in different formats (JSON, XML, etc.).
- Include more detailed breakdowns for CPU and memory statistics.
- Enable remote system monitoring via SSH.

## Contributing

Feel free to contribute to the project by submitting issues or pull requests. Any ideas for improvements or additional features are welcome!
