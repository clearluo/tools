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
	id       = "201107889"
	name     = "Magic Watch 2"
	maxPrice = 860
	addPrice = 3 // 加价间隔
	cookie   = "shshshfpa=b8cc7f40-5852-c878-2008-202cb7b564aa-1576751256; shshshfpb=b8cc7f40-5852-c878-2008-202cb7b564aa-1576751256; mba_muid=1576642527839566679144; 3AB9D23F7A4B3C9B=MX3IRQ3GTOJTZWGF3WHRPHINBUVMCNIWU5CQR7J25IYYFJO5JJ57QBXR3KLXSC56VGZEDIZ3AURLP5S44YW3HXVQYY; shshshfp=e3c1b507580913054efdb15be6c61cc9; areaId=16; ipLoc-djd=16-1303-3484-0; __jdv=104464258|dbd.jinyiweb.com|-|referral|-|1579397542537; __jda=104464258.1576642527839566679144.1576642528.1579600241.1579660993.33; __jdc=104464258; __tak=6e970d662aebe7e73ce4edfb635f2f798148e555dd56c387dd71f40b93e314b6f2168d7445e9c12a331a92bcb81115ad84f6473bcf9011145bb203e7364b314f2764dad7a3de63dbbdb758d064385ac0; thor=A6CD0F6C27D13A138EBD4D18B8A97E3C532040F908F7342FDA073258B976394890E5FEE40B78F0E4B4F25E2E6BF60262A576894A4892C2927D6AE74959ED2A0FA604BD4DB15F980D6FB261B8FEEA846082D438B84830EB1EC5D16C6285F4C7E74ABDADDC5D703D4B322A7D23E98A71DEB5D0D61DC72080F6888937828FAD72BC; pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; __jdb=104464258.4.1576642527839566679144|33.1579660993"
	token    = "rog9ibjwl70et0e54bd15796609923194i30~NmZeSyVEbFNSdHB0cldcAnh3BQpmRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaIx8/FV4dZWEYQwtub35rGmBRYUMCcCV0cAVbC3h1Bgs3X24QVnkhLiVUWAB9cgYOZgA7FD9jaxFmCB5fEWYNZHMANx0QJBtvaD1PWj4waxprOnQCBi0rYzQABEIsLRlbPgsKU08dZT0qPU8IEWYYWSQFIhgML2opIRUMWyFrBQhgV2ZGW3l+EWZNMRA9MGsaazp0GRc1Nz5+PU8eEWZHUQ1EbC1Bc3dhfE1bHn1oBxRlXwpTHmNrbyEFCUEqZg0aaFA3REFtZS43Q1cQPilTXmQTOklSM3V+NRQBC3wmQEI+VDJCB3YoLi8HBQYqLERVc0p0EkF7ZSsgWB5RI31ECjpTJwIOOSR/M1NcB3pyAQxjU2RGVnlzeykSBxBhZlNLP0RsU1o4ITdwUg4Bb2gVUSJEbFNSY2tvLggOEHdmDgFoXnQM|~1579661053706~1~20200120~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIxIiwibGUiOiIxMDAiLCJjdCI6IjAiLCJkdCI6ImkifX0=~1~-231~miix|doei:,1,1,0,0,1,1000,-1000,1000,-1000;dmei:,1,1,1,1000,-1000,1000,-1000,1000,-1000;emc:;emcf:;ivli:;iivl:;ivcvj:;scvje:;1579661053223,1579661053702,0,0,0,0,0,0,0,0,0;jcnm"
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
