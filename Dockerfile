FROM node:17.7-bullseye AS ui-builder
WORKDIR /ui
COPY package.json package.json
COPY package-lock.json package-lock.json
COPY Makefile Makefile
RUN make dep-ui
COPY . .
RUN make build-ui

FROM golang:1.18-bullseye AS api-builder
WORKDIR /api
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
RUN make dep-api
COPY . .
COPY --from=ui-builder /ui/frontend/build ./frontend/build
RUN make build

FROM python:3.9-slim
RUN apt-get update \
    && apt-get install --no-install-recommends -y make \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /readability
COPY requirements.txt requirements.txt
COPY Makefile Makefile
RUN make dep-readability
WORKDIR /app
COPY --from=api-builder /api/bin/hackernews /app/hackernews
RUN rm -rf /readability

EXPOSE 8080
CMD ["/app/hackernews"]