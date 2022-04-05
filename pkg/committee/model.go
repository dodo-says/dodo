package committee

type Committee struct {
	// The required unique name for this committee.
	Name string
	// The optional description for this committee.
	Description string
}

func NewCommittee(name string, description string) *Committee {
	return &Committee{Name: name, Description: description}
}

type Member struct {
	// The required unique name for this member.
	Name string
	// The required name of which committee belongs to.
	CommitteeName string
	// The optional description for this member.
	Description string
	// The required public key for this member.
	PublicKey []byte
}

func NewMember(name string, committeeName string, description string, publicKey []byte) *Member {
	return &Member{Name: name, CommitteeName: committeeName, Description: description, PublicKey: publicKey}
}
