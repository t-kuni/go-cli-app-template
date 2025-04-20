package external

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rotisserie/eris"
	"github.com/samber/do"
	"github.com/t-kuni/go-cli-app-template/domain/infrastructure/external"
	"github.com/t-kuni/go-cli-app-template/domain/model"
)

// TodoClient はJSONPlaceholderからTodoを取得するクライアントの実装です
type TodoClient struct {
	baseURL string
}

// NewTodoClient はTodoClientのインスタンスを生成します
func NewTodoClient(i *do.Injector) (external.ITodoClient, error) {
	return &TodoClient{
		baseURL: "https://jsonplaceholder.typicode.com",
	}, nil
}

// FetchTodo はJSONPlaceholderからTodoを取得します
func (c *TodoClient) FetchTodo(id int) (*model.Todo, error) {
	url := fmt.Sprintf("%s/todos/%d", c.baseURL, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, eris.Wrap(err, "HTTPリクエストに失敗しました")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, eris.New(fmt.Sprintf("HTTPリクエストが失敗しました: ステータスコード %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, eris.Wrap(err, "レスポンスボディの読み取りに失敗しました")
	}

	var todo model.Todo
	if err := json.Unmarshal(body, &todo); err != nil {
		return nil, eris.Wrap(err, "JSONのパースに失敗しました")
	}

	return &todo, nil
}
