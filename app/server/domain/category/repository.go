package category

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// 创建商品分类
func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成商品分类表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
}

// 创建商品分类
func (r *Repository) Create(c *Category) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 通过名称查询商品分类
func (r *Repository) GetByName(name string) (*Category, error) {
	var category *Category
	err := r.db.Where("IsActive = ?", 1).Where(Category{Name: name}).First(&category).Error
	if err != nil {
		return nil, ErrCategoryNotFound
	}
	return category, nil
}

// 批量创建商品分类
func (r *Repository) BulkCreate(categories []*Category) (int, error) {
	var count int64
	err := r.db.Create(&categories).Count(&count).Error
	return int(count), err
}

// 获得分页内分类
func (r *Repository) GetAllInPage(pageIndex, pageSize int) ([]Category, int, int) {
	var categories []Category
	var t []Category
	var count int64

	r.db.Where("IsActive = ?", 1).Find(&t).Count(&count)
	r.db.Where("IsActive = ?", 1).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories)

	return categories, int(count), int(count)/pageSize + 1
}

func (r *Repository) GetAll() ([]Category, int) {
	var categories []Category
	var count int64

	r.db.Where("IsActive = ?", 1).Find(&categories).Count(&count)

	return categories, int(count)
}

// 根据分类名称删除
func (r *Repository) Delete(name string) error {
	currentCategory, err := r.GetByName(name)
	if err != nil {
		return err
	}
	currentCategory.IsActive = false

	err = r.db.Save(currentCategory).Error
	return err
}
