package creacion_datos

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ArchivoProductos(currentTime string) {

	urlProducts := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"
	response, err := http.Get(urlProducts+ "?date=" + currentTime)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		csv := strings.Replace(string(data), "'", ",", -1)
		// Write to file
		err = ioutil.WriteFile("archivos/productos.csv", []byte(csv), 0644)
		fmt.Printf("Archivo de productos generado.\n")
		if err != nil {
			panic(err)
		}
	}
}