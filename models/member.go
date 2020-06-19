package models

import (
	"time"
)

type Member struct {
	Mid      uint
	Groupid  int
	Mobile   string
	Account  string
	Password string
	Salt     string
	Ctime    time.Time
}

func (member *Member) First(account string) *Member {

	orm.Where("account=?", account).First(member).RecordNotFound()
	return member
}
