// @author: Nicolas Vasquez Murcia
package main
import ("fmt"
	"log"
    "encoding/json"
	"database/sql"
	"net/http"
	"math/rand"
	"strconv"
	_"github.com/lib/pq"
    "github.com/gorilla/mux"
	)
//Conexion con CockroachDB
var dbCKDB *sql.DB

// Tipo de objeto receta
type Recipe struct{
	id int `json:"id,omitempty"`
	name string `json:"name,omitempty"`
	imgURL string `json:"imgURL,omitempty"`
	description string `json:"description,omitempty"`
	ingredients string `json:"ingredients,omitempty"`
}

var recipes []Recipe

/*
* Metodo principal para ejecutar
*/ 
func main() {
    fmt.Println("Recipes Truora")
    fmt.Println("Nicolas Vasquez Murcia")

    //Conectarse a CockroachDB en puerto 26257 en la base de datos Cook
    db, err := sql.Open("postgres", "postgresql://truora@localhost:26257/Cook?sslmode=disable")
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }
    
    dbCKDB = db
    // Crear la tabla de recetas
    if _, err := db.Exec(
        "create table if not exists Cook.Recipes(id INT PRIMARY KEY, name STRING, imgURL STRING, description STRING, ingredients STRING)"); err != nil {
        fmt.Println("FAIL")
        log.Fatal(err)
    }

    //Iniciar servicios
    router := mux.NewRouter()
    router.HandleFunc("/recipes/", readAll).Methods("GET")
	router.HandleFunc("/recipe/{id}", read).Methods("GET")
	router.HandleFunc("/search/{name}", search).Methods("GET")
	router.HandleFunc("/recipe/{id}", update).Methods("POST")
	router.HandleFunc("/recipe/", create).Methods("PUT")
    router.HandleFunc("/recipe/{id}", deleteR).Methods("DELETE")
    //Exponer servicios y escuchar en el puesto 8000
    log.Fatal(http.ListenAndServe(":8000", router))
}
/*
*Metodo para realizar PUT de una nueva receta
*/
func create(writer http.ResponseWriter, request *http.Request) {
	//param:= mux.Vars(request)
	var rec Recipe
	_ = json.NewDecoder(request.Body).Decode(&rec)
    fmt.Println(request.Body)
    fmt.Println(request)
	random:= rand.New(rand.NewSource(999))
	rec.id =random.Intn(999999)
	var idString = strconv.Itoa(rec.id)
	recipes = append(recipes, rec)
	
	// Insert two rows into the "accounts" table.
    if rec.name != ""{
        if _, err := dbCKDB.Exec(
        "INSERT INTO Cook.Recipes VALUES ("+idString+",'"+rec.name+"','"+rec.imgURL+"','"+rec.description+"','"+rec.ingredients+"')"); err != nil {
        log.Fatal(err)
        }    
    }
    json.NewEncoder(writer).Encode(recipes)
}

/*
* Metodo para realizar POST de una receta y actualizarla
*/
func update(writer http.ResponseWriter, request *http.Request) {
	param:= mux.Vars(request)
	idU := param["id"]
	var rec Recipe
	_ = json.NewDecoder(request.Body).Decode(&rec)
    fmt.Println(rec)
	recid, err := strconv.Atoi(idU)
	if err!= nil {
		log.Fatal(err)
	}
	rec.id=recid
	recipes = append(recipes, rec)

	if _, err := dbCKDB.Exec(
        "UPDATE Cook.Recipes SET name="+"'"+rec.name+"'"+", description='"+rec.description+"'"+
        ", imgurl='"+rec.imgURL+"', ingredients='"+rec.ingredients+"' where id="+idU); err != nil {
        log.Fatal(err)
    }
	json.NewEncoder(writer).Encode(recipes)
}

/*
*Metodo para hacer GET de todas las recetas
*/
func readAll(writer http.ResponseWriter, request *http.Request) {

    rows, err := dbCKDB.Query("SELECT * FROM Cook.Recipes;")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    recipes = nil
    for rows.Next() {
        var id int
        var imgURL,name, description,ingredients string
        if err := rows.Scan(&id, &name, &imgURL, &description, &ingredients); err != nil {
            log.Fatal(err)
        }
        var rec = Recipe{id: id, name:name, imgURL: imgURL, description:description, ingredients:ingredients}
        recipes=append(recipes, rec)
    }
    if recipes != nil {
        json.NewEncoder(writer).Encode(recipes)
    } else {
        
    }
}

/*
*Metodo para hacer GET de una receta por ID
*/
func read(writer http.ResponseWriter, request *http.Request) {
	param:=mux.Vars(request)
	idR:=param["id"]
    rows, err := dbCKDB.Query("SELECT* FROM Cook.Recipes WHERE id="+idR)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {

        var id int
        var imgURL,name, description,ingredients string
        if err := rows.Scan(&id, &name, &imgURL, &description, &ingredients); err != nil {
            log.Fatal(err)
        }
        var rec = Recipe{id: id, name:name, imgURL: imgURL, description:description, ingredients:ingredients}
        fmt.Println(rec)
        recipes = append(recipes, rec)
    }
    json.NewEncoder(writer).Encode(recipes)
}

/*
*Metodo para hacer GET de recetas por nombre
*/
func search(writer http.ResponseWriter, request *http.Request) {
	param:=mux.Vars(request)
	nameR:=param["name"]
    rows, err := dbCKDB.Query("SELECT * FROM Cook.Recipes WHERE name LIKE '%"+nameR+"%'")

    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    recipes=nil
    for rows.Next() {
        var id int
        var imgURL,name, description,ingredients string
        if err := rows.Scan(&id, &imgURL, &name, &description, &ingredients); err != nil {
            log.Fatal(err)
        }
        var rec = Recipe{id: id, name:name, imgURL: imgURL, description:description, ingredients:ingredients}
        recipes=append(recipes, rec)
        fmt.Println(rec)
    }
    json.NewEncoder(writer).Encode(recipes)
}

/*
* Metodo para hacer DELETE de una receta por ID
*/
func deleteR(writer http.ResponseWriter, request *http.Request) {
	param:=mux.Vars(request)
	idD:= param["id"]


	if _, err := dbCKDB.Exec(
        "Delete * from Cook.Recipes where id ="+idD); err != nil {
        log.Fatal(err)
    }
    json.NewEncoder(writer).Encode(recipes)
}