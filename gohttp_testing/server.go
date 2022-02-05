package gohttp_testing

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"sync"

	"github.com/getmiranda/go-httpclient/core"
)

var (
	MockupServer = mockServer{
		mocks:      make(map[string]*Mock),
		httpClient: &httpClientMock{},
	}
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex
	mocks       map[string]*Mock
	httpClient  core.HttpClient
}

// Start sets the enviroment to send all client requests
// to the mockup server.
func (m *mockServer) Start() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = true
}

// Stop stop sending requests to the mockup server
func (m *mockServer) Stop() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = false
}

// AddMockups add a mock to the mockup server.
func (m *mockServer) AddMock(mock *Mock) {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	key := m.getMockKey(mock.Method, mock.URL, mock.RequestBody)
	m.mocks[key] = mock
}

// IsEnabled check whether the mock environment is enabled or not.
func (m *mockServer) IsEnabled() bool {
	return m.enabled
}

// GetMockedClient gets the http mock client.
func (m *mockServer) GetMockedClient() core.HttpClient {
	return m.httpClient
}

// DeleteMocks delete all mocks in every new test case to ensure a clean environment.
func (m *mockServer) DeleteMocks() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.mocks = make(map[string]*Mock)
}

func (m *mockServer) cleanBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\n", "")
	body = strings.ReplaceAll(body, "\t", "")
	return body
}

func (m *mockServer) getMockKey(method, url, body string) string {
	hasher := md5.New()
	hasher.Write([]byte(method + url + m.cleanBody(body)))
	key := hex.EncodeToString(hasher.Sum(nil))
	return key
}
