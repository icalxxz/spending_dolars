package models

import "time"

type User struct {
	UID       string    `json:"uid" firestore:"uid"`
	Email     string    `json:"email" firestore:"email"`
	Name      string    `json:"name" firestore:"name"`
	Role      string    `json:"role" firestore:"role"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
}
