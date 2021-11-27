# Bookstore
This project is a backend API developed using microservice architecture for a book store. Each Service communicates with each other using gRPC protocol, and their endpoints were built with REST and everything is running on a Docker container.

To start the application, run the following command:
``make install``

# Architecture
The structure was built using ``Clean Architecture`` - The application is divided in layers which uses adapters to communicate with external/internal layers (ports and adapters).

# Services
Each service has its own database. They're also linked to an Api Gateway (Kong) which is also responsible for the service discovery and authentication

## gRPC
This protocol is being used to communicate between the microservices
#### Generate proto file
```protoc --go_out=plugins=grpc:build/ build/server/inventory/grpc/proto/inventory.proto```