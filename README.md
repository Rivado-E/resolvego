# DNS Server Project

A custom DNS server capable of handling requests from tools like `dig`. This server implements core DNS functionalities and can be extended with advanced features like query forwarding, caching, and DNSSEC.

---

## Features

### Basic Features
- Respond to DNS queries over UDP and TCP.
- Parse DNS queries and generate valid DNS responses.
- Support for common DNS record types:
  - A (IPv4)
  - AAAA (IPv6)
  - MX (Mail Exchange)
  - CNAME (Canonical Name)
  - NS (Name Server)

### Advanced Features
- Query forwarding to upstream DNS servers.
- Recursive and non-recursive query handling.
- Response caching for improved performance.
- Zone file parsing for custom domain configurations.
- Support for EDNS to handle large packets.
- Optional DNSSEC implementation for secure responses.

### Security and Monitoring
- Rate limiting to prevent abuse.
- Logging for incoming queries and responses.
- Metrics export for monitoring (e.g., Prometheus integration).

