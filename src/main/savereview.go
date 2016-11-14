package main

import (
	"bytes"
	"encoding/json"
	//	"fmt"
	//	"io/ioutil"
	//	"log"
	//	"strconv"
	"strings"
	//	"time"

		"../beans"
	//	"../db"
	"../proxyurl"
	//	rx "../readfile/readexcel"
	//	we "../writefile/writeexcel"
	//	wj "../writefile/writejson"
	"github.com/bitly/go-simplejson"
	//	"github.com/henrylee2cn/surfer"
	"github.com/tealeg/xlsx"
)

const (
	excelPath = "/home/zl/reviews/"
)

func main() {
	//readHotelInfo()

}

func readHotelInfo() {

	xlFile, err := xlsx.OpenFile(excelPath + "hotel_new.xlsx")

	if err != nil {
		panic(err)
	}
	sheet := xlFile.Sheets[0]
	for _, row := range sheet.Rows {

		url := row.Cells[1].Value
		println(url)
		if url != "" {
			openMurl(url, row)

			xlFile.Save("/home/zl/gworkspace/sp/hotel_new_11_11.xlsx")
		}

	}

}


func readFile(myhotelId string) {
	
			if err == nil {
				reviewObjectId := hs[0]
				for index, dp := range rsp.DpList {
					la := ""
					switch dp.ServicePoint {
					case "1":
						la = "1"

					case "2":
						la = "3"

					default:
						la = "2"

					}
					src := "3"
					if index%2 == 1 {
						src = "4"
					}
					haspic := "1"
					if dp.DpImgUrl == nil {
						haspic = "2"
					}
					urls := ""
					for _, pu := range dp.DpImgUrl {
						urls += pu.ImgUrl + ","
					}
					cUpdateTime := time.Now().Format("2006-01-02 15:04:05")
					db.InsertReview(&beans.Review{
						reviewObjectId,
						"8",
						"1705",
						reviewObjectId,
						dp.DpUserName,
						"3",
						"",
						la,
						src,
						"2",
						"1",
						"",
						haspic,
						urls,
						dp.ZanCount,
						"0",
						"",
						dp.DpContent,
						"1",
						"2",
						cUpdateTime,
						cUpdateTime,
						"1",
						"",
					})
				}
				//StoreId := hs[0]
				//GoodReviewNum := rsp.GoodNum
				//BadReviewNum := rsp.BadNum
				//CategoryId := "1075"
				//ReviewNum := rsp.TotalNum
				//HasPicReviewNum := rsp.HasImgNum
				//rn, _ := strconv.Atoi(ReviewNum)
				//gr, _ := strconv.Atoi(GoodReviewNum)
				//br, _ := strconv.Atoi(BadReviewNum)
				//hpr, _ := strconv.Atoi(HasPicReviewNum)

				//d := rsp.DegreeLevel
				//var AverageScore float64
				//if d != "" {
				//	ss := strings.Split(d, "%")
				//	temp, _ := strconv.Atoi(ss[0])
				//	AverageScore = (float64(temp) / 100) * 5
				//}

				//UpdateTime, CreateTime := time.Now().Unix(), time.Now().Unix()

				//db.InsertReviewObj(&beans.ReviewObjectVo{
				//	reviewObjectId, //	评论对象ID
				//	6,              //	对象类型ID[1:”商品”,2:”券”,3:”商户(门店)”,4:”票(乐园/秀场)”,5:”电影”]
				//	"",             //	城市ID
				//	"0",            //	广场ID
				//	StoreId,        //	门店ID
				//	CategoryId,     //	类目ID
				//	rn,             //	评论数
				//	AverageScore,   //	综合评分
				//	0,              //	评分1人数
				//	0,              //	评分2人数
				//	0,              //	评分3人数
				//	0,              //	评分4人数
				//	0,              //	评分5人数
				//	gr,             //	好评数
				//	br,             //	差评数
				//	hpr,            //	有图评论数
				//	0,              //	喜欢标识[暂时无用]
				//	CreateTime,     //	创建时间
				//	UpdateTime,     //	修改时间
				//	0,              //	版本号
				//	"",
				//})
			}
		
	}
}



func saveConment(page, hotelid, myhotelName string) {
	resp, err := surfer.Download(&surfer.DefaultRequest{
		Url: "http://www.ly.com/hotel/handler/gethotelcomments.json?pagesize=10&sorttype=0&tagid=1&page=" + page + "&hotelid=" + hotelid,
	})
	var rsp beans.TC
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &rsp)
	if err != nil {
		panic(err)
	}
	c := rsp.Response.Header.RspCode
	if c == "0000" {
		totalPage, _ := strconv.Atoi(rsp.Response.Body.PageInfo.TotalPage)
		p, _ := strconv.Atoi(page)
		js, err := json.Marshal(rsp.Response.Body)
		if err == nil {
			wj.WriteResult(string(js), excelPath+myhotelName+"_"+page)
		}
		if totalPage > 0 {
			saveConment(strconv.Itoa(p+1), hotelid, myhotelName)
			totalPage--
		}
	}
	defer resp.Body.Close()

}

func openMurl(url string, row *xlsx.Row) {
	body, err := proxyurl.Open(url, false)
	if err == nil {
		bodystr := string(body)
		star := strings.Index(bodystr, "initialPageData")
		rs := bodystr[star+18:]
		end := strings.Index(string(rs), "</script>")
		b := rs[0:end]
		println(b)
		js, err1 := simplejson.NewJson([]byte(b))
		if err1 == nil {

			pageData := js.Get("pageData")
			facilitys := pageData.Get("Facilitys")
			cs := pageData.Get("CtripHotelDescription")
			className, c, yearOpen, yearFix, cityName, hotelOrderCount, dpTotalNum, address, faci := "n/a", "n/a", "n/a", "n/a", "n/a", "n/a", "n/a", "n/a", "n/a"
			f := bytes.Buffer{}
			if facilitys != nil {
				facilitysArray, errs := facilitys.Array()
				if errs != nil {
					panic(errs)
				}

				for _, v := range facilitysArray {
					if m, ok := v.(map[string]interface{}); ok {
						f.WriteString(m["FacilityKey"].(string))
						f.WriteString(",")
					}

				}
				faci = f.String()
			}
			if cs != nil {
				cs, errs := pageData.Get("CtripHotelDescription").Array()

				if errs != nil {
					panic(errs)
				}
				if cs != nil && len(cs) > 0 {
					c = cs[0].(string)
				}
			}

			pageDataMap, errs := pageData.Map()
			if errs != nil {
				panic(errs)
			}
			if pageDataMap["YearOpen"] != nil {
				yearOpen = pageDataMap["YearOpen"].(json.Number).String()
			}
			if pageDataMap["YearFix"] != nil {
				yearFix = pageDataMap["YearFix"].(json.Number).String()
			}
			if pageDataMap["CityName"] != nil {
				cityName = pageDataMap["CityName"].(string)
			}
			if pageDataMap["HotelOrderCount"] != nil {
				hotelOrderCount = pageDataMap["HotelOrderCount"].(json.Number).String()
			}
			if pageDataMap["DpTotalNum"] != nil {
				dpTotalNum = pageDataMap["DpTotalNum"].(json.Number).String()
			}
			if pageDataMap["Address"] != nil {
				address = pageDataMap["Address"].(string)
			}
			if pageDataMap["ClassName"] != nil {
				className = pageDataMap["ClassName"].(string)
			}

			row.AddCell().Value = className
			row.AddCell().Value = c
			row.AddCell().Value = yearOpen
			row.AddCell().Value = yearFix
			row.AddCell().Value = cityName
			row.AddCell().Value = hotelOrderCount
			row.AddCell().Value = dpTotalNum
			row.AddCell().Value = address
			row.AddCell().Value = faci
		}
	}

}
