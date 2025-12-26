# NS Lookup Tool

A lightweight command-line utility for performing DNS lookups. This tool queries Google's public DNS (8.8.8.8) to retrieve various DNS records for a given domain.

## Features

- **A Records**: Retrieve the IP address (Default).
- **CNAME Records**: Get the canonical domain name.
- **NS Records**: Find the authoritative name servers.
- **TXT Records**: View associated text records.

## Usage

Run the tool by providing a domain name. You can specify the record type using the `--type` flag.

### Basic Lookup (A Record)
```bash
go run main.go example.com
```

### Specify Record Type
```bash
# Check CNAME
go run main.go example.com --type CNAME

# Check Name Servers
go run main.go example.com --type NS

# Check TXT Records
go run main.go example.com --type TXT
```
