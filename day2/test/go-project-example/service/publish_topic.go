package service

import "github.com/Moonlight-Zhao/go-project-example/repository"

func PublishTopic(topic *repository.Topic) error {
	err := repository.NewTopicDaoInstance().CreateTopic(*topic)
	if err != nil {
		return err
	}

	return nil
}
