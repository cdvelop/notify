package notify

func Registered(IAM publisher, L listen_to) {
	if IAM != nil {
		notify_store.we_are_publishers[IAM.MainName()] = IAM
	}

	if L != nil {
		notify_store.we_are_listeners[L.MainName()] = L
	}
}

func (s *store) PublishersList() map[string]publisher {
	return notify_store.we_are_publishers
}

func (s *store) ListenerList() map[string]listen_to {
	return notify_store.we_are_listeners
}
