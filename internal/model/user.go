package model

import (
	"encoding/json"
	"log"

	"github.com/pgvector/pgvector-go"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID        string          `json:"user_id" gorm:"type:varchar(255);not null;unique_index"` // 用户分布式 ID
	Username      string          `json:"username" gorm:"type:varchar(255);not null;unique"`      // 用户名,唯一性，用户注册幂等性
	Password      string          `json:"password" gorm:"type:varchar(255);not null"`             // 用户密码（加密后）
	Like          datatypes.JSON  `json:"like" gorm:"type:jsonb;not null"`                        // 用户喜好，存储为 JSON 格式
	LikeEmbedding pgvector.Vector `json:"like_embedding" gorm:"type:vector(768)"`                 // 喜好的词嵌入向量值
}

func (u *User) TableName() string {
	return "users"
}

// CreateUser 创建新用户
func CreateUser(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

// GetUserByUserID 根据 userID 获取用户信息
func GetUserByUserID(db *gorm.DB, userID string) (*User, error) {
	var user User
	err := db.Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserIDByUsername 根据用户名获取 userID
func GetUserIDByUsername(db *gorm.DB, username string) (string, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.UserID, nil
}

// 将 datatypes.JSON 转换为 []string
func (u *User) GetLikeList() []string {
	var likes []string
	err := json.Unmarshal(u.Like, &likes)
	if err != nil {
		log.Printf("Failed to unmarshal Like field: %v", err)
	}
	return likes
}
