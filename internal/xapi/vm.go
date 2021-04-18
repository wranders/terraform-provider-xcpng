package xapi

type VM struct {
	client *Client
}

// VMRecord (incomplete)
type VMRecord struct {
	UUID string
}

type VMRef string

func (vm *VM) GetAll(session SessionRef) ([]VMRef, error) {
	var recordMap []VMRef
	err := vm.client.rpc.Call(
		&recordMap,
		"VM.get_all",
		session,
	)
	return recordMap, err
}
