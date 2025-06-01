# TinyCache 🚀 – A Distributed Cache Scaling Benchmark in Go

**TinyCache** is a hands-on journey through distributed system design using Go — starting from a basic in-memory cache to a geo-distributed, observable, high-throughput cache layer.

It’s inspired by real-world systems like Redis, CDN edge caches, and enterprise key-value stores. This project is structured in progressive levels to simulate production bottlenecks and architectural trade-offs as you scale from 10 RPS to 1M RPS.

---

## 💡 Why TinyCache?

TinyCache was born out of a personal need — I’ve often struggled with scaling databases and backend systems under real-world pressure. Questions like *"Why does this API slow down under load?"*, *"How do I prevent cascading failures?"*, and *"What’s the cleanest way to scale reads without sacrificing consistency?"* have followed me through many projects.

Instead of reading endless theory, this project is my attempt to learn by building, breaking, and iterating.

Each level of TinyCache simulates realistic backend challenges, while giving me space to master Go, concurrency, metrics, and infra design.

---

## 📊 Level-Based Scaling Plan (10 RPS → 1M RPS)

### 🟢 Level 1: In-Memory Cache (10 RPS)

* Thread-safe in-memory cache using `map + RWMutex`
* Supports all Go data types using `interface{}`
* TTL-based expiration for keys
* Basic REST API for `GET`, `SET`, `DELETE`
* Stores only in RAM

📌 **What this answers**: Can I build a safe, basic cache that expires keys and handles concurrency?

---

### 🔵 Level 2: Observability & Metrics (50 RPS)

* Prometheus integration with `/metrics` endpoint
* Gin HTTP metrics exposed via `promhttp`
* Local load testing (e.g. `hey`, `vegeta`, `ab`)

📌 **What this answers**: How does my cache behave under moderate traffic? What’s the hit/miss rate, latency, and throughput?

---

### 🟡 Level 3: Persistence Layer (100 RPS)

* Optional PostgreSQL backend (write-through)
* Compare in-memory latency vs DB fallback
* Connection pooling, error handling

📌 **What this answers**: Can I persist cache data safely? How does fallback to DB impact performance?

---

### 🟠 Level 4: Eviction & Compression (500 RPS)

* Implement LRU eviction policy
* Add size limits to key/value entries
* Gzip compression for large values

📌 **What this answers**: How do I control memory footprint as the cache grows?

---

### 🟣 Level 5: Sharding & Routing (1K RPS)

* Split cache into multiple shards
* Use consistent hashing to assign keys
* Add peer-to-peer routing logic

📌 **What this answers**: How do I scale writes and reads across multiple cache nodes?

---

### 🔴 Level 6: Replication & Failover (5K RPS)

* Read replicas per shard
* Leader election via Raft or custom heartbeat
* Retry on leader failure

📌 **What this answers**: Can my cache tolerate node failures while staying consistent?

---

### [TBD] Level 7: Fault Tolerance & Chaos (50K RPS)
---

### [TBD] Level 8: Global Distribution (1M RPS)

---

## 🛠 Tooling
[TBD]
---
