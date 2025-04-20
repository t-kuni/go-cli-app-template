package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/api"
	"github.com/t-kuni/go-cli-app-template/domain/model"
)

// TodoClient はJSONPlaceholderからTodoを取得するクライアントの実装です
type TodoClient struct {
	baseURL string
}

// NewTodoClient はTodoClientのインスタンスを生成します
func NewTodoClient(i *do.Injector) (api.ITodoClient, error) {
	return &TodoClient{
		baseURL: "https://jsonplaceholder.typicode.com",
	}, nil
}

// FetchTodo はJSONPlaceholderからTodoを取得します
func (c *TodoClient) FetchTodo(id int) (*model.Todo, error) {
	url := fmt.Sprintf("%s/todos/%d", c.baseURL, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("HTTPリクエストに失敗しました: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTPリクエストが失敗しました: ステータスコード %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("レスポンスボディの読み取りに失敗しました: %w", err)
	}

	var todo model.Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		return nil, fmt.Errorf("JSONのパースに失敗しました: %w", err)
	}

	return &todo, nil
}
