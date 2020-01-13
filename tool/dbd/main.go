package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Ret struct {
	Code    int         `json:"code"`
	Data    *Data       `json:"data"`
	List    interface{} `json:"list"`
	Message string      `json:"message"`
}
type Data struct {
	AuctionId         int         `json:"auctionId"` // 商品id
	AuctionRecordId   int         `json:"auctionRecordId"`
	CurrentPrice      float64     `json:"currentPrice"` // 当前价格
	Num               int         `json:"num"`
	CurrentBidder     string      `json:"currentBidder"` // 当前最高价用户
	BidderNickName    string      `json:"bidderNickName"`
	BidderImage       interface{} `json:"bidderImage"`
	Status            interface{} `json:"status"`
	OfferPrice        interface{} `json:"offerPrice"`
	ActualEndTime     int64       `json:"actualEndTime"` // 结束时间
	DelayCount        int         `json:"delayCount"`
	VirtualDelayCount int         `json:"virtualDelayCount"`
	SpectatorCount    interface{} `json:"spectatorCount"`
	UserNum           interface{} `json:"userNum"`
}

const (
	userId   = "t***w"
	id       = "200445993"
	name     = "MagicWatch2"
	maxPrice = 999
	addPrice = 5 // 加价间隔
	cookie   = "shshshfp=8b3cbd79d85c35055d31fdaaad749f33; shshshfpa=b8cc7f40-5852-c878-2008-202cb7b564aa-1576751256; shshshfpb=b8cc7f40-5852-c878-2008-202cb7b564aa-1576751256; __jdv=104464258|direct|-|none|-|1578034862589; mba_muid=1576642527839566679144; 3AB9D23F7A4B3C9B=MX3IRQ3GTOJTZWGF3WHRPHINBUVMCNIWU5CQR7J25IYYFJO5JJ57QBXR3KLXSC56VGZEDIZ3AURLP5S44YW3HXVQYY; __jda=104464258.1576642527839566679144.1576642528.1578736734.1578874542.15; __jdc=104464258; __tak=048eff3fa2b5f384a24707749124ef65611095ba2ce80301f3aa7af39d0f078071d25e90b9ff8999dd043f5ebb21bfbaf0dfc73ef9ba5c4816f6bec328292952b7528fd77aeda3ed59d7238503b7c1e6; thor=A6CD0F6C27D13A138EBD4D18B8A97E3C4500B17A8DE81BCF36D68E25785D899098B1DCDA754E8074C88D2343AD720CE2210460179B215B95FEF5FA1315F61ADFB84F7FDE8709545B3752B2D46698B534AEEF82BE8EA439399ECC1FF8350F70F75A49469AAA9F119A2D2F4E317F1EC9B232A7DB34257DCCF7C01BD9EF2934CCC9; pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; __jdb=104464258.4.1576642527839566679144|15.1578874542"
	token    = "agq7n3usqez520p99oc1578874541099bwk0~NmZeSyVEbFNSdHB1fFZZBHtwAwFhRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaDUR6LUEnG29+PU8DenMFWWAEb0UHd354dVBeVHR8Vg1pADUQVnR1fXJQWwUrKVJkc0oKUwoyKhFmWzEQKyVbSzQ6dF0/Yy8+MD1PCBFmRF09CngBAig3LC1PDl0gGBUUDUQmHz9jfRFmTgxHLjBeVz9LMhQXIC4ha1NdAnlwBQ5lUApTTx1lPTA9TwgRZl9MJRYlSz9jaxFmEQRub35rGmNUeklPd2t9aFFBBnQYFUVzSnQUByU0KmZbTwY1cAMafUQ1AkF7ZSh2VQlIPTRPAD8Xb0YMIis4PFZfVD8mTlQzBDUIUy13dS5VGhBhZlQaa0QwFVoyJCN9El9ZeDVEVSkFZAZRcHJ6cldZAHh2AA1pUmAcECtlYWYFHlxvfhVUOAA7ElInLG9oQwRBb34VCXNKdBsKImV3ZlhUC3VmSg==|~1578874664955~1~20191203~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIxIiwibGUiOiIxMDAiLCJjdCI6IjAiLCJkdCI6ImkifX0=~1~~00v4|1d25-cq,21;1d9-er,2m;1d11-hl,3f;1d4w-j4,3t;doei:,1,1,0,0,1,1000,-1000,1000,-1000;dmei:,1,1,1,1000,-1000,1000,-1000,1000,-1000;emc:,d:5;emcf:,d:5;ivli:;iivl:;ivcvj:;scvje:;1578874664487,1578874664953,0,0,4,4,0,1,0,0,0;jh2u"
)

func init() {
}

func main() {
	isSleep := true
	for {
		if isSleep {
			time.Sleep(time.Second * 2)
		}

		ret, err := getPrice()
		if err != nil {
			continue
		}
		remainTime := (ret.ActualEndTime - time.Now().UnixNano()/1000000) / 1000
		fmt.Printf("No:%v     当前价格:%-6v出价者:%-10v剩余时间:%-8v最高价格:%-10v\n", id+"-"+name, ret.CurrentPrice, ret.CurrentBidder, remainTime, maxPrice)

		if ret.CurrentPrice > maxPrice {
			fmt.Println("超出最高价，退出竞拍")
			break
		}
		if ret.CurrentBidder == userId {
			//fmt.Println("最高价为本人，不出价")
			continue
		}
		if remainTime < 0 {
			fmt.Println("竞拍结束")
			break
		}
		if remainTime < 5 {
			isSleep = false
			bidPrice(ret.CurrentPrice + addPrice)
		} else {
			//bidPrice(ret.CurrentPrice + addPrice)
		}
	}
	if ret, err := getPrice(); err != nil {
		fmt.Println("次轮竞拍失败")
	} else {
		remainTime := (ret.ActualEndTime - time.Now().UnixNano()/1000000) / 1000
		fmt.Printf("No:%v     当前价格:%-6v出价者:%-10v剩余时间:%-10v最高价格:%-10v\n", id+"-"+name, ret.CurrentPrice, ret.CurrentBidder, remainTime, maxPrice)
	}
}

func getPrice() (*Data, error) {
	apiUrl := "https://used-api.jd.com/auctionRecord/getCurrentAndOfferNum?auctionId=" + id //+ "&callback=__jp17"
	data := url.Values{}
	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()
	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Referer", "https://sell.paipai.com/auction-list/"+id)
	r.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	r.Header.Add("cookie", cookie)
	r.Header.Add("content-type", "application/x-www-form-urlencoded")
	//r.Header.Add("content-length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	retByte, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("resp:", string(retByte))
	ret := &Ret{}
	err = json.Unmarshal(retByte, ret)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if ret.Code != 200 {
		err := fmt.Errorf("%v:%v", ret.Code, ret.Message)
		fmt.Println(err)
		return nil, err
	}
	return ret.Data, nil
}

// code=302,message=同一用户不能连续出价
// code=304,message=客官，您出价太低了，再加点吧！
func bidPrice(price float64) int {
	apiUrl := "https://used-api.paipai.com/auctionRecord/offerPrice"
	data := url.Values{}
	data.Set("trackId", "0d1b9bc80e5339d06cd27454b25c4b66")
	data.Set("eid", "MX3IRQ3GTOJTZWGF3WHRPHINBUVMCNIWU5CQR7J25IYYFJO5JJ57QBXR3KLXSC56VGZEDIZ3AURLP5S44YW3HXVQYY")
	data.Set("auctionId", id)
	data.Set("token", token)
	data.Set("price", fmt.Sprintf("%v", price))
	data.Set("entryid", "")
	data.Set("address", "16-1303-48712-48759")
	data.Set("initFailed", "false")
	u, _ := url.ParseRequestURI(apiUrl)
	//u.Path = resource
	urlStr := u.String()

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Referer", "https://sell.paipai.com/auction-detail/"+id)
	r.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	r.Header.Add("cookie", cookie)
	r.Header.Add("content-type", "application/x-www-form-urlencoded")
	//r.Header.Add("content-length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	defer resp.Body.Close()
	retByte, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("resp:", string(retByte))
	ret := &Ret{}
	err = json.Unmarshal(retByte, ret)
	if err != nil {
		fmt.Println(err)
		return -1
	}
	if ret.Code != 200 {
		err := fmt.Errorf("%v:%v", ret.Code, ret.Message)
		fmt.Println(err)
		return ret.Code
	}
	fmt.Println("出价成功:", price)
	return ret.Code
}
