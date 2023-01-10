package models

import "time"

/*GraboPost es el formato o estructura que tendrá nuestro Post en la BD */
type GraboPost struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	PostImg string    `bson:"postimg" json:"postimf,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
