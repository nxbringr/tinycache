# 🐳 TinyCache Docker Guide

Run your local TinyCache instance in a container with just a few commands.

---

## 🔧 Step 1: Build the Docker Image

From the root of your project (where the `Dockerfile` is located):

```bash
docker build -t tinycache .
```

---

## ▶️ Step 2: Run the Container

```bash
docker run -p 8080:8080 tinycache
```

Your cache server will be accessible at [http://localhost:8080](http://localhost:8080).

---

## 🔁 Step 3: Sample API Usage

### Set a Key

```bash
curl -X POST http://localhost:8080/cache/hello -d "world"
```

### Get a Key

```bash
curl http://localhost:8080/cache/hello
```

### Delete a Key

```bash
curl -X DELETE http://localhost:8080/cache/hello
```

### Get Full Cache

```bash
curl http://localhost:8080/cache
```

---

Happy caching! 🚀