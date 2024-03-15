package user

import (
	"log"

	"gorm.io/gorm"
)

// Repository 结构体
type Repository struct {
	db *gorm.DB
}

// 实例化
func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&User{})
	if err != nil {
		log.Print(err)
	}
}

// 创建用户
func (r *Repository) Create(u *User) error {
	result := r.db.Create(u)

	return result.Error
}

// 根据用户名查询用户
func (r *Repository) GetByName(name string) (User, error) {
	var user User
	err := r.db.Where("UserName = ?", name).Where("IsDeleted = ?", 0).First(&user, "UserName = ?", name).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *Repository) GetByID(uid uint) (User, error) {
	var user User
	err := r.db.Where("ID = ?", uid).Where("IsDeleted = ?", 0).First(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// 更新用户
func (r *Repository) Update(u *User) error {
	return r.db.Save(&u).Error
}
