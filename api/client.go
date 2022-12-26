package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	// DefaultToken = "78WfWjphakysAZ4uParrmPgSxDPCUAXqw3Q7aHYH"
	DefaultToken = "WJHdk5dNVR6sFb5NZAABGG9JGZTNXzScQG4WfQ8t"
)

type Client struct {
	token  string
	client *http.Client
}

func closeBody(body io.ReadCloser) {
	io.Copy(io.Discard, body)
	body.Close()
}

// Returns a Client object that will use the provided Quizlet session cookie to
// interact with the Quizlet API.
func NewClient(token string, client *http.Client) Client {
	return Client{token, client}
}

// Returns a Client object using a default HTTP client with a timeout of 10s.
func NewDefaultClient() Client {
	return NewClient(DefaultToken, &http.Client{
		Timeout: 30 * time.Second,
	})
}

func (c Client) headers() http.Header {
	return http.Header{
		"Host":            []string{"api.quizlet.com"},
		"Accept":          []string{"*/*"},
		"User-Agent":      []string{"QuizletIOS/6.2.1 (QuizletBuild/3; iPhone12,8; iOS 13.5.1; Scale/2.0)"},
		"Accept-Language": []string{"en-us"},
		"Authorization":   []string{"Bearer " + c.token},
	}
}

func (c Client) get(url string) (*http.Response, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header = c.headers()
	return c.client.Do(request)
}

func (c Client) ResolveURL(path string) (GenericResponse, error) {
	response, err := c.get(fmt.Sprintf("https://api.quizlet.com/3.6/resolve-url?url=%s", url.QueryEscape(path)))
	if err != nil {
		return GenericResponse{}, err
	}
	defer closeBody(response.Body)

	var result GenericResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return GenericResponse{}, err
	}
	return result, nil
}

func (c Client) SetsByID(id string) ([]Set, error) {
	params := url.Values{}
	params.Add("filters[creatorId]", id)
	params.Add("include[set][]", "subjectClassification_ae86c4b")
	params.Add("perPage", "2000")

	response, err := c.get(fmt.Sprintf("https://api.quizlet.com/3.6/sets?%s", params.Encode()))
	if err != nil {
		return nil, err
	}
	defer closeBody(response.Body)

	var result GenericResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	} else if len(result.Responses) == 0 {
		return nil, errors.New("no sets")
	}
	return result.Responses[0].Models.Set, nil
}

func (c Client) SetsByName(name string) ([]Set, error) {
	resolved, err := c.ResolveURL("/" + name)
	if err != nil {
		return nil, err
	} else if len(resolved.Responses) == 0 || len(resolved.Responses[0].Models.User) == 0 {
		return nil, fmt.Errorf("no user with username %s found", name)
	}

	return c.SetsByID(strconv.Itoa(resolved.Responses[0].Models.User[0].ID))
}

func (c Client) UserByID(id string) (User, error) {
	response, err := c.get(fmt.Sprintf("https://api.quizlet.com/3.6/users/%s", id))
	if err != nil {
		return User{}, err
	}
	defer closeBody(response.Body)

	var result GenericResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return User{}, err
	} else if len(result.Responses) == 0 || len(result.Responses[0].Models.User) != 1 {
		return User{}, fmt.Errorf("no user with ID %s found", id)
	}
	return result.Responses[0].Models.User[0], nil
}

func (c Client) Terms(id string) ([]Term, error) {
	params := url.Values{}
	params.Add("filters[setId]", id)
	params.Add("include[term][]", "definitionImage")
	params.Add("include[term][]", "wordCustomAudio")
	params.Add("include[term][]", "definitionCustomAudio")
	params.Add("include[term][set][]", "creator")
	params.Add("perPage", "2000")

	response, err := c.get(fmt.Sprintf("https://api.quizlet.com/3.6/terms?%s", params.Encode()))
	if err != nil {
		return nil, err
	}
	defer closeBody(response.Body)

	var result GenericResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	} else if len(result.Responses) == 0 {
		return nil, fmt.Errorf("no set with ID %s found", id)
	}
	return result.Responses[0].Models.Term, err
}

func (c Client) SearchSets(query string) ([]Set, error) {
	response, err := c.get(fmt.Sprintf("https://api.quizlet.com/3.6/sets/search?perPage=30&query=%s&showNumCreatedSets=1", url.QueryEscape(query)))
	if err != nil {
		return nil, err
	}
	defer closeBody(response.Body)

	var result GenericResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Responses[0].Models.Set, nil
}

func (c Client) SearchUsers(query string) ([]User, error) {
	response, err := c.get(fmt.Sprintf("https://api.quizlet.com/3.6/users/search?perPage=30&query=%s&showNumCreatedSets=1", url.QueryEscape(query)))
	if err != nil {
		return nil, err
	}
	defer closeBody(response.Body)

	var result GenericResponse
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result.Responses[0].Models.User, nil
}
