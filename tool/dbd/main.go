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
	id       = "226257630"
	name     = "wifi放大器"
	maxPrice = 26
	addPrice = 1 // 加价间隔
	cookie   = "unpl=V2_ZzNtbUsAFhBxWEZdfEsJUWIBEg1LUENCcVhCVygcDlJmBEZfclRCFnQUR11nGV0UZwAZXUFcRxRFCEdkexhdBGYCEFpBU3NMJVZGV3lFFzVXABJtQlZzFXEAQFV5GloDZgMbX0tVQxV9CE9SeilsAlczIlxyVnMURUMoVTYZWA1hAhBeRFFCFXwKT1Z7GVQFbgUTbUNnQA%3d%3d; __jdv=76161171|www.hao123.com|t_1000003625_hao123mz|tuiguang|8fd45a186cde431a861f5a52b4cf06e3|1597032770776; __jdu=1597032770183831191870; areaId=16; ipLoc-djd=16-1303-3484-0; PCSYCityID=CN_350000_350100_350102; shshshfpa=21c228a4-ce59-2efe-bcdf-707d4eb77027-1597032772; shshshfpb=wS54E7rGhryotayEdEKpyKg%3D%3D; 3AB9D23F7A4B3C9B=LLYWXM24PGCI2ZS6DR2JLHX57NF4PD3ONXFO3UOJ2RRB3CUCMRQOMLKU34JIGJBCPWINTPWL5DXCFU4QIFPZBJP7CU; cid=NXpVMzUyNm9SNzUwN3dMMDcwOHRBNzE4MXhOMjEyMnhHNzc0M2NSMzcyNHlUNTk4; shshshfp=f29d17d857ae68206626f19fb892a433; pinId=O6R_B8s_J3E; pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; _tp=gqQK%2Btk9D0R7L9OT%2B7S%2Bsw%3D%3D; _pst=tolsw; __tak=46a8344fbc7a53adaa9336865bab68cff1e2d43aca76908876cd2900616cd2642570fc50ec6feeba8b88334f55dcbcc6034a810adbb7eeabc660eac43b84f91e9868217088bc885f0d53b00023f5a5f4; wlfstk_smdl=wyg4xgmpzgyqdqaraqeqkmhrv5l9kuhg; TrackID=1yg20uMgEGQALk2RWOiloFbRSN_cP8Ire2SeyGK5L3ZmYi99quXvDHpBZWYEARbn-9Wk75juGhR95F0TixbvtK-hjIz_ilk6wi-nI2kUghcc; thor=A6CD0F6C27D13A138EBD4D18B8A97E3C97149A1EF040FE64E7FFA5549429A7C8061222A42FC46CC7773C939E196F496C9142EC377AEB03D71BC898A556E909E08D11DD2D392BF96F5B131375B83BD2C98402A2FAB4224DB736F7AD70A690E7E6920D8E8D266F16934F2FECD9A3CB9140765EB226C166BF1481BE572AA5040680; ceshi3.com=201; logining=1; __jda=148612534.1597032770183831191870.1597032770.1597290215.1597627704.7; __jdc=148612534; __jdb=148612534.6.1597032770183831191870|7.1597627704"
	token    = "504azbc1mc1sdzz84re15976277037347sp4~NmZeSyVEbFNSdH56clNaBXlyBwpoRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaZFFgQlMdZWEYQwtub35rGjcEYEUCInd5IFZcUXshAgtlVDNBB3ByeSECWFB/JQZaNB8gFz9jaxFmCB5fEWYNZHMANx0QJBtvaD1PWj4waxprOnQBAig3LC1PB1ZjJ1hVDUR6LUExKRFmWzEQYiVCWyUPOR9OJSI5JQgBHX92AQpkUWBCUx1lYRhDHUYRZg1kcw4iBRMyfRFmTTEQPS1rGms6dENRbXBhck1dHn1oAAANRCtTT2MiKSASChB3ZlxKO1Z0XUEiNG9+QwpGLHxASWQVbxwHKyIiKQkaRXpyX1A4Hi9HGjt2On0OHlAiJRUUcwV0S0F0Kzt9VRkGfy5PUWYTZAMaNDYmdVRUBX1wDw9mXmFFWzQsOSZDQRApN1kaa0Q8GVYjci89Vk8eby1EGmtEZ1NPYy0kJ0NXEHR9Dg9zGw==|~1597627747813~1~20200318~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIxIiwibGUiOiIxMDAiLCJjdCI6IjAiLCJkdCI6ImkifX0=~2~311~uii7|1d2s-f0,fq,f0,fr;1d3-ez,fs,ez,fs;1dr-ey,fx,ey,fx;1d3g-f7,ea,f7,f;1da-f7,ds,f7,3j;1da-f7,dl,f7,21;1de-f7,d6,f7,1l;1dj-f7,cp,f7,15;1d6-f7,cm,f7,11;1d1t-f4,bn,f5,3;1dp-f3,bk,f3,0;1d20-es,b2,el,3y;1db-ek,at,ee,3p;1d8-ee,am,e7,3j;1d8-ec,aj,e5,3g;1de-e2,aa,dv,37;1d9-dv,a1,do,2x;1d1h-cs,8y,4k,1f;1db-cl,8u,4e,1c;1de-ck,8u,4d,1c;1dr-ck,8w,4d,1d;1dc-ck,91,4d,1i;1dc-co,9e,4g,8;1d6-cp,9j,4i,d;1d1q-di,bq,i,q;1d2n-es,dk,8,8;1d2p-et,dj,a,7;1dd-eu,di,a,5;1dp-ev,di,b,5;1d2m-ex,dh,e,4;1d1m-f6,dl,m,8;1d2x-g4,f1,1m,l;1d2n-gm,g8,3m,58;1d2o-gj,gk,39,b;bd8a-gf,gu,35,k;doei:,1,1,0,0,1,1000,-1000,1000,-1000;dmei:,1,1,1,1000,-1000,1000,-1000,1000,-1000;emc:,d:78;emmm:,d:37-0;emcf:,d:78;ivli:;iivl:;ivcvj:;scvje:;ewhi:;1597627746025,1597627747812,0,0,35,35,0,44,0,0,0;b7v5"
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
		remainTime := float64(ret.ActualEndTime-time.Now().UnixNano()/1000000) / 1000.0
		fmt.Printf("No:%v     当前价格:%-6v出价者:%-10v剩余时间:%-8.3f接受最高价格:%-10v\n", id+"-"+name, ret.CurrentPrice, ret.CurrentBidder, remainTime, maxPrice)
		if ret.CurrentPrice > maxPrice {
			fmt.Println("超出最高价，退出竞拍")
			break
		}
		if remainTime < -1 {
			fmt.Println("竞拍结束")
			break
		}
		if ret.CurrentBidder == userId {
			//fmt.Println("最高价为本人，不出价")
			continue
		}
		//bidPrice(ret.CurrentPrice + addPrice)
		if remainTime < 10 {
			sleepTime = time.Second
			continue
		}
		if remainTime < 3 {
			sleepTime = time.Millisecond * 500
			bidPrice(ret.CurrentPrice + addPrice)
		}
		if remainTime < 1 {
			isSleep = false
			price := ret.CurrentPrice + addPrice
			for price < maxPrice {
				go bidPrice(price)
				price += addPrice
				time.Sleep(time.Millisecond * 5)
			}
		}
	}
	if ret, err := getPrice(); err != nil {
		fmt.Println("竞拍失败")
	} else {
		remainTime := float64(ret.ActualEndTime-time.Now().UnixNano()/1000000) / 1000.0
		if remainTime < 0 {
			fmt.Printf("竞拍已结束，用户%v以%v元的价格竞拍成功\n", ret.CurrentBidder, ret.CurrentPrice)
		} else {
			fmt.Printf("竞拍结束，当前价超出你能接受的最高价，当前价:%v\n", ret.CurrentPrice)
		}
	}
	select {}
}

func getPrice() (*Data, error) {
	//start := time.Now()
	//defer func() {
	//	fmt.Println("查询价格用时：", time.Since(start))
	//}()
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
		err := fmt.Errorf("出价失败:%v:%v", ret.Code, ret.Message)
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
