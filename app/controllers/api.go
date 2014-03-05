package controllers

import (
	// "encoding/json"
	// "fantastic/app/models"
	// "bytes"
	"fmt"
	// . "github.com/jbrukh/bayesian"
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
	"io/ioutil"
	// "os"
	// "strconv"
	"strings"

	// "strings"
	"crypto/sha1"
	"encoding/xml"
	"sort"
	"time"
)

type Api struct {
	*revel.Controller
	revmgo.MongoController
}

const (
	TOKEN    = "weixin"
	Text     = "text"
	Location = "location"
	Image    = "image"
	Link     = "link"
	Event    = "event"
	Music    = "music"
	News     = "news"
)

type item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string
	Description string
	PicUrl      string
	Url         string
}

type msgBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

type Request struct {
	XMLName                xml.Name `xml:"xml"`
	msgBase                         // base struct
	Location_X, Location_Y float32
	Scale                  int
	Label                  string
	PicUrl                 string
	MsgId                  int
}

type Response struct {
	XMLName xml.Name `xml:"xml"`
	msgBase
	ArticleCount int     `xml:",omitempty"`
	Articles     []*item `xml:"Articles>item,omitempty"`
	FuncFlag     int
}

func (c *Api) Index(signature, timestamp, nonce, echostr string) revel.Result {
	if signatureSha1(timestamp, nonce) == signature {
		return c.RenderText(echostr)
	} else {
		fmt.Println("-------------")
		return c.RenderText("false")
	}
}

func DecodeRequest(data []byte) (req *Request, err error) {
	req = &Request{}
	err = xml.Unmarshal(data, req)
	req.CreateTime *= time.Second
	return req, err
}

func (c *Api) Post() revel.Result {
	defer c.Request.Body.Close()
	body, _:= ioutil.ReadAll(c.Request.Body)
	var wreq *Request
	wreq, _ = DecodeRequest(body)
	
	wresp, _ := dealwith(wreq)
	
	data, _ := wresp.Encode()

	fmt.Println(string(data))

	return c.RenderText(string(data))
}

func dealwithText(keyword string, resp *Response) {

	if keyword == "help" {
		resp.Content = "welcome!"
	} else {
		resp.Content = "亲，已经收到您的消息, 将尽快回复您."
	}
}

func dealwithEvent(req *Request, resp *Response) {
	if req.Content == "subscribe" {
		resp.Content = "welcome!"
	}
}

func dealwithImage(req *Request, resp *Response) {
	var a item
	a.Description = "雅蠛蝶。。。^_^^_^1024你懂的"
	a.Title = "雅蠛蝶图文测试"
	a.PicUrl = "http://static.yaliam.com/gwz.jpg"
	a.Url = "http://blog.csdn.net/songbohr"

	resp.MsgType = News
	resp.ArticleCount = 1
	resp.Articles = append(resp.Articles, &a)
	resp.FuncFlag = 1
}

func dealwith(req *Request) (resp *Response, err error) {
	resp = NewResponse()
	resp.ToUserName = req.FromUserName
	resp.FromUserName = req.ToUserName
	resp.MsgType = Text

	if req.MsgType == Event {
		dealwithEvent(req, resp)
	} else if req.MsgType == Text {
		keyword := strings.Trim(strings.ToLower(req.Content), " ")
		dealwithText(keyword, resp)
	} else if req.MsgType == Image {
		dealwithImage(req, resp)
	} else {
		resp.Content = "暂时还不支持其他的类型"
	}
	return resp, nil
}
func (c *Api) Put() revel.Result {
	return c.RenderText("tf")

}

func signatureSha1(timestamp, nonce string) string {
	strs := sort.StringSlice{TOKEN, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func NewResponse() (resp *Response) {
	resp = &Response{}
	resp.CreateTime = time.Duration(time.Now().Unix())
	return
}

func (resp Response) Encode() (data []byte, err error) {
	resp.CreateTime = time.Duration(time.Now().Unix())
	data, err = xml.Marshal(resp)
	return
}
