package models
import (
  "encoding/xml"
  "time"
)

const (
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

type item struct {
  XMLName     xml.Name `xml:"item"`
  Title       string
  Description string
  PicUrl      string
  Url         string
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

func NewResponse() (resp *Response) {
  resp = &Response{}
  resp.CreateTime = time.Duration(time.Now().Unix())
  return
}
