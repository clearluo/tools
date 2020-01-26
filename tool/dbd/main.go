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
	id       = "200589607"
	name     = "MagicWatch2"
	maxPrice = 1399
	addPrice = 3 // 加价间隔
	cookie   = "pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; __jdc=104464258; __jdv=104464258|direct|-|none|-|1578745305644; 3AB9D23F7A4B3C9B=BRVN5TORUPQAVMNAMKSTAH3HZHZSISSEHQ65YZ5LYHLTDM7J4QMUR6SIM7GSQHOJMPXTWSI6CBA5KD63KKTKDCJLJI; __tak=7281f1f9690c8d37994c246dadfbc89937c21cdffa36f375bb764f672079a3401aecc016b9d282835763ce01d71b745ebfaeb2ddb5878470d3c07e7ba95c6ec6dde61c3338a44cd8a5cca88bf251da25; __jda=104464258.15787453056431122413037.1578745306.1579002824.1579093435.5; thor=A6CD0F6C27D13A138EBD4D18B8A97E3CF72E408050B9B590D3C90896260F95CAB72BC5DDEEF1B01E8ED32F5DEEB0EB5884C83C6119818FB843F735871ACC7F11472B6BF44D1E67ACFF2F8061C19861517C86AEC220338A8C50D236243673903F0931A8F239FDC0427A305B01413A5898C4EBB0BA750D312A0678A3D89C6E8563; __jdb=104464258.4.15787453056431122413037|5.1579093435"
	token    = "g1mtsitn8ep5wa0zxhz1578745305108poq3~NmZeSyVEbFNSdHB0dFheBHt9AghhRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaaVBgRlIdZWEYQwtub35rGmBSYURbc3J+IgNcU3snDl5mA2RFVnIifyFXWQZ5IgZdJQskBT9jaxFmCB5fEWYNZHMANx0QJBtvaD1PWj4waxprOnQCBi0rYzQABEIsLRlbPgsKU08dZT0qPU8IEWYYWSQFIhgML2opIRUMWyFrBQhhU25HVXZ2EWZNMRA9MGsaazp0GRc1Nz5+PU8eEWZHUQ1EbC1Bc3Vhc01bHn1oBxRlXwpTHmNrbyEFCUEqZg0aPQQ3REFtZS43Q1cQeT1BUWcKZxILJit7PQkdR3smQgxoXzIeUiMyfT4KBwQ0NFZWc0p0EkF7ZSE1WFtHOXIEX2YcOggbMHUsPgdcB3p8AAxkVWZEUnF/NHVZXhBhZlNLP0RsUxsmIi4sFBVab2gVUSJEbFNSY2tvLggOEHdmDgFoXnQM|~1579093669851~1~20191203~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIxIiwibGUiOiIxMDAiLCJjdCI6IjAiLCJkdCI6ImkifX0=~1~~40v4|1d3l-g8,82;1d3-gi,82;1d1u-j6,8q;1d3-ja,8u;1dq-jh,97;1d5-jj,9d;1d9-jj,9i;1d13-jf,am;1dt-iz,bt;doei:,1,0,0,0,0,1000,-1000,1000,-1000;dmei:,1,0,0,1000,-1000,1000,-1000,1000,-1000;emc:,d:11;emcf:,d:11;ivli:;iivl:;ivcvj:;scvje:;1579093669497,1579093669849,0,0,9,9,0,2,0,0,0;pb9r"
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
			start := time.Now()
			bidPrice(ret.CurrentPrice + addPrice)
			fmt.Println("出价用时:", time.Since(start))
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
