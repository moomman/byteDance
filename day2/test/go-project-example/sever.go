package main

import (
	"github.com/Moonlight-Zhao/go-project-example/handler"
	"github.com/Moonlight-Zhao/go-project-example/repository"
	"github.com/Moonlight-Zhao/go-project-example/service"
	"github.com/Moonlight-Zhao/go-project-example/util"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()

	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := handler.QueryPageInfo(topicId)
		c.JSON(200, data)
	})

	r.POST("/community/post/do", func(c *gin.Context) {
		uid, _ := c.GetPostForm("uid")
		topicId, _ := c.GetPostForm("topic_id")
		content, _ := c.GetPostForm("content")
		data := handler.PublishPost(uid, topicId, content)
		c.JSON(200, data)
	})

	r.POST("/topic/add", func(context *gin.Context) {
		var topicParam repository.Topic
		err := context.ShouldBindJSON(&topicParam)
		if err != nil {
			return
		}
		var UnixTime int64
		UnixTime, err = strconv.ParseInt(context.PostForm("create_time"), 10, 64)
		if err != nil {
			return
		}
		topicParam.CreateTime = time.Unix(UnixTime, 0)

		err = service.PublishTopic(&topicParam)
		if err != nil {
			log.Println("insert err :", err)
		}
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init() error {
	if err := repository.Init(); err != nil {
		return err
	}
	if err := util.InitLogger(); err != nil {
		return err
	}
	return nil
}
