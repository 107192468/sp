package db

import (
	"database/sql"
	"fmt"

	"../beans"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func MustConnectDB() {
	if err := connectDatabase("hotel_comment", "hotel_comment", "10.209.44.14", "10044", "hotel_comment"); err != nil {
		panic(err)
	}
}

func connectDatabase(username, password, host, port, database string) (err error) {

	uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	if db, err = sql.Open("mysql", uri); err != nil {
		return
	}

	err = db.Ping()
	return
}

func InsertReview(review *beans.Review) {

	stmt, err := db.Prepare("INSERT INTO review(" +
		"review_object_id,object_type_id," +
		"category_id,store_id," +
		"user_id,current_roster_type," +
		"type,source," +
		"status,order_flag," +
		"order_no,pic_flag," +
		"pic_urls,like_num," +
		"reply_num,title," +
		"content,is_valid," +
		"is_anonym,version," +
		"remark,create_time)VALUES(?,?,?,?,?,?,?,?,?,?," +
		"?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		&review.ReviewObjectId,
		&review.ObjectTypeId,
		&review.CategoryId,
		&review.StoreId,
		&review.UserId,
		"3",
		&review.Type_,
		&review.Source,
		&review.Status,
		&review.OrderFlag,
		&review.OrderNo,
		&review.PicFlag,
		&review.PicUrls,
		&review.LikeNum,
		&review.ReplyNum,
		&review.Title,
		&review.Content,
		&review.IsValid,
		&review.IsAnonym,
		&review.Version,
		&review.Remark,
		&review.CreateTime,
	)
	if err != nil {
		fmt.Println()
	}
}

func InsertReviewObj(reviewObj *beans.ReviewObjectVo) {

	stmt, err := db.Prepare("INSERT INTO review_object(id,object_type_id,city_id," +
		"plaza_id,category_id," +
		"store_id,review_num," +
		"average_score,score_1_num," +
		"score_2_num,score_3_num," +
		"score_4_num,score_5_num," +
		"good_review_num,bad_review_num," +
		"like_flag,version," +
		"remark,create_time) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		&reviewObj.ReviewObjectId,
		&reviewObj.ObjectTypeId,
		&reviewObj.CityId,
		&reviewObj.PlazaId,
		&reviewObj.CategoryId,
		&reviewObj.StoreId,
		&reviewObj.ReviewNum,
		&reviewObj.AverageScore,
		&reviewObj.Score1Num,
		&reviewObj.Score2Num,
		&reviewObj.Score3Num,
		&reviewObj.Score4Num,
		&reviewObj.Score5Num,
		&reviewObj.GoodReviewNum,
		&reviewObj.BadReviewNum,
		&reviewObj.LikeFlag,
		&reviewObj.Version,
		&reviewObj.Remark,
		&reviewObj.CreateTime,
	)
	if err != nil {
		panic(err)
	}
}
func InsertAverageScore(vs *beans.ItemAverageScores) {
	stmt, err := db.Prepare("INSERT INTO item_average_score(review_object_id,tag_id," +
		"tag_name,tag_alias," +
		"average_score" +
		") VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		&vs.ReviewObjectVo,
		&vs.TagId,
		&vs.TagName,
		&vs.TagAlias,
		&vs.AvergeScore,
	)
	if err != nil {
		fmt.Println()
	}
}
