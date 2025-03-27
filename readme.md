### **Over-Engineered Calculator 🫢**

A **simple** yet *over-engineered* calculator API built with **Go**, **Chi router**, and **Docker**, deployed using **Fly.io**.  
This project was created as a coding challenge and explores API development, dependency injection, vertical slicing, and containerized deployment.
In the bottom of this readme you can find an explanation of what I would change had I spent more time on this project.

---

## **🚀 Features**
- **Calculate expressions** via an API (`/api/v1/calculate`)
- **Retrieve history** of calculations (`/api/v1/get_history`)
- **OpenAPI documentation** generated with Swaggo
- **Dockerized & deployed on Fly.io**

---

## **🛠️ Setup & Usage**

### **1️⃣ Build & Run Locally**
Make sure you have **Go** installed. Then, you can run:

```sh
go mod tidy  # Install dependencies
go run main.go
```

The API will be available at `http://localhost:8080`.

---

### **2️⃣ Running with Docker**

#### **Build the Docker image**
```sh
docker build -t over-engineered-calculator:latest .
```

#### **Run the container**
```sh
docker run -p 8080:8080 over-engineered-calculator:latest
```

---

## **📚 API Documentation (Swagger)**
After running the server, access Swagger UI at:  
👉 `http://localhost:8080/swagger/index.html`

---

### **4️⃣ Access API on Fly.io**
The deployed API will be available at:  
👉 `https://over-engineered-calculator.fly.dev`

---

## **⏳ What would I change with more time ?**