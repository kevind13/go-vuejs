package info_dgraph

import (
	"fmt"
	"io/ioutil"
	//"net/http"
	"strconv"
	"os"
	"encoding/csv"
	"bytes"
	"strings"


	"context"
    "encoding/json"
	

    "log"


    "github.com/dgraph-io/dgo/v2"
    "github.com/dgraph-io/dgo/v2/protos/api"
    "google.golang.org/grpc"

)

type Comprador struct {
	Uid   string   `json:"uid,omitempty"`
	IDComprador    string   `json:"idcomprador,omitempty"`
	Name  string   `json:"name,omitempty"`
	Age   string      `json:"age,omitempty"`
	Hace  []Transaccion `json:"hace,omitempty"`
	DType []string `json:"dgraph.type,omitempty"`
}

type Producto struct {
	Uid   	string   `json:"uid,omitempty"`
	Name	string   `json:"name,omitempty"`
	Price   string    `json:"price,omitempty"`
	DType   []string `json:"dgraph.type,omitempty"`
}

type Transaccion struct {
	Uid      string   `json:"uid,omitempty"`
	BuyerID  string   `json:"buyerid,omitempty"`
	IP       string      `json:"ip,omitempty"`
	Device   string      `json:"device,omitempty"`
	Tiene    []Producto `json:"tiene"`
	DType   []string `json:"dgraph.type,omitempty"`
}



// Estructuras para leer los archivos JSON y CSV.

type Compradores struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	ID   string `json:"id"`
}

type Productos struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	ID    string `json:"id"`
}

type Transacciones struct {
	ID      string `json:"id"`
	BuyerID string `json:"buyerid"`
	IP      string `json:"ip"`
	Device  string `json:"device"`
}

type Buyers struct {
	Buyers []Compradores `json:"buyers"`
}

type ProductosList struct {
	ProductosList []ProductosList `json:"productoslist"`
}

type TransaccionesList struct {
	TransaccionesList []Transacciones `json:"transaccioneslist"`
}



func LlenarDGraph() (map[string]string, map[string]string,map[string]string){

	dicCompradores 		:= make(map[string]string)
	dicTransacciones  	:= make(map[string]string)
	dicProductos 		:= make(map[string]string)

	ctx := context.TODO()
	conn, err := grpc.Dial("localhost:9080", grpc.WithInsecure())
    if err != nil {
		log.Fatal("failed to dial ", err)
    }
    defer conn.Close()

    dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))


	// Se agrega los compradores a la base de datos
	txn1 := dgraphClient.NewTxn()

	jsonFile, err := os.Open("archivos/compradores.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var buyers Buyers
	json.Unmarshal(byteValue, &buyers)

	for i := 0; i < len(buyers.Buyers); i++ {
		id := "_:"+buyers.Buyers[i].ID
		name := buyers.Buyers[i].Name
		age := strconv.FormatInt(int64(buyers.Buyers[i].Age), 10)
		//data := "'" + id + "'" + "," + "'" + name + "'" + "," + age


		p :=  Comprador{
			Uid: id,
			IDComprador: buyers.Buyers[i].ID,
			Name:  name,
			Age:  age,
			DType: []string{"Comprador"},
		}
		pb, err := json.Marshal(p)
		if err != nil {
			log.Fatal(err)
		}	
 
			
		mu := &api.Mutation{
			SetJson: pb,
		}

		res, err := txn1.Mutate(ctx, mu)
    	if err != nil {
			log.Fatal("failed to mutate ", err)
		}
		dicCompradores[buyers.Buyers[i].ID] = res.Uids[buyers.Buyers[i].ID]
		
	}
	fmt.Printf("Compradores agregados\n")

	txn1.Commit(ctx)


	
    

	// Se agrega los productos a la base de datos
	csvProductos, err := os.Open("archivos/productos.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Se abrio el csv de productos.")
	defer csvProductos.Close()
	
    csvProductosLines, err := csv.NewReader(csvProductos).ReadAll()
    if err != nil {
        fmt.Println(err)
    }    	

	
	for i, line := range csvProductosLines {
		//data := "'" + line[0] + "'" + "," + "'" + line[1] + "'" + "," + "'" + line[2] + "'" + "," + "'" + line[3] + "'"
		//fmt.Printf("Transaccion %v: %v \n",i,data)
		produc :=  Producto{
			Uid: "_:"+line[0],
			Name:  line[1],
			Price:  line[2],
			DType: []string{"Producto"},
		}

		producb, err := json.Marshal(produc)
		if err != nil {
			log.Fatal(err)
		}	
		
		
		
		muTrans := &api.Mutation{
			SetJson: producb,
		}

		txn2 := dgraphClient.NewTxn()

		res2, err := txn2.Mutate(ctx, muTrans)
    	if err != nil {
			log.Fatal("failed to mutate ", err)
		}

		dicProductos[line[0]] = res2.Uids[line[0]]

		if i%1000 == 0{
			fmt.Printf("%v Productos agregados.\n",i)
			//fmt.Println("Breaking out of loop")
            //break // break here
		}
		txn2.Commit(ctx)
	}



	// Se agrega las transacciones a la base de datos
	csvTransacciones, err := os.Open("archivos/transacciones.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvTransacciones.Close()
	
    csvTransaccionesLines, err := csv.NewReader(csvTransacciones).ReadAll()
    if err != nil {
        fmt.Println(err)
    }    	

	
	for i, line := range csvTransaccionesLines {
		//data := "'" + line[0] + "'" + "," + "'" + line[1] + "'" + "," + "'" + line[2] + "'" + "," + "'" + line[3] + "'"
		//fmt.Printf("Transaccion %v: %v \n",i,data)
		trans :=  Transaccion{
			Uid: "_:"+line[0],
			BuyerID:  line[1],
			IP:  line[2],
			Device: line[3],
			DType: []string{"Transaccion"},
		}

		transb, err := json.Marshal(trans)
		if err != nil {
			log.Fatal(err)
		}	
		
		
		
		muTrans := &api.Mutation{
			SetJson: transb,
		}

		txn2 := dgraphClient.NewTxn()

		res2, err := txn2.Mutate(ctx, muTrans)
    	if err != nil {
			log.Fatal("failed to mutate ", err)
		}

		dicTransacciones[line[0]] = res2.Uids[line[0]]

		var rdfs bytes.Buffer

		arrayProductsString := line[4]
		arrayProducts := strings.Split(arrayProductsString[1:len(arrayProductsString)-1], ",")
		for _, produ := range arrayProducts{
			//fmt.Printf("%v\n",arrayProductsString)
			//fmt.Println(dicProductos[produ])
			fmt.Fprintf(&rdfs, "<%v> <tiene> <%v> .\n", dicTransacciones[line[0]], dicProductos[produ])
			muRelations := &api.Mutation{
				SetNquads: rdfs.Bytes(),
			}
	
			_, errr := txn2.Mutate(ctx, muRelations)
			if errr != nil {
				log.Fatal("failed to mutate ", errr)
			}
		}

		
		fmt.Fprintf(&rdfs, "<%v> <hace> <%v> .\n", dicCompradores[line[1]], dicTransacciones[line[0]])

		muRelations := &api.Mutation{
			SetNquads: rdfs.Bytes(),
		}
		rdfs.Reset()

		_, errr := txn2.Mutate(ctx, muRelations)
    	if errr != nil {
			log.Fatal("failed to mutate ", errr)
		}

		if i%1000 == 0{
			fmt.Printf("%v Transacciones agregadas.\n",i)
			//fmt.Println("Breaking out of loop")
            //break // break here
		}
		txn2.Commit(ctx)
	}
	
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Transacciones agregadas...")
	}
	



	return dicCompradores, dicTransacciones, dicProductos
}