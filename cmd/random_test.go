package cmd

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockDoType func(req *http.Request) (*http.Response, error)

type MockClient struct {
	MockDo MockDoType
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return m.MockDo(req)
}

func TestGetRandomJokeNoTermPositive(t *testing.T) {
	var jsonResponse = `{"id":"MZga2gFlysc","joke":"What is the least spoken language in the world?\r\nSign Language","status":200}`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	getRandomJoke()
}

func TestEvaluateCallNoTermPositive(t *testing.T) {
	var jsonResponse = `{"id":"MZga2gFlysc","joke":"What is the least spoken language in the world?\r\nSign Language","status":200}`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	evaluateCall("", nil)
}

func TestGetRandomJokeNoTermBadJSON(t *testing.T) {
	var jsonResponse = `"id":"MZga2gFlysc","joke":"What is the least spoken language in the world?\r\nSign Language","status":200}`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	getRandomJoke()
}

func TestGetRandomJokeNoTermBadCall(t *testing.T) {
	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       nil,
			}, errors.New("Mock Error")
		},
	}

	getRandomJoke()
}

func TestEvaluateCallWithTermPositive(t *testing.T) {
	var jsonResponse = `{"current_page":1,"limit":20,"next_page":1,"previous_page":1,"results":[{"id":"xc21Lmbxcib","joke":"How did the hipster burn the roof of his mouth? He ate the pizza before it was cool."},{"id":"GlGBIY0wAAd","joke":"How much does a hipster weigh? An instagram."},{"id":"NRuHJYgaUDd","joke":"How many hipsters does it take to change a lightbulb? Oh, it's a really obscure number. You've probably never heard of it."}],"search_term":"hipster","status":200,"total_jokes":3,"total_pages":1}`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	evaluateCall("hipster", nil)
}

func TestGetRandomJokeWithTermPositive(t *testing.T) {
	var jsonResponse = `{"current_page":1,"limit":20,"next_page":1,"previous_page":1,"results":[{"id":"xc21Lmbxcib","joke":"How did the hipster burn the roof of his mouth? He ate the pizza before it was cool."},{"id":"GlGBIY0wAAd","joke":"How much does a hipster weigh? An instagram."},{"id":"NRuHJYgaUDd","joke":"How many hipsters does it take to change a lightbulb? Oh, it's a really obscure number. You've probably never heard of it."}],"search_term":"hipster","status":200,"total_jokes":3,"total_pages":1}`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	getRandomJokeWithTerm("hipster")
}

func TestGetRandomJokeWithTermBadJSON(t *testing.T) {
	var jsonResponse = `"current_page":1,"limit":20,"next_page":1,"previous_page":1,"results":[{"id":"xc21Lmbxcib","joke":"How did the hipster burn the roof of his mouth? He ate the pizza before it was cool."},{"id":"GlGBIY0wAAd","joke":"How much does a hipster weigh? An instagram."},{"id":"NRuHJYgaUDd","joke":"How many hipsters does it take to change a lightbulb? Oh, it's a really obscure number. You've probably never heard of it."}],"search_term":"hipster","status":200,"total_jokes":3,"total_pages":1}`

	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))

	Client = &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}

	getRandomJokeWithTerm("hipster")
}
