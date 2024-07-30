package service

type ViewPrize struct {
	ID             int64  `gorm:"column:id" json:"id" form:"id"`
	Name           string `gorm:"column:name" json:"name" form:"name"`
	Pic            string `gorm:"column:pic" json:"pic" form:"pic"`
	Link           string `gorm:"column:link" json:"link" form:"link"`
	Type           int64  `gorm:"column:type" json:"type" form:"type"`
	Data           string `gorm:"column:data" json:"data" form:"data"`
	Total          int64  `gorm:"column:total" json:"total" form:"total"`
	Left           int64  `gorm:"column:left" json:"left" form:"left"`
	IsUse          int64  `gorm:"column:is_use" json:"is_use" form:"is_use"`
	Probability    int64  `gorm:"column:probability" json:"probability" form:"probability"`
	ProbabilityMax int64  `gorm:"column:probability_max" json:"probability_max" form:"probability_max"`
	ProbabilityMin int64  `gorm:"column:probability_min" json:"probability_min" form:"probability_min"`
}
type InitPrizeList struct {
	ViewPrizeList []*ViewPrize `json:"view_prize_list"`
}
type GetPrizeInfoRsp struct {
	PrizeTypeNum int   `json:"prize_type_num"`
	PrizeTotal   int64 `json:"prize_total"`
}
