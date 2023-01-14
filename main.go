package main

import (
	"database/sql"
	"fmt"
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
	http.HandleFunc("/insertar", Insertar)
	http.HandleFunc("/eliminar", Eliminar)
	http.HandleFunc("/editar", Editar)
	http.HandleFunc("/actualizar", Actualizar)
	log.Println("servidor corriendo...")
	http.ListenAndServe(":8080", nil)
}

type Empleado struct {
	Id     int
	Nombre string
	Correo string
}

func Inicio(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola GusGus")
	conexionEstablecida := conexionBD()
	consultarRegistros, err := conexionEstablecida.Query("SELECT * FROM empleados")

	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}
	listaEmpleado := []Empleado{}

	for consultarRegistros.Next() {
		var id int
		var nombre, correo string
		err = consultarRegistros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo

		listaEmpleado = append(listaEmpleado, empleado)
	}

	fmt.Println(listaEmpleado)

	plantillas.ExecuteTemplate(w, "inicio", listaEmpleado)
}

func Registrar(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hola GusGus")
	plantillas.ExecuteTemplate(w, "registrar", nil)
}

func Insertar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		conexionEstablecida := conexionBD()
		insertarRegistros, err := conexionEstablecida.Prepare("INSERT INTO empleados(nombre,correo) VALUES (?,?)")

		if err != nil {
			panic(err.Error())
		}

		insertarRegistros.Exec(nombre, correo)

		http.Redirect(w, r, "/", 301)
	}
}

func Eliminar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)
	conexionEstablecida := conexionBD()
	eliminarRegistro, err := conexionEstablecida.Prepare("DELETE FROM empleados WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	eliminarRegistro.Exec(idEmpleado)

	http.Redirect(w, r, "/", 301)
}

func Editar(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)
	conexionEstablecida := conexionBD()
	consultarRegistros, err := conexionEstablecida.Query("SELECT * FROM empleados WHERE id=?", idEmpleado)
	if err != nil {
		panic(err.Error())
	}

	empleado := Empleado{}

	for consultarRegistros.Next() {
		var id int
		var nombre, correo string
		err = consultarRegistros.Scan(&id, &nombre, &correo)
		if err != nil {
			panic(err.Error())
		}
		empleado.Id = id
		empleado.Nombre = nombre
		empleado.Correo = correo
	}
	fmt.Println(empleado)

	plantillas.ExecuteTemplate(w, "editar", empleado)

}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nombre := r.FormValue("nombre")
		correo := r.FormValue("correo")
		conexionEstablecida := conexionBD()
		actualizarRegistros, err := conexionEstablecida.Prepare("UPDATE empleados SET nombre=?, correo=? WHERE id = ?")

		if err != nil {
			panic(err.Error())
		}

		actualizarRegistros.Exec(nombre, correo, id)

		http.Redirect(w, r, "/", 301)
	}
}
