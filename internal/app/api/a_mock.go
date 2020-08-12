package api

import (
	"gin-admin-template/internal/app/ginplus"
	"gin-admin-template/internal/app/schema"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"math/rand"
	"time"
)

// SysSet 注入Sys
var MockSet = wire.NewSet(wire.Struct(new(Mock), "*"))

// Demo 示例程序
type Mock struct{}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn((max+1)-min) + min
}

// UsersQuery 模拟用户查询数据
func (a *Mock) UserQuery(c *gin.Context) {

	gofakeit.Seed(time.Now().UnixNano())

	type user struct {
		Id         string   `json:"id"`
		Username   string   `json:"username"`
		Name       string   `json:"name"`
		Department string   `json:"department"`
		Starttime  string   `json:"starttime"`
		State      int      `json:"state"`
		Sex        int      `json:"sex"`
		Age        int      `json:"age"`
		Email      string   `json:"email"`
		Areacode   []string `json:"areacode"`
		Areaname   string   `json:"areaname"`
	}

	var userList []user

	for i := 0; i < 10; i++ {
		user := user{
			gofakeit.UUID(),
			gofakeit.FirstName(),
			gofakeit.Username(),
			gofakeit.State(),
			gofakeit.Date().Format("2006-01-02 15:04:05"),
			gofakeit.Number(0, 1),
			gofakeit.Number(1, 2),
			gofakeit.Number(16, 80),
			gofakeit.Email(),
			[]string{"hangzhou"},
			"杭州",
		}
		userList = append(userList, user)
	}

	ginplus.ResPage(c, userList, &schema.PaginationResult{Total: 40, Current: 1, PageSize: 10})
}
