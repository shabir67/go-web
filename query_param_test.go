package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func SayHello(w http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello Strangers")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}

}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "htttp://localhost:8080/hello?name=Shobir", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	require.Equal(t, "Hello Shobir", bodyString, "Result must be 'Hello Gais'")

}

func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("first_name")
	lastname := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)

}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?first_name=Muhammad&last_name=Shobir", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	require.Equal(t, "Hello Muhammad Shobir", bodyString, "Result must be 'Hello Gais'")

}
