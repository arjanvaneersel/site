package meetup

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const baseURL = "https://api.meetup.com/"

type Option func(*url.Values)

func WithDesc(val bool) Option {
	return func(v *url.Values) {
		v.Set("desc", strconv.FormatBool(val))
	}
}

func WithFields(val ...string) Option {
	return func(v *url.Values) {
		v.Set("fields", strings.Join(val, ","))
	}
}

type scrollType string

func (s scrollType) String() string {
	return string(s)
}

const (
	RecentPast   scrollType = "recent_past"
	NextUpcoming scrollType = "next_upcoming"
	FutureOrPast scrollType = "future_or_past"
)

func WithScroll(val scrollType) Option {
	return func(v *url.Values) {
		v.Set("scroll", val.String())
	}
}

func WithStatus(val ...string) Option {
	return func(v *url.Values) {
		v.Set("status", strings.Join(val, ","))
	}
}

type API struct {
	key string
}

func (a *API) call(path string, method string, v url.Values, body io.Reader) (io.ReadCloser, error) {
	req, err := http.NewRequest(method, baseURL+path, body)
	if err != nil {
		return nil, err
	}

	if v == nil {
		v = url.Values{}
	}
	v.Set("sign", "true")
	v.Set("key", a.key)
	req.URL.RawQuery = v.Encode()

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

func (a *API) GetEvents(group string, opts ...Option) ([]Event, error) {
	events := []Event{}

	params := url.Values{}
	for _, o := range opts {
		o(&params)
	}

	res, err := a.call(fmt.Sprintf("%s/events", group), http.MethodGet, params, nil)
	if err != nil {
		return nil, err
	}

	defer res.Close()
	if err := json.NewDecoder(res).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}

func New() (*API, error) {
	key := os.Getenv("MEETUP_KEY")
	if key == "" {
		return nil, fmt.Errorf("MEETUP_KEY not set")
	}
	return &API{
		key: key,
	}, nil
}
