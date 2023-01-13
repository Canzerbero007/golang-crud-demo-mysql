package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	Usuario := "root"
	Contrasenia := "123456"
	Nombre := "demogo"

	conexion, err := sql.Open(Driver, Usuario+":"+Contrasenia+"@tcp(127.0.0.1)/"+Nombre)
	if err != nil {
		panic(err.Error())

	}
	return conexion
}

var plantillas = template.Must(template.ParseGlob("c:/go/src/sistema/plantillas/*"))

func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/registrar", Registrar)
	log.Println("servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola GusGus")
	conexionEstablecida := conexionBD()
	insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES ('LULU','LULU@GMAIL.COM')")

	if err != nil {
		panic(err.Error())
	}

	insertarRegistros.Exec()
	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func Registrar(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola GusGus")
	plantillas.ExecuteTemplate(w, "registrar", nil)
}
