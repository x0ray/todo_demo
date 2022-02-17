package main

import (
	"context"
	"fmt"

	api "github.com/x0ray/todo_demo/todoapi"
)

var todos map[int]string
var nextTodo int

type todoServer struct {
	api.UnimplementedTodoServer
}

func main() {
	fmt.Printf("Todo Server")
}

func (todoServer) CreateTodo(ctx context.Context, in *api.TodoItem) (*api.TodoItem, error) {
	nextTodo += 1
	item := nextTodo
	text := in.GetText()
	todos[item] = text
	resp := api.TodoItem{Id: item, Text: text}
	return resp, nil
}

func (todoServer) ReadTodos(context.Context, *api.VoidNoParams) (*api.TodoItems, error) {
	var resp []api.TodoItems
	for k, v := range todos {
		item := &api.TodoItem{Id: k, Text: v}
		resp.Items = append(resp.Items, item)
	}
	return resp, nil
}
