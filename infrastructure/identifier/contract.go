package identifier

type IdentifierContract interface {
	GenerateID() (ID string)
}
