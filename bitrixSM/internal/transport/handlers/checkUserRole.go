package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CheckUserRole(id int64) (string, error) {
	url := fmt.Sprintf("example.url/users/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var role string
	if err := json.Unmarshal(body, &role); err != nil {
		return "", err
	}
	return role, nil
}
