package model

import "time"

type Note struct {
	NoteID      int       `json:"note_id"`
	Text        string    `json:"text"`
	ReadOnly    bool      `json:"read_only"`
	ExpiredTime time.Time `json:"expired_time"`
}
