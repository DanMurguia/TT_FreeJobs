package bd

import (
	"context"
	"time"

	"github.com/DanMurguia/TT_FreeJobs/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoPostsSeguidores lee los posts de mis seguidores */
func LeoPostsSeguidores(ID string, pagina int) ([]models.DevuelvoPostsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("FreeJobs")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "post",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "post",
		}})
	condiciones = append(condiciones, bson.M{"$unwind": "$post"})
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"post.fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DevuelvoPostsSeguidores
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
