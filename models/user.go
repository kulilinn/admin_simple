package models

import "time"

// User 用户模型
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser 创建新用户
func NewUser(name, email, role string) *User {
	now := time.Now()
	return &User{
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// GetMockUsers 获取模拟用户数据
func GetMockUsers() []*User {
	return []*User{
		{
			ID:        1,
			Name:      "张三",
			Email:     "zhangsan@example.com",
			Role:      "admin",
			CreatedAt: time.Now().AddDate(0, -1, 0),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "李四",
			Email:     "lisi@example.com",
			Role:      "user",
			CreatedAt: time.Now().AddDate(0, -2, 0),
			UpdatedAt: time.Now(),
		},
		{
			ID:        3,
			Name:      "王五",
			Email:     "wangwu@example.com",
			Role:      "user",
			CreatedAt: time.Now().AddDate(0, -3, 0),
			UpdatedAt: time.Now(),
		},
	}
}

