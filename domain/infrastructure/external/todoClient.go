//go:generate mockgen -source=$GOFILE -destination=${GOFILE}_mock.go -package=$GOPACKAGE
package external

import (
	"github.com/t-kuni/go-cli-app-template/domain/model"
)

// ITodoClient はTodoを取得するためのクライアントインターフェースです
type ITodoClient interface {
	// FetchTodo はJSONPlaceholderからTodoを取得します
	FetchTodo(id int) (*model.Todo, error)
}
