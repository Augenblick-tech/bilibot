package client

import (
	"bytes"
	"io"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sync"
	"time"
)

var UA = []string{
	"Opera/9.30 (Nintendo Wii; U; ; 2047-7; en)",
	"Mozilla/5.0 (Linux; Android 4.0.4; BNTV400 Build/IMM76L) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.111 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.155 Safari/537.36 OPR/31.0.1889.174",
	"Mozilla/5.0 (X11; Linux 3.8-6.dmz.1-liquorix-686) KHTML/4.8.4 (like Gecko) Konqueror/4.8",
}

type Visitor struct {
	UserAgent         string
	client            *http.Client
	RequestCallbacks  []RequestCallback
	ResponseCallbacks []ResponseCallback
	lock              *sync.RWMutex
}

type RequestCallback func(*Request)

type ResponseCallback func(*Response)

type Request struct {
	URL     *url.URL
	Headers *http.Header
	Method  string
	Body    io.Reader
}

type Response struct {
	StatusCode int
	Body       []byte
	Request    *Request
	Headers    *http.Header
}

func NewVisitor() *Visitor {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	return &Visitor{
		UserAgent: getRandomUA(),
		client: &http.Client{
			Jar:     jar,
			Timeout: 10 * time.Second,
		},
		RequestCallbacks:  []RequestCallback{},
		ResponseCallbacks: []ResponseCallback{},
		lock:              &sync.RWMutex{},
	}
}

func (v *Visitor) OnRequest(f RequestCallback) {
	v.lock.Lock()
	v.RequestCallbacks = append(v.RequestCallbacks, f)
	v.lock.Unlock()
}

func (v *Visitor) OnResponse(f ResponseCallback) {
	v.lock.Lock()
	v.ResponseCallbacks = append(v.ResponseCallbacks, f)
	v.lock.Unlock()
}

func (v *Visitor) handleOnRequest(r *Request) {
	for _, f := range v.RequestCallbacks {
		f(r)
	}
}

func (v *Visitor) handleOnResponse(r *Response) {
	for _, f := range v.ResponseCallbacks {
		f(r)
	}
}

func (v *Visitor) SetCookies(URL string, cookies []*http.Cookie) error {
	u, err := url.Parse(URL)
	if err != nil {
		return err
	}
	v.client.Jar.SetCookies(u, cookies)
	return nil
}

func (v *Visitor) Cookies(URL string) []*http.Cookie {
	u, err := url.Parse(URL)
	if err != nil {
		return nil
	}
	return v.client.Jar.Cookies(u)
}

func (v *Visitor) Visit(URL string) error {
	return v.fetch(URL, "GET", nil, nil)
}

func (v *Visitor) Post(URL string, requestData []byte) error {
	return v.fetch(URL, "POST", bytes.NewReader(requestData), nil)
}

func (v *Visitor) fetch(URL, method string, requestData io.Reader, hdr http.Header) error {
	if hdr == nil {
		hdr = http.Header{"User-Agent": []string{v.UserAgent}}
	}
	req, err := http.NewRequest(method, URL, requestData)
	if err != nil {
		return err
	}
	req.Header = hdr

	request := &Request{
		URL:     req.URL,
		Headers: &req.Header,
		Method:  req.Method,
		Body:    req.Body,
	}

	v.handleOnRequest(request)

	if method == "POST" && req.Header.Get("Content-Type") == "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "*/*")
	}

	resp, err := v.client.Do(req)
	if err != nil {
		return err
	}

	// if resp.Request != nil {
	// 	req = resp.Request
	// }

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	response := &Response{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    &resp.Header,
	}

	v.handleOnResponse(response)

	return nil
}

func getRandomUA() string {
	rand.Seed(time.Now().Unix())
	return UA[rand.Intn(len(UA))]
}
