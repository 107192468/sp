package beans

type TC struct {
	Response Response `json:"response"`
}
type Response struct {
	Header Header `json:"header"`
	Body   Body   `json:"body"`
}

type PageInfo struct {
	PageSize   string `json:"pageSize"`
	Page       string `json:"page"`
	TotalPage  string `json:"totalPage"`
	TotalCount string `josn:"totalCount"`
}
type Body struct {
	IsCache             string            `json:"isCache"`
	PageInfo            PageInfo          `json:"pageInfo"`
	DianpingInfo        []string          `json:"dianpingInfo"`
	TotalNum            string            `json:"totalNum"`
	DegreeLevel         string            `json:"degreeLevel"`
	GoodNum             string            `json:"goodNum"`
	MidNum              string            `json:"midNum"`
	BadNum              string            `json:"badNum"`
	HasImgNum           string            `json:"hasImgNum"`
	StarNum             string            `json:"starNum"`
	UKeyList            []string          `json:"uKeyList"`
	IsShowConinsicon    string            `json:"isShowConinsicon"`
	ServiceScoreAvgList []ServiceScoreAvg `json:"serviceScoreAvgList"`
	DpList              []Dp              `json:"dpList"`
	RulesContent        string            `json:"rulesContent"`
	DpConinRemark       string            `json:"dpConinRemark"`
	IsCanEnter          string            `json:"isCanEnter"`
	DpTagList           []DbTag           `json:"dpTagList"`
	IsCanDianPingBefore string            `json:"isCanDianPingBefore"`
	ExternalUrl         string            `json:"externalUrl"`
}
type Header struct {
	RspType string `json:"rspType"`
	RspCode string `json:"rspCode"`
	RspDesc string `json:"rspDesc"`
	RspTime string `json:"rspTime"`
}
type Dp struct {
	DpImpressionList   []DpImpression `json:"dpImpressionList"`
	DpUserName         string         `json:"dpUserName"`
	MemberHeaderImgUrl string         `json:"memberHeaderImgUrl"`
	MemberGradeUrl     string         `json:"memberGradeUrl"`
	HomeId             string         `json:"homeId"`
	DpId               string         `json:"dpId"`
	DpGuid             string         `json:"dpGuid"`
	DpDate             string         `json:"dpDate"`
	LineAccess         string         `json:"lineAccess"`
	DpContent          string         `json:"dpContent"`
	DpTripPurpose      string         `json:"dpTripPurpose"`
	DpTripMode         string         `json:"dpTripMode"`
	ZanCount           string         `json:"zanCount"`
	IsPraised          string         `json:"isPraised"`
	IsMyComments       string         `json:"isMyComments"`
	DpImgUrl           []DpImg        `json:"dpImgUrl"`
	ServicePoint       string         `json:"servicePoint"` //1好评,2中评3,差评
	IsElite            string         `json:"isElite"`
	ProductPriceDesc   string         `json:"productPriceDesc"`
	HighLightList      []string       `json:"highLightList" `
	HighLightColor     string         `json:"highLightColor"`
	SubList            []string       `json:"subList"`
	PakList            []string       `json:"pakList"`
	CommentType        string         `json:"commentType"`
	DPItemId           string         `json:"DPItemId"`
	DPItemName         string         `json:"DPItemName"`
	DPUserLevel        string         `json:"DPUserLevel"`
	RoomTypeName       string         `json:"roomTypeName"`
	RoomTypeId         string         `json:"roomTypeId"`
	Mobile             string         `json:"mobile"`
	DPPrize_HongBao    string         `json:"DPPrize_HongBao"`
	DPPrize_JiangJin   string         `json:"DPPrize_JiangJin"`
	DPPrize_YouBi      string         `json:"DPPrize_YouBi"`
	DPSite             string         `json:"DPSite"`
	CsReplyList        []CsReply      `json:"csReplyList"`
}

type ServiceScoreAvg struct {
	ServiceName string `json:"serviceName"`
	Score       string `json:"score"`
}
type DbTag struct {
	TagId    string `json:"tagId"`
	TagName  string `json:"tagName"`
	TagNum   string `json:"tagNum"`
	TagColor string `json:"tagColor"`
	TagType  string `json:"tagType"`
	TagTest  string `json:"tagTest"`
}
type DpImpression struct {
}
type DpImg struct {
	imgId          string `json:"imgId"`
	imgItemId      string `json:"imgItemId"`
	ImgUrl         string `json:"imgUrl"`
	OriginalImgUrl string `json:"originalImgUrl"`
	SmallImgUrl    string `json:"smallImgUrl"`
}
type CsReply struct {
	replyContent string `json:"replyContent"`
}
type ReviewObjectVo struct {
	ReviewObjectId  string  //	评论对象ID
	ObjectTypeId    int64   //	对象类型ID[1:”商品”,2:”券”,3:”商户(门店)”,4:”票(乐园/秀场)”,5:”电影”]
	CityId          string  //	城市ID
	PlazaId         string  //	广场ID
	StoreId         string  //	门店ID
	CategoryId      string  //	类目ID
	ReviewNum       int     //	评论数
	AverageScore    float64 //	综合评分
	Score1Num       int32   //	评分1人数
	Score2Num       int32   //	评分2人数
	Score3Num       int32   //	评分3人数
	Score4Num       int32   //	评分4人数
	Score5Num       int32   //	评分5人数
	GoodReviewNum   int     //	好评数
	BadReviewNum    int     //	差评数
	HasPicReviewNum int     //	有图评论数
	LikeFlag        int32   //	喜欢标识[暂时无用]
	CreateTime      int64   //	创建时间
	UpdateTime      int64   //	修改时间
	Version         int32   //	版本号
	Remark          string  //	备注
}
type ReviewTagValueVo struct {
	ReviewId     string //	评论ID
	TagId        int64  //	属性ID
	TagValue     string //	属性值
	TagValueType int32  //	属性值类型	[1:”文本”, 2:”评分”,3:”子评分”,4:”标签”]
	Status       int32  //	评论状态	[1:”待审核”,2:”已通过”,3:”已驳回”]
	UpdateTime   int64  //修改时间
}
type Review struct {
	//ReviewId           string             //	评论ID
	ReviewObjectId    string //	评论对象ID
	ObjectTypeId      string //	评论对象类型ID	[1:”商品”,2:”券”,3:”商户(门店)”,4:”票(乐园/秀场)”,5:”电影”]
	CategoryId        string //	目录ID
	StoreId           string //	门店ID
	UserId            string //	用户ID
	CurrentRosterType string //	当前名单类型	[1:”黑名单”,2:”灰名单”,3:”白名单”]
	Ip                string //	IP地址
	Type_             string //	评论类型	[1:”好评”,2:”差评”]
	Source            string //	评论来源	[1:”WEB”,2:”WAP”,3:”Android”,4:”IOS”]
	Status            string //	评论状态	[1:”待审核”,2:”已通过”,3:”已驳回”]
	OrderFlag         string //	订单标识	[1:”线上”,2”线下”]
	OrderNo           string //	订单号
	PicFlag           string //	图片标识	[1:”有图”,2”无图”]
	PicUrls           string //	图片地址(多个地址用逗号隔开)
	LikeNum           string //	喜欢数(即点赞数)
	ReplyNum          string //	回复数
	Title             string //	评论标题
	Content           string //	评论内容
	IsValid           string //	是否有效	[1:”有效”,2:”无效”]
	IsAnonym          string //        是否匿名 1,2
	CreateTime        string //	创建时间
	UpdateTime        string //	修改时间
	Version           string //	版本号
	Remark            string //	备注
	//ReviewObject       ReviewObjectVo     //	评论对象
	//ReviewTagValueList []ReviewTagValueVo //	评论属性值列表
}
type ItemAverageScores struct {
	ReviewObjectVo string
	TagId          string
	TagName        string
	TagAlias       string
	AvergeScore    string
}
