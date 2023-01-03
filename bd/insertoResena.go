package bd

import (
	"context"
	"time"

	"github.com/DanMurguia/TT_FreeJobs/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoReseña graba una Reseña en la BD */
func InsertoResena(t models.GraboResena) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("FreeJobs")
	col := db.Collection("resena")

	registro := bson.M{
		"userid":       t.UserID,
		"userresenaid": t.UserResenaID,
		"mensaje":      t.Mensaje,
		"calificacion": t.Calificacion,
		"fecha":        t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
