package notify

type store struct {
	we_are_publishers map[string]publisher
	we_are_listeners  map[string]listen_to
}

var notify_store = store{
	we_are_publishers: map[string]publisher{},
	we_are_listeners:  map[string]listen_to{},
}
