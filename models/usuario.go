package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Usuario Define la estructura de un usuario
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempty"`
	Apellido        string             `bson:"apelligo" json:"apellido,omitempty"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	SitioWeb        string             `bson:"sitioweb" json:"sitioweb,omitempty"`
}
