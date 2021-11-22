package models

/*Relacion modelo para grabar la relacion de un usuario con otro */
type Relation struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}