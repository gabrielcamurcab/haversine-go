# Algoritmo de Indicação de Locais com GoLang

Este repositório contém o código-fonte para um algoritmo de indicação de locais similar ao do Google Maps, desenvolvido em GoLang. O projeto utiliza a fórmula de Haversine para calcular distâncias entre pontos geográficos e ordenar locais com base na proximidade do usuário.

## Funcionalidades
- Cálculo de distâncias entre coordenadas geográficas usando a fórmula de Haversine.
- Ordenação de locais com base na proximidade.
- Estruturação em GoLang para eficiência e desempenho.

## Pré-requisitos
- Go 1.16 ou superior instalado.

## Como executar o projeto
1. Clone este repositório:
    ```sh
    git clone https://github.com/gabrielcamurcab/haversine-go.git
    cd haversine-go
    ```

2. Execute o projeto:
    ```sh
    go run main.go
    ```

#### Exemplos de uso
```go
package main

import (
    "fmt"
)

func main() {
    client := Client{
        Latitude:  -23.5505,
        Longitude: -46.6333,
    }

    restaurants := []Restaurant{
        {Name: "Restaurante do João", Latitude: -23.5489, Longitude: -46.6388},
        {Name: "Lanchonete da Maria", Latitude: -23.5510, Longitude: -46.6340},
        {Name: "Cafeteria Chique", Latitude: -23.5530, Longitude: -46.6400},
        {Name: "Pizzaria Gourmet", Latitude: -23.5600, Longitude: -46.6500},
        {Name: "Churrascaria Boi na Brasa", Latitude: -23.5700, Longitude: -46.6300},
        {Name: "Sushi Bar", Latitude: -23.5800, Longitude: -46.6200},
    }

    sortedRestaurants := findClosestDistributors(client, restaurants)

    fmt.Println("Restaurantes ordenados por proximidade:")
    for _, restaurant := range sortedRestaurants {
        fmt.Printf("%s: %.2f km\n", restaurant.Name, restaurant.Distance)
    }
}
```

#### Contribuições
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

#### Licença
Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

#### Contato
Para mais informações, entre em contato através do [LinkedIn](https://www.linkedin.com/in/gabrielcamurca/).

---

Assista ao vídeo tutorial completo no YouTube e veja como desenvolver este algoritmo passo a passo:
[Ir para o vídeo](https://youtu.be/1lb-U0Dvcb0)

---

#GoLang #DesenvolvimentoWeb #GoogleMaps #Algoritmos #Geolocalização #Programação #DesenvolvimentoDeSoftware #Tech #Inovação