# Load Balancer

A load balancer is a crucial component in system design that distributes incoming network traffic across multiple servers. Its main purpose is to ensure that no single server is overburdened with too many requests, which helps improve the performance, reliability, and availability of applications. This is a simple mock load balancer implemented in **Go** with backend servers running on **Flask**. This setup distributes incoming requests to multiple backend servers in a round-robin fashion, ensuring load distribution and fault tolerance.


## Features
- **Round-robin request distribution**
- **Health checks** to skip unhealthy servers
- **Dockerized setup** for easy deployment


## Prerequisites
- **Docker**
- **Docker Compose**

## Setup & Installation
1. Clone the repository:
   ```sh
   git clone <https://github.com/avd1729/Load-Balancer.git>
   cd load-balancer
   ```

2. Build and start the services:
   ```sh
   docker-compose up --build
   ```

3. The load balancer will be available at:
   ```
   http://localhost:8080
   ```

## Testing
- Check the status of backend servers:
  ```sh
  curl http://localhost:5001/status
  curl http://localhost:5002/status
  curl http://localhost:5003/status
  ```
- Send requests to the load balancer:
  ```sh
  curl http://localhost:8080/
  ```
  This request will be forwarded to one of the backend servers.

## Load Balancer (Go)
### **Key Features:**
- Uses **round-robin** to distribute requests across backend servers
- Skips servers that fail health checks (`/status` endpoint)
- Forwards requests while preserving headers and query parameters


## Backend Servers (Flask)
Each backend server is a simple **Flask** application that runs on a different port.


## Stop & Cleanup
To stop the services and remove containers, run:
```sh
docker-compose down --volumes
```


## License
This project is open-source and available under the **Apache License**.
