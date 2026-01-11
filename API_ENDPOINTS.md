# Hexathon Backend API Endpoints

This document lists all the available API endpoints in the Hexathon backend.

## Base URL
All API endpoints are prefixed with `/api/v1`

## General Endpoints

### Root
- **GET** `/` - Welcome message
  - Response: "Welcome to Hexathon API!🎉"

### Health Check
- **GET** `/ping` - Health check endpoint
  - Response: `{"detail": "pong"}`

## Authentication Endpoints

Base path: `/api/v1/auth`

- **POST** `/api/v1/auth/register` - Register a new team
  - Body: `{"name": string, "password": string, "member_ids": string[]}`
  - Returns: Team details with JWT token

- **POST** `/api/v1/auth/login` - Login a team
  - Body: `{"name": string, "password": string}`
  - Returns: Team details with JWT token

## Categories Endpoints

Base path: `/api/v1/categories`

**Authentication required for all endpoints**

### Public Endpoints (Team authentication required)
- **GET** `/api/v1/categories/` - Get all categories
- **GET** `/api/v1/categories/:id` - Get a specific category by ID

### Admin Endpoints (Admin authentication required)
- **POST** `/api/v1/categories/` - Create a new category
  - Body: `{"name": string, "photo_url": string, "description": string, "max_items": int}`
- **PUT** `/api/v1/categories/:id` - Update a category
  - Body: `{"name": string, "photo_url": string, "description": string, "max_items": int}`
- **DELETE** `/api/v1/categories/:id` - Delete a category

## Items Endpoints

Base path: `/api/v1/items`

**Authentication required for all endpoints**

### Public Endpoints (Team authentication required)
- **GET** `/api/v1/items/` - Get all items
  - Query params: `category` (optional) - Filter items by category ID
- **GET** `/api/v1/items/:id` - Get a specific item by ID

### Admin Endpoints (Admin authentication required)
- **POST** `/api/v1/items/` - Create a new item
  - Body: `{"name": string, "photo_url": string, "description": string, "price": int, "category_id": string}`
- **PUT** `/api/v1/items/:id` - Update an item
  - Body: `{"name": string, "photo_url": string, "description": string, "price": int, "category_id": string}`
- **DELETE** `/api/v1/items/:id` - Delete an item

## Cart Endpoints

Base path: `/api/v1/carts`

**Authentication required for all endpoints**

- **GET** `/api/v1/carts/` - Get current team's cart
- **POST** `/api/v1/carts/checkout` - Checkout the cart
- **POST** `/api/v1/carts/:itemId` - Add an item to the cart
- **DELETE** `/api/v1/carts/:itemId` - Remove an item from the cart

## Teams Endpoints

Base path: `/api/v1/teams`

**Authentication required for all endpoints**

### Team Endpoints (Team authentication required)
- **GET** `/api/v1/teams/me` - Get current team's information

### Admin Endpoints (Admin authentication required)
- **GET** `/api/v1/teams/` - Get all teams
  - Query params: `checked_out` (optional) - Filter by checkout status (true/false)
- **GET** `/api/v1/teams/:name` - Get a specific team by name
- **POST** `/api/v1/teams/` - Create a new team
  - Body: `{"name": string, "member_ids": string[]}`
- **POST** `/api/v1/teams/admin` - Create an admin team
  - Body: `{"name": string, "password": string}`
- **POST** `/api/v1/teams/:name/regeneratePassword` - Regenerate password for a team
- **PUT** `/api/v1/teams/:name` - Update a team
  - Body: `{"name": string, "password": string, "member_ids": string[], "ps_generations": int, "ps_confirmed": bool}`
- **DELETE** `/api/v1/teams/:name` - Delete a team
- **POST** `/api/v1/teams/checkout` - Checkout all teams who haven't checked out
- **POST** `/api/v1/teams/confirmProblemStatement` - Confirm problem statements for all teams

## Problem Statements Endpoints

Base path: `/api/v1/problemStatements`

**Authentication required for all endpoints**

### Team Endpoints (Team authentication required)
- **GET** `/api/v1/problemStatements/team` - Get problem statement for current team
  - Query params: `type` (optional) - Set to "one_liner" for one-liner format
- **POST** `/api/v1/problemStatements/team` - Generate a problem statement for current team
- **POST** `/api/v1/problemStatements/confirm` - Confirm the selected problem statement

### Admin Endpoints (Admin authentication required)
- **GET** `/api/v1/problemStatements/` - Get all problem statements
- **GET** `/api/v1/problemStatements/:id` - Get a specific problem statement by ID
- **POST** `/api/v1/problemStatements/` - Create a new problem statement
  - Body: `{"name": string, "one_liner": string, "description": string}`
- **PUT** `/api/v1/problemStatements/:id` - Update a problem statement
  - Body: `{"name": string, "one_liner": string, "description": string}`
- **DELETE** `/api/v1/problemStatements/:id` - Delete a problem statement

## Submissions Endpoints

Base path: `/api/v1/submissions`

**Authentication required for all endpoints**

### Team Endpoints (Team authentication required)
- **GET** `/api/v1/submissions/me` - Get current team's submission
- **POST** `/api/v1/submissions/submit` - Submit or update a submission
  - Body: `{"figmaURL": string, "docURL": string}`

### Admin Endpoints (Admin authentication required)
- **GET** `/api/v1/submissions/` - Get all submissions
- **DELETE** `/api/v1/submissions/:id` - Delete a submission

## Participants Endpoints

Base path: `/api/v1/participants`

**Admin authentication required for all endpoints**

- **GET** `/api/v1/participants/` - Get all participants (not checked in)
  - Query params: `q` (optional) - Search query
- **GET** `/api/v1/participants/checkedin` - Get all checked-in participants
  - Query params: `q` (optional) - Search query
- **GET** `/api/v1/participants/all` - Get all participants (both checked in and not)
- **GET** `/api/v1/participants/:id` - Get a specific participant by ID
- **POST** `/api/v1/participants/` - Create a new participant
  - Body: `{"name": string, "email": string, "reg_no": string, "checked_in": bool}`
- **PUT** `/api/v1/participants/:id` - Update a participant
  - Body: `{"name": string, "email": string, "reg_no": string, "checked_in": bool}`
- **DELETE** `/api/v1/participants/:id` - Delete a participant

## Authentication & Authorization

### JWT Authentication
Most endpoints require JWT authentication. Include the JWT token in the Authorization header:
```
Authorization: Bearer <token>
```

### User Roles
- **Team**: Regular team accounts with limited access
- **Admin**: Administrator accounts with full access to all endpoints

### Middleware
- `JWTAuthMiddleware`: Validates JWT token and authenticates the user
- `IsAdminMiddleware`: Ensures the authenticated user has admin role

## Response Formats

### Success Responses
Most successful requests return JSON with relevant data and appropriate HTTP status codes:
- `200 OK`: Successful GET, PUT, POST requests
- `201 Created`: Successful resource creation
- `202 Accepted`: Successful update
- `204 No Content`: Successful deletion

### Error Responses
Error responses include a `detail` field with error message:
```json
{
  "detail": "Error message"
}
```

Common error status codes:
- `400 Bad Request`: Invalid request body or parameters
- `401 Unauthorized`: Missing or invalid authentication
- `403 Forbidden`: Insufficient permissions
- `404 Not Found`: Resource not found
- `409 Conflict`: Resource conflict
- `500 Internal Server Error`: Server error
