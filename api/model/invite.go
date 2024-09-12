package model

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base32"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"github.com/google/uuid"
	"github.com/maatko/society/internal/server"
)

const (
	INVITE_CODE_LENGTH = 15
	INVITE_CODE_START  = 5
)

type Invite struct {
	ID        int
	Code      string
	CreatedBy *User
	UsedBy    *User
}

func GenerateInviteCode(length int) (string, error) {
	// Generate UUID
	uuidString := uuid.New().String()

	// Generate random bytes
	randomBytes := make([]byte, length/2)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Combine UUID and random bytes
	data := uuidString + hex.EncodeToString(randomBytes)

	// Hash the combined data for improved randomness and uniqueness
	hasher := sha256.New()
	_, err = io.WriteString(hasher, data)
	if err != nil {
		return "", err
	}
	hashedData := hasher.Sum(nil)

	// Encode the hashed data in base32
	hashedBase32 := base32.StdEncoding.EncodeToString(hashedData)

	// Trim or pad the result to the required length
	if len(hashedBase32) > length {
		hashedBase32 = hashedBase32[:length]
	} else if len(hashedBase32) < length {
		// Pad with characters if the length is shorter
		padding := make([]byte, length-len(hashedBase32))
		_, err := rand.Read(padding)
		if err != nil {
			return "", err
		}
		hashedBase32 += base32.StdEncoding.EncodeToString(padding)
	}

	return hashedBase32, nil
}

func NewInvite(user *User) (*Invite, error) {
	code, err := GenerateInviteCode(INVITE_CODE_LENGTH)
	if err != nil {
		return nil, err
	}

	result, err := server.DataBase().Exec("INSERT INTO invite (created_by, code) VALUES (?, ?)", user.ID, code)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return GetInviteByID(int(id))
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

func GetInviteByID(id int) (*Invite, error) {
	row := server.DataBase().QueryRow("SELECT code, created_by, used_by FROM invite WHERE id=?", id)

	var code string
	var createdBy sql.NullInt64
	var usedBy sql.NullInt64

	if err := row.Scan(&code, &createdBy, &usedBy); err != nil {
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
