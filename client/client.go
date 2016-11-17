package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ddliu/go-httpclient"
)

const sessionKey = "YTSESSIONID"
const principalKey = "jetbrains.charisma.main.security.PRINCIPAL"

// Issue models the response for a single isse from the API
type Issue struct {
	ID       string `json:"id"`
	EntityID string `json:"entityId"`
	Field    []struct {
		Name    string      `json:"name"`
		Value   interface{} `json:"value"`
		ValueID []string    `json:"valueId,omitempty"`
		Color   struct {
			Bg string `json:"bg"`
			Fg string `json:"fg"`
		} `json:"color,omitempty"`
	} `json:"field"`
	Comment []struct {
		ID             string `json:"id"`
		Author         string `json:"author"`
		AuthorFullName string `json:"authorFullName"`
		IssueID        string `json:"issueId"`
		// ParentID interface{} `json:"parentId"`
		Deleted             bool   `json:"deleted"`
		Text                string `json:"text"`
		ShownForIssueAuthor bool   `json:"shownForIssueAuthor"`
		Created             int64  `json:"created"`
		// Updated interface{} `json:"updated"`
		// PermittedGroup interface{} `json:"permittedGroup"`
		Replies []interface{} `json:"replies"`
	} `json:"comment"`
	Tag []struct {
		Value    string `json:"value"`
		CSSClass string `json:"cssClass"`
	} `json:"tag"`
}

// YouTrackClient exposes methods to issue requests to the YouTrack REST API
type YouTrackClient struct {
	baseURL   string
	session   string
	principal string
	expires   time.Time
}

// NewYouTrackClient returns an authenticated client to operate against the YouTrack API
func NewYouTrackClient(host, login, password string) YouTrackClient {
	client := YouTrackClient{baseURL: host + "/rest/"}

	client.login(login, password)

	return client
}

// GetIssue retrieves the details for the prvided issue ID
func (c *YouTrackClient) GetIssue(id string) (Issue, error) {
	res, err := httpclient.
		WithCookie(&http.Cookie{Name: sessionKey, Value: c.session}).
		WithCookie(&http.Cookie{Name: principalKey, Value: c.principal}).
		WithHeader("Accept", "application/json").
		Get(c.baseURL+"issue/"+id, nil)

	if err != nil {
		return Issue{}, err
	}

	c.setCredsFromCookies(res.Cookies())

	var issue Issue
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&issue)
	if err != nil {
		return Issue{}, err
	}

	return issue, nil
}

// CommandIssue sends a command string to apply to the given issue ID
func (c *YouTrackClient) CommandIssue(id, command, comment string) (string, error) {
	url := c.baseURL + "issue/" + id + "/execute"

	var params = make(map[string]string)

	params["command"] = command
	if comment != "" {
		params["comment"] = comment
	}

	res, err := httpclient.
		WithCookie(&http.Cookie{Name: sessionKey, Value: c.session}).
		WithCookie(&http.Cookie{Name: principalKey, Value: c.principal}).
		WithHeader("Accept", "application/json").
		Post(url, params)

	if err != nil {
		return "", err
	}

	c.setCredsFromCookies(res.Cookies())

	body, err := res.ToString()
	if err != nil {
		return "", err
	}

	return body, nil

}

// SearchIssues gets a list of issues which match the provided query string.
func (c *YouTrackClient) SearchIssues(filter string) (string, error) {
	url := c.baseURL + "issue"

	res, err := httpclient.
		WithCookie(&http.Cookie{Name: sessionKey, Value: c.session}).
		WithCookie(&http.Cookie{Name: principalKey, Value: c.principal}).
		WithHeader("Accept", "application/json").
		Get(url, map[string]string{filter: filter})

	if err != nil {
		return "", err
	}

	c.setCredsFromCookies(res.Cookies())

	body, err := res.ToString()
	if err != nil {
		return "", err
	}

	return body, nil

}

func (c *YouTrackClient) login(login, password string) error {
	res, err := httpclient.Post(c.baseURL+"user/login", map[string]string{
		"login":    login,
		"password": password,
	})

	if err != nil {
		return err
	}

	c.setCredsFromCookies(res.Cookies())

	return nil
}

func (c *YouTrackClient) setCredsFromCookies(cookies []*http.Cookie) {
	for _, cookie := range cookies {
		switch cookie.Name {
		case sessionKey:
			c.session = cookie.Value
		case principalKey:
			c.principal = cookie.Value
			c.expires = cookie.Expires
		}
	}
}
