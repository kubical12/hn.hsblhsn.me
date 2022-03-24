dep-go:
	@go mod download -x

dep-py:
	@pip3 install -r requirements.txt


build-go:
	@go build -o bin/hn-api ./api/cmd/hn-api



proto-go:
	@cd api/internal/grpc/readabilityclient && \
		protoc \
			--proto_path=../protos \
			--go_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
			../protos/readability.proto		
	
proto-py:
	@cd api/internal/grpc/readabilityserver && \
		python3 \
			-m grpc_tools.protoc \
			--proto_path=../protos \
			--python_out=. \
			--grpc_python_out=. \
			../protos/readability.proto	

proto:
	@make proto-go proto-py


run-grpc:
	@cd api/internal/grpc/readabilityserver && \
		python3 readability_server.py

