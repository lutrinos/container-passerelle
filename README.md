# Container Passerelle

A basic docker container used to pass ping securely through a proxy.

In a compose file, add the following

```yaml
services:
  passerelle:
    image: ghcr.io/lutrinos/container-passerelle:latest
    volumes:
      - ./config.yaml:/config.yaml:ro
    restart: unless-stopped
```

With a list of services you want to ping in `config.yaml`
```yaml
- id: service1
  url: http://service_1:8080/ping
- id: service2
  url: http://service_2:3000/healthcheck
```

Then all requests made to `/service1` will have the status code of a request made to `http://service_1:8080/ping`.