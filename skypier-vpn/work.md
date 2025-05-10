# Skypier VPN Testing Guide

This guide provides step-by-step instructions for testing the Skypier VPN application without requiring a subscription.

## Table of Contents
1. [Building the Application](#1-building-the-application)
2. [Configuration Setup](#2-configuration-setup)
3. [Network Configuration](#3-network-configuration)
4. [Service Setup](#4-service-setup)
5. [Testing the Node](#5-testing-the-node)
6. [API Testing](#6-api-testing)
7. [Monitoring](#7-monitoring)
8. [Troubleshooting](#8-troubleshooting)
9. [Cleanup](#9-cleanup)

## 1. Building the Application

```bash
# Clone the repository
git clone -b v2.0 https://github.com/thealonemusk/WarpNet.git
cd skypier-vpn

# Build both the client and node binaries
go build -o skypier-vpn cmd/skypier-vpn/main.go
go build -o skypier-vpn-node cmd/skypier-vpn-node/main.go
```

## 2. Configuration Setup

```bash
# Create the config directory
sudo mkdir -p /etc/skypier

# Create and edit the config file
sudo nano /etc/skypier/config.json
```

Add this content to the config file:
```json
{
    "nickname": "TestNode",
    "debug": true,
    "privateKey": "",
    "advertisePrivateAddresses": false,
    "swaggerEnabled": true,
    "DHTDiscovery": true
}
```

## 3. Network Configuration

### Enable IP Forwarding
```bash
# Enable IP forwarding
echo "net.ipv4.ip_forward = 1" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p
```

### Set up NAT Rules

Choose either iptables or UFW method:

####  iptables
```bash
# Install iptables-persistent
sudo apt-get install iptables-persistent

# Add NAT rules
sudo iptables -t nat -A POSTROUTING -s 10.1.1.0/24 -o eth0 -j MASQUERADE
sudo iptables -A FORWARD -i eth0 -o skypier0 -m state --state RELATED,ESTABLISHED -j ACCEPT
sudo iptables -A FORWARD -s 10.1.1.0/24 -o eth0 -j ACCEPT

# Save the rules
sudo netfilter-persistent save
```


## 4. Service Setup

### Create Systemd Service
```bash
# Create the service file
sudo nano /etc/systemd/system/skypier-vpn-node.service
```

Add this content:
```ini
[Unit]
Description=Skypier VPN Node
After=network.target

[Service]
ExecStart=/home/thealonemusk/WarpNet/skypier-vpn/skypier-vpn-node
Restart=always
RestartSec=10
User=root
Group=root
Environment=PATH=/usr/bin:/usr/local/bin
Environment=NODE_ENV=production

[Install]
WantedBy=multi-user.target
```

### Start the Service
```bash
# Enable and start the service
sudo systemctl daemon-reload
sudo systemctl enable skypier-vpn-node
sudo systemctl start skypier-vpn-node

# Check the status
sudo systemctl status skypier-vpn-node
```

## 5. Testing the Node

### Basic Testing
```bash
# Check the logs
sudo journalctl -u skypier-vpn-node -f

# Get your node's IP address
hostname -I | awk '{print $1}'
```

### Web Interface
- Open your browser and go to: `http://localhost:8081`
- Or use your WSL IP: `http://<WSL_IP>:8081`

## 6. API Testing

When debug mode is enabled, you can test the following endpoints:

### Swagger UI
- Access: `http://localhost:8081/swagger/index.html`

### Key Endpoints
- `GET /api/v0/me` - Get node details
- `GET /api/v0/id` - Get peer ID
- `GET /api/v0/status` - Check VPN status
- `GET /api/v0/connected_peers_count` - Get connected peers

## 7. Monitoring

### System Monitoring
```bash
# Check the VPN interface
ip addr show skypier0

# Check routing table
ip route show

# Check NAT rules
sudo iptables -t nat -L -n -v
```

### Peer Discovery
```bash
# Check DHT discovery logs
sudo journalctl -u skypier-vpn-node -f | grep "DHT"

# Check connected peers
curl http://localhost:8081/api/v0/connected_peers_count
```

## 8. Troubleshooting

### Common Issues and Solutions

#### Service Status
```bash
# Check if the service is running
sudo systemctl status skypier-vpn-node

# Check if the port is listening
sudo netstat -tulpn | grep 8081
```

#### Network Configuration
```bash
# Check if IP forwarding is enabled
cat /proc/sys/net/ipv4/ip_forward

# Check if the TUN interface is created
ip addr show skypier0
```

#### Logs
```bash
# Check system logs for errors
sudo journalctl -u skypier-vpn-node -n 50
```

## 9. Cleanup

### Remove Service and Configuration
```bash
# Stop the service
sudo systemctl stop skypier-vpn-node

# Disable the service
sudo systemctl disable skypier-vpn-node

# Remove the service file
sudo rm /etc/systemd/system/skypier-vpn-node.service

# Reload systemd
sudo systemctl daemon-reload
```

### Remove Network Configuration
```bash
# Remove NAT rules (if using iptables)
sudo iptables -t nat -D POSTROUTING -s 10.1.1.0/24 -o eth0 -j MASQUERADE
sudo netfilter-persistent save
```

## Notes

- All commands requiring root access should be run with `sudo`
- Replace `eth0` with your actual network interface name
- The VPN interface name might be different on your system (e.g., `utun0` on macOS)
- Make sure to adjust paths according to your system configuration
- Keep the debug mode enabled during testing for better visibility into the system's operation 