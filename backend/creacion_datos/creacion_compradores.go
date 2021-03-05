package creacion_datos

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func ArchivoCompradores(currentTime string) {

	urlBuyers := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
	response, err := http.Get(urlBuyers+ "?date=" + currentTime)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		// Write to file
		processed := "{\"Buyers\":" + string(data) + "}"
		err = ioutil.WriteFile("archivos/compradores.json", []byte(processed), 0644)
		fmt.Printf("Archivo de compradores generado.\n")
		if err != nil {
			panic(err)
		}
	}
}