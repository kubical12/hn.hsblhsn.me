# Development

This project is written in Go, Typescript, React and Python.

- **Go** does all the weight lifting on backend.
- **React** and **Typescript** is used for the frontend.
- **Python** is used for the converting unstructured html page to a readable html page.
- **GraphQL** is used for the api.
- **gRPC** is used for the internal interprocess communication.
- **Docker** is used for the deployment.

## Go.

- [gqlgen](https://gqlgen.com) is used to generate the graphql server.
- [goquery](https://pkg.go.dev/github.com/PuerkitoBio/goquery) for parsing and processing the html page.
- [opengraph](https://pkg.go.dev/github.com/otiai10/opengraph/v2) is used for parsing the opengraph meta data.
- [zap](https://pkg.go.dev/go.uber.org/zap) is used for logging.
- [fx](https://pkg.go.dev/go.uber.org/fx) is used for dependency injection.
- [embed](https://go.dev/blog/go1.16) is used for embedding the static contents.

## React, Typescript, Vite.

- [baseui](https://baseweb.design) is used for the ui components.
- [styletron](https://styletron.org) is used for the styling.
- [tailwindcss](https://tailwindcss.com) is used for the styling.
- [react-router-dom](https://www.npmjs.com/package/react-router-dom) is used for the routing.
- [react-helmet-async](https://www.npmjs.com/package/react-helmet-async) is used for the meta data.
- [vite](https://vitejs.dev) is used for building and developing the frontend..
- [@apollo/client](https://www.apollographql.com) is used for the sending graphql queries.

## Python

- [readability-lxml](https://pypi.org/project/readability-lxml/) is used for the converting the html page to a readable
  html page.

## gRPC and Protobuf

- [protobuf](https://developers.google.com/protocol-buffers/) and [gRPC](https://grpc.io) is used for the generating the
  gRPC server and client.

## How things work.

The React client sends graphql queries to the Go server. The Go server sends http requests to Hacker News API. After
getting the response, the Go server sends http request to any third party content link received from Hacker News API. It
then makes a gRPC call to the Python server to convert the html page to a readable html page. After everything is done,
It sends the response back to the React client. You see the result!