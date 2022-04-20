package ast

func IdentifiersToStrings(identifiers []Expression) []string {
	strings := make([]string, 0, len(identifiers))
	for _, identifier := range identifiers {
		id := identifier.(*Identifier)
		strings = append(strings, id.Name)
	}
	return strings
}
