dep-api:
	@go mod download -x
dep-readability:
	@pip3 install -r requirements.txt
dep-ui:
	@npm install
dep:
	@make dep-api dep-readability dep-ui


build-api:
	@go build -o bin/hn-api ./api/cmd/hn-api
build-ui:
	@npm run build
build-ui-server: build-ui
	@go build -o bin/hn-ui ./api/cmd/hn-ui
build:
	@go build -o bin/hn ./api/cmd/hn



docker-build-api:
	@docker build -t hn-api -f ./dockerfiles/api.Dockerfile .
docker-build-ui:
	@docker build -t hn-ui -f ./dockerfiles/api.Dockerfile .
docker-build:
	@docker build -t hn -f ./dockerfiles/Dockerfile .



proto-api:
	@cd api/internal/grpc/readabilityclient && \
		protoc \
			--proto_path=../protos \
			--go_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
			../protos/readability.proto		
proto-readability:
	@cd api/internal/grpc/readabilityserver && \
		python3 \
			-m grpc_tools.protoc \
			--proto_path=../protos \
			--python_out=. \
			--grpc_python_out=. \
			../protos/readability.proto	
proto:
	@make proto-api proto-readability

