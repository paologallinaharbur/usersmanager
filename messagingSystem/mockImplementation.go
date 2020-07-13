package messagingSystem

// This mock is useful when running tests without the need of having a testing endpoint of the messaging system, in this case no
// need to add any logic to it
type MockMessageQueue struct {
}

func (sm *MockMessageQueue) StartRoutine() {
}

func (sm *MockMessageQueue) AddMessageToQueue(m interface{}) {
}
