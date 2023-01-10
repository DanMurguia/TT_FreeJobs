package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/DanMurguia/TT_FreeJobs/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertoPost graba el Post en la BD */
func InsertoPost(t models.GraboPost) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("FreeJobs")
	col := db.Collection("post")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"postimg": t.postimg,
		"fecha":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
