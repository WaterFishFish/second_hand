package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"second_hand/DAO"
	"second_hand/logic"
	"strconv"
	"sync"
	"time"
)

var mutex sync.Mutex

func postImg(c *gin.Context) {

}

func GetGoodsList(c *gin.Context) {

	var info DAO.ExaggerateInfo
	// 获取参数的值
	token := c.Query("token")
	name := logic.GetUserNameFromTokenString(token)
	current := c.Query("current")
	pageSize := c.Query("pagesize")
	search := c.Query("search")
	total := c.Query("total")
	class := c.Query("class")
	sortType := c.Query("sortType")
	AD := c.Query("AD")

	time1 := time.Now()
	parseCurrent, _ := strconv.ParseInt(current, 10, 32)
	parsePagesize, _ := strconv.ParseInt(pageSize, 10, 32)
	parseTotal, _ := strconv.ParseInt(total, 10, 32)
	info.UserName = name
	info.Current = int(parseCurrent)
	info.PageSize = int(parsePagesize)
	info.Search = search
	info.Total = int(parseTotal)
	info.Class = class
	info.Time = time1
	info.SortType = sortType
	info.AD = AD

	fmt.Println("Class:", info.Class)
	fmt.Println("info:", info)
	fmt.Println("userName in ReadGoodsList:", name)

	var sellList []DAO.Sell
	var getTotal *gorm.DB
	//插入info到表中,使用save，有则更新，无则插入
	if info.Class == "全部" {
		getTotal = logic.DB.Debug().Find(&sellList)
	} else {
		getTotal = logic.DB.Where("class = ?", info.Class).Find(&sellList)
	}
	if len(sellList) == 0 {
		fmt.Println("该分类没有书籍")
		return
	}
	info.Total = (int(getTotal.RowsAffected) + info.PageSize) / info.PageSize
	fmt.Println("current", info.Current)
	fmt.Println("(info.Current-1)*info.PageSize :", (info.Current-1)*info.PageSize)
	err := logic.SaveInfos(&info)
	var right int
	if info.Current < info.Total {
		right = info.Current * info.PageSize
	} else {
		right = int(getTotal.RowsAffected) + 1
	}

	c.JSON(http.StatusOK, gin.H{
		"list":     sellList[(info.Current-1)*info.PageSize : right],
		"total":    info.Total,
		"current":  info.Current,
		"class":    info.Class,
		"pagesize": info.PageSize,
		"AD":       info.AD,
		"sortType": info.SortType,
	})
	if err != nil {
		fmt.Println(err)
	}

}

func GetDetailInfo(c *gin.Context) {
	sid := c.Query("id")

	parseSid, _ := strconv.ParseInt(sid, 10, 64)
	detailInfo := logic.GetSellById(parseSid)
	c.JSON(http.StatusOK, gin.H{
		"src":         detailInfo.Src,
		"description": detailInfo.Description,
		"price":       detailInfo.Price,
	})

}
