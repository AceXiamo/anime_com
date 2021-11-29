package controllers

import (
	"anime_com/connections"
	"anime_com/global"
	"anime_com/models"
	"github.com/kataras/iris/v12"
	"time"
)

func FindAnimeCom(ctx iris.Context) {
	sqlStr := "select * from anime_com"
	var data []models.AnimeCom
	rows, error := connections.Db.Query(sqlStr)
	if error != nil {
		global.Result(ctx, 500, error)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var anime models.AnimeCom
		var time1 time.Time
		var time2 time.Time
		err := rows.Scan(&anime.Id, &anime.AnimeCover, &anime.AnimeTitle, &anime.AnimeContent, &anime.MyComment, &time1, &time2)
		if err != nil {
			global.Result(ctx, 500, err)
			return
		}
		anime.VisitTime = time1.Format("2006-01-02 15:04:05")
		anime.CreateTime = time2.Format("2006-01-02 15:04:05")
		data = append(data, anime)
	}
	ctx.JSON(data)
}

func AddAnimeCom(ctx iris.Context) {
	var anime models.AnimeCom
	ctx.ReadJSON(&anime)
	anime.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "insert into anime_com(anime_cover, anime_title, anime_content, my_comment, visit_time, create_time) values (?,?,?,?,?,?)"
	ret, err := connections.Db.Exec(sqlStr, &anime.AnimeCover, &anime.AnimeTitle, &anime.AnimeContent, &anime.MyComment, &anime.VisitTime, &anime.CreateTime)
	if err != nil {
		global.Result(ctx, 500, err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		global.Result(ctx, 500, err)
		return
	}
	if n < 1 {
		global.Result(ctx, 500, "新增失败!")
		return
	}
	global.Result(ctx, 200, "add success!")
}

func UpdateAnimeCom(ctx iris.Context) {
	var anime models.AnimeCom
	ctx.ReadJSON(&anime)
	sqlStr := "update anime_com set anime_cover=?,anime_title=?,anime_content=?,my_comment=?,visit_time=? where id=?"
	ret, err := connections.Db.Exec(sqlStr, &anime.AnimeCover, &anime.AnimeTitle, &anime.AnimeContent, &anime.MyComment, &anime.VisitTime, &anime.Id)
	if err != nil {
		global.Result(ctx, 500, err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		global.Result(ctx, 500, err)
		return
	}
	if n < 1 {
		global.Result(ctx, 500, "没有该记录!")
		return
	}
	global.Result(ctx, 200, "update success!")
}

func DelAnimeCom(ctx iris.Context) {
	var anime models.AnimeCom
	ctx.ReadQuery(&anime)
	sqlStr := "delete from anime_com where id = ?"
	ret, error := connections.Db.Exec(sqlStr, anime.Id)
	if error != nil {
		global.Result(ctx, 500, error)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		global.Result(ctx, 500, err)
		return
	}
	if n < 1 {
		global.Result(ctx, 500, "没有该记录!")
		return
	}
	global.Result(ctx, 200, "del success!")
}
