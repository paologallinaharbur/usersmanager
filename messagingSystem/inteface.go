package messagingSystem

//MessageQueue stub of an interface of a messaging system, the API call should not stop waiting for the massage to be delivered
type MessageQueue interface {
	AddMessageToQueue(interface{})
	StartRoutine()
}
