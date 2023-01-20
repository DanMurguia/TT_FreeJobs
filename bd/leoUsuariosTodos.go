package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/DanMurguia/TT_FreeJobs/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
LeoUsuariosTodos Lee los usuarios registrados en el sistema, si se recibe "R" en quienes

	trae solo los que se relacionan conmigo
*/
func LeoUsuariosTodos(ID string, coordenadasActual primitive.M, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("FreeJobs")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{}
	if tipo == "new" {
		query = bson.M{
			"$or": []bson.M{
				{"nombre": bson.M{"$regex": `(?i)` + search}},
				{"biografia": bson.M{"$regex": `(?i)` + search}},
				{"ubicacion": bson.M{"$regex": `(?i)` + search}},
			},
			"coordenadas": bson.M{
				"$near": bson.M{
					"$geometry":    coordenadasActual,
					"$maxDistance": 15000,
				},
			},
		}
	} else if tipo == "follow" {
		query = bson.M{
			"$or": []bson.M{
				{"nombre": bson.M{"$regex": `(?i)` + search}},
				{"biografia": bson.M{"$regex": `(?i)` + search}},
				{"ubicacion": bson.M{"$regex": `(?i)` + search}},
			},
		}
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Print(err)
		return results, false
	}
	fmt.Print("cursor del mal\n")
	fmt.Print(cur)
	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false
		fmt.Print("\nUsuario\n")
		fmt.Print(s)
		encontrado, _ = ConsultoRelacion(r)
		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.Phone = ""
			s.IsOfer = false
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""
			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		return results, false
	}
	cur.Close(ctx)
	return results, true
}
