package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoResenas es la estructura con la que devolveremos las reseñas */
type DevuelvoResenas struct {
	ID           primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID       string             `bson:"userid" json:"userId,omitempty"`
	UserResenaID string             `bson:"userresenaid" json:"userresenaid,omitempty"`
	Mensaje      string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha        time.Time          `bson:"fecha" json:"fecha,omitempty"`
	Calificacion int                `bson:"calificacion" json:"calificacion,omitempty"`
}
