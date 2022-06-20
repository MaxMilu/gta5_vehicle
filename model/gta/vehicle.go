package gta

import "my_qor_test/model"

type Vehicle struct {
	model.BaseModel
	ID              int     //
	Category        string  // 类别英文名
	CategoryName    string  // 类别中文名
	Brand           string  // 品牌
	Name            string  // 名称
	URL             string  //
	Type            string  // 载具类别
	Conditional     string  // 获取条件
	Speed           float64 // 速度
	Acceleration    float64 // 加速度
	Braking         float64 // 刹车
	Handling        float64 // 操控
	TopSpeed        bool    //
	TopAcceleration bool    //
	TopBraking      bool    //
	TopHandling     bool    //
	ForSale         bool    // 可购买
	Website         string  // 购买网站
	Cost            string  // 价格
	Seats           int     // 座位数量
	Personal        bool    //
	Premium         bool    //
	Moddable        bool    // 可改装
	SuperModdable   bool    // 可超级改装
	Sellable        bool    // 可出售
	SellPrice       string  // 出售价格
	MainImageURL    string  // 主图片
	MainImage       []byte  //
	ActionImageURL  string  // 副图片
	ActionImage     []byte  //
}

func (Vehicle) TableName() string {
	return "gta5_vehicles"
}
