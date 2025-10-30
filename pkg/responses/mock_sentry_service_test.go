package responses_test

type MockSentryService struct {
	CapturedErrors []error
	CallCount      int
}

func (m *MockSentryService) CaptureException(err error) {
	m.CapturedErrors = append(m.CapturedErrors, err)
	m.CallCount++
}

func (m *MockSentryService) WasCalled() bool {
	return m.CallCount > 0
}

func (m *MockSentryService) Reset() {
	m.CapturedErrors = []error{}
	m.CallCount = 0
}
