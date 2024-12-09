package main

import (
	"time"
	"github.com/google/uuid"
	"github.com/danielhallinan88/rssagg/internal/database"
)

type User struct {
	ID		uuid.UUID `json:"id"`
	CreatedAt	time.Time `json:"create_at"`
	UpdatedAt	time.Time `json:"update_at"`
	Name		string `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:		dbUser.Name,
	}
}
