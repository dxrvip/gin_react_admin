package service

import (
	"goVueBlog/models"
	"sort"
)

type CategoryTree struct {
	ID        uint           `json:"id"`
	ParentID  *uint          `json:"parentId,omitempty"`
	Name      string         `json:"name"`
	Order     int            `json:"order"`
	IsEnabled bool           `json:"isEnabled"`
	Children  []CategoryTree `json:"children,omitempty"`
}

var shopCategoryService *ShopCategoryService

type ShopCategoryService struct {
	BaseService
}

func NewCategoriesService() *ShopCategoryService {
	if shopCategoryService == nil {
		return &ShopCategoryService{
			BaseService: NewBaseApi(&models.ProductCategory{}),
		}
	}
	return shopCategoryService
}

// 添加树形结构构建方法
func BuildCategoryTree(categories []models.ProductCategory) []CategoryTree {
	categoryMap := make(map[uint]*CategoryTree)
	// 第一遍遍历创建映射
	for _, c := range categories {
		categoryMap[c.ID] = &CategoryTree{
			ID:        c.ID,
			ParentID:  c.ParentID,
			Name:      c.Name,
			Order:     c.Order,
			IsEnabled: c.IsEnabled,
			Children:  []CategoryTree{},
		}
	}

	// 第二遍遍历构建树
	var tree []CategoryTree
	for _, node := range categoryMap {
		if node.ParentID == nil {
			tree = append(tree, *node)
		} else {
			if parent, exists := categoryMap[*node.ParentID]; exists {
				parent.Children = append(parent.Children, *node)
			}
		}
	}

	// 排序逻辑
	sort.SliceStable(tree, func(i, j int) bool {
		return tree[i].Order < tree[j].Order
	})
	for i := range tree {
		sort.SliceStable(tree[i].Children, func(j, k int) bool {
			return tree[i].Children[j].Order < tree[i].Children[k].Order
		})
	}

	return tree
}

// 修改 ShopCategoryService 的 List 方法
func (s *ShopCategoryService) ListTree(categories []models.ProductCategory) ([]CategoryTree, error) {

	return BuildCategoryTree(categories), nil
}
