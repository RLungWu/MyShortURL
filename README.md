# MyShortURL

## Introcution
MyShortURL is a simple, efficient Short URL generator built using Gin framework and Redis. This project allows you to set up a robust URL shortening service easily using Docker, suitable for both development and production environments.
## Feature
- Quick Setup: Get up and running with just a few commands.
- High Performance: Leverages the fast Gin framework and Redis caching for optimal performance.
- Docker Integration: Fully containerized for easy deployment and scaling.
## Requirements
- Docker and Docker Compose
- Redis server (included in the docker-compose file)

## Quickstart/Demo
```
docker-compose up
```

## API
### Creating a Short URL

**POST** `/api/v1/shorten`

You can create a new short URL by sending a POST request to this endpoint. Include the original URL in the body of your request.

#### Request:

- **Endpoint**: `localhost:3000/api/v1/shorten`
- **Method**: POST
- **Body**:
  ```json
  {
    "url": "https://example.com"
  }

## Contributing
Contributions are welcome! Please refer to our contributing guidelines for more information on how to participate in the project.