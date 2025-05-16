# Th3-Sh0p API

A Go-based REST API service for an AI-powered image generation and management platform. This API provides endpoints for image generation, user credit management, and payment processing.

## Features

- **Image Generation**: Generate AI images from text prompts
- **Image Pack Purchases**: Purchase predefined image packs
- **User Credit System**: Track and manage user credits for image generation
- **Google Authentication**: User authentication via Google OAuth
- **Stripe Integration**: Secure payment processing for image pack purchases

## API Endpoints

### Image Generation
- `POST /v1/image`: Generate an image from a text prompt
- `GET /v1/images`: List generated images with pagination
- `GET /v1/images/pages`: Get total number of image pages

### User Management
- `GET /v1/google-profile`: Get user profile information via Google OAuth
- `GET /v1/user-credit`: Check user's available image credits

### Payment Processing
- `POST /v1/image-pack`: Purchase image packs (pack_1, pack_2, pack_3)
- `GET /v1/pub-key`: Get Stripe publishable key for client-side integration

## Technology Stack

- **Language**: Go 1.23
- **API Documentation**: Swagger/OpenAPI 2.0
- **Database**: MySQL
- **Authentication**: Google OAuth
- **Payment Processing**: Stripe
- **Containerization**: Docker

## Dependencies

Major dependencies include:
- github.com/go-openapi/* - OpenAPI/Swagger tooling
- github.com/stripe/stripe-go - Stripe payment integration
- google.golang.org/api - Google OAuth integration
- github.com/go-sql-driver/mysql - MySQL database driver

## Getting Started

1. Clone the repository
2. Install Go 1.23 or later
3. Install dependencies:
   ```bash
   go mod download
   ```
4. Configure environment variables (see Configuration section)
5. Run the server:
   ```bash
   go run cmd/th3-sh0p-api-server/main.go
   ```

## Configuration

The following environment variables need to be configured:
- Database connection details
- Google OAuth credentials
- Stripe API keys
- Server host and port

## Docker Support

A Dockerfile is included for containerized deployment. Build and run using:
```bash
docker build -t th3-sh0p-api .
docker run -p 8080:8080 th3-sh0p-api
```

## API Documentation

Full API documentation is available in the `swagger.yml` file. You can view the interactive documentation by running the server and visiting the Swagger UI endpoint.

## License

[Add appropriate license information]

## Contributing

[Add contribution guidelines if applicable] 