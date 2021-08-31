package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)


// UserInfo 用户信息
type UserInfo struct {
	gorm.Model
	Name string
	Gender string
	Hobby string
	//Age          sql.NullInt64
	//Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
	//AnimalID int64 `gorm:"primary_key"`

}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:mj123456@tcp(123.56.16.1:3306)/rokyou?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	// 自动迁移
	db.AutoMigrate(&UserInfo{})
	var str = "q212121"
	u1 := UserInfo{gorm.Model{}, "七米22", "男", "篮球","mac@12.cim","dsfd",&str,123,"3222222222",1111}
	//u2 := UserInfo{gorm.Model{}, "沙河娜扎22", "女", "足球"}
	// 创建记录
	db.Create(&u1)
	//db.Create(&u2)

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {

		// 查询
		var uu []UserInfo
		db.Find(&uu, "hobby=?", "足球")
		fmt.Printf("%#v\n", uu)
		c.JSON(http.StatusOK, uu)
	})
	r.GET("/user", func(c *gin.Context) {

		// 查询
		var u = new(UserInfo)
		db.First(u)
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, u)
	})
	r.POST("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"sddddddddddd": "22"})
	})
	r.PUT("/user", func(c *gin.Context) {
		// 更新
		var u = new(UserInfo)
		db.Model(&u).Update("hobby", "双色球")
		c.JSON(http.StatusOK, u)
	})
	r.DELETE("/user", func(c *gin.Context) {
		var u = new(UserInfo)
		// 删除
		db.Delete(&u)
		c.JSON(http.StatusOK, gin.H{"sddddddddddd": "22"})
	})

	r.Run(":7070")
}
