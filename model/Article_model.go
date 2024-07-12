package model

import (
	"BlogAPI/pkg/database"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Article_ID int    `json:"id"`
	Author_ID  int    `json:"author_id"`
	Content    string `json:"content"`
	//Status     string`json:"id"`
}

func (*Article) GetAll() ([]Article, error) {
	var List []Article
	var err error
	db := database.GetInstance()
	strQuery := "SELECT id, author_id, content FROM `Table`;"
	rows, err := db.Query(strQuery)
	if err != nil {
		fmt.Println("Can't Query")
		log.Fatal(err)
		return List, err
	}

	// Article struct
	var (
		articles_id int
		author_id   int
		content     string
	)

	for rows.Next() {
		err = rows.Scan(&articles_id, &author_id, &content)
		if err != nil {
			log.Fatal(err)
			return List, err
		}
		List = append(List, Article{Article_ID: articles_id, Author_ID: author_id, Content: content})
	}
	return List, err
}

func (x *Article) GetByID(c *gin.Context) error {
	db := database.GetInstance()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	strQuery := "SELECT id, author_id, content FROM `Table` WHERE id = ?"
	res, err := db.Query(strQuery, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var (
		articles_id int
		author_id   int
		content     string
	)

	for res.Next() {
		err = res.Scan(&articles_id, &author_id, &content)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	x.Article_ID = articles_id
	x.Author_ID = author_id
	x.Content = content
	return nil
}
