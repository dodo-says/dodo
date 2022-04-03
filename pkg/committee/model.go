package committee

type Committee struct {
	// The required unique name for this committee.
	Name string
	// The optional description for this committee.
	Description string
}

type Member struct {
	// The required unique name for this member.
	Name string
	// The optional description for this member.
	Description string
	// The required public key for this member.
	PublicKey []byte
}
