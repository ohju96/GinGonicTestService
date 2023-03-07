package repository

import "sample/ent"

type NoticeRepository interface {
}

type noticeRepository struct {
	client *ent.Client
}

func NewNoticeRepository(client *ent.Client) NoticeRepository {
	return &noticeRepository{client: client}
}
