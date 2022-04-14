dep-backend:
	@go mod download
dep-readability:
	@pip3 install -r requirements.txt
dep-frontend:
	@npm ci
dep:
	@make dep-backend dep-readability dep-frontend


lint-backend:
	@golangci-lint run --fix ./...
lint-frontend:
	@npm run lint
lint:
	@make lint-backend lint-frontend


build-backend:
	@go build ./backend/...
build-frontend:
	@npm run build
build:
	@go build -o bin/hackernews ./cmd/hackernews



docker-build:
	@docker build -t hackernews -f ./Dockerfile .



gql:
	@go run github.com/99designs/gqlgen

proto-backend:
	@cd backend/graph/internal/grpc/readabilityclient && \
		protoc \
			--proto_path=../protos \
			--go_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
			../protos/readability.proto		
proto-readability:
	@cd backend/graph/internal/grpc/readabilityserver && \
		python3 \
			-m grpc_tools.protoc \
			--proto_path=../protos \
			--python_out=. \
			--grpc_python_out=. \
			../protos/readability.proto	
proto:
	@make proto-backend proto-readability

