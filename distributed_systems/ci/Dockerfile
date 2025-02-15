FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .

ARG IMAGE_TITLE
RUN make "build/cmd/${IMAGE_TITLE}"

FROM gcr.io/distroless/static:nonroot
ARG IMAGE_TITLE
WORKDIR /app

COPY --from=builder --chown=nonroot "/app/build/cmd/${IMAGE_TITLE}" "/app"

ENTRYPOINT ["/app/run"]