# layered-architecture-go
This is a sample Go package structure designed with a layered architecture following DDD (Domain-Driven Design) principles.

## Directory Structure
```
.
|-- /cmd
|   |-- /myapp
|       |-- main.go
|-- /domain
|   |-- /model
|   |   |-- user.go
|   |-- /repository
|   |   |-- user_repository.go
|   |-- /service
|       |-- user_service.go
|-- /application
|   |-- /usecase
|   |   |-- user_usecase.go
|   |-- /event
|   |   |-- user_created_event.go
|   |-- /dto
|   |   └── user_dto.go
|-- /infrastructure
│   ├── persistence/
│   │   ├── entity/
│   │   │   └── user_entity.go
│   │   └── repository/
│   │       └── user_repository_impl.go
│   └── mysql/
│       └── mysql_config.go
|-- /interface
│   ├── handler/ 
│   │   └── user_handler.go
|   |-- /middleware
|       |-- auth_middleware.go
|-- go.mod
|-- go.sum
```

## Package Descriptions

### `cmd`
- **Description**: Contains the application's entry points. Typically used for the main function that starts the application.

### `domain`
- **Description**: Represents the core business logic and entities of the application. It encapsulates the business rules and domain models.
  - `model`: Contains the domain entities and value objects.
  - `repository`: Defines interfaces for accessing data sources, such as databases or external services.
  - `service`: Implements domain-specific business logic that doesn't naturally fit within an entity or value object.

### `application`
- **Description**: Contains application-specific business rules and orchestrates the flow of data.
  - `usecase`: Orchestrates the flow of data to and from the entities and communicates between the domain and the interface layers.
  - `event`: Contains events that are triggered by certain actions within the application, allowing for decoupled and reactive architectures.
  - `dto`: Contains Data Transfer Objects used for data communication within the application layer.

### `infrastructure`
- **Description**: Provides concrete implementations of the repository interfaces, external services, and other technical concerns.
  - `persistence`: Contains implementations for accessing and storing data, typically database integrations.
    - `entity`: Contains structures reflecting the database table structures.
    - `repository`: Contains the concrete implementations of the repository interfaces.
  - `mysql`: Specific implementations and configurations for a MySQL database.

## `interface`
- **Description**: Handles the interaction between the user and the application.
  - `handler`: Contains the logic for handling specific API endpoints.
  - `middleware`: Implements cross-cutting concerns like authentication, logging, and error handling.
