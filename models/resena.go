package models

/*Reseña captura del Body, el mensaje que nos llega */
type Resena struct {
	UserResenaID string `bson:"userresenaid" json:"userresenaid"`
	Mensaje      string `bson:"mensaje" json:"mensaje"`
	Calificacion int    `bson:"calificacion" json:"calificacion"`
}
