package main

import (
	"log"
	"fmt"
	//"strconv"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Device1 struct {
	gorm.Model
	nombre string
	kind  uint  //`gorm:"default:18"`
  }
     
 
func main() {
	db, err := gorm.Open("postgres", "host=hansken.db.elephantsql.com port=5432 user=vtwlajng dbname=vtwlajng password=	Sty59HjeuNLpFjjhRA5HNok1gHc58lVs")

	if err != nil {
		log.Print(err)

	}
	 

	db.CreateTable(&Device1{nombre:"" , kind: 99})
	db.AutoMigrate(Device1{})
	//var juan = Admin5{Name:"juan" , Age: 45}db.Create(&admin5)
	//db.NewRecord(juan)
	log.Print("Conectado!")

	user := r.Context().Value("user").(models.User)

			err := json.NewDecoder(r.Body).Decode(&ente)
			if err != nil {
				Message = fmt.Sprintf("Error al leer los datos de los entes a registrar: %s", err)
				Code = http.StatusBadRequest
				commons.DisplayMessage()
				return
			}

		

			db = configuration.GetConnection()
			defer db.Close()

			err = db.Create(&user).Error
			if err != nil {
				Message = fmt.Sprintf("Error al  crear el registro: %s", err)
				Code = http.StatusBadRequest
				commons.DisplayMessage(w, m)
				return
			}

			m.Message = "Ente creado con éxito"
			m.Code = http.StatusOK
			commons.DisplayMessage(w, m)

	defer db.Close()
}