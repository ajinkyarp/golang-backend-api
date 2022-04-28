package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	Roll_no string `json:"roll_no"`
	Name    string `json:"name"`
	Marks   uint16 `json:"marks"`
	Address string `json:"address"`
}

// albums slice to seed record album data.
var student_info = []album{
	{Roll_no: "1", Name: "Arun", Marks: 100, Address: "sadashiv peth, pune-030"},
	{Roll_no: "2", Name: "Rahul", Marks: 75, Address: "shanivar peth, pune-030"},
	{Roll_no: "3", Name: "Vikas", Marks: 50, Address: "shivajinagar, pune-004"},
	{Roll_no: "4", Name: "Sumit", Marks: 88, Address: "hadpsar, pune-028"},
}

func main() {
	router := gin.Default()
	router.GET("/student_info", getAlbums)
	router.GET("/student_info/:id", getAlbumByID)
	router.POST("/student_info", postAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, student_info)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	student_info = append(student_info, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range student_info {
		if a.Roll_no == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student_info not found"})
}
