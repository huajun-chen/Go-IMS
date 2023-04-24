package model

import (
	"gorm.io/gorm"
	"time"
)

// GoodsCategory 商品分类表
type GoodsCategory struct {
	ID           uint           `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CategoryName string         `json:"category_name" gorm:"size:16;not null;default:'默认分类';comment:'分类名称'"`
	ParentID     int            `json:"parent_id" gorm:"index;default:0;comment:'父级分类ID，默认为0，表示该分类没有父级分类'"`
}

// TableName 自定义表名
// 参数：
//		无
// 返回值：
//		string：表名
func (GoodsCategory) TableName() string {
	return "goods_category"
}

// Goods 商品表
type Goods struct {
	ID             uint           `json:"id" gorm:"primarykey"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	GoodsName      string         `json:"goods_name" gorm:"size:32;not null;comment:'商品名称'"`
	CategoryID     int            `json:"category_id" gorm:"not null;index;comment:'所属分类ID'"`
	GoodsUnit      string         `json:"goods_unit" gorm:"size:1;not null;comment:'商品单位（件、个、套、台、份...）'"`
	GoodsWeight    int            `json:"goods_weight"gorm:"comment:'商品重量'"`
	WholesalePrice int            `json:"wholesale_price" gorm:"comment:'批发价'"`
	RetailPrice    int            `json:"retail_price" gorm:"comment:'零售价'"`
	NoteInfo       string         `json:"note_info" gorm:"size:256;comment:'备注信息'"`
	ImageURL       string         `json:"image_url" gorm:"size:256;comment:'商品图片URL地址'"`
}

// TableName 自定义表名
// 参数：
//		无
// 返回值：
//		string：表名
func (Goods) TableName() string {
	return "goods"
}
