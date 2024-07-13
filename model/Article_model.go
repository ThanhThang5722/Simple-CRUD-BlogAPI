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
type Article_creation struct {
	Author_ID int    `json:"author_id"`
	Content   string `json:"content"`
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

func (atc *Article_creation) InsertToDB() error {
	db := database.GetInstance()
	strQuery :=
		"INSERT INTO `Table` (author_id, content) Values(?,?);"
	_, err := db.Query(strQuery, strconv.Itoa(atc.Author_ID), atc.Content)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (*Article) CreateItem(List []Article_creation) error {
	for _, atc := range List {
		if err := atc.InsertToDB(); err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

func (*Article) DeleteArticleByID(id int) error {
	db := database.GetInstance()
	strQuery := "DELETE FROM `Table` WHERE id = ?"
	_, err := db.Query(strQuery, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (a *Article) DeleteArticles(Ids []int) error {
	for _, id := range Ids {
		err := a.DeleteArticleByID(id)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
	return nil
}

type Article_updating struct {
	Content string `json:"content"`
}

func (a *Article_updating) UpdateByID(ctx *gin.Context) error {
	db := database.GetInstance()
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	strQuery := "Update `Table` SET content = ? WHERE id = ?"
	_, err = db.Query(strQuery, a.Content, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
