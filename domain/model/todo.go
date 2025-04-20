package model

// Todo はJSONPlaceholderから取得するTodoアイテムの構造体です
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
