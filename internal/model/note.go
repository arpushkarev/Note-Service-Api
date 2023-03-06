package model

import "database/sql"

// NoteInfo structure
type NoteInfo struct {
	Title  string `db:"title"`
	Text   string `db:"text"`
	Author string `db:"author"`
}

// UpdateNoteInfo structure
type UpdateNoteInfo struct {
	Title  sql.NullString `db:"title"`
	Text   sql.NullString `db:"text"`
	Author sql.NullString `db:"author"`
}

// Note structure
type Note struct {
	ID   int64     `db:"id"`
	Info *NoteInfo `db:""`
}
