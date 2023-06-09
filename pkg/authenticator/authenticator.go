package authenticator

import (
	"github.com/maciejgaleja/gosimple/pkg/keyvalue"
)

type HashedUsername string
type HashedPassword string

type Entry struct {
	Username HashedUsername
	Password HashedPassword
}

type Authenticator struct {
	db keyvalue.Store
}

func NewAuthenticator(db keyvalue.Store) (*Authenticator, error) {
	a := Authenticator{db: db}
	return &a, nil
}

func (a *Authenticator) Register(e Entry) error {
	return a.db.Set(keyvalue.Key(e.Username), e)
}

func (a *Authenticator) Verify(e Entry) (bool, error) {
	var re Entry
	err := a.db.Get(keyvalue.Key(e.Username), &re)
	if err != nil {
		return false, nil
	}
	return (re.Username == e.Username && re.Password == e.Password), nil
}
