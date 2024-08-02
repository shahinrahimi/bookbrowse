# BOOKBROWSE

## Overview
BookBrowse is a comprehensive full-stack application designed to showcase my skills in both frontend and backend development. This project serves as a prototype and portfolio piece, demonstrating my expertise in building scalable and maintainable web applications using modern technologies.

## Features
- Frontend: Utilizes the Templ templating engine for rendering views, combined with Tailwind CSS for responsive design and styling.
- Backend: Built with Go (Golang), leveraging Gorilla Mux for routing and SQLite for database management.
- API: Includes a robust RESTful API for managing books, authors, and genres, with full CRUD functionality.
- Middleware: Implements custom middleware for logging, request validation, and rate limiting.
- Development and Production Modes: Static assets are served from the public directory in development and embedded in the binary for production.

## Key Files
- main.go: Entry point of the application, setting up routes, middleware, and server configurations.
- handlers/: Contains handlers for various routes and functionalities.
- middlewares/: Custom middleware implementations.
- models/: Database models for books, authors, and genres.
- stores/: Data access layer for interacting with the database.
- views/: Templ files for rendering the frontend.
- public/: Static assets for the application.

## Running the Poject
### Quick Run
```sh
make run
```
### Development
```sh
make css

```
```sh
make templ
```

```sh
air
```

### Production
```sh
make build
```

## Contact
For any questions or inquiries, feel free to reach out via my GitHub profile.

