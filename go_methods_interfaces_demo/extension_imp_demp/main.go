package extension_imp_demp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type customWriter struct{}
type GithubResponse []struct {
	FullName string `json:"full_name"`
}

func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GithubResponse
	err = json.Unmarshal(p, &resp)
	if err != nil {
		fmt.Println("Unmarshal Error:", err)
		return 0, err
	}
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), err
}
func main() {
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	writer := customWriter{}
	io.Copy(writer, resp.Body)
}
