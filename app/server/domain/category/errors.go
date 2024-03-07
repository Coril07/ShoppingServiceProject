package category

import (
	"errors"
)

var (
	ErrCategoryExistWithName = errors.New("商品分类已经存在")
	ErrCategoryNotFound      = errors.New("类别未找到")
)
