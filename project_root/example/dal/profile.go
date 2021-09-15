package dal

import "gorm.io/gorm"

func TempGormFix()  {
    println(gorm.Config{})
}