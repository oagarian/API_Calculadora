package main

import (
	"log"
	"net/http"
	"strconv"
	"fmt"
	"mux"
)

func calculo(w http.ResponseWriter, req *http.Request ){
	vars := mux.Vars(req)
	calc := vars["param"]
	n1 := vars["n1"]
	n2 := vars["n2"]

	number_one, error1 := strconv.ParseFloat(n1, 64)
	if error1 != nil {
		log.Fatal(error1)
	}
	number_two, error2 := strconv.ParseFloat(n2, 64)
	if error2 != nil {
		log.Fatal(error2)
	}

	var result float64
	switch calc {
		case "soma": 
			result = number_one + number_two
			
		case "subtracao": 
			result = number_one - number_two
		case "divisao": 
			result = number_one / number_two
		case "multiplicacao": 
			result = number_one * number_two
	}	
	
	var float_to_int int = int(result)
	int_to_string := strconv.Itoa(float_to_int)
	fmt.Fprintf(w, "%s", int_to_string)
	

}


func main(){
	rout := mux.NewRouter()
	rout.HandleFunc("/{param}/{n1}/{n2}/", calculo).Methods("GET")
	http.ListenAndServe(":8080", rout)
}
