package creacion_datos

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

)

func ArchivoTransacciones(currentTime string) {

	urlTransactions := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions"
	response, err := http.Get(urlTransactions + "?date=" + currentTime)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		raw := string(data)

		

		// example: 
		//			#000060370d268727f22222.123.167.115android(626b3de3,f16e846a,8d191cc2,3fbde956,b9f2e6d)
		//			#00006036e880�dd39959f�188.156.205.158�mac�(8a78119e)��

		
		ini_replacer := strings.NewReplacer("(", "\"[", ")", "]\"","\x00",",","\r","\r","#","\n")
		output := ini_replacer.Replace(raw)

		final_replacer := strings.NewReplacer(",,\r","\r",",,\n","\n")
		final := final_replacer.Replace(output[1:len(output)-2])
		
		/*
		processed := ""
		commaCount := 0
		for _, chr := range raw {
			actual := string(chr)
			if actual == "(" {
				processed += "\"["
			} else if actual == ")" {
				processed += "]\""
			} else if actual == "\x00" {
				processed += ","
				commaCount += 1
			} else if actual != "\r" && actual != "#" {
				processed += actual
			}
			if commaCount == 6 {
				processed = processed[:len(processed)-2]
				processed += "\n"
				commaCount = 0
			}
		}
		*/
		// Write to file
		err = ioutil.WriteFile("archivos/transacciones.csv", []byte(final), 0644)
		fmt.Printf("Archivo de transacciones generado\n")
		if err != nil {
			panic(err)
		}
	}
}