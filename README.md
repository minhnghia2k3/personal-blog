<img height="250" src="./ui/assets/golang.png" width="250"/>

# Personal Blog

Welcome to **minhnghia2k3's Blog**,
a personal blogging platform built with Go and server-side rendered using Go's `html/template`.
This application provides a platform for publishing blog posts, managing categories,
and interacting with users through authentication and commenting features.
It also includes robust security, monitoring, and deployment functionalities.

## Features

- **Blog Management**: Create, update, and delete blog posts with support for categories and rich content.
- **User Authentication**: Sign up, log in, log out, and reset password functionality with OAuth2 (Google login).
- **Comments & Likes**: Users can engage with blog posts through comments and likes.
- **Blog History**: Track changes to blog posts with the ability to restore previous versions.
- **Rate Limiting**: Control the number of requests a user can make to protect against abuse.
- **Logging & Monitoring**: Detailed logging for error tracking and performance monitoring with Prometheus metrics.
- **CORS**: Secure cross-origin requests.
- **Email Notifications**: Send email notifications to subscribers when new blogs are published.
- **Graceful Shutdown**: Ensure smooth termination of services.
- **Containerized Deployment**: Easily deploy with Docker.

## Project Structure
```
    ├── bin
    ├── cmd
        └──web
    ├── internal
        └──config           # Configuration files
        └──handlers         # HTTP handlers
        └──middlewares      # Middleware logic (e.g., authentication, logging)
        └──models           # Business logic and data models
        └──routes           # API and web routes
        └──templates        # HTML templates for SSR
        └──model            # Data representation models
    ├── pkg                  # Reusable packages and libraries
    ├── ui
       └──html             # Static HTML pages
       └──static           # CSS, JS, Images
    ├── utils                # Utility functions and helpers
```

## API Endpoints

### System
| Method     | Endpoint                      | Description                                             |
|------------|-------------------------------|---------------------------------------------------------|
| **GET**    | `/healthcheck`                | Check application health status.                        |

### Info
| Method     | Endpoint                      | Description                                             |
|------------|-------------------------------|---------------------------------------------------------|
| **GET**    | `/`                           | View the homepage with top blogs.                       |
| **GET**    | `/me`                         | View the information of the author.                     |

### Users
| Method     | Endpoint                      | Description                                             |
|------------|-------------------------------|---------------------------------------------------------|
| **GET**    | `/user/signup`                | View sign-up form.                                      |
| **POST**   | `/user/signup`                | Register a new user account.                            |
| **GET**    | `/user/login`                 | View login form.                                        |
| **POST**   | `/user/login`                 | Log in to the user account.                             |
| **POST**   | `/user/logout`                | Log out the current user.                               |
| **GET**    | `/user/forgot-password`       | View forgot password form.                              |
| **POST**   | `/user/forgot-password`       | Reset password.                                         |
| **GET**    | `/user/google_login`          | Redirect to Google authentication URL.                  |
| **POST**   | `/user/google_login`          | Log in using Google OAuth.                              |
| **POST**   | `/user/google_callback`       | Handle Google OAuth callback.                           |

### Blogs
| Method     | Endpoint                      | Description                                             |
|------------|-------------------------------|---------------------------------------------------------|
| **GET**    | `/blogs`                      | Fetch all blogs (supports pagination, filtering).       |
| **GET**    | `/blogs/{title}/{id}`         | View a specific blog by title and ID.                   |
| **GET**    | `/blogs/create`               | View create blog form.                                  |
| **POST**   | `/blogs/create`               | Submit a new blog.                                      |
| **GET**    | `/blogs/update/{id}`          | View update blog form.                                  |
| **PATCH**  | `/blogs/update/{id}`          | Partially update a blog (save draft to `blog_history`). |
| **DELETE** | `/blogs/delete/{id}`          | Delete a blog and related images (CASCADE).             |

### Comments
| Method     | Endpoint                      | Description                                             |
|------------|-------------------------------|---------------------------------------------------------|
| **GET**    | `/comments/{id}/blogs`        | Fetch all comments on a blog.                           |
| **POST**   | `/comments/create/{id}/blogs` | Post a comment on a blog.                               |
| **PATCH**  | `/comments/update/{id}`       | Update a comment.                                       |
| **DELETE** | `/comments/delete/{id}`       | Delete a comment.                                       |

### Likes
| Method     | Endpoint                      | Description                                             |
|------------|-------------------------------|---------------------------------------------------------|
| **GET**    | `/likes/{id}/blogs`           | Fetch all likes on a blog.                              |
| **POST**   | `/likes/add/{id}/blogs`       | Like a blog post.                                       |
| **DELETE** | `/likes/delete/{id}/blogs`    | Unlike a blog post.                                     |

### Categories
| Method     | Endpoint                      | Description                                             |
|------------|-------------------------------|---------------------------------------------------------|
| **GET**    | `/categories`                 | Fetch all categories.                                   |
| **GET**    | `/categories/{id}`            | Fetch a specific category by ID.                        |
| **POST**   | `/categories`                 | Create a new category.                                  |
| **PATCH**  | `/categories/{id}`            | Update a category.                                      |
| **DELETE** | `/categories/{id}`            | Delete a category.                                      |

### Images
| Method     | Endpoint                      | Description                                             |
|------------|-------------------------------|---------------------------------------------------------|
| **POST**   | `/images/create`              | Upload an image (local or cloud).                       |
| **GET**    | `/images/{name}`              | View an image.                                          |
| **DELETE** | `/images/delete/{id}`         | Delete an image.                                        |

## Setup Instructions

1. **Install dependencies:**

```bash
go mod tidy
```

2. **Setup the environment:**

- Configure your .env file for database connection, Google OAuth credentials, etc.

3. **Run database migrations:**

```bash
make migrate-up
```

4. **Run the application:**

```bash
go run cmd/web/main.go
```

or

```bash
docker-compose up --build
```

## **Deployment**

You can easily deploy the application using Docker. Ensure your environment variables are correctly configured for the
production environment.

```bash
docker-compose -f docker-compose.prod.yml up --build
```

## **Monitoring and Metrics**
The application integrates Prometheus for monitoring key metrics. You can view the metrics at the `/metrics` endpoint.

## Todos
- [x] Design web routes
- [x] Project setup 
- [x] Create file structure
- [x] Setup Git version control
- [x] Setup Go server on local & first healthcheck handler
- [x] Setup Building, Versioning
- [x] Logging
- [x] Setup Postgres database
- [ ] SQL migrations
- [ ] User model
- [ ] Blog model...
- [ ] Email service
- [ ] Store image on cloud service
- [ ] User activation
- [ ] Role-based authorization
- [ ] Rate limiting
- [ ] Graceful shutdown
- [ ] CORS
- [ ] Metrics

## License
This project is licensed under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0) - see the LICENSE file for details.