# WarpNet VPN Connection Guide

## Building the Project

1. Clone the repository:
```bash
git clone https://github.com/thealonemusk/WarpNet.git
cd WarpNet
```

2. Build the project:
```bash
go mod tidy
go build -o WarpNet
```

## Generating Configuration

You have two options to generate the network configuration:

### Option 1: Generate Config File
```bash
./WarpNet -g > config.yaml
```

## Prerequisites
- WarpNet binary built and available
- config.yaml file or WarpNetTOKEN generated
- Root/sudo access

## Performance Tuning
Before starting the VPN, set the recommended system parameter for better performance:
```bash
sudo sysctl -w net.core.rmem_max=2500000
```

## Starting the VPN

### On First Machine (Node A)
1. Navigate to the WarpNet directory
2. Run the following command (using config file):
```bash
sudo WarpNetCONFIG=config.yaml ./WarpNet --address 10.1.0.11/24
```
OR (using token):
```bash
sudo WarpNetTOKEN=<your-token> ./WarpNet --address 10.1.0.11/24
```

### On Second Machine (Node B)
1. Copy the same `config.yaml` file or use the same token
2. Run the following command (with a different IP):
```bash
# Using config file
sudo WarpNetCONFIG=config.yaml ./WarpNet --address 10.1.0.12/24

# OR using token
sudo WarpNetTOKEN=<your-token> ./WarpNet --address 10.1.0.12/24
```

### On Additional Machines
For each additional machine:
1. Copy the same `config.yaml` file or use the same token
2. Run with a unique IP address:
```bash
# Using config file
sudo WarpNetCONFIG=config.yaml ./WarpNet --address 10.1.0.13/24  # increment the last number

# OR using token
sudo WarpNetTOKEN=<your-token> ./WarpNet --address 10.1.0.13/24  # increment the last number
```

## Verifying the Connection

1. Check if the VPN interface is created:
```bash
ip addr show edgevpn0
```

2. Test connectivity between nodes:
```bash
# From Node A
ping 10.1.0.12  # to reach Node B
ping 10.1.0.13  # to reach Node C

# From Node B
ping 10.1.0.11  # to reach Node A
ping 10.1.0.13  # to reach Node C
```

## Troubleshooting

1. If you see "ioctl: operation not permitted":
   - Make sure you're running the command with sudo
   - Check if you have the necessary permissions

2. If connection takes time:
   - The initial connection might take up to 5 minutes
   - This is normal as the DHT (Distributed Hash Table) needs to bootstrap

3. If you see buffer size warnings:
   - Make sure you've run the sysctl command mentioned in Prerequisites
   - The warning won't prevent the VPN from working, but performance might be affected

## Notes
- Each machine must have a unique IP address in the 10.1.0.0/24 range
- The config.yaml file or token must be identical on all machines
- Keep the config.yaml file and token secure as they contain network credentials 