# -Document-Crud-Operations

A minimal Gin + GORM + PostgreSQL app, containerized with Docker Compose.

## 🛠️ Prerequisites

* **Docker** & **Docker Compose** installed on your machine.
* (Optional) A local install of Go 1.24+ if you want to run outside Docker.

## 🚀 Getting Started

1. **Clone the repo**

   ```bash
   git clone git@github.com:MaTb3aa/Project-Base-Training.git
   cd Project-Base-Training
   ```

2. **Build and run the Docker containers**

   ```bash
   docker-compose up --build
   ```

3. **Verify the server is running**

   ```bash
   curl http://localhost:8080/ping
   ```

## 📚 API Usage

All endpoints assume the server is accessible at `http://localhost:8080` and use JSON.

### Create a Document

**Endpoint:** `POST /documents/create`

```bash
curl -X POST http://localhost:8080/documents/create \
  -H "Content-Type: application/json" \
  -d '{"title":"My Title","author":"Author Name","content":"Some content"}'
```

### List All Documents

**Endpoint:** `GET /documents`

```bash
curl http://localhost:8080/documents
```

### Get a Document by ID

**Endpoint:** `GET /documents/:id`

```bash
curl http://localhost:8080/documents/1
```

### Update a Document by ID

**Endpoint:** `PUT /documents/update/:id`

```bash
curl -X PUT http://localhost:8080/documents/update/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Updated Title","author":"New Author","content":"Updated content"}'
```

### Delete a Document by ID

**Endpoint:** `DELETE /documents/delete/:id`

```bash
curl -X DELETE http://localhost:8080/documents/delete/1
```

