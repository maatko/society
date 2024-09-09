package model

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/maatko/society/internal/server"
)

type Invite struct {
	ID        int
	Code      string
	CreatedBy *User
	UsedBy    *User
}

func GetInviteByCode(code string) (*Invite, error) {
	row := server.DataBase().QueryRow("SELECT id, created_by, used_by FROM invite WHERE code=?", code)

	var id int
	var createdBy sql.NullInt64
	var usedBy sql.NullInt64

	if err := row.Scan(&id, &createdBy, &usedBy); err != nil {
		log.Println(err)
		return nil, err
	}

	var from *User = nil
	var to *User = nil

	if createdBy.Valid {
		user, err := GetUserByID(int(createdBy.Int64))
		if err == nil {
			from = user
		}
	}

	if createdBy.Valid {
		user, err := GetUserByID(int(usedBy.Int64))
		if err == nil {
			to = user
		}
	}

	return &Invite{
		ID:        id,
		Code:      code,
		CreatedBy: from,
		UsedBy:    to,
	}, nil
}

func (inv *Invite) Update() error {
	var createdBy string = "NULL"
	if inv.CreatedBy != nil {
		createdBy = fmt.Sprintf("%d", inv.CreatedBy.ID)
	}

	var usedBy string = "NULL"
	if inv.UsedBy != nil {
		usedBy = fmt.Sprintf("%d", inv.UsedBy.ID)
	}

	_, err := server.DataBase().Exec("UPDATE invite SET code=?, created_by=?, used_by=? WHERE id=?", inv.Code, createdBy, usedBy, inv.ID)
	if err != nil {
		return err
	}
	return nil
}
