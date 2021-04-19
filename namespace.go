package dumps

// Namespace single namespace
type Namespace struct {
	ID        int    `json:"id"`
	Case      string `json:"case"`
	Canonical string `json:"canonical"`
	Local     string `json:"*"`
}

type namespacesResponse struct {
	Query struct {
		Namespaces map[int]Namespace
	}
}
