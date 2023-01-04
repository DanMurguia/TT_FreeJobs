package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoPostsSeguidores es la estructura con la que devolveremos los posts */
type DevuelvoPostsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationId,omitempty"`
	Post              struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		PostImg string    `bson:"postimg" json:"postimg,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}
