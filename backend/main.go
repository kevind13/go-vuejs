package main

import (

	"fmt"
	//"io/ioutil"
	"net/http"
	"strconv"
	"time"
	//"os"
	//"encoding/csv"
	//"bytes"

	"context"
    //"encoding/json"
	

    "log"


    "github.com/dgraph-io/dgo/v2"
    "github.com/dgraph-io/dgo/v2/protos/api"
    "google.golang.org/grpc"

	"github.com/kevind13/backend/creacion_datos"
	"github.com/kevind13/backend/info_dgraph"


	"github.com/go-chi/chi"
)

var dicCompradores map[string]string
var dicTransacciones map[string]string
var dicProductos map[string]string
var currentTime string 


func init() { 
	dicCompradores,dicTransacciones,dicProductos = InitApp("")
}

func InitApp(currentTime string) (map[string]string,map[string]string,map[string]string){
	if currentTime == "" {
		currentTime = strconv.FormatInt(time.Now().Unix(), 10)
	  }
	ctx := context.TODO()
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
    if err != nil {
		log.Fatal("failed to dial ", err)
    }
    defer conn.Close()

    dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	// Borrar todo para hacer las pruebas
	op := &api.Operation{DropAll: true}
	if err := dgraphClient.Alter(ctx, op); err != nil {
		log.Fatal(err)
	}

	
	fmt.Printf("Dia actual en unix format: %v \n",currentTime)

    // Generacion del json de compradores
	fmt.Println("Compradores")
	creacion_datos.ArchivoCompradores(currentTime)
	
	// Generacion de csv de productos
	fmt.Println("Productos")
	creacion_datos.ArchivoProductos(currentTime)
	
	// Generacion de csv de transacciones
	fmt.Println("Transacciones")
	creacion_datos.ArchivoTransacciones(currentTime)


	dicCompradores, dicTransacciones,dicProductos := info_dgraph.LlenarDGraph()
	fmt.Printf("%v%v%v\n", dicTransacciones[""],dicCompradores["ded2ec99"],dicProductos[""])
    return dicCompradores, dicTransacciones,dicProductos

}

func main() {

	ctx := context.TODO()
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
    if err != nil {
		log.Fatal("failed to dial ", err)
    }
    defer conn.Close()

    dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))



	r := chi.NewRouter()

	r.Get("/load/{day}", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		

		t, _ := time.Parse("2006-01-02",chi.URLParam(r, "day"))
		currentTime := strconv.FormatInt(t.Unix(), 10)
		
		fmt.Printf("Dia en unix format: %v \n",currentTime)

    	
		dicCompradores, dicTransacciones,dicProductos = InitApp(currentTime)

		fmt.Printf("%v%v%v\n", dicTransacciones[""],dicCompradores["ded2ec99"],dicProductos[""])


		response := "Datos del dia "+chi.URLParam(r, "day")+"cargados en DGraph"
		w.Write([]byte(response))
		
	})
	


	r.Get("/allbuyers", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		q := `query all ($a: string){
			all(func: type($a)) {
				name
				idcomprador
				age
				uid  
		  	}
		}`

		txn3 := dgraphClient.NewTxn()
		defer txn3.Commit(ctx)
		
	
		res, err := txn3.QueryWithVars(ctx, q, map[string]string{"$a": "Comprador"})
		if err != nil {
			fmt.Println("Error con query:", err)
		}
		response := res.Json
		//fmt.Printf("%s\n", res.Json)
		w.Write([]byte(response))
	})


	r.Get("/buyer/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		
		q := `query all ($a: string){
			all(func: uid($a)) {
				name
				idcomprador
				age
				uid
				hace{
					ip
					device
					uid
					tiene{
						name
					}
				}
		  	}
		}`

		txn3 := dgraphClient.NewTxn()
		defer txn3.Commit(ctx)
		
		fmt.Println(dicCompradores[chi.URLParam(r, "id")])
		res, err := txn3.QueryWithVars(ctx, q, map[string]string{"$a": dicCompradores[chi.URLParam(r, "id")]})
		if err != nil {
			fmt.Println("Error con query:", err)
		}
		response := res.Json
		//fmt.Printf("%s\n", res.Json)
		w.Write([]byte(response))

	})


	r.Get("/allproducts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		q := `query all ($a: string){
			all(func: type($a)) {
				name
				price  
		  	}
		}`

		txn3 := dgraphClient.NewTxn()
		defer txn3.Commit(ctx)
		
	
		res, err := txn3.QueryWithVars(ctx, q, map[string]string{"$a": "Producto"})
		if err != nil {
			fmt.Println("Error con query:", err)
		}
		response := res.Json
		//fmt.Printf("%s\n", res.Json)
		w.Write([]byte(response))
	})

	r.Get("/sameip/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		q := `query all ($a: string){
			var(func: uid($a)) { 
				INI as name
				age
				hace{
					ESTE as ip 
				}
			  }
			all(func: type(Comprador)) @filter(NOT eq(name,val(INI)))@cascade{
				name
				age
				idcomprador
				uid
				  hace @filter(eq(ip,val(ESTE))){
				  ip
				  tiene{
					  name
					  price
				  }
				}
			}
				
		}`

		txn3 := dgraphClient.NewTxn()
		defer txn3.Commit(ctx)
		
	
		res, err := txn3.QueryWithVars(ctx, q, map[string]string{"$a": dicCompradores[chi.URLParam(r, "id")]})
		if err != nil {
			fmt.Println("Error con query:", err)
		}
		response := res.Json
		//fmt.Printf("%s\n", res.Json)
		w.Write([]byte(response))
	})


	
	  
	/*
	q := `query all ($a: string){
			all(func: uid($a)) {
				name
				age
				uid  
				hace{
					buyerid
					uid
					ip
					device
			  	}
		  	}
		}`
	*/

	fmt.Printf("%v%v%v\n", dicTransacciones["7f62c691"],dicCompradores[""],dicProductos[""])
	// Server
	err333 := http.ListenAndServe(":3000", r)
	if err333 != nil {
		fmt.Println("ListenAndServe:", err333)
	}

	

}