package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de usuario de la base de MongoDB */
type Usuario struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre            string             `bson:"nombre" json:"nombre,omitempty"`
	Apellidos         string             `bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento   time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email             string             `bson:"email" json:"email"`
	Password          string             `bson:"password" json:"password,omitempty"`
	Phone             string             `bson:"phone" json:"phone,omitempty"`
	Avatar            string             `bson:"avatar" json:"avatar,omitempty"`
	Banner            string             `bson:"banner" json:"banner,omitempty"`
	Biografia         string             `bson:"biografia" json:"biografia,omitempty"`
	Ubicacion         string             `bson:"ubicacion" json:"ubicacion,omitempty"`
	Coordenadas       primitive.M        `bson:"coordenadas" json:"coordenadas,omitempty"`
	CoordenadasActual primitive.M        `bson:"coordenadasActual" json:"coordenadasActual,omitempty"`
	IsOfer            bool               `bson:"isOfer" json:"isOfer,omitempty"`
}
