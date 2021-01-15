package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"CURD/database"
)

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Content string `json:"body"`
}

func Read(c *gin.Context) {
	db := database.DBConn()

	if nil == db {
		log.Println("Info: Can't connect to dabase!")
		return
	}

	query := "SELECT id, title, body FROM post WHERE id = " + c.Param("id")
	rows, err := db.Query(query)

	log.Println("Info: Execute query '", query, "'")

	if nil != err {
		log.Println("Error: ", err)

		c.JSON(500, gin.H {
			"messages": "Story not found",
		});
	}

	post := Post{}
	if rows != nil {
		for rows.Next() {
			var id int
			var title, body string

			err = rows.Scan(&id, &title, &body)
			if nil != err {
				panic (err.Error())
			}

			post.Id = id
			post.Title = title
			post.Content = body
		}

		c.JSON(200, post)
		defer db.Close()
	} else {
		log.Println("Info: Empty data!")

	}
}

func Create(c *gin.Context) {
	db := database.DBConn()

	if nil == db {
		log.Println("Create function -> Error: Can't connect to database")
	}

	type CreatePost struct {
		Title string `form:"title json:"title" binding:"required"`
		Body string `form:"body json:"body" binding:"required"`
	}

	var json CreatePost

	if err := c.ShouldBindJSON(&json); nil == err {
		query := "INSER INTO post(title, body) VALUES(?, ?)"
		insPost, err := db.Prepare(query)
		if nil != err {
			log.Println("Create function -> Info: Prepare query")

			c.JSON(500, gin.H {
				"mesasges": err,
			})
		}

		insPost.Exec(json.Title, json.Body) 

		log.Println ("Create function -> Info: Execute query")

		c.JSON(200, gin.H {
			"messages": "inserted",
		})
	} else {
		log.Println ("Create function -> Error: ", err)
		c.JSON(500, gin.H {"error": err.Error()})
	}

	defer db.Close()
}

func Update(c *gin.Context) {
	db := database.DBConn()

	if nil == db {
		log.Println("Update function -> Error: Can't connect to database")
	}

	type UpdatePost struct {
		Title string `form:"title" json:"title" binding:"required"`
		Body string `form:"body" json:"body" binding:"required"`
	}

	var json UpdatePost
	if err := c.ShouldBindJSON(&json); nil == err {	
		edit, err := db.Prepare ("UPDATE post SET title=?, body=? WHERE id = " + c.Param("id"))
		if nil != err {
			panic (err.Error())
		}
		edit.Exec(json.Title, json.Body)

		c.JSON(200, gin.H {
			"messages": "edited",
		})
	} else {
		c.JSON (500, gin.H {"error": err.Error()})
	}

	defer db.Close()
}



func Delete (c *gin.Context) {
	db := database.DBConn()

	if nil == db {
		log.Println ("Delete function -> Error: Can't connect to database")
	}

	delete, err := db.Prepare ("DELETE FROM post WHERE id = ?")
	if nil != err {
		panic (err.Error())
	}

	delete.Exec (c.Param("id"))
	c.JSON (200, gin.H {
		"messages": "deleted",
	})

	defer db.Close()
}