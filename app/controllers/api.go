package controllers

import (
	// "encoding/json"
	// "fantastic/app/models"
	// "bytes"
	"fmt"
	// . "github.com/jbrukh/bayesian"
	"github.com/jgraham909/revmgo"
	"github.com/revel/revel"
	// "io/ioutil"
	// "os"
	// "strconv"
	// "strings"

	// "strings"
	"crypto/sha1"
	// "encoding/xml"
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

type msgBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}

type Badge struct {
	Master int `json:"master"`
	Note   int `json:"note"`
}

func (c *Api) Index(signature, timestamp, nonce, echostr string) revel.Result {
	revel.INFO.Println(signature)
	revel.INFO.Println(timestamp)
	revel.INFO.Println(nonce)
	revel.INFO.Println(echostr)
	revel.INFO.Println(signatureSha1(timestamp, nonce))
	if signatureSha1(timestamp, nonce) == signature {
		return c.RenderText(echostr)
	} else {
		return c.RenderText("false")
	}
}

func (c *Api) Post() revel.Result {
	return c.RenderText("tf")
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
