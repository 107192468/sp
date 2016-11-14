package readhtml

import (
	"bytes"
	"io"
	"log"

	"regexp"
	"strings"

	//	"../db"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/xquery/html"
	"golang.org/x/net/html"
)

func ReadHotelSectionListAndFacility(body io.Reader) (string, string) {

	htmltc, xmlerr := html.Parse(body)
	if xmlerr != nil {
		log.Fatal(xmlerr)
	}
	q := "n/a"
	detail := "n/a"
	xpath := `//*[@id="HotelSectionList"]`
	node := htmlquery.FindOne(htmltc, xpath)

	section := htmlquery.InnerText(node)
	hotel_details := `//*[@id="hotel-details"]/div[1]/ul/li`

	details := bytes.Buffer{}
	htmlquery.FindEach(htmltc, hotel_details, func(i int, node *html.Node) {

		lable := htmlquery.FindOne(node, "//label")

		details.WriteString(htmlquery.InnerText(lable))
		details.WriteString(":")
		span := "//span"
		htmlquery.FindEach(node, span, func(j int, spanNode *html.Node) {

			s := htmlquery.FindOne(spanNode, "//span")
			details.WriteString(htmlquery.InnerText(s))
			details.WriteString(",")
		})
	})
	detail = details.String()
	if section != "" {
		q = section
	}
	return q, detail

}
func readarea(url string) {
	p, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	} else {
		p.Find("div.area-common").Each(func(i int, s *goquery.Selection) {
			s.Find("div.select-single").Each(func(j int, sel *goquery.Selection) {
				data_id, isd := sel.Attr("data-id")
				if isd {
					print(" ")
					print(data_id)
				}
				title, ist := sel.Attr("title")
				if ist {
					print(" ")
					print(title)
				}
			})
		})
	}
}
func readBusinessSectionId(url string) {
	p, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	} else {
		p.Find("div.selects-wrap").Each(func(i int, s *goquery.Selection) {
			data_modul, ok := s.Attr("data-module")
			if ok && data_modul == "txtBusinessSectionId" {

				s.Find("div.select-single").Each(func(j int, sel *goquery.Selection) {

					data_id, isd := sel.Attr("data-id")
					if isd {
						print(" ")
						print(data_id)
					}
					title, ist := sel.Attr("title")
					if ist {
						print(" ")
						print(title)
					}

				})
				print(" ")

			}

		})

	}
}

func phtml(url, dataId string) {
	doc, err := htmlquery.Load(url)
	if err != nil {
		panic(err)
	}
	area := `//*[@id="filterbox"]/li[1]/div[2]`

	areaNode := htmlquery.FindOne(doc, area)
	div := `//div`
	htmlquery.FindEach(areaNode, div, func(j int, divNode *html.Node) {

		htmlquery.FindEach(divNode, "//div[@class='selects-wrap options-wrap clearfix']", func(j int, selectNode *html.Node) {

			for _, attr := range selectNode.Attr {
				if attr.Val == "txtBusinessSectionId" {
					if n := htmlquery.FindOne(selectNode, "//div[@class='select-single']"); n != nil {
						id := ""
						bname := ""
						for _, nattr := range n.Attr {
							if nattr.Key == "data-id" {
								id = nattr.Val
								print(id)
								print("  ")
							}
							if nattr.Key == "title" {
								bname = nattr.Val
								println(bname)
							}
						}
						//						db.InsertBusiness(id, bname, dataId)
					}
				}
			}

		})

	})
	return
}

func ReadHtmlBusiness() {
	//	db.MustConnectDB()
	tl := `<h3 class="city-key-A">A</h3><ul class="sort-city clearfix"><li data-id="36" data-name="安庆">安庆</li><li data-id="150" data-name="安阳">安阳</li><li data-id="322" data-name="阿坝">阿坝</li><li data-id="245" data-name="鞍山">鞍山</li><li data-id="112" data-name="安顺">安顺</li><li data-id="311" data-name="安康">安康</li><li data-id="351" data-name="阿克苏">阿克苏</li><li data-id="3114" data-name="阿勒泰">阿勒泰</li></ul><h3 class="city-key-B">B</h3><ul class="sort-city clearfix"><li data-id="53" data-name="北京">北京</li><li data-id="99" data-name="北海">北海</li><li data-id="261" data-name="包头">包头</li><li data-id="139" data-name="保定">保定</li><li data-id="37" data-name="蚌埠">蚌埠</li><li data-id="283" data-name="滨州">滨州</li><li data-id="312" data-name="宝鸡">宝鸡</li><li data-id="367" data-name="保山">保山</li><li data-id="52" data-name="亳州">亳州</li><li data-id="98" data-name="百色">百色</li><li data-id="213" data-name="白山">白山</li><li data-id="246" data-name="本溪">本溪</li><li data-id="353" data-name="巴音郭楞">巴音郭楞</li><li data-id="260" data-name="巴彦淖尔">巴彦淖尔</li><li data-id="323" data-name="巴中">巴中</li><li data-id="122" data-name="保亭">保亭</li><li data-id="113" data-name="毕节">毕节</li><li data-id="63" data-name="白银">白银</li></ul><h3 class="city-key-C">C</h3><ul class="sort-city clearfix"><li data-id="324" data-name="成都">成都</li><li data-id="199" data-name="长沙">长沙</li><li data-id="394" data-name="重庆">重庆</li><li data-id="221" data-name="常州">常州</li><li data-id="214" data-name="长春">长春</li><li data-id="4569" data-name="长白山">长白山</li><li data-id="39" data-name="池州">池州</li><li data-id="141" data-name="承德">承德</li><li data-id="198" data-name="常德">常德</li><li data-id="200" data-name="郴州">郴州</li><li data-id="300" data-name="长治">长治</li><li data-id="140" data-name="沧州">沧州</li><li data-id="40" data-name="滁州">滁州</li><li data-id="38" data-name="巢湖">巢湖</li><li data-id="77" data-name="潮州">潮州</li><li data-id="100" data-name="崇左">崇左</li><li data-id="368" data-name="楚雄">楚雄</li><li data-id="262" data-name="赤峰">赤峰</li><li data-id="247" data-name="朝阳">朝阳</li><li data-id="355" data-name="昌吉">昌吉</li><li data-id="345" data-name="昌都">昌都</li><li data-id="123" data-name="昌江">昌江</li><li data-id="124" data-name="澄迈">澄迈</li></ul><h3 class="city-key-D">D</h3><ul class="sort-city clearfix"><li data-id="248" data-name="大连">大连</li><li data-id="78" data-name="东莞">东莞</li><li data-id="369" data-name="大理">大理</li><li data-id="301" data-name="大同">大同</li><li data-id="249" data-name="丹东">丹东</li><li data-id="285" data-name="东营">东营</li><li data-id="326" data-name="德阳">德阳</li><li data-id="168" data-name="大庆">大庆</li><li data-id="371" data-name="迪庆">迪庆</li><li data-id="284" data-name="德州">德州</li><li data-id="138" data-name="儋州">儋州</li><li data-id="370" data-name="德宏">德宏</li><li data-id="169" data-name="大兴安岭">大兴安岭</li><li data-id="125" data-name="定安">定安</li><li data-id="325" data-name="达州">达州</li></ul><h3 class="city-key-E">E</h3><ul class="sort-city clearfix"><li data-id="263" data-name="鄂尔多斯">鄂尔多斯</li><li data-id="182" data-name="恩施">恩施</li><li data-id="181" data-name="鄂州">鄂州</li></ul><h3 class="city-key-F">F</h3><ul class="sort-city clearfix"><li data-id="54" data-name="福州">福州</li><li data-id="79" data-name="佛山">佛山</li><li data-id="41" data-name="阜阳">阜阳</li><li data-id="101" data-name="防城港">防城港</li><li data-id="250" data-name="抚顺">抚顺</li><li data-id="234" data-name="抚州">抚州</li></ul><h3 class="city-key-G">G</h3><ul class="sort-city clearfix"><li data-id="80" data-name="广州">广州</li><li data-id="102" data-name="桂林">桂林</li><li data-id="114" data-name="贵阳">贵阳</li><li data-id="235" data-name="赣州">赣州</li><li data-id="327" data-name="甘孜">甘孜</li><li data-id="329" data-name="广元">广元</li><li data-id="328" data-name="广安">广安</li><li data-id="103" data-name="贵港">贵港</li><li data-id="271" data-name="固原">固原</li></ul><h3 class="city-key-H">H</h3><ul class="sort-city clearfix"><li data-id="383" data-name="杭州">杭州</li><li data-id="395" data-name="香港">香港</li><li data-id="170" data-name="哈尔滨">哈尔滨</li><li data-id="42" data-name="合肥">合肥</li><li data-id="45" data-name="黄山">黄山</li><li data-id="127" data-name="海口">海口</li><li data-id="264" data-name="呼和浩特">呼和浩特</li><li data-id="222" data-name="淮安">淮安</li><li data-id="82" data-name="惠州">惠州</li><li data-id="384" data-name="湖州">湖州</li><li data-id="265" data-name="呼伦贝尔">呼伦贝尔</li><li data-id="201" data-name="衡阳">衡阳</li><li data-id="286" data-name="菏泽">菏泽</li><li data-id="142" data-name="邯郸">邯郸</li><li data-id="81" data-name="河源">河源</li><li data-id="313" data-name="汉中">汉中</li><li data-id="202" data-name="怀化">怀化</li><li data-id="252" data-name="葫芦岛">葫芦岛</li><li data-id="143" data-name="衡水">衡水</li><li data-id="44" data-name="淮南">淮南</li><li data-id="184" data-name="黄石">黄石</li><li data-id="183" data-name="黄冈">黄冈</li><li data-id="279" data-name="海西">海西</li><li data-id="171" data-name="鹤岗">鹤岗</li><li data-id="172" data-name="黑河">黑河</li><li data-id="104" data-name="河池">河池</li><li data-id="43" data-name="淮北">淮北</li><li data-id="105" data-name="贺州">贺州</li><li data-id="151" data-name="鹤壁">鹤壁</li><li data-id="357" data-name="和田">和田</li><li data-id="278" data-name="海南藏族">海南藏族</li><li data-id="372" data-name="红河">红河</li></ul><h3 class="city-key-J">J</h3><ul class="sort-city clearfix"><li data-id="287" data-name="济南">济南</li><li data-id="385" data-name="嘉兴">嘉兴</li><li data-id="386" data-name="金华">金华</li><li data-id="238" data-name="九江">九江</li><li data-id="303" data-name="晋中">晋中</li><li data-id="288" data-name="济宁">济宁</li><li data-id="83" data-name="江门">江门</li><li data-id="236" data-name="吉安">吉安</li><li data-id="215" data-name="吉林">吉林</li><li data-id="68" data-name="酒泉">酒泉</li><li data-id="237" data-name="景德镇">景德镇</li><li data-id="186" data-name="荆州">荆州</li><li data-id="153" data-name="焦作">焦作</li><li data-id="302" data-name="晋城">晋城</li><li data-id="253" data-name="锦州">锦州</li><li data-id="66" data-name="嘉峪关">嘉峪关</li><li data-id="84" data-name="揭阳">揭阳</li><li data-id="185" data-name="荆门">荆门</li><li data-id="152" data-name="济源">济源</li><li data-id="173" data-name="鸡西">鸡西</li><li data-id="174" data-name="佳木斯">佳木斯</li><li data-id="67" data-name="金昌">金昌</li></ul><h3 class="city-key-K">K</h3><ul class="sort-city clearfix"><li data-id="373" data-name="昆明">昆明</li><li data-id="154" data-name="开封">开封</li><li data-id="358" data-name="喀什">喀什</li><li data-id="359" data-name="克拉玛依">克拉玛依</li></ul><h3 class="city-key-L">L</h3><ul class="sort-city clearfix"><li data-id="374" data-name="丽江">丽江</li><li data-id="155" data-name="洛阳">洛阳</li><li data-id="330" data-name="乐山">乐山</li><li data-id="223" data-name="连云港">连云港</li><li data-id="346" data-name="拉萨">拉萨</li><li data-id="69" data-name="兰州">兰州</li><li data-id="291" data-name="临沂">临沂</li><li data-id="387" data-name="丽水">丽水</li><li data-id="144" data-name="廊坊">廊坊</li><li data-id="290" data-name="聊城">聊城</li><li data-id="46" data-name="六安">六安</li><li data-id="289" data-name="莱芜">莱芜</li><li data-id="331" data-name="凉山">凉山</li><li data-id="107" data-name="柳州">柳州</li><li data-id="55" data-name="龙岩">龙岩</li><li data-id="166" data-name="漯河">漯河</li><li data-id="304" data-name="临汾">临汾</li><li data-id="203" data-name="娄底">娄底</li><li data-id="130" data-name="陵水">陵水</li><li data-id="347" data-name="林芝">林芝</li><li data-id="342" data-name="泸州">泸州</li><li data-id="115" data-name="六盘水">六盘水</li><li data-id="106" data-name="来宾">来宾</li><li data-id="216" data-name="辽源">辽源</li><li data-id="254" data-name="辽阳">辽阳</li><li data-id="305" data-name="吕梁">吕梁</li></ul><h3 class="city-key-M">M</h3><ul class="sort-city clearfix"><li data-id="396" data-name="澳门">澳门</li><li data-id="333" data-name="绵阳">绵阳</li><li data-id="175" data-name="牡丹江">牡丹江</li><li data-id="47" data-name="马鞍山">马鞍山</li><li data-id="86" data-name="梅州">梅州</li><li data-id="85" data-name="茂名">茂名</li><li data-id="332" data-name="眉山">眉山</li></ul><h3 class="city-key-N">N</h3><ul class="sort-city clearfix"><li data-id="224" data-name="南京">南京</li><li data-id="388" data-name="宁波">宁波</li><li data-id="239" data-name="南昌">南昌</li><li data-id="108" data-name="南宁">南宁</li><li data-id="225" data-name="南通">南通</li><li data-id="56" data-name="南平">南平</li><li data-id="57" data-name="宁德">宁德</li><li data-id="334" data-name="南充">南充</li><li data-id="156" data-name="南阳">南阳</li><li data-id="335" data-name="内江">内江</li><li data-id="376" data-name="怒江">怒江</li></ul><h3 class="city-key-P">P</h3><ul class="sort-city clearfix"><li data-id="58" data-name="莆田">莆田</li><li data-id="336" data-name="攀枝花">攀枝花</li><li data-id="167" data-name="濮阳">濮阳</li><li data-id="157" data-name="平顶山">平顶山</li><li data-id="240" data-name="萍乡">萍乡</li><li data-id="255" data-name="盘锦">盘锦</li><li data-id="378" data-name="普洱">普洱</li><li data-id="72" data-name="平凉">平凉</li></ul><h3 class="city-key-Q">Q</h3><ul class="sort-city clearfix"><li data-id="292" data-name="青岛">青岛</li><li data-id="59" data-name="泉州">泉州</li><li data-id="145" data-name="秦皇岛">秦皇岛</li><li data-id="131" data-name="琼海">琼海</li><li data-id="393" data-name="衢州">衢州</li><li data-id="117" data-name="黔南">黔南</li><li data-id="87" data-name="清远">清远</li><li data-id="116" data-name="黔东南">黔东南</li><li data-id="177" data-name="齐齐哈尔">齐齐哈尔</li><li data-id="377" data-name="曲靖">曲靖</li><li data-id="118" data-name="黔西南">黔西南</li><li data-id="109" data-name="钦州">钦州</li><li data-id="187" data-name="潜江">潜江</li><li data-id="73" data-name="庆阳">庆阳</li><li data-id="176" data-name="七台河">七台河</li><li data-id="132" data-name="琼中">琼中</li></ul><h3 class="city-key-R">R</h3><ul class="sort-city clearfix"><li data-id="293" data-name="日照">日照</li><li data-id="349" data-name="日喀则">日喀则</li></ul><h3 class="city-key-S">S</h3><ul class="sort-city clearfix"><li data-id="321" data-name="上海">上海</li><li data-id="226" data-name="苏州">苏州</li><li data-id="91" data-name="深圳">深圳</li><li data-id="133" data-name="三亚">三亚</li><li data-id="256" data-name="沈阳">沈阳</li><li data-id="146" data-name="石家庄">石家庄</li><li data-id="389" data-name="绍兴">绍兴</li><li data-id="241" data-name="上饶">上饶</li><li data-id="88" data-name="汕头">汕头</li><li data-id="90" data-name="韶关">韶关</li><li data-id="189" data-name="十堰">十堰</li><li data-id="227" data-name="宿迁">宿迁</li><li data-id="337" data-name="遂宁">遂宁</li><li data-id="60" data-name="三明">三明</li><li data-id="190" data-name="随州">随州</li><li data-id="218" data-name="松原">松原</li><li data-id="158" data-name="三门峡">三门峡</li><li data-id="188" data-name="神农架">神农架</li><li data-id="204" data-name="邵阳">邵阳</li><li data-id="48" data-name="宿州">宿州</li><li data-id="179" data-name="绥化">绥化</li><li data-id="159" data-name="商丘">商丘</li><li data-id="89" data-name="汕尾">汕尾</li><li data-id="217" data-name="四平">四平</li><li data-id="178" data-name="双鸭山">双鸭山</li><li data-id="306" data-name="朔州">朔州</li><li data-id="314" data-name="商洛">商洛</li><li data-id="350" data-name="山南">山南</li><li data-id="272" data-name="石嘴山">石嘴山</li></ul><h3 class="city-key-T">T</h3><ul class="sort-city clearfix"><li data-id="390" data-name="台州">台州</li><li data-id="343" data-name="天津">天津</li><li data-id="307" data-name="太原">太原</li><li data-id="294" data-name="泰安">泰安</li><li data-id="228" data-name="泰州">泰州</li><li data-id="147" data-name="唐山">唐山</li><li data-id="219" data-name="通化">通化</li><li data-id="74" data-name="天水">天水</li><li data-id="49" data-name="铜陵">铜陵</li><li data-id="266" data-name="通辽">通辽</li><li data-id="119" data-name="铜仁">铜仁</li><li data-id="257" data-name="铁岭">铁岭</li><li data-id="315" data-name="铜川">铜川</li><li data-id="363" data-name="吐鲁番">吐鲁番</li></ul><h3 class="city-key-W">W</h3><ul class="sort-city clearfix"><li data-id="391" data-name="温州">温州</li><li data-id="192" data-name="武汉">武汉</li><li data-id="229" data-name="无锡">无锡</li><li data-id="295" data-name="威海">威海</li><li data-id="364" data-name="乌鲁木齐">乌鲁木齐</li><li data-id="50" data-name="芜湖">芜湖</li><li data-id="296" data-name="潍坊">潍坊</li><li data-id="135" data-name="万宁">万宁</li><li data-id="110" data-name="梧州">梧州</li><li data-id="136" data-name="文昌">文昌</li><li data-id="316" data-name="渭南">渭南</li><li data-id="273" data-name="吴忠">吴忠</li><li data-id="267" data-name="乌海">乌海</li><li data-id="268" data-name="乌兰察布">乌兰察布</li><li data-id="137" data-name="五指山">五指山</li><li data-id="75" data-name="武威">武威</li><li data-id="379" data-name="文山">文山</li></ul><h3 class="city-key-X">X</h3><ul class="sort-city clearfix"><li data-id="317" data-name="西安">西安</li><li data-id="61" data-name="厦门">厦门</li><li data-id="230" data-name="徐州">徐州</li><li data-id="281" data-name="西宁">西宁</li><li data-id="206" data-name="湘西">湘西</li><li data-id="195" data-name="襄阳">襄阳</li><li data-id="380" data-name="西双版纳">西双版纳</li><li data-id="194" data-name="咸宁">咸宁</li><li data-id="205" data-name="湘潭">湘潭</li><li data-id="318" data-name="咸阳">咸阳</li><li data-id="51" data-name="宣城">宣城</li><li data-id="160" data-name="新乡">新乡</li><li data-id="269" data-name="锡林郭勒">锡林郭勒</li><li data-id="162" data-name="许昌">许昌</li><li data-id="308" data-name="忻州">忻州</li><li data-id="242" data-name="新余">新余</li><li data-id="196" data-name="孝感">孝感</li><li data-id="161" data-name="信阳">信阳</li><li data-id="4580" data-name="兴城">兴城</li><li data-id="148" data-name="邢台">邢台</li><li data-id="270" data-name="兴安盟">兴安盟</li><li data-id="193" data-name="仙桃">仙桃</li></ul><h3 class="city-key-Y">Y</h3><ul class="sort-city clearfix"><li data-id="232" data-name="扬州">扬州</li><li data-id="297" data-name="烟台">烟台</li><li data-id="274" data-name="银川">银川</li><li data-id="197" data-name="宜昌">宜昌</li><li data-id="231" data-name="盐城">盐城</li><li data-id="258" data-name="营口">营口</li><li data-id="310" data-name="运城">运城</li><li data-id="319" data-name="延安">延安</li><li data-id="209" data-name="岳阳">岳阳</li><li data-id="220" data-name="延边">延边</li><li data-id="207" data-name="益阳">益阳</li><li data-id="338" data-name="雅安">雅安</li><li data-id="320" data-name="榆林">榆林</li><li data-id="111" data-name="玉林">玉林</li><li data-id="244" data-name="鹰潭">鹰潭</li><li data-id="92" data-name="阳江">阳江</li><li data-id="381" data-name="玉溪">玉溪</li><li data-id="309" data-name="阳泉">阳泉</li><li data-id="243" data-name="宜春">宜春</li><li data-id="208" data-name="永州">永州</li><li data-id="339" data-name="宜宾">宜宾</li><li data-id="93" data-name="云浮">云浮</li><li data-id="180" data-name="伊春">伊春</li><li data-id="366" data-name="伊犁">伊犁</li></ul><h3 class="city-key-Z">Z</h3><ul class="sort-city clearfix"><li data-id="163" data-name="郑州">郑州</li><li data-id="97" data-name="珠海">珠海</li><li data-id="392" data-name="舟山">舟山</li><li data-id="210" data-name="张家界">张家界</li><li data-id="96" data-name="中山">中山</li><li data-id="233" data-name="镇江">镇江</li><li data-id="299" data-name="淄博">淄博</li><li data-id="62" data-name="漳州">漳州</li><li data-id="211" data-name="株洲">株洲</li><li data-id="95" data-name="肇庆">肇庆</li><li data-id="94" data-name="湛江">湛江</li><li data-id="120" data-name="遵义">遵义</li><li data-id="149" data-name="张家口">张家口</li><li data-id="3105" data-name="中卫">中卫</li><li data-id="298" data-name="枣庄">枣庄</li><li data-id="165" data-name="驻马店">驻马店</li><li data-id="76" data-name="张掖">张掖</li><li data-id="340" data-name="资阳">资阳</li><li data-id="341" data-name="自贡">自贡</li><li data-id="164" data-name="周口">周口</li>`
	//	harr := strings.Split(tl, "</h3>")
	regexStr := "<li[\\S\\s]+?\\</li\\>"
	reg := regexp.MustCompile(regexStr)

	for _, s := range reg.FindAllString(tl, -1) {
		datas := s[strings.Index(s, `"`)+1:]
		dataId := datas[:strings.Index(datas, `"`)]
		names := datas[strings.Index(datas, `"`)+1:]
		name := names[strings.Index(names, `>`)+1 : strings.Index(names, `<`)]
		print(dataId)
		print(" ")
		print(name)
		readarea("http://www.ly.com/SearchList.aspx?txtCityId=" + dataId)
		println(" ")
	}

}

func ReadHtmlSegment() {
	//	db.MustConnectDB()
	tl := `<h3 class="city-key-A">A</h3><ul class="sort-city clearfix"><li data-id="36" data-name="安庆">安庆</li><li data-id="150" data-name="安阳">安阳</li><li data-id="322" data-name="阿坝">阿坝</li><li data-id="245" data-name="鞍山">鞍山</li><li data-id="112" data-name="安顺">安顺</li><li data-id="311" data-name="安康">安康</li><li data-id="351" data-name="阿克苏">阿克苏</li><li data-id="3114" data-name="阿勒泰">阿勒泰</li></ul><h3 class="city-key-B">B</h3><ul class="sort-city clearfix"><li data-id="53" data-name="北京">北京</li><li data-id="99" data-name="北海">北海</li><li data-id="261" data-name="包头">包头</li><li data-id="139" data-name="保定">保定</li><li data-id="37" data-name="蚌埠">蚌埠</li><li data-id="283" data-name="滨州">滨州</li><li data-id="312" data-name="宝鸡">宝鸡</li><li data-id="367" data-name="保山">保山</li><li data-id="52" data-name="亳州">亳州</li><li data-id="98" data-name="百色">百色</li><li data-id="213" data-name="白山">白山</li><li data-id="246" data-name="本溪">本溪</li><li data-id="353" data-name="巴音郭楞">巴音郭楞</li><li data-id="260" data-name="巴彦淖尔">巴彦淖尔</li><li data-id="323" data-name="巴中">巴中</li><li data-id="122" data-name="保亭">保亭</li><li data-id="113" data-name="毕节">毕节</li><li data-id="63" data-name="白银">白银</li></ul><h3 class="city-key-C">C</h3><ul class="sort-city clearfix"><li data-id="324" data-name="成都">成都</li><li data-id="199" data-name="长沙">长沙</li><li data-id="394" data-name="重庆">重庆</li><li data-id="221" data-name="常州">常州</li><li data-id="214" data-name="长春">长春</li><li data-id="4569" data-name="长白山">长白山</li><li data-id="39" data-name="池州">池州</li><li data-id="141" data-name="承德">承德</li><li data-id="198" data-name="常德">常德</li><li data-id="200" data-name="郴州">郴州</li><li data-id="300" data-name="长治">长治</li><li data-id="140" data-name="沧州">沧州</li><li data-id="40" data-name="滁州">滁州</li><li data-id="38" data-name="巢湖">巢湖</li><li data-id="77" data-name="潮州">潮州</li><li data-id="100" data-name="崇左">崇左</li><li data-id="368" data-name="楚雄">楚雄</li><li data-id="262" data-name="赤峰">赤峰</li><li data-id="247" data-name="朝阳">朝阳</li><li data-id="355" data-name="昌吉">昌吉</li><li data-id="345" data-name="昌都">昌都</li><li data-id="123" data-name="昌江">昌江</li><li data-id="124" data-name="澄迈">澄迈</li></ul><h3 class="city-key-D">D</h3><ul class="sort-city clearfix"><li data-id="248" data-name="大连">大连</li><li data-id="78" data-name="东莞">东莞</li><li data-id="369" data-name="大理">大理</li><li data-id="301" data-name="大同">大同</li><li data-id="249" data-name="丹东">丹东</li><li data-id="285" data-name="东营">东营</li><li data-id="326" data-name="德阳">德阳</li><li data-id="168" data-name="大庆">大庆</li><li data-id="371" data-name="迪庆">迪庆</li><li data-id="284" data-name="德州">德州</li><li data-id="138" data-name="儋州">儋州</li><li data-id="370" data-name="德宏">德宏</li><li data-id="169" data-name="大兴安岭">大兴安岭</li><li data-id="125" data-name="定安">定安</li><li data-id="325" data-name="达州">达州</li></ul><h3 class="city-key-E">E</h3><ul class="sort-city clearfix"><li data-id="263" data-name="鄂尔多斯">鄂尔多斯</li><li data-id="182" data-name="恩施">恩施</li><li data-id="181" data-name="鄂州">鄂州</li></ul><h3 class="city-key-F">F</h3><ul class="sort-city clearfix"><li data-id="54" data-name="福州">福州</li><li data-id="79" data-name="佛山">佛山</li><li data-id="41" data-name="阜阳">阜阳</li><li data-id="101" data-name="防城港">防城港</li><li data-id="250" data-name="抚顺">抚顺</li><li data-id="234" data-name="抚州">抚州</li></ul><h3 class="city-key-G">G</h3><ul class="sort-city clearfix"><li data-id="80" data-name="广州">广州</li><li data-id="102" data-name="桂林">桂林</li><li data-id="114" data-name="贵阳">贵阳</li><li data-id="235" data-name="赣州">赣州</li><li data-id="327" data-name="甘孜">甘孜</li><li data-id="329" data-name="广元">广元</li><li data-id="328" data-name="广安">广安</li><li data-id="103" data-name="贵港">贵港</li><li data-id="271" data-name="固原">固原</li></ul><h3 class="city-key-H">H</h3><ul class="sort-city clearfix"><li data-id="383" data-name="杭州">杭州</li><li data-id="395" data-name="香港">香港</li><li data-id="170" data-name="哈尔滨">哈尔滨</li><li data-id="42" data-name="合肥">合肥</li><li data-id="45" data-name="黄山">黄山</li><li data-id="127" data-name="海口">海口</li><li data-id="264" data-name="呼和浩特">呼和浩特</li><li data-id="222" data-name="淮安">淮安</li><li data-id="82" data-name="惠州">惠州</li><li data-id="384" data-name="湖州">湖州</li><li data-id="265" data-name="呼伦贝尔">呼伦贝尔</li><li data-id="201" data-name="衡阳">衡阳</li><li data-id="286" data-name="菏泽">菏泽</li><li data-id="142" data-name="邯郸">邯郸</li><li data-id="81" data-name="河源">河源</li><li data-id="313" data-name="汉中">汉中</li><li data-id="202" data-name="怀化">怀化</li><li data-id="252" data-name="葫芦岛">葫芦岛</li><li data-id="143" data-name="衡水">衡水</li><li data-id="44" data-name="淮南">淮南</li><li data-id="184" data-name="黄石">黄石</li><li data-id="183" data-name="黄冈">黄冈</li><li data-id="279" data-name="海西">海西</li><li data-id="171" data-name="鹤岗">鹤岗</li><li data-id="172" data-name="黑河">黑河</li><li data-id="104" data-name="河池">河池</li><li data-id="43" data-name="淮北">淮北</li><li data-id="105" data-name="贺州">贺州</li><li data-id="151" data-name="鹤壁">鹤壁</li><li data-id="357" data-name="和田">和田</li><li data-id="278" data-name="海南藏族">海南藏族</li><li data-id="372" data-name="红河">红河</li></ul><h3 class="city-key-J">J</h3><ul class="sort-city clearfix"><li data-id="287" data-name="济南">济南</li><li data-id="385" data-name="嘉兴">嘉兴</li><li data-id="386" data-name="金华">金华</li><li data-id="238" data-name="九江">九江</li><li data-id="303" data-name="晋中">晋中</li><li data-id="288" data-name="济宁">济宁</li><li data-id="83" data-name="江门">江门</li><li data-id="236" data-name="吉安">吉安</li><li data-id="215" data-name="吉林">吉林</li><li data-id="68" data-name="酒泉">酒泉</li><li data-id="237" data-name="景德镇">景德镇</li><li data-id="186" data-name="荆州">荆州</li><li data-id="153" data-name="焦作">焦作</li><li data-id="302" data-name="晋城">晋城</li><li data-id="253" data-name="锦州">锦州</li><li data-id="66" data-name="嘉峪关">嘉峪关</li><li data-id="84" data-name="揭阳">揭阳</li><li data-id="185" data-name="荆门">荆门</li><li data-id="152" data-name="济源">济源</li><li data-id="173" data-name="鸡西">鸡西</li><li data-id="174" data-name="佳木斯">佳木斯</li><li data-id="67" data-name="金昌">金昌</li></ul><h3 class="city-key-K">K</h3><ul class="sort-city clearfix"><li data-id="373" data-name="昆明">昆明</li><li data-id="154" data-name="开封">开封</li><li data-id="358" data-name="喀什">喀什</li><li data-id="359" data-name="克拉玛依">克拉玛依</li></ul><h3 class="city-key-L">L</h3><ul class="sort-city clearfix"><li data-id="374" data-name="丽江">丽江</li><li data-id="155" data-name="洛阳">洛阳</li><li data-id="330" data-name="乐山">乐山</li><li data-id="223" data-name="连云港">连云港</li><li data-id="346" data-name="拉萨">拉萨</li><li data-id="69" data-name="兰州">兰州</li><li data-id="291" data-name="临沂">临沂</li><li data-id="387" data-name="丽水">丽水</li><li data-id="144" data-name="廊坊">廊坊</li><li data-id="290" data-name="聊城">聊城</li><li data-id="46" data-name="六安">六安</li><li data-id="289" data-name="莱芜">莱芜</li><li data-id="331" data-name="凉山">凉山</li><li data-id="107" data-name="柳州">柳州</li><li data-id="55" data-name="龙岩">龙岩</li><li data-id="166" data-name="漯河">漯河</li><li data-id="304" data-name="临汾">临汾</li><li data-id="203" data-name="娄底">娄底</li><li data-id="130" data-name="陵水">陵水</li><li data-id="347" data-name="林芝">林芝</li><li data-id="342" data-name="泸州">泸州</li><li data-id="115" data-name="六盘水">六盘水</li><li data-id="106" data-name="来宾">来宾</li><li data-id="216" data-name="辽源">辽源</li><li data-id="254" data-name="辽阳">辽阳</li><li data-id="305" data-name="吕梁">吕梁</li></ul><h3 class="city-key-M">M</h3><ul class="sort-city clearfix"><li data-id="396" data-name="澳门">澳门</li><li data-id="333" data-name="绵阳">绵阳</li><li data-id="175" data-name="牡丹江">牡丹江</li><li data-id="47" data-name="马鞍山">马鞍山</li><li data-id="86" data-name="梅州">梅州</li><li data-id="85" data-name="茂名">茂名</li><li data-id="332" data-name="眉山">眉山</li></ul><h3 class="city-key-N">N</h3><ul class="sort-city clearfix"><li data-id="224" data-name="南京">南京</li><li data-id="388" data-name="宁波">宁波</li><li data-id="239" data-name="南昌">南昌</li><li data-id="108" data-name="南宁">南宁</li><li data-id="225" data-name="南通">南通</li><li data-id="56" data-name="南平">南平</li><li data-id="57" data-name="宁德">宁德</li><li data-id="334" data-name="南充">南充</li><li data-id="156" data-name="南阳">南阳</li><li data-id="335" data-name="内江">内江</li><li data-id="376" data-name="怒江">怒江</li></ul><h3 class="city-key-P">P</h3><ul class="sort-city clearfix"><li data-id="58" data-name="莆田">莆田</li><li data-id="336" data-name="攀枝花">攀枝花</li><li data-id="167" data-name="濮阳">濮阳</li><li data-id="157" data-name="平顶山">平顶山</li><li data-id="240" data-name="萍乡">萍乡</li><li data-id="255" data-name="盘锦">盘锦</li><li data-id="378" data-name="普洱">普洱</li><li data-id="72" data-name="平凉">平凉</li></ul><h3 class="city-key-Q">Q</h3><ul class="sort-city clearfix"><li data-id="292" data-name="青岛">青岛</li><li data-id="59" data-name="泉州">泉州</li><li data-id="145" data-name="秦皇岛">秦皇岛</li><li data-id="131" data-name="琼海">琼海</li><li data-id="393" data-name="衢州">衢州</li><li data-id="117" data-name="黔南">黔南</li><li data-id="87" data-name="清远">清远</li><li data-id="116" data-name="黔东南">黔东南</li><li data-id="177" data-name="齐齐哈尔">齐齐哈尔</li><li data-id="377" data-name="曲靖">曲靖</li><li data-id="118" data-name="黔西南">黔西南</li><li data-id="109" data-name="钦州">钦州</li><li data-id="187" data-name="潜江">潜江</li><li data-id="73" data-name="庆阳">庆阳</li><li data-id="176" data-name="七台河">七台河</li><li data-id="132" data-name="琼中">琼中</li></ul><h3 class="city-key-R">R</h3><ul class="sort-city clearfix"><li data-id="293" data-name="日照">日照</li><li data-id="349" data-name="日喀则">日喀则</li></ul><h3 class="city-key-S">S</h3><ul class="sort-city clearfix"><li data-id="321" data-name="上海">上海</li><li data-id="226" data-name="苏州">苏州</li><li data-id="91" data-name="深圳">深圳</li><li data-id="133" data-name="三亚">三亚</li><li data-id="256" data-name="沈阳">沈阳</li><li data-id="146" data-name="石家庄">石家庄</li><li data-id="389" data-name="绍兴">绍兴</li><li data-id="241" data-name="上饶">上饶</li><li data-id="88" data-name="汕头">汕头</li><li data-id="90" data-name="韶关">韶关</li><li data-id="189" data-name="十堰">十堰</li><li data-id="227" data-name="宿迁">宿迁</li><li data-id="337" data-name="遂宁">遂宁</li><li data-id="60" data-name="三明">三明</li><li data-id="190" data-name="随州">随州</li><li data-id="218" data-name="松原">松原</li><li data-id="158" data-name="三门峡">三门峡</li><li data-id="188" data-name="神农架">神农架</li><li data-id="204" data-name="邵阳">邵阳</li><li data-id="48" data-name="宿州">宿州</li><li data-id="179" data-name="绥化">绥化</li><li data-id="159" data-name="商丘">商丘</li><li data-id="89" data-name="汕尾">汕尾</li><li data-id="217" data-name="四平">四平</li><li data-id="178" data-name="双鸭山">双鸭山</li><li data-id="306" data-name="朔州">朔州</li><li data-id="314" data-name="商洛">商洛</li><li data-id="350" data-name="山南">山南</li><li data-id="272" data-name="石嘴山">石嘴山</li></ul><h3 class="city-key-T">T</h3><ul class="sort-city clearfix"><li data-id="390" data-name="台州">台州</li><li data-id="343" data-name="天津">天津</li><li data-id="307" data-name="太原">太原</li><li data-id="294" data-name="泰安">泰安</li><li data-id="228" data-name="泰州">泰州</li><li data-id="147" data-name="唐山">唐山</li><li data-id="219" data-name="通化">通化</li><li data-id="74" data-name="天水">天水</li><li data-id="49" data-name="铜陵">铜陵</li><li data-id="266" data-name="通辽">通辽</li><li data-id="119" data-name="铜仁">铜仁</li><li data-id="257" data-name="铁岭">铁岭</li><li data-id="315" data-name="铜川">铜川</li><li data-id="363" data-name="吐鲁番">吐鲁番</li></ul><h3 class="city-key-W">W</h3><ul class="sort-city clearfix"><li data-id="391" data-name="温州">温州</li><li data-id="192" data-name="武汉">武汉</li><li data-id="229" data-name="无锡">无锡</li><li data-id="295" data-name="威海">威海</li><li data-id="364" data-name="乌鲁木齐">乌鲁木齐</li><li data-id="50" data-name="芜湖">芜湖</li><li data-id="296" data-name="潍坊">潍坊</li><li data-id="135" data-name="万宁">万宁</li><li data-id="110" data-name="梧州">梧州</li><li data-id="136" data-name="文昌">文昌</li><li data-id="316" data-name="渭南">渭南</li><li data-id="273" data-name="吴忠">吴忠</li><li data-id="267" data-name="乌海">乌海</li><li data-id="268" data-name="乌兰察布">乌兰察布</li><li data-id="137" data-name="五指山">五指山</li><li data-id="75" data-name="武威">武威</li><li data-id="379" data-name="文山">文山</li></ul><h3 class="city-key-X">X</h3><ul class="sort-city clearfix"><li data-id="317" data-name="西安">西安</li><li data-id="61" data-name="厦门">厦门</li><li data-id="230" data-name="徐州">徐州</li><li data-id="281" data-name="西宁">西宁</li><li data-id="206" data-name="湘西">湘西</li><li data-id="195" data-name="襄阳">襄阳</li><li data-id="380" data-name="西双版纳">西双版纳</li><li data-id="194" data-name="咸宁">咸宁</li><li data-id="205" data-name="湘潭">湘潭</li><li data-id="318" data-name="咸阳">咸阳</li><li data-id="51" data-name="宣城">宣城</li><li data-id="160" data-name="新乡">新乡</li><li data-id="269" data-name="锡林郭勒">锡林郭勒</li><li data-id="162" data-name="许昌">许昌</li><li data-id="308" data-name="忻州">忻州</li><li data-id="242" data-name="新余">新余</li><li data-id="196" data-name="孝感">孝感</li><li data-id="161" data-name="信阳">信阳</li><li data-id="4580" data-name="兴城">兴城</li><li data-id="148" data-name="邢台">邢台</li><li data-id="270" data-name="兴安盟">兴安盟</li><li data-id="193" data-name="仙桃">仙桃</li></ul><h3 class="city-key-Y">Y</h3><ul class="sort-city clearfix"><li data-id="232" data-name="扬州">扬州</li><li data-id="297" data-name="烟台">烟台</li><li data-id="274" data-name="银川">银川</li><li data-id="197" data-name="宜昌">宜昌</li><li data-id="231" data-name="盐城">盐城</li><li data-id="258" data-name="营口">营口</li><li data-id="310" data-name="运城">运城</li><li data-id="319" data-name="延安">延安</li><li data-id="209" data-name="岳阳">岳阳</li><li data-id="220" data-name="延边">延边</li><li data-id="207" data-name="益阳">益阳</li><li data-id="338" data-name="雅安">雅安</li><li data-id="320" data-name="榆林">榆林</li><li data-id="111" data-name="玉林">玉林</li><li data-id="244" data-name="鹰潭">鹰潭</li><li data-id="92" data-name="阳江">阳江</li><li data-id="381" data-name="玉溪">玉溪</li><li data-id="309" data-name="阳泉">阳泉</li><li data-id="243" data-name="宜春">宜春</li><li data-id="208" data-name="永州">永州</li><li data-id="339" data-name="宜宾">宜宾</li><li data-id="93" data-name="云浮">云浮</li><li data-id="180" data-name="伊春">伊春</li><li data-id="366" data-name="伊犁">伊犁</li></ul><h3 class="city-key-Z">Z</h3><ul class="sort-city clearfix"><li data-id="163" data-name="郑州">郑州</li><li data-id="97" data-name="珠海">珠海</li><li data-id="392" data-name="舟山">舟山</li><li data-id="210" data-name="张家界">张家界</li><li data-id="96" data-name="中山">中山</li><li data-id="233" data-name="镇江">镇江</li><li data-id="299" data-name="淄博">淄博</li><li data-id="62" data-name="漳州">漳州</li><li data-id="211" data-name="株洲">株洲</li><li data-id="95" data-name="肇庆">肇庆</li><li data-id="94" data-name="湛江">湛江</li><li data-id="120" data-name="遵义">遵义</li><li data-id="149" data-name="张家口">张家口</li><li data-id="3105" data-name="中卫">中卫</li><li data-id="298" data-name="枣庄">枣庄</li><li data-id="165" data-name="驻马店">驻马店</li><li data-id="76" data-name="张掖">张掖</li><li data-id="340" data-name="资阳">资阳</li><li data-id="341" data-name="自贡">自贡</li><li data-id="164" data-name="周口">周口</li>`
	harr := strings.Split(tl, "</h3>")
	regexStr := "<li[\\S\\s]+?\\</li\\>"
	reg := regexp.MustCompile(regexStr)

	for _, s := range reg.FindAllString(tl, -1) {
		datas := s[strings.Index(s, `"`)+1:]
		dataId := datas[:strings.Index(datas, `"`)]
		names := datas[strings.Index(datas, `"`)+1:]
		name := names[strings.Index(names, `>`)+1 : strings.Index(names, `<`)]
		println(dataId, name)
		//		db.InsertCiyt(dataId, name)

	}
	i := 0
	for _, h := range harr {
		liarr := strings.Split(h, "</li>")
		i = len(liarr) + i
	}
	println(i)

}
