dep-api:
	@go mod download
dep-readability:
	@pip3 install -r requirements.txt
dep-ui:
	@npm install
dep:
	@make dep-api dep-readability dep-ui


build-api:
	@go build -o bin/hackernews ./cmd/hackernews
build-ui:
	@npm run build
build:
	@go build -o bin/hackernews ./cmd/hackernews



docker-build:
	@docker build -t hackernews -f ./Dockerfile .



proto-api:
	@cd backend/internal/grpc/readabilityclient && \
		protoc \
			--proto_path=../protos \
			--go_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
			../protos/readability.proto		
proto-readability:
	@cd backend/internal/grpc/readabilityserver && \
		python3 \
			-m grpc_tools.protoc \
			--proto_path=../protos \
			--python_out=. \
			--grpc_python_out=. \
			../protos/readability.proto	
proto:
	@make proto-api proto-readability

