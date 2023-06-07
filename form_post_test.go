package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstname := request.PostForm.Get("first_name")
	lastname := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "%s %s", firstname, lastname)

}

func TestFormPost(t *testing.T) {
	requsetBody := strings.NewReader("first_name=Muhammad&last_name=Shobir")
	request := httptest.NewRequest("POST", "http://localhost:8080", requsetBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recoreder := httptest.NewRecorder()

	FormPost(recoreder, request)

	response := recoreder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}
