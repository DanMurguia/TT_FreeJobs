package bd

import (
	"context"
	"time"

	"github.com/DanMurguia/TT_FreeJobs/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModificoRegistro permite modificar el perfil del usuario */
func ModificoPost(u models.Post, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("FreeJobs")
	col := db.Collection("post")

	registro := make(map[string]interface{})
	if len(u.Mensaje) > 0 {
		registro["mensaje"] = u.Mensaje
	}
	if len(u.PostImg) > 0 {
		registro["postimg"] = u.PostImg
	}
	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
