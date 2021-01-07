package fetch

import "testing"

type testResponse struct {
	UserID   int    `json:"userId"`
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

func Test(t *testing.T) {
	var response testResponse
	err := FetchJSON("https://jsonplaceholder.typicode.com/todos/1", "GET", nil, &response, FetchOptions{})
	if err != nil {
		t.Error(err)
	}

	// Check response
	if !(response.ID == 1 && response.UserID == 1 && !response.Complete && response.Title == "delectus aut autem") {
		t.Error("Response incorrect")
	}
}
