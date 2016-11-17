package client

import (
	"net/http"
	"time"

	"github.com/ddliu/go-httpclient"
)

const sessionKey = "YTSESSIONID"
const principalKey = "jetbrains.charisma.main.security.PRINCIPAL"

// YouTrackClient is used to interact with YouTrack API at BaseURL
type YouTrackClient struct {
	BaseURL   string
	Session   string
	Principal string
	Expires   time.Time
}

// NewYouTrackClient returns an authenticated YouTrackClinet at host using the
// provided login and password
func NewYouTrackClient(host, login, password string) YouTrackClient {
	client := YouTrackClient{BaseURL: host + "/rest/"}

	err := client.login(login, password)
	if err != nil {
		panic(err)
	}

	return client
}

// GetIssue returns the string representation of a YouTrack issue
// Note: this string will contain XML
func (client *YouTrackClient) GetIssue(id string) (string, error) {
	res, err := httpclient.WithCookie(&http.Cookie{
		Name:  sessionKey,
		Value: client.Session,
	}).WithCookie(&http.Cookie{
		Name:  principalKey,
		Value: client.Principal,
	}).Get(client.BaseURL+"issue/"+id, nil)

	if err != nil {
		return "", err
	}

	client.setCredsFromCookies(res.Cookies())

	body, err := res.ToString()
	if err != nil {
		return "", err
	}

	return body, nil
}

// CommandIssue will apply command onto the issue denoted by id while optionally
// adding a comment of comment, if not empty.
func (client *YouTrackClient) CommandIssue(id, command, comment string) (string, error) {
	url := client.BaseURL + "issue/" + id + "/execute"

	var params = make(map[string]string)

	params["command"] = command
	if comment != "" {
		params["comment"] = comment
	}

	res, err := httpclient.WithCookie(&http.Cookie{
		Name:  sessionKey,
		Value: client.Session,
	}).WithCookie(&http.Cookie{
		Name:  principalKey,
		Value: client.Principal,
	}).Post(url, params)

	if err != nil {
		return "", err
	}

	client.setCredsFromCookies(res.Cookies())

	body, err := res.ToString()
	if err != nil {
		return "", err
	}

	return body, nil

}

func (client *YouTrackClient) login(login, password string) error {
	res, err := httpclient.Post(client.BaseURL+"user/login", map[string]string{
		"login":    login,
		"password": password,
	})

	if err != nil {
		return err
	}

	client.setCredsFromCookies(res.Cookies())

	return nil
}

func (client *YouTrackClient) setCredsFromCookies(cookies []*http.Cookie) {
	for _, cookie := range cookies {
		switch cookie.Name {
		case sessionKey:
			client.Session = cookie.Value
		case principalKey:
			client.Principal = cookie.Value
			client.Expires = cookie.Expires
		}
	}
}
