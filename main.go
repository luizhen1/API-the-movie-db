package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title       string  `json:"title"`
	Overview    string  `json:"overview"`
	PosterPath  string  `json:"poster_path"`
	VoteAverage float64 `json:"vote_average"`
}

type Response struct {
	Results []Movie `json:"results"`
}

func main() {
	// Configura o roteador Gin
	router := gin.Default()

	// Define a rota para o endpoint "/top20"
	router.GET("/top20", func(c *gin.Context) {
		// Faz a requisição para o The Movie Database
		url := fmt.Sprintf("https://api.themoviedb.org/3/movie/top_rated?api_key=7a82ba8583c8f6e3642a771e3888f585&language=pt-BR&page=1")
		resp, err := http.Get(url)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		defer resp.Body.Close()

		// Decodifica a resposta em um struct
		var result Response

		err = json.NewDecoder(resp.Body).Decode(&result)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Retorna os 20 primeiros filmes
		if len(result.Results) >= 20 {
			c.JSON(http.StatusOK, result.Results[:20])
		} else {
			c.JSON(http.StatusOK, result.Results)
		}
	})

	// Inicia o servidor
	router.Run(":8080")
}
