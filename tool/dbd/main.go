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
	id       = "200590046"
	name     = "MagicWatch2"
	maxPrice = 900
	addPrice = 5 // 加价间隔
	cookie   = "shshshfpa=b8cc7f40-5852-c878-2008-202cb7b564aa-1576751256; shshshfpb=b8cc7f40-5852-c878-2008-202cb7b564aa-1576751256; __jdv=104464258|direct|-|none|-|1578034862589; mba_muid=1576642527839566679144; 3AB9D23F7A4B3C9B=MX3IRQ3GTOJTZWGF3WHRPHINBUVMCNIWU5CQR7J25IYYFJO5JJ57QBXR3KLXSC56VGZEDIZ3AURLP5S44YW3HXVQYY; shshshfp=e3c1b507580913054efdb15be6c61cc9; areaId=16; ipLoc-djd=16-1303-3484-0; __jdc=104464258; __tak=3abcee8b2e24588670bf5a131f4c8892fd4ad34146e8732717e121ab22b10584feb8b824eb5722adce847eb15324814d875b636e3659ff2a9ec56ee11664ad9f6b25b8704d6586cb96c41341926130c3; __jda=104464258.1576642527839566679144.1576642528.1579049397.1579054173.20; thor=A6CD0F6C27D13A138EBD4D18B8A97E3C62A6F403197B591A1AFFF8FC4B42DF94247A82A6D78E88DC9F90945E40FCB265C9D16B4753A3C6F31059A5F85E04EBC60C672D676A204094FAA4D23C50B624B6660CC7C216544E9714BFB963BD0AA561B4879696FF97E6A39237F80B18FDB1A666910DD82DD0C4703575F2C1BBF0E3B8; pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; __jdb=104464258.15.1576642527839566679144|20.1579054173"
	token    = "ak8p224mjwbtbxwuwbm15790493966694zcg~NmZeSyVEbFNSdHB0dFRaCnlxAw1mRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaIx8/FV4dZWEYQwtub35rGmBRYUMCcCV0cAVbC3h1Bgs3X24QVnkhLiVUWAB9cgYOZgA7FD9jaxFmCB5fEWYNZHMANx0QJBtvaD1PWj4waxprOnQCBi0rYzQABEIsLRlbPgsKU08dZT0qPU8IEWYYWSQFIhgML2opIRUMWyFrBQhhU29BU3VxEWZNMRA9MGsaazp0GRc1Nz5+PU8eEWZHUQ1EbC1Bc3VhfE1bHn1oBxRlXwpTHmNrbyEFCUEqZg0aaFA3SEFtZS43Q1cQeHEATjkQIAIUdSMsLxUbWXgsA1lhUTVGVHkvK3wFWgB+K0NJc0p0EkF7ZSsgWB5RI31ECjpTJwIOOSR/M1NcB3pyAQxjU2RGVnlzeykSBxBhZlNLP0RsUwcoITdwUhVeb2gVUSJEbFNSY2tvLggOEHdmDgFoXnQM|~1579057846361~1~20191203~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIxIiwibGUiOiIxMDAiLCJjdCI6IjAiLCJkdCI6ImkifX0=~1~~30v4|1d6w-hb,8b;doei:,1,1,0,0,1,1000,-1000,1000,-1000;dmei:,1,1,1,1000,-1000,1000,-1000,1000,-1000;emc:,d:1;emcf:,d:1;ivli:;iivl:;ivcvj:;scvje:;1579057845453,1579057846357,0,0,1,1,0,0,0,0,0;uk67"
)

func init() {
}

func main() {
	isSleep := true
	sleepTime := time.Second * 3
	for {
		if isSleep {
			time.Sleep(sleepTime)
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
		if remainTime < 0 {
			fmt.Println("竞拍结束")
			break
		}
		if ret.CurrentBidder == userId {
			//fmt.Println("最高价为本人，不出价")
			continue
		}
		if remainTime < 10 {
			sleepTime = time.Second
		}
		if remainTime < 2 {
			isSleep = false
			price := ret.CurrentPrice + 3
			retCode := 0
			for price < maxPrice {
				retCode = bidPrice(price)
				if retCode == 302 {
					time.Sleep(time.Millisecond * 50) // 休眠1/50秒
					continue                          // 连续出价不抬价
				}
				price += 3
			}

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
		//fmt.Println(err.Error())
		return -1
	}
	defer resp.Body.Close()
	retByte, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("resp:", string(retByte))
	ret := &Ret{}
	err = json.Unmarshal(retByte, ret)
	if err != nil {
		//fmt.Println(err)
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
