package main

import (
  "github.com/gofiber/fiber/v2"
  "strconv"
  "log"
  "fmt"
  "encoding/json"
)

//Handler inicial informativo
func handler(c *fiber.Ctx) error {
  fmt.Fprint(c, "API Canculadora \n Use: \n /soma/{n1}/{n2} \n /subtracao/{n1}/{n2} \n /divisao/{n1}/{n2} \n /multiplicacao/{n1}/{n2}")
  return nil
}

//Função de soma
func soma(c *fiber.Ctx) error {
  number_one := c.Params("n1")
  number_two := c.Params("n2")

  n1, err := strconv.Atoi(number_one)
  n2, err := strconv.Atoi(number_two)

  if err != nil {
    log.Fatal(err)
  }

  result := n1 + n2
  result_string := strconv.Itoa(result)
  json.NewEncoder(c).Encode(map[string]string{
    "result": fmt.Sprintf("%s", result_string),
    })  
  return nil
}

//Função de subtração
func subtracao(c *fiber.Ctx) error {
  number_one := c.Params("n1")
  number_two := c.Params("n2")

  n1, err := strconv.Atoi(number_one)
  n2, err := strconv.Atoi(number_two)

  if err != nil {
    log.Fatal(err)
  }

  result := n1 - n2
  result_string := strconv.Itoa(result)
  json.NewEncoder(c).Encode(map[string]string{
    "result": fmt.Sprintf("%s", result_string),
    }) 
  return nil
}

//Função de divisão
func divisao(c *fiber.Ctx) error {
  number_one := c.Params("n1")
  number_two := c.Params("n2")

  n1, err := strconv.Atoi(number_one)
  n2, err := strconv.Atoi(number_two)

  if err != nil {
    log.Fatal(err)
  }
  if n2 != 0 {
    result := n1 / n2
    result_string := strconv.Itoa(result)
    json.NewEncoder(c).Encode(map[string]string{
    "result": fmt.Sprintf("%s", result_string),
    }) 
    
  } else {
    fmt.Fprintf(c, "Infinito")
  }
  return nil
}


//Função de multiplicação
func multiplicacao(c *fiber.Ctx) error {
  number_one := c.Params("n1")
  number_two := c.Params("n2")

  n1, err := strconv.Atoi(number_one)
  n2, err := strconv.Atoi(number_two)

  if err != nil {
    log.Fatal(err)
  }

  result := n1 * n2
  result_string := strconv.Itoa(result)
  json.NewEncoder(c).Encode(map[string]string{
    "result": fmt.Sprintf("%s", result_string),
    }) 
  return nil
}

//Função principal
func main() {
    app := fiber.New()
    app.Get("/", handler)
    app.Get("/soma/:n1/:n2", soma)  
    app.Get("/subtracao/:n1/:n2", subtracao) 
    app.Get("/divisao/:n1/:n2", divisao) 
    app.Get("/multiplicacao/:n1/:n2", multiplicacao) 
    app.Listen(":8080")
}
