############################
# STEP 1 build executable binary
############################
FROM golang:1.20-alpine3.18 as builder

RUN apk --no-cache update && apk --no-cache add git tzdata

# Create appuser.
ENV USER=appuser
ENV UID=10001

RUN adduser \
  --disabled-password \
  --gecos "" \
  --home "/nonexistent" \
  --shell "/sbin/nologin" \
  --no-create-home \
  --uid "${UID}" \
  "${USER}"

# Setting timezone
ENV TZ=Asia/Jakarta
RUN ln -s /usr/share/zoneinfo/$TZ /etc/localtime

# Set default working directory of container
WORKDIR $GOPATH/src/github.com/azkanurhuda/multi-finance-golang-clean-architecture

# Copy all
COPY . .

# Copy env file
RUN mkdir -p /go/bin
RUN cp app.yaml /go/bin

# Copy the database/migration directory
#COPY ./database/migration /go/bin/database/migration
RUN mkdir -p /go/bin/database/migration
RUN cp ./database/migration/* /go/bin/database/migration

# Build an excutable app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags="-w -s" -o /go/bin/multi-finance ./cmd/web

############################
# STEP 2 build a small image
############################
FROM scratch AS base

# Add Maintainer info
LABEL maintainer="Azka <nurhudaazka@gmail.com>"

# Setting timezone
ENV TZ=Asia/Jakarta
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/localtime /etc/localtime

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy excutable app and env from builder stage to base stage
COPY --from=builder /go/bin/multi-finance /go/bin/multi-finance
COPY --from=builder /go/bin/app.yaml app.yaml

# Set default user
USER appuser:appuser

# Expose app port
EXPOSE 3000
EXPOSE 4000

# Run app
ENTRYPOINT ["/go/bin/multi-finance"]