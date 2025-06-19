
## Architecture Diagram

```mermaid
graph TD
    A["Client"] -->|HTTP| B["API Gateway"]
    B -->|gRPC| C["Auth Service (Go + gRPC)"]
    B -->|gRPC| D["Shortener Service (Go + gRPC)"]
    B -->|gRPC| E["Redirector Service (Go + gRPC)"]
    C <-->|gRPC| D
    D <-->|gRPC| E
    E <-->|gRPC| C
```



## Install protobuf

```bash
brew install protobuf
```

## Install go
```bash
brew install go
```


## Install bun
```bash
curl -fsSL https://bun.sh/install | bash 
```

