package note_v1

import desc "github.com/arpushkarev/note-service-api/pkg/note_v1"

type Note struct {
	desc.UnimplementedNoteV1Server
}

func NewNote() *Note {
	return &Note{}
}

func GotNote() *Note {
	return &Note{}
}

func GotList() *Note {
	return &Note{}
}
