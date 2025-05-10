# WarpNet

A decentralized private network built using libp2p that creates an immutable trusted blockchain addressable p2p network.

## Features

- Create secure VPN connections between p2p peers
- Automatically assign IPs to nodes
- Embedded DNS server for internal/external IP resolution
- Create trusted zones to prevent network access if token is leaked
- Act as a reverse proxy for TCP services
- Send files securely over p2p
- Blockchain-based immutable ledger

## Installation

### Option 1: Using the installer script
```bash
curl -sfL https://raw.githubusercontent.com/thealonemusk/WarpNet/main/install.sh | sh -
```

### Option 2: Manual installation
1. Download the latest release from the [releases page](https://github.com/thealonemusk/WarpNet/releases)
2. Extract the binary to your desired location
3. Make it executable: `chmod +x WarpNet`

## Quick Start

1. Generate a configuration file:
```bash
# Generate a new config file
WarpNet -g > config.yaml

# OR generate a portable token
WarpNetTOKEN=$(WarpNet -g -b)
```

2. Start the VPN:
```bash
# On Node A
WarpNetCONFIG=config.yaml WarpNet --address 10.1.0.11/24

# On Node B
WarpNetCONFIG=config.yaml WarpNet --address 10.1.0.12/24
```

## Development Setup

1. Clone the repository:
```bash
git clone https://github.com/thealonemusk/WarpNet.git
cd WarpNet
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the development server:
```bash
# Generate config
go run main.go -g > config.yaml

# Start the API server
WarpNetCONFIG=config.yaml go run main.go api
```

## Performance Tuning

For better network performance, run:
```bash
sudo sysctl -w net.core.rmem_max=2500000
```
