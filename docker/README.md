# Docker

You can create a Stallion server on docker from the image below.

### Image
```shell
amirhossein21/stallion:v1.1.1
```

### Docker run
To run image on a single container:
```shell
docker run -p 7025:7025 -d stallion amirhossein21/stallion:v1.1.1
```

### Docker compose
Docker compose for stallion server:
```yaml
version: "3.9"
services:
  stallion-server:
    image: amirhossein21/stallion:v1.1.1
    ports:
      - "7025:7025"
    environment:
      SERVER_PORT: 7025
```

## Docker hub
You can check docker hub [repository](https://hub.docker.com/repository/docker/amirhossein21/stallion) for more information and image lists.
