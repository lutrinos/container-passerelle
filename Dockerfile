# --- Build ---
FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /passerelle .

# --- Runtime ---
FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=build /passerelle /passerelle
USER nonroot
ENTRYPOINT ["/passerelle"]