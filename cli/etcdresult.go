package cli

// Node struct
type Node struct {
	Nodes
	CreatedIndex  int    `json:"createdIndex"`
	Key           string `json:"key"`
	ModifiedIndex int    `json:"modifiedIndex"`
	Value         string `json:"value"`
}

// Nodes struct
type Nodes struct {
	Nodes [] struct {
		Node
	} `json:"nodes"`
}
// EtcdResult output
type EtcdResult struct {
	Action string `json:"action"`
	Node Node
}
