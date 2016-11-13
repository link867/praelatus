package models

// Label is a label used on tickets
type Label struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (l *Label) String() string {
	return jsonString(l)
}
