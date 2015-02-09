package user

import (
	"fmt"
	"time"

	"github.com/bearded-web/bearded/pkg/pagination"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id"`
	Nickname string        `json:"nickname"`
	Email    string        `json:"email"`
	Password string        `json:"-"` // password hash in passlib format: $hashAlgo[$values]$hexdigest_hash$
	Avatar   string        `json:"avatar,omitempty"`

	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
}

func (u *User) String() string {
	return fmt.Sprintf("%x %s", u.Id, u.Email)
}

// Id in hex
func (u *User) IdStr() string {
	return fmt.Sprintf("%x", string(u.Id))
}

type UserList struct {
	pagination.Meta `json:",inline"`
	Results         []*User `json:"results"`
}
