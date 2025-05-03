# Skypier VPN - Build and Run Instructions

## Prerequisites

- Go 1.16 or higher
- WSL (Windows Subsystem for Linux) with Ubuntu 22.04
- Git
- Make (optional, for using Makefile)

## Step-by-Step Build Process

### 1. Setting Up WSL

1. Open Windows Terminal or PowerShell
2. Launch WSL:
```bash
wsl
```

### 2. Installing Dependencies

1. Update package lists:
```bash
sudo apt update
```

2. Install required packages:
```bash
sudo apt install -y git golang-go make
```

3. Verify Go installation:
```bash
go version
```

### 3. Building the Project

1. Clone the repository:
```bash
git clone https://github.com/SkyPierIO/skypier-vpn.git
cd skypier-vpn
```

2. Build the project:
```bash
go build -o skypier-vpn cmd/skypier-vpn/main.go
```

Expected output:
```
[No output if successful]
```

3. Make the binary executable:
```bash
chmod +x skypier-vpn
```

### 4. Running the Project

1. Start the server with root privileges:
```bash
sudo ./skypier-vpn
```

Expected output:
```
───────────────────────────────────────────────────
  ____    _                      _               
 / ___|  | | __  _   _   _ __   (_)   ___   _ __ 
 \___ \  | |/ / | | | | | '_ \  | |  / _ \ | '__|
  ___) | |   <  | |_| | | |_) | | | |  __/ | |   
 |____/  |_|\_\  \__, | | .__/  |_|  \___| |_|   
                 |___/  |_|                      
───────────────────────────────────────────────────
2025/05/03 23:38:11 Checking if directory /etc/skypier exists
2025/05/03 23:38:11 Generating identity...
2025/05/03 23:38:11 initializing DHT...
[Connection logs to bootstrap peers]
2025/05/03 23:38:14 ┌────────────────────────────────────────────────────┐
2025/05/03 23:38:14 │ VPN UI available at http://skypier.localhost:8081/ │
2025/05/03 23:38:14 └────────────────────────────────────────────────────┘
```

### 5. Accessing the Web Interface

1. Get your WSL IP address:
```bash
hostname -I | awk '{print $1}'
```

2. Access the web interface in your Windows browser:
```
http://<WSL_IP>:8081/
```

3. Access Swagger documentation:
```
http://<WSL_IP>:8081/swagger/index.html
```

## Troubleshooting

### Common Build Issues

1. **Permission Denied Errors**:
   - Ensure you're running build commands with proper permissions
   - Check that the target directory is writable

2. **Go Module Issues**:
   - If you see module-related errors, run:
   ```bash
   go mod tidy
   go mod download
   ```

3. **WSL Network Issues**:
   - If you can't access the web interface, verify WSL IP address
   - Check Windows Defender Firewall settings
   - Ensure WSL network bridge is properly configured

### Common Runtime Issues

1. **IPv6 Disable Error**:
   - Error message: `Failed to disable IPv6: permission denied`
   - Solution: Run the application with `sudo`

2. **Directory Creation Error**:
   - Error message: `Failed to create directory /etc/skypier`
   - Solution: Run the application with `sudo`

3. **Connection Issues**:
   - If the web interface is not accessible, verify:
     - Server is running with `sudo`
     - Correct WSL IP address is being used
     - Port 8081 is not blocked

## Development

### Project Structure

- `cmd/skypier-vpn/`: Main application entry point
- `pkg/`: Core packages and libraries
- `docs/`: Documentation and API specifications

### Building for Development

1. Install development dependencies:
```bash
go mod download
```

2. Build with debug information:
```bash
go build -gcflags="-N -l" -o skypier-vpn cmd/skypier-vpn/main.go
```

### Testing

Run the test suite:
```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the LICENSE file for details. 