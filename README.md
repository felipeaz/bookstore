# Bookstore
This project is a backend API developed using microservice architecture for a book store. Each Service communicates with each other using gRPC protocol, and their endpoints were built with REST and everything is running on a Docker container.

To start the application, run the following command:

``make install``

Before start using it, open another terminal on the project folder ``/bookstore`` and run the api gateway setup

``make gateway``

# Architecture
The structure was built using ``Clean Architecture`` - The application is divided in layers which uses adapters to communicate with external/internal layers (ports and adapters).

All the tools that are not related to the application domain such as MySQL, Redis, Logger and Auth are defined at infra layer level, outside the domain. The reason for that is, the application domain should be responsible to its own functionality and must run without coupling.

The application domain layer knows only the interface of infra tools, those are used as adapters.

```
./bookstore/
├── build
│   └── docker    
│   │   └── images
│   │   │   └── go
│   │   │       └── Dockerfile
│   └── server    
│       └── domain_item
│           └── router
│           │   └── build
│           │   │   └── domain_routes
│           │   └── router.go
│           └── router
│               └── server
│                   └── server.go
├── cmd
│   └── domain_item
│       └── main.go
├── internal
│   └── app
│   │   └── constants
│   │    │   ├── errors
│   │    │   │   └── errors_constants.go
│   │    ├── database
│   │    │   └── database_interface
│   │    ├── domain
│   │    │   └── domain_service
│   │    │       └── domain_item
│   │    │           └── handler
│   │    │           │   └── domain_item_handler.go
│   │    │           └── model
│   │    │           │   └── domain_item_model.go
│   │    │           │   └── converter
│   │    │           │       └── domain_item_model_converter.go
│   │    │           └── module
│   │    │           │   └── domain_item_module.go
│   │    │           └── repository
│   │    │               └── interface
│   │    │               │    └── domain_item_repository_interface.go
│   │    │               └── mock
│   │    │               │    └── domain_item_repository_mock.go
│   │    │               └── domain_item_repository.go
│   │    └── logger
│   │    │   └── service_logger.go
│   │    └── middleware
│   │        └── middleware.go
├── scripts
│   └── shell_script_file.sh
├── vendor
│   └── dependencies
├── .env_example
├── docker-compose.yml
├── go.mod
│   └── go.sum
├── makefile
└── README.md
```

### Directories

| Dir |Content|
| --- | --- |
| cmd | The cmd contains the application services main files. |
| build | Build contains the code that will build the system  |
| infra | The infra layer has all tools logic such as MySQL connection, Redis etc  |
| internal | Internal contains the Domain layer  |
| handler | The handlers are responsible of redirecting the user requests. |
| module | The modules will process user requests manipulating the data and sending them to the repository. |
| repository | Persistence Layer. |
| constants | The constants folder contains all constants used by the service.  |
| pkg | The PKG folder contains all lib codes used on the application. |

### Flow
```User Request > Login (Api Gateway - JWT Auth) > Domain Handler > Domain Module > Domain Repository```

# Services
Each service has its own database & cache. They're also linked to an Api Gateway (Kong) which is also responsible for the service discovery and authentication

## Documentation
### Account Service
https://documenter.getpostman.com/view/7958753/UVJeFbyS

### Inventory
https://documenter.getpostman.com/view/7958753/UVJeFbu9

### Orders
https://documenter.getpostman.com/view/7958753/UVJeFbyU

## gRPC
This protocol is being used to communicate between the microservices
#### Generate proto file
```protoc --go_out=plugins=grpc:build/ build/server/inventory/grpc/proto/inventory.proto```