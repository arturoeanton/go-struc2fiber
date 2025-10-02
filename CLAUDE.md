# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Development
```bash
# Run the application
go run main.go

# Build the application
go build -o app main.go

# Install dependencies
go mod download
go mod tidy

# Format code
go fmt ./...

# Run linter (install golangci-lint if not available)
golangci-lint run

# Run tests (when tests are added)
go test ./...
go test -v ./pkg/...  # verbose test output for specific packages
```

## Architecture Overview

This is a Go REST API application using the Fiber web framework with a layered architecture pattern. The application provides a generic CRUD API with automatic validation using YAML schemas.

### Core Technologies
- **Fiber v2**: High-performance web framework
- **GORM**: ORM for database operations (configured for SQLite)
- **YAML-based validation**: Custom struct validator using schema files

### Project Structure

The application follows a clean architecture with clear separation of concerns:

1. **Handler Layer** (`pkg/handlers/`): Generic HTTP handlers using Go generics
   - Provides standard CRUD operations (GetAll, GetByID, Create, Update, Delete)
   - Integrates validation through function composition
   - Returns appropriate HTTP status codes and error messages

2. **Service Layer** (`pkg/services/`): Business logic layer
   - Thin abstraction over repository layer
   - Handles business rules and orchestration

3. **Repository Layer** (`pkg/repositories/`): Data access layer
   - Generic repository pattern using GORM
   - Supports preloading, transactions, and various query methods
   - Global `DB` variable holds the database connection

4. **Model Layer** (`pkg/model/`): Domain entities
   - GORM models with proper tags for JSON serialization
   - Complex models with relationships (e.g., InternalUser with Skills, Connects, DayOffs)
   - Custom JSON marshaling for special cases (ExtendedProps)

5. **Validator** (`pkg/validator/`): YAML-based validation system
   - Load validation rules from YAML files
   - Support for various types: string, number, boolean, array, object
   - Constraints: required, min/max, pattern matching, enums

### Key Design Patterns

1. **Generic Handlers/Services/Repositories**: Heavy use of Go generics to avoid code duplication
2. **Dependency Injection**: Services injected into handlers, repositories into services
3. **Schema-based Validation**: External YAML files define validation rules (see `schemas/` directory)

### Database

- Uses SQLite by default (`db.sqlite`)
- Connection configured through environment variable `DB_SOURCE`
- Models use GORM tags for table/column mapping

### API Endpoints

The main.go sets up these endpoint patterns:
- `GET /{resource}` - List all resources
- `GET /{resource}/:id` - Get single resource
- `POST /{resource}` - Create new resource (with validation)
- `PUT /{resource}/:id` - Update resource (with validation)
- `DELETE /{resource}/:id` - Delete resource

Currently configured resources:
- `/internal_users`
- `/skill`

### Environment Variables

Handled by `pkg/commons/env.go`:
- `DB_SOURCE`: Database connection string (default: "db.sqlite")

### Important Implementation Notes

1. **Validation**: Controllers use `GetValidator()` to load YAML schemas and pass them to handlers via function composition
2. **ID Handling**: All models have an `int64` ID field that is set automatically after creation
3. **Error Messages**: Generic error messages use the handler's Name() method
4. **Repository Flags**: `FlagLog` can be toggled for debugging database queries