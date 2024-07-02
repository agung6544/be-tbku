package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ayam struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" example:"123456789"`
	Jenis          string             `bson:"jenis,omitempty" json:"jenis,omitempty" example:"Ayam Saigon"`
	Umur		   string             `bson:"umur,omitempty" json:"umur,omitempty" example:"123456789"`
	Bobot          string             `bson:"bobot,omitempty" json:"bobot,omitempty" example:"123456789"`
	Tinggi         string             `bson:"tinggi,omitempty" json:"tinggi,omitempty" example:"123456789"`
	Jenis_Kelamin  string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty" example:"Jantan"`
	Harga		   string             `bson:"harga,omitempty" json:"harga,omitempty" example:"12.000"`
}

type Order struct {
	ID           	  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Ayam 	          Ayam               `bson:"ayam,omitempty" json:"ayam,omitempty"`
	Nama_Pemesan 	  string 			 `bson:"nama_pemesan,omitempty" json:"nama_pemesan,omitempty"`
	Alamat 			  string 			 `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Tanggal_Pemesanan primitive.DateTime `bson:"tanggal_pemesanan,omitempty" json:"tanggal_pemesanan,omitempty" swaggertype:"string"`
}

type ReqAyam struct {
	Jenis          string             `bson:"jenis,omitempty" json:"jenis,omitempty" example:"Ayam Saigon"`
	Umur		   string             `bson:"umur,omitempty" json:"umur,omitempty" example:"123456789"`
	Bobot          string             `bson:"bobot,omitempty" json:"bobot,omitempty" example:"123456789"`
	Tinggi         string             `bson:"tinggi,omitempty" json:"tinggi,omitempty" example:"123456789"`
	Jenis_Kelamin  string             `bson:"jenis_kelamin,omitempty" json:"jenis_kelamin,omitempty" example:"Jantan"`
	Harga		   string             `bson:"harga,omitempty" json:"harga,omitempty" example:"12.000"`
}

type ReqOrder struct {
	Ayam 	          Ayam               `bson:"ayam,omitempty" json:"ayam,omitempty"`
	Nama_Pemesan 	  string 			 `bson:"nama_pemesan,omitempty" json:"nama_pemesan,omitempty" example:"Dimas"`
	Alamat 			  string 			 `bson:"alamat,omitempty" json:"alamat,omitempty" example:"Bandung"`
	Tanggal_Pemesanan primitive.DateTime `bson:"tanggal_pemesanan,omitempty" json:"tanggal_pemesanan,omitempty" swaggertype:"string" example:"2024-12-03"`
}