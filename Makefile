.PHONY: generate-todo

generate-todo:
	protoc -I todoapi/ todoapi/todo.proto --go_out=./ --go-grpc_out=./