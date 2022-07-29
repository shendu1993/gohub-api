package models

const (
	//categories table
	CategoryStatusNormal  = 1 //正常的状态
	CategoryStatusDeleted = 0 //软删除状态

	//topics table
	TopicStatusPublished = 2 //已经发布
	TopicStatusPending   = 1 //待发布、待审核
	TopicStatusDeleted   = 0 //软删除状态
)
