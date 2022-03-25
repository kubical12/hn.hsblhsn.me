FROM node:17.7-bullseye AS ui-builder
WORKDIR /ui
COPY package.json package.json
COPY package-lock.json package-lock.json
COPY Makefile Makefile
RUN make dep-ui
COPY . .
RUN make build-ui

FROM golang:1.17.7-bullseye AS api-builder
WORKDIR /api
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile
RUN make dep-api
COPY . .
COPY --from=ui-builder /ui/embedded/dist ./embedded/dist
RUN make build

FROM python:3.9-bullseye
WORKDIR /readability
COPY requirements.txt requirements.txt
COPY Makefile Makefile
RUN make dep-readability
COPY --from=api-builder /api/bin/hn /app/hn

EXPOSE 8080
CMD ["/app/hn"]