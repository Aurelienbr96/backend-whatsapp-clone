FROM golang:1.23-bookworm as build

# Add non root user
RUN useradd -u 1001 nonroot

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=cache,target=/root/.cache/go-build \
  go mod download
COPY . .
RUN cd cmd && go build \
-ldflags="-linkmode external -extldflags -static" \
-tags netgo \
-o api-golang

FROM scratch


ENV GIN_MODE release
WORKDIR /
# Copy the passwd file
COPY --from=build /etc/passwd /etc/passwd
# Copy the binary from the build stage
COPY --from=build /app/cmd/api-golang api-golang
COPY --from=build /app/config /config
# Use nonroot user
USER nonroot

EXPOSE 8080

CMD ["/api-golang"]