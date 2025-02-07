package tests

import (
	"net/http/httptest"
	"os"
	"strings"

	"groxy/application"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// MockUpstreamSenderはUpstreamSenderインターフェースのモック実装
type MockUpstreamSender struct {
	Response string
	Error    error
}

func (m *MockUpstreamSender) Send(request string) (string, error) {
	return m.Response, m.Error
}

var _ = Describe("RequestProcessor", func() {
	var (
		processor *application.RequestProcessor
	)

	BeforeEach(func() {
		mockSender := &MockUpstreamSender{
			Response: "mock response",
			Error:    nil,
		}
		processor = application.NewRequestProcessor(mockSender)
	})

	When("Process method", func() {
		It("should process the request and return a mock response", func() {
			req := httptest.NewRequest("POST", "/", strings.NewReader("body=test body"))
			os.Setenv("HTTP_TEST_HEADER", "test-value")

			response, err := processor.Process(req)
			Expect(err).To(BeNil())
			Expect(response).To(Equal("mock response"))
		})
	})
})
