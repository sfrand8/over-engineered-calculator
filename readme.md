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
- **"Bearer token" authentication** for API access - just add a "Bearer <random UUID>" to the Authorizaiton header of the request and it will be accepted.  

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

## **📚 API Documentation (Swagger) & Postman**
OpenAPI documentation is generated using **Swaggo**.

```
swag init
```

After running the server, access Swagger UI at:  
👉 `http://localhost:8080/swagger/index.html`

A postman collection is also created and found in the `docs` folder along with the swagger.json/yaml files.

Generate the documentation with swaggo:

For testing generate a go client from the swagger spec using [openapi-generator](https://openapi-generator.tech/docs/generators/go/)

Run the following command in the root directory of the project:
```
 openapi-generator generate -i ./docs/swagger.json -g go -o ./pkg/http-client --skip-validate-spec
```
I didn't find a way that worked to make the go.mod file to be generated correctly, so I added had to manually change it after generation.
Howerever there should be an option to set it, it just didn't work for me for some reason.


### **4️⃣ Access API on Fly.io**
The deployed API will be available at:  
👉 `https://over-engineered-calculator.fly.dev`

---

## **⏳ What would I change with more time ?**
- Create an actual calculator instead of just importing a library that does it for me.
- Add open telemetry to monitor the API with a middleware to track the requests and responses.
- Add a database to store the history of calculations and make it persistent so it isn't removed after each deployment.
- Could also add a user authentication system to restrict access to the API as right now it's just anyone with a uuid.
