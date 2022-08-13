package requests

import (
    "github.com/gin-gonic/gin"
    "github.com/thedevsaddam/govalidator"
)

type OrderRequest struct {
    // Name        string `valid:"name" json:"name"`
    // Description string `valid:"description" json:"description,omitempty"`
    ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`

    Name     string `valid:"name" json:"name"`
    Email    string `gorm:"type:varchar(255);index;default:null"`
    Phone    string `gorm:"type:varchar(20);index;default:null"`
    Password string `gorm:"type:varchar(255)"`
}

func OrderSave(data interface{}, c *gin.Context) map[string][]string {

    rules := govalidator.MapData{
        "name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:order,name"},
        // "description": []string{"min_cn:3", "max_cn:255"},
    }
    messages := govalidator.MapData{
        "name": []string{
            "required:名称为必填项",
            "min_cn:名称长度需至少 2 个字",
            "max_cn:名称长度不能超过 8 个字",
            "not_exists:名称已存在",
        },
        // "description": []string{
        //     "min_cn:描述长度需至少 3 个字",
        //     "max_cn:描述长度不能超过 255 个字",
        // },
    }
    return validate(data, rules, messages)
}