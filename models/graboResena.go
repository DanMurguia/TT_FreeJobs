package models

import "time"

/*GraboReseña es el formato o estructura que tendrá nuestra reseña en la BD */
type GraboResena struct {
	UserID       string    `bson:"userid" json:"userid,omitempty"`
	UserResenaID string    `bson:"userresenaid" json:"userresenaid,omitempty"`
	Mensaje      string    `bson:"mensaje" json:"mensaje,omitempty"`
	Calificacion int       `bson:"calificacion" json:"calificacion,omitempty"`
	Fecha        time.Time `bson:"fecha" json:"fecha,omitempty"`
}
