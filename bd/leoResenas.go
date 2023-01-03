package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/DanMurguia/TT_FreeJobs/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoResena lee las reseñas de un perfil */
func LeoResena(ID string, pagina int64) ([]*models.DevuelvoResenas, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("FreeJobs")
	col := db.Collection("resena")

	var resultados []*models.DevuelvoResenas

	condicion := bson.M{
		"userresenaid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultados, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.DevuelvoResenas
		err := cursor.Decode(&registro)
		fmt.Println(err)

		if err != nil {
			return resultados, false
		}
		resultados = append(resultados, &registro)
	}
	return resultados, true
}
