# Build the Go Binary.
FROM golang:1.22 as build_pokemon-srv
ENV CGO_ENABLED 0
ARG BUILD_REF

# Create the service directory and the copy the module files first and then
# download the dependencies. If this doesn't change, we won't need to do this
# again in future builds.
# RUN mkdir /service
# COPY go.* /service/
# WORKDIR /service
# RUN go mod download

# Copy the source code into the container.
COPY . /service

# Build the service binary.
WORKDIR /service/app/pokemon-srv
RUN go build -ldflags "-X main.build=${BUILD_REF}"


# Run the Go Binary in Alpine.
FROM alpine:3.19
ARG BUILD_DATE
ARG BUILD_REF
RUN addgroup -g 1000 -S pokemon && \
    adduser -u 1000 -h /service -G pokemon -S pokemon
COPY --from=build_pokemon-srv --chown=pokemon:pokemon /service/app/pokemon-srv/pokemon-srv /service/pokemon-srv
WORKDIR /service
USER pokemon
CMD ["./pokemon-srv"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
      org.opencontainers.image.title="pokemon-srv" \
      org.opencontainers.image.authors="Manuel de la Peña <mdelapenya@gmail.com>" \
      org.opencontainers.image.source="https://github.com/mdelapenya" \
      org.opencontainers.image.revision="${BUILD_REF}" \
      org.opencontainers.image.vendor="mdelapenya"