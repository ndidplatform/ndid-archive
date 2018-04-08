package apitest

import "fmt"

// RelyingParty represents a RP
type RelyingParty interface {
	SendRequestAsync(params SendRequestAsyncParams) SendRequestAsyncResult
}

// SendRequestAsyncResult is the result of calling a send_request_async
type SendRequestAsyncResult struct {
	RequestID string
}

// SendRequestAsyncParams represents SendRequestAsync parameters
type SendRequestAsyncParams struct {
	ReferenceID string
}

type mockRelyingParty struct {
	requestIDMap map[string]string
}

// NewMockRelyingParty creates a mock relying party
// that doesn't actually make any request to any server
func NewMockRelyingParty() RelyingParty {
	m := new(mockRelyingParty)
	m.requestIDMap = make(map[string]string)
	return m
}

// Performs a fake async request
func (m *mockRelyingParty) SendRequestAsync(params SendRequestAsyncParams) SendRequestAsyncResult {
	requestID := m.requestIDMap[params.ReferenceID]
	if requestID == "" {
		requestID = fmt.Sprintf("req%d", len(m.requestIDMap))
		m.requestIDMap[params.ReferenceID] = requestID
	}
	return SendRequestAsyncResult{requestID}
}
