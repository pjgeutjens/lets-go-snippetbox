package forms

type errors map[string][]string

// Add function adds error message for a given field to the map
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get function retrieves the first error for a given field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
