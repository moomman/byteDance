package repository

import (
	"github.com/Moonlight-Zhao/go-project-example/util"
	"log"
	"sync"
	"time"
)

type Topic struct {
	Id         int64     `gorm:"column:id"`
	UserId     int64     `gorm:"column:user_id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Topic) TableName() string {
	return "topic"
}

type TopicDao struct {
}

var topicDao *TopicDao
var topicOnce sync.Once

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}
func (*TopicDao) CreateTopic(topic Topic) error {
	err := db.Create(&topic).Error
	if err != nil {
		log.Println("insert error: ", err)
		return err
	}
	return nil
}

func (*TopicDao) QueryTopicById(id int64) (*Topic, error) {
	var topic Topic
	err := db.Where("id = ?", id).Find(&topic).Error
	if err != nil {
		util.Logger.Error("find topic by id err:" + err.Error())
		return nil, err
	}
	return &topic, nil
}
