# URL Shortener

A simple URL shortener service built using Golang and the Gin web framework. This application provides the ability to shorten URLs, redirect to the original URLs, and track the most shortened domains.

## Features

- **Shorten URL**: Accepts a long URL and returns a shortened version of it.
- **Redirection**: Allows redirection to the original URL by accessing the shortened URL.
- **Metrics**: Returns the top 3 most shortened domains.
  
## Endpoints

### 1. **POST /shorten**
- **Description**: Shortens a provided URL.
- **Request Body**:
  ```json
  {
    "url": "https://example.com"
  }
Response:
json
Copy code
{
  "shortened_url": "http://localhost:8080/abc123"
}
2. GET /:shortURL
Description: Redirects to the original URL using the shortened URL.
Example: Access http://localhost:8080/abc123 to be redirected to the original URL.
3. GET /metrics
Description: Returns the top 3 domains that have been shortened the most times.
Response:
json
Copy code
{
  "Udemy": 6,
  "YouTube": 4,
  "Wikipedia": 2
}
Getting Started
Prerequisites
Go (version 1.18 or higher)
Docker (Optional, for running the app in a container)
Installation
Clone the repository:

bash
Copy code
git clone https://github.com/your-username/url-shortener.git
cd url-shortener
Install dependencies:

bash
Copy code
go mod tidy
Run the application:

bash
Copy code
go run main.go
The server will start on http://localhost:8080.

Docker (Optional)
To run the application inside a Docker container, follow these steps:

Build the Docker image:

bash
Copy code
docker build -t url-shortener .
Run the Docker container:

bash
Copy code
docker run -p 8080:8080 url-shortener
The application will be accessible at http://localhost:8080 inside the container.
