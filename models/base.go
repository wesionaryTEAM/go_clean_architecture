package models

import "clean-architecture/lib"

type Base struct {
	ID lib.BinaryUUID `json:"id"`
}
