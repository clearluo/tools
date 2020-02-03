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
	AuctionId       int         `json:"auctionId"` // 商品id
	AuctionRecordId int         `json:"auctionRecordId"`
	CurrentPrice    float64     `json:"currentPrice"` // 当前价格
	Num             int         `json:"num"`
	CurrentBidder   string      `json:"currentBidder"` // 当前最高价用户
	BidderNickName  string      `json:"bidderNickName"`
	BidderImage     interface{} `json:"bidderImage"`
	Status          interface{} `json:"status"`
	OfferPrice      interface{} `json:"offerPrice"`
	ActualEndTime   int64       `json:"actualEndTime"` // 结束时间
	//DelayCount        int         `json:"delayCount"`
	//VirtualDelayCount int         `json:"virtualDelayCount"`
	//SpectatorCount    interface{} `json:"spectatorCount"`
	//UserNum           interface{} `json:"userNum"`
}

const (
	userId   = "t***w"
	id       = "201775014"
	name     = "MagicWatch2"
	maxPrice = 907
	addPrice = 3 // 加价间隔
	cookie   = "__jdc=104464258; __jdv=104464258|direct|-|none|-|1580090591192; __tak=a3be736a5c0b3fe7ca7b462cb347026fc0163058be5b215362489884e018f562661a3193aa20ff30615c1991ffcf4301eb2580927fbc8f42f0f978c7d88fb6e1f71192699604a35fa78914194748a0a7; pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; 3AB9D23F7A4B3C9B=ZYRNLXJSY4BFD7XORCJM4R53XEGVRVTM27WG4UKM6M2L6NBV6CVML6D2J5GCAT2QXQSOMLBNKY5LIOMNOE5JE2VDOU; __jda=104464258.15800905911911676925759.1580090591.1580610518.1580626242.4; thor=A6CD0F6C27D13A138EBD4D18B8A97E3CC69DC4C2A4C6FAFF527E3033B2D17DC0AD26C296E2D6A87B78C4CB00419B20DF723ADFCE992A797EBB573C5192B28419C9E40FB2CFE7C5AD7914C0E59B6292DA1E6443C42CB8BFDA31A2E9959D1113BDE02DC999FB348E23BB1BE40A4A978826A4C901C5D2E8A782D1173B678ED68537; __jdb=104464258.4.15800905911911676925759|4.1580626242"
	token    = "jdp1060uuv2r6e4prjy1580090590560yady~NmZeSyVEbFNSdH99clNbAHp3AAtmRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaIx8/FV4dZWEYQwtub35rGjICZRVWdHZ9JwRZBHxzAw9jUTRHUSIhfiFZDwcrdQ9baV4yAD9jaxFmCB5fEWYNZHMANx0QJBtvaD1PWj4waxprOnQCBi0rYzQABEIsLRlbPgsKU08dZT0qPU8IEWYYWSQFIhgML2opIRUMWyFrBQhgUWFEU3BzEWZNMRA9MGsaazp0GRc1Nz5+PU8eEWZHUQ1EbC1Bc3Vhc01bHn1oBxRlXwpTHmNrbyEFCUEqZg0aaBA3SEFtZS43Q1cQPD5VQT8RLhsIdHJ7LBQDXygnQwwkEjRBDTQwdHwKFQE/MAdWc0p0EkF7ZX0vVgtYfjwPVGcBIgIJJC19fBFcB3V0BwFhU29BVnd3KyoQFBBhZlNLP0RsUwZ5NTc+FgQDb2gVUSJEbFNSY2tvLggOEHdmDgFoXnQM|~1580626274121~1~20200120~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIwIiwibGUiOiIxMDAiLCJjdCI6ImkiLCJkdCI6IjI4NTAwIn19~1~158~vi6f|2d8f;doei:,1,0,0,0,0,1000,-1000,1000,-1000;dmei:,1,0,0,1000,-1000,1000,-1000,1000,-1000;emc:;emcf:;ivli:;iivl:;ivcvj:;scvje:;1580626273735,1580626274120,0,0,1,1,0,0,0,0,0;qz86"
	)

var (
	queryClient = &http.Client{}
	queryHead   = &http.Request{}
	priceData   = url.Values{} // 出价的body
	priceClient = &http.Client{}
	priceHead   = &http.Request{}
	priceUrlStr = ""
)

func init() {

	quryApiUrl := "https://used-api.jd.com/auctionRecord/getCurrentAndOfferNum?auctionId=" + id //+ "&callback=__jp17"
	data := url.Values{}
	uTmp, _ := url.ParseRequestURI(quryApiUrl)
	urlStr := uTmp.String()
	queryHead, _ = http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	queryHead.Header.Add("Referer", "https://sell.paipai.com/auction-list/"+id)
	queryHead.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	queryHead.Header.Add("cookie", cookie)
	queryHead.Header.Add("content-type", "application/x-www-form-urlencoded")

	priceData.Set("trackId", "0d1b9bc80e5339d06cd27454b25c4b66")
	priceData.Set("eid", "MX3IRQ3GTOJTZWGF3WHRPHINBUVMCNIWU5CQR7J25IYYFJO5JJ57QBXR3KLXSC56VGZEDIZ3AURLP5S44YW3HXVQYY")
	priceData.Set("auctionId", id)
	priceData.Set("token", token)
	//priceData.Set("price", fmt.Sprintf("%v", price))
	priceData.Set("entryid", "")
	priceData.Set("address", "16-1303-48712-48759")
	priceData.Set("initFailed", "false")

	apiUrl := "https://used-api.paipai.com/auctionRecord/offerPrice"
	//priceData.Set("price", fmt.Sprintf("%v", price))
	u, _ := url.ParseRequestURI(apiUrl)
	priceUrlStr = u.String()
	priceHead, _ = http.NewRequest("POST", priceUrlStr, strings.NewReader(priceData.Encode())) // URL-encoded payload
	priceHead.Header.Add("Referer", "https://sell.paipai.com/auction-detail/"+id)
	priceHead.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	priceHead.Header.Add("cookie", cookie)
	priceHead.Header.Add("content-type", "application/x-www-form-urlencoded")

}

func main() {
	isSleep := true
	sleepTime := time.Second * 5
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
			bidPrice(ret.CurrentPrice + addPrice)
		}
		if remainTime < 2 {
			isSleep = false
			price := ret.CurrentPrice + addPrice
			retCode := 0
			for price < maxPrice {
				retCode = bidPrice(price)
				if retCode == 302 {
					continue // 连续出价不抬价
				}
				price += addPrice
			}

		} else {
			//start := time.Now()
			//bidPrice(ret.CurrentPrice + addPrice)
			//fmt.Println("出价用时:", time.Since(start))
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
	resp, err := queryClient.Do(queryHead)
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
func bidPriceOld(price float64) (retCode int) {
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
	urlStr := u.String()
	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Referer", "https://sell.paipai.com/auction-detail/"+id)
	r.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	r.Header.Add("cookie", cookie)
	r.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := client.Do(r)
	if err != nil {
		//fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	retByte, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("resp:", string(retByte))
	ret := &Ret{}
	err = json.Unmarshal(retByte, ret)
	if err != nil {
		//fmt.Println(err)
		return
	}
	retCode = ret.Code
	if retCode != 200 {
		err := fmt.Errorf("%v:%v", ret.Code, ret.Message)
		fmt.Println(err)
		return
	}
	fmt.Println("出价成功:", price)
	return
}

func bidPrice(price float64) (retCode int) {
	getReqHead(price)
	resp, err := priceClient.Do(priceHead)
	if err != nil {
		//fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	retByte, _ := ioutil.ReadAll(resp.Body)
	ret := &Ret{}
	err = json.Unmarshal(retByte, ret)
	if err != nil {
		//fmt.Println(err)
		return
	}
	retCode = ret.Code
	if retCode != 200 {
		err := fmt.Errorf("%v:%v", ret.Code, ret.Message)
		fmt.Println(err)
		return
	}
	fmt.Println("出价成功:", price)
	return
}

func getReqHead(price float64) {
	priceData.Set("price", fmt.Sprintf("%v", price))
	priceHead, _ = http.NewRequest("POST", priceUrlStr, strings.NewReader(priceData.Encode())) // URL-encoded payload
	priceHead.Header.Add("Referer", "https://sell.paipai.com/auction-detail/"+id)
	priceHead.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	priceHead.Header.Add("cookie", cookie)
	priceHead.Header.Add("content-type", "application/x-www-form-urlencoded")
}
