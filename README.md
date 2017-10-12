# cluster-examples

## Setup
Startup consul in Docker:
```
 make setup-infra
 ```
 
 ## Build Docker Images
 ```
 make build-all
 ```
 
 ## Start nodes
 Start the seed node in Docker:
 ```
 make run-seed
 ```
 Start the member node in Docker:
 ```
 make run-member
 ```
