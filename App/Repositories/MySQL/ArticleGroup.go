package MySQL

import "time"

type ArticleGroup struct {
	Id          int        `gorm:"column:id" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Created_at  *time.Time `gorm:"column:created_at" json:"created_at"`
	Updated_at  *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (self ArticleGroup) TableName() string {
	return "blog_article_group"
}
