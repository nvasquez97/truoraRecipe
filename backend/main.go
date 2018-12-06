// @author: Nicolas Vasquez Murcia
package main
import ("fmt"
    "time"
	"log"
    "encoding/json"
	"database/sql"
	"net/http"
	"math/rand"
	"strconv"
    "github.com/gorilla/handlers"
	_"github.com/lib/pq"
    "github.com/gorilla/mux"
	)
//Conexion con CockroachDB
var dbCKDB *sql.DB

// Tipo de objeto receta
type Recipe struct{
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	ImgURL string `json:"imgURL,omitempty"`
	Description string `json:"description,omitempty"`
	Ingredients string `json:"ingredients,omitempty"`
    Instructions string `json:"instructions,omitempty"`
}

var recipes []Recipe

/*
* Metodo principal para ejecutar
*/ 
func main() {
    fmt.Println("Recipes Truora")
    fmt.Println("Nicolas Vasquez Murcia")

    //Conectarse a CockroachDB en puerto 26257 en la base de datos Cook con el usuario truora
    db, err := sql.Open("postgres", "postgresql://truora@localhost:26257/Cook?sslmode=disable")
    if err != nil {
        log.Fatal("Error connecting to the database: ", err)
    }
    
    dbCKDB = db
    // Crear la tabla de recetas
    if _, err := db.Exec(
        "create table if not exists Cook.Recipes(id INT PRIMARY KEY, name STRING, imgURL STRING, description STRING, ingredients STRING, instructions STRING)"); err != nil {
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
    //Exponer servicios y escuchar en el puesto 8000 y permitir CORS para localhost:9000 (Frontend Server)
    log.Fatal(http.ListenAndServe(":8000", 
        handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}),
        handlers.AllowedHeaders([]string{"Content-Type"}),
        handlers.AllowedOrigins([]string{"http://localhost:9000"}))(router)))
}
/*
*Metodo para realizar PUT de una nueva receta
*/
func create(writer http.ResponseWriter, request *http.Request) {
	var rec Recipe
	_ = json.NewDecoder(request.Body).Decode(&rec)
    fmt.Println(request.Body)
    fmt.Println(request)
	random:= rand.New(rand.NewSource(time.Now().UnixNano()))
	rec.Id =random.Intn(999999)
	var idString = strconv.Itoa(rec.Id)
	recipes = append(recipes, rec)
	
	// Insert two rows into the "accounts" table.
    if rec.Name != ""{
        if _, err := dbCKDB.Exec(
        "INSERT INTO Cook.Recipes VALUES ("+idString+",'"+rec.Name+"','"+rec.ImgURL+"','"+rec.Description+"','"+rec.Ingredients+"','"+rec.Instructions+"')"); err != nil {
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
	recid, err := strconv.Atoi(idU)
	if err!= nil {
		log.Fatal(err)
	}
	rec.Id=recid
	recipes = append(recipes, rec)
    if rec.Name!= ""{
        if _, err := dbCKDB.Exec(
        "UPDATE Cook.Recipes SET name="+"'"+rec.Name+"'"+", description='"+rec.Description+"'"+
        ", imgurl='"+rec.ImgURL+"', ingredients='"+rec.Ingredients+"', instructions='"+rec.Instructions+"' where id="+idU); err != nil {
        log.Fatal(err)
        }
        json.NewEncoder(writer).Encode(rec)    
    } else {
        rec = Recipe{Id: 1, Name:"Hello", ImgURL: "URL", Description:"description", Ingredients:"ingredients", Instructions:"instructions"}
        //json.NewEncoder(writer).Encode(rec)
        writer.WriteHeader(http.StatusInternalServerError)
        writer.Write([]byte("Hubo un error actualizando la receta porque los datos están vacíos"))
    }
	
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
        var imgURL,name, description,ingredients, instructions string
        if err := rows.Scan(&id, &name, &imgURL, &description, &ingredients, &instructions); err != nil {
            log.Fatal(err)
        }
        var rec = Recipe{Id: id, Name:name, ImgURL: imgURL, Description:description, Ingredients:ingredients, Instructions:instructions}
        recipes=append(recipes, rec)
    }
    if recipes != nil {
        json.NewEncoder(writer).Encode(recipes)
    } else {
        writer.Write([]byte("No hay recetas en el momento"))
    }
}

/*
*Metodo para hacer GET de una receta por Id
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
        var imgURL,name, description,ingredients, instructions string
        if err := rows.Scan(&id, &name, &imgURL, &description, &ingredients, &instructions); err != nil {
            log.Fatal(err)
        }
        var rec = Recipe{Id: id, Name:name, ImgURL: imgURL, Description:description, Ingredients:ingredients, Instructions:instructions}
        json.NewEncoder(writer).Encode(rec)
    }

    //json.NewEncoder(writer).Encode(rec)
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
        var imgURL,name, description,ingredients, instructions string
        if err := rows.Scan(&id, &imgURL, &name, &description, &ingredients, &instructions); err != nil {
            log.Fatal(err)
        }
        var rec = Recipe{Id: id, Name:name, ImgURL: imgURL, Description:description, Ingredients:ingredients, Instructions:instructions}
        recipes=append(recipes, rec)
    }
    json.NewEncoder(writer).Encode(recipes)
}

/*
* Metodo para hacer DELETE de una receta por Id
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