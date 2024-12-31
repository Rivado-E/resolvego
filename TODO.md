---
id: TODO
aliases: []
tags: []
---

# DNS Server TODOs

## 1. Basic UDP Server
- [ ] Create a UDP server listening on port 53.
- [ ] Handle incoming UDP packets.
- [ ] Parse raw data from incoming requests.
- [ ] Send back a basic hardcoded response.

---

## 2. DNS Query Parsing
- [ ] Parse the DNS header section.
- [ ] Extract the question section (domain name, query type, and class).
- [ ] Handle queries with multiple questions (optional).

---

## 3. Simple Response Generation
- [ ] Construct a DNS response packet.
- [ ] Include a valid header with response flags.
- [ ] Echo the question section from the query.
- [ ] Create an answer section with hardcoded records.

---

## 4. Handling Different Record Types
- [ ] Add support for `A` (IPv4) records.
- [ ] Add support for `AAAA` (IPv6) records.
- [ ] Add support for `MX` (Mail Exchange) records.
- [ ] Add support for `CNAME` (Canonical Name) records.
- [ ] Add support for `NS` (Name Server) records.

---

## 5. Query Forwarding
- [ ] Forward unresolved queries to an upstream DNS server (e.g., 8.8.8.8).
- [ ] Relay the response back to the client.
- [ ] Handle recursive queries from clients.

---

## 6. Recursive and Non-Recursive Queries
- [ ] Parse the "Recursion Desired" (RD) flag in the query.
- [ ] Implement recursive resolution by querying other DNS servers.

---

## 7. Zone File Parsing
- [ ] Parse standard zone file format.
- [ ] Serve records from the zone file.
- [ ] Add support for SOA, A, AAAA, MX, CNAME, and NS records.

---

## 8. Response Caching
- [ ] Implement caching for responses.
- [ ] Evict cached records based on TTL.
- [ ] Serve cached responses for repeated queries.

---

## 9. Advanced Query Types
- [ ] Add support for `TXT` (Text) records.
- [ ] Add support for `PTR` (Reverse Lookup) records.
- [ ] Add support for `SRV` (Service Locator) records.
- [ ] Add support for `ANY` queries.

---

## 10. TCP Query Support
- [ ] Handle DNS queries over TCP.
- [ ] Implement support for large packets exceeding UDP size limits.

---

## 11. EDNS (Extension Mechanisms for DNS)
- [ ] Parse EDNS options in incoming queries.
- [ ] Increase packet size limit for UDP queries (up to 4096 bytes).
- [ ] Handle additional EDNS features if necessary.

---

## 12. DNSSEC (Optional)
- [ ] Respond with DNSSEC signatures (RRSIG).
- [ ] Validate DNSSEC-enabled queries.

---

## 13. Logging and Metrics
- [ ] Log all incoming queries and responses.
- [ ] Record query types, source IPs, and response times.
- [ ] Export server metrics (e.g., Prometheus) for monitoring.

---

## 14. Rate Limiting and Security
- [ ] Implement rate limiting per client IP.
- [ ] Add support for blacklisting/whitelisting clients.
- [ ] Gracefully handle malformed or malicious queries.

---

## 15. Full DNS Resolver
- [ ] Implement iterative resolution starting from root servers.
- [ ] Query authoritative servers for specific zones.
- [ ] Support the full DNS hierarchy (root, TLDs, authoritative servers).

