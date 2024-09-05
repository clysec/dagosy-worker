# Dagosy Worker Rest API
This is the worker API for the Dagosy project. It's meant to be a universal Rest API that can perform a wide variety of tasks, together with a plugin system to extend the functionality with custom endpoints/tasks.

The worker is designed to be stateless, which means all information required to perform a task is passed in the request.

## API Documentation
The API documentation is generated using [Swagger](https://swagger.io/). To generate the documentation, run the following command:
```
swag init -g main.go --output docs --parseDependency --md .
```

## Getting Started
To get started, run the following command:
```
go run main.go
```

## Included Endpoints
### Files
The files endpoints are used to interact with a variety of file systems. The endpoints are based on the [rclone](https://rclone.org/) library, and support all remote filesystems that rclone supports.

### Domain
The domain endpoints are used to interact with domain registrars, and currently include Whois and Nameserver lookup.

### DNS
The DNS endpoints are used to interact with DNS servers, and currently include DNS lookup and Reverse DNS lookup. DNS Requests are forwarded to a number of DNS over HTTPS providers.

## Development
### Adding your own endpoints
TBD

## TODO
### Files
- [ ] Add support for bulk rename operations

