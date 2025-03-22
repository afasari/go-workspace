.PHONY: proto
proto:  ## Generate Protobuf file
	protoc -I todo/ todo/todo.proto --go_out=plugins=grpc:todo

.PHONY: tidy
tidy: ## Run golang mod tidy
	go mod tidy
	
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'