package users

import (
	"fmt"
	"wyaoo/models"
	"wyaoo/models/bases"
)

type FranchiseePort struct {
	models.LogicBaseModel
	FranchiseeID  uint               `gorm:"not null;default:0"`                              // 加盟商ID
	Type          string             `gorm:"not null"`                                        // 类型（1：始发港；2：中转港；3：目的港）
	PortCode      string             `gorm:"not null"`                                        // 港口代码
	Port          bases.BasePortInfo `gorm:"foreignkey:PortCode;association_foreignkey:Code"` // 港口信息
	Name          string             `gorm:"not null"`
	TransportType string             `gorm:"not null"` // 运输方式（Air：空运；Ocean：海运；Land：陆运）
}

func (FranchiseePort) TableName() string {
	return "user_franchisee_ports"
}

func (p FranchiseePort) Label() string {
	return fmt.Sprintf("%v[%v]", p.PortCode, p.Name)
}
