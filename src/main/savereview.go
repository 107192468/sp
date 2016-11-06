package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	//	"regexp"
	"bytes"
	"time"

	"../beans"
	"../db"
	"../proxyurl"
	rx "../readfile/readexcel"
	"../readfile/readhtml"
	we "../writefile/writeexcel"
	wj "../writefile/writejson"
	"github.com/antchfx/xquery/html"
	"github.com/bitly/go-simplejson"
	"github.com/henrylee2cn/surfer"
	"github.com/tealeg/xlsx"
	"golang.org/x/net/html"
)

const (
	excelPath = "/home/zl/reviews/"
)

func main() {
	readHotelInfo()
	//	readReview()
	//	readExcel()
	//	db.MustConnectDB()
	//	readFile()
	//for k,v := range am{
	//	dolink("1",k,v[1],v[2])
	//}
	//	saveConment("1", "170221", "91", "101881")
}

//func dolink(page, hotelid, cityId, myhotelid string) {
//	resp, err := surfer.Download(&surfer.DefaultRequest{
//		Url: "http://www.ly.com/hotel/handler/gethotelcomments.json?pagesize=10&sorttype=0&tagid=1&page=" + page + "&hotelid="+hotelid,
//	})
//	defer resp.Body.Close()
//	var rsp beans.Response
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}
//	err = json.Unmarshal(body, &rsp)
//	if err != nil {
//		panic(err)
//	}
//	c := rsp.Header.RspCode
//	if c != "0000" {
//		panic("bad Response")
//	}
//	totalPage := rsp.Body.PageInfo.TotalPage
//	p, _ := strconv.Atoi(page)
//	reviewObjectId := hotelid
//	StoreId := hotelid
//	GoodReviewNum := rsp.Body.GoodNum
//	BadReviewNum := rsp.Body.BadNum
//	CategoryId := "1075"
//	ReviewNum := rsp.Body.TotalNum
//	HasPicReviewNum := rsp.Body.HasImgNum
//	d := rsp.Body.DegreeLevel
//	var AverageScore float64
//	if d != "" {
//		ss := strings.Split(d, "%")
//		temp, _ := strconv.Atoi(ss[0])
//		AverageScore = math.Floor(float64(temp)/100) * 5
//	}
//
//	UpdateTime, CreateTime := time.Now().Unix(), time.Now().Unix()
//	db.InsertReviewObj(&beans.ReviewObjectVo{
//		reviewObjectId,  //	评论对象ID
//		6,               //	对象类型ID[1:”商品”,2:”券”,3:”商户(门店)”,4:”票(乐园/秀场)”,5:”电影”]
//		cityId,          //	城市ID
//		"0",             //	广场ID
//		StoreId,         //	门店ID
//		CategoryId,      //	类目ID
//		ReviewNum,       //	评论数
//		AverageScore,    //	综合评分
//		0,               //	评分1人数
//		0,               //	评分2人数
//		0,               //	评分3人数
//		0,               //	评分4人数
//		0,               //	评分5人数
//		GoodReviewNum,   //	好评数
//		BadReviewNum,    //	差评数
//		HasPicReviewNum, //	有图评论数
//		0,               //	喜欢标识[暂时无用]
//		CreateTime,      //	创建时间
//		UpdateTime,      //	修改时间
//		0,               //	版本号
//		"",
//	})
//
//	if totalPage > 0 {
//		dolink(strconv.Itoa(p+1), hotelid, cityId, myhotelid)
//		totalPage--
//	}
//
//}

//func saveConment(page, hotelid, cityId, myhotelid string) {
//	resp, err := surfer.Download(&surfer.DefaultRequest{
//		Url: "http://www.ly.com/hotel/handler/gethotelcomments.json?pagesize=10&sorttype=0&tagid=1&page=" + page + "&hotelid=" + hotelid,
//	})
//	var rsp beans.TC
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}

//	err = json.Unmarshal(body, &rsp)
//	if err != nil {
//		panic(err)
//	}
//	c := rsp.Response.Header.RspCode
//	if c == "0000" {
//		totalPage, _ := strconv.Atoi(rsp.Response.Body.PageInfo.TotalPage)
//		p, _ := strconv.Atoi(page)
//		js, err := json.Marshal(rsp.Response.Body)
//		if err == nil {
//			wj.WriteResult(string(js), excelPath+myhotelid+"_"+cityId+"_"+page)
//		}
//		if totalPage > 0 {
//			saveConment(strconv.Itoa(p+1), hotelid, cityId, myhotelid)
//			totalPage--
//		}
//	}
//	defer resp.Body.Close()

//}

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

//func readExcel() {
//	excelFileName := "/Users/zhangling/Documents/rs.xlsx"
//	xlFile, err := xlsx.OpenFile(excelFileName)

//	if err == nil {
//		sheet1 := xlFile.Sheets[0]
//		if sheet1 == nil {
//			return
//		}

//		for i, row := range sheet1.Rows {
//			if i < 1 {
//				continue
//			}
//			cell := row.Cells
//			myhotelid := cell[2].Value
//			if len(cell) == 5 {
//				hid := cell[3].Value
//				city := cell[4].Value
//				saveConment("1", hid, city, myhotelid)
//			}

//		}
//	}

//}

func readFile() {
	path := "/Users/zhangling/Documents/file/"
	files, err := ioutil.ReadDir(path)
	//m := make(map[string]int, len(files))
	var rsp beans.Body
	if err == nil {
		for _, file := range files {
			filename := file.Name()
			hs := strings.Split(filename, "_")
			b, _ := ioutil.ReadFile(path + filename)
			//if m[hs[0]] == 1 {
			//	continue
			//} else {
			//	m[hs[0]] = 1
			//}
			fmt.Println(string(b))
			json.Unmarshal(b, &rsp)

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
						"6",
						"1075",
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
}

//func readExcel() {
//	am = make(map[string][]string)
//	excelFileName := "/Users/zhangling/Documents/zl.xlsx"
//	xlFile, err := xlsx.OpenFile(excelFileName)
//
//	if err == nil {
//		m := xlFile.Sheet
//		print(m["s1"])
//		sheet1 := xlFile.Sheets[0]
//		if sheet1 == nil {
//			return
//		}
//
//		for i, row := range sheet1.Rows {
//
//			if i < 1 {
//				continue
//			}
//			if i == 1 {
//				v, _ := row.Cells[1].String()
//				println(v)
//			} else {
//				hname := row.Cells[1].Value
//				row.AddCell().Value,row.AddCell().Value=getHotelId(hname)
//			}
//			xlFile.Save("/Users/zhangling/Documents/rs.xlsx")
//		}
//	}
//
//}
func getHotelId(hname string) (string, string) {
	resp, err := surfer.Download(&surfer.DefaultRequest{
		Url: "http://m.ly.com/AjaxHelper/GlobalSearchHandler.ashx?keyword=" + hname,
	})
	if err == nil {
		b, err := ioutil.ReadAll(resp.Body)
		js, err := simplejson.NewJson(b)
		if err == nil {
			docs := js.Get("Response")
			if docs != nil {
				docs = docs.Get("result")
				if docs != nil {
					docs = docs.Get("doc")
					if docs != nil {
						arr, err := docs.Array()
						if err == nil {
							for i, v := range arr {
								if i == 1 {
									if m, ok := v.(map[string]interface{}); ok {
										return m["id"].(string), m["cityid"].(string)
									}

								}
							}
						}

					}
				}
			}
		}

		defer resp.Body.Close()
	}
	return "", ""
}

//func writeExcel(){
//	am = make(map[string][]string)
//	excelFileName := "/Users/zhangling/Documents/rs.xlsx"
//	xlFile := xlsx.NewFile()
//	sheet, err :=xlFile.AddSheet("data")
//		if err == nil {
//			return
//		}
//		for k, v := range am {
//			nr := sheet.AddRow()
//			if i<1{
//				continue
//			}
//
//			nr.AddCell().Value = am[k]
//			nr.AddCell().Value = row.Cells[1].Value
//			nr.AddCell().Value = row2.Cells[1].Value
//			nr.AddCell().Value = row2.Cells[2].Value
//			nr.AddCell().Value = row2.Cells[3].Value
//			nr.AddCell().Value = row2.Cells[4].Value
//			nr.AddCell().Value = row2.Cells[5].Value
//			xlFile.Save(excelFileName)
//			//if i < 1 {
//			//	continue
//			//}
//			//if i == 1 {
//			//	v, _ := row.Cells[1].String()
//			//	println(v)
//			//} else {
//			//	hname := row.Cells[1].Value
//			//	hid := row.Cells[2].Value
//			//	getHotelId(hname, hid)
//			//}
//
//		}
//
//}
func readHotelInfo() {
	xl := we.CreateNewExcel("/home/zl/gworkspace/sp/hotel_new_06.xlsx")
	for i, row := range rx.ReadExcelRow("/home/zl/gworkspace/sp/hotel_new.xlsx", 0) {
		log.Println(i)
		url := rx.ReadExcelCell(row.Cells, 1)

		if url != "" {
			s := strings.Split(url, "_")
			if len(s) == 2 && s[1] != "" {
				nurls := "http://www.ly.com/HotelInfo-" + s[1]

				body, err := proxyurl.Open(nurls, false)

				if err == nil {
					reader := strings.NewReader(string(body))

					q, detail := readhtml.ReadHotelSectionListAndFacility(reader)
					var rs []string={q,detail}
					xl.WriteNewExcelRow(rs)
				}

			}

			//			facilityKey, cd, yearOpen, yearFix := openUrl(url + "?pageName=hotelDescription")

			//			row.AddCell().Value = facilityKey
			//			row.AddCell().Value = cd
			//			row.AddCell().Value = yearOpen
			//			row.AddCell().Value = yearFix
			//			xlFile.Save("/home/zl/gworkspace/sp/hotel_new_11_6.xlsx")
		}

	}

	//	openUrl("http://www.ly.com/HotelInfo-93781.html")

}

func readReview() {
	//	excelFileName := "/Users/zhangling/Documents/hotel_new.xlsx"
	excelFileName := "/home/zl/gworkspace/sp/hotel_new.xlsx"
	xlFile, err1 := xlsx.OpenFile(excelFileName)

	if err1 != nil {
		panic(err1)
	}
	sheet1 := xlFile.Sheets[0]
	if sheet1 == nil {
		return
	}

	for i, row := range sheet1.Rows {
		if i < 154 {
			continue
		}
		cell := row.Cells

		url := cell[1].Value
		fileName := cell[0].Value
		if url != "" {
			s := strings.Split(url, "_")
			if len(s) == 2 && s[1] != "" {
				ids := strings.Split(s[1], ".")
				println(ids[0])
				saveConment("1", ids[0], fileName)
			}
		}

	}

	//	openUrl("http://www.ly.com/HotelInfo-93781.html")

}
func openUrl(url string) (string, string) {
	body, err := proxyurl.Open(url, false)
	if err == nil {
		reader := strings.NewReader(string(body))
		htmltc, xmlerr := html.Parse(reader)
		if xmlerr != nil {
			log.Fatal(xmlerr)
		}

		xpath := `//*[@id="HotelSectionList"]`
		node := htmlquery.FindOne(htmltc, xpath)
		q := htmlquery.InnerText(node)
		hotel_details := `//*[@id="hotel-details"]/div[1]/ul/li`
		details := bytes.Buffer{}
		htmlquery.FindEach(htmltc, hotel_details, func(i int, node *html.Node) {

			lable := htmlquery.FindOne(node, "//label")
			//			fmt.Println(htmlquery.InnerText(lable))
			details.WriteString(htmlquery.InnerText(lable))
			details.WriteString(":")
			span := "//span"
			htmlquery.FindEach(node, span, func(j int, spanNode *html.Node) {
				fmt.Println(j)
				s := htmlquery.FindOne(spanNode, "//span")
				details.WriteString(htmlquery.InnerText(s))
				details.WriteString(",")
			})
		})
		return q, details.String()
	}
	return "n/a", "n/a"
}

func openMurl(url string) (string, string, string, string) {
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
			//			page, errs := js.Get("pageData").Map()
			//			if errs != nil {
			//				panic(errs)
			//			}
			//			if page["BusinessCircle"] != nil {
			//				return page["BusinessCircle"].(string)
			//			}
			facilitys, errs := js.Get("pageData").Get("Facilitys").Array()
			if errs != nil {
				panic(errs)
			}
			f := bytes.Buffer{}
			for _, v := range facilitys {
				if m, ok := v.(map[string]interface{}); ok {
					f.WriteString(m["FacilityKey"].(string))
					f.WriteString(",")
				}

			}
			cs, errs := js.Get("pageData").Get("CtripHotelDescription").Array()
			c := ""
			if errs != nil {
				panic(errs)
			}
			if cs != nil && len(cs) > 0 {
				c = cs[0].(string)
			}
			yearOpen := ""
			yearFix := ""
			pageData, errs := js.Get("pageData").Map()

			if errs != nil {
				panic(errs)
			}
			if pageData["YearOpen"] != nil {
				yearOpen = pageData["YearOpen"].(json.Number).String()
			}
			if pageData["YearFix"] != nil {
				yearFix = pageData["YearFix"].(json.Number).String()
			}
			//			if page["BusinessCircle"] != nil {
			//				return page["BusinessCircle"].(string)
			//			}

			return f.String(), c, yearOpen, yearFix
		}
	}
	return "n/a", "n/a", "n/a", "n/a"
}
