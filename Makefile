# Generate go code
.PHONY: api
api:
	kratos proto client ./proto/users/users.proto -p ./proto/lib
	rm -rf openapi.yaml

# Generate swaager json files
.PHONY: docs
docs:
	protoc --proto_path=. \
           --proto_path=./proto/lib \
           --proto_path=${GOPATH}/src \
           --openapiv2_out=./docs/swagger \
           --openapiv2_opt=json_names_for_fields=false \
           ./proto/users/users.proto

.PHONY: all
all: api docs
