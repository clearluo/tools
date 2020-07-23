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
	id       = "222578385"
	name     = "Master2S"
	maxPrice = 239
	addPrice = 3 // 加价间隔
	cookie   = "pinId=O6R_B8s_J3E; shshshfpa=d9396e71-77e5-0d91-f212-08f43a6787fb-1594431682; __jdu=1594431679963119011653; shshshfpb=pCb1F%20bWkBQmteQYOsYxeqw%3D%3D; user-key=7f4db9d7-1636-495f-b16a-75f63b45a938; cn=0; areaId=16; ipLoc-djd=16-1303-3483-0; unpl=V2_ZzNtbUdRFxIhC0JXLhxeUGJUG1VKAkdFIQ8VV3pLXwdnARdUclRCFnQUR1JnGVsUZwMZWEZcRhZFCEdkexhdBGYCEFpBU3NMJVZGV3lFFzVXABJtQlZzFXEAQ1F5GVsDbwsVX0JXSxR1CEFQeSlsAlczIlxyVnMURUMoVTYZWA1iBhBdRVFLHXIKRlRzGFwFYAcQbUNnQA%3d%3d; __jdv=76161171|www.hao123.com|t_1000003625_hao123mz|tuiguang|47e6e252d43d4f899d5ae6b20c231348|1594871142180; shshshfp=f29d17d857ae68206626f19fb892a433; __tak=aada616649dad71f7d4061bcae371e1f212e01934ead4e9e7fe971fbdd54654e00ba03954778bc5bcc73aa5d9dd01c470f0dd36bfdd393effc92f95cd52ba177959fa20c5e81a2dbf511b5bdfcbdac7e; wlfstk_smdl=6hhn199kwn8dp2odte6dsyntv31cmxms; TrackID=1naol_ya94MMBq0koYo7_C-JwaoL9ZZp_GYjneZ7w6OrbPSs59zuhf94zgKhm62LIWIyJ__ymmj9QY1xzj7VWvK79v1IibYws9fpPx-_JYmo; thor=A6CD0F6C27D13A138EBD4D18B8A97E3C417181ACDD0E70F7C23AC31477C84BAFFC6261DBFFD1FBF32759B482B6C59DCD64FBCC3CFA66281F68CE4FD8AB0E8E23CBE9F5028E7EF2587EFA047E75D32B6B921AC92960E89C8374C594B22AE64E80F442D0A80096A75697A7FA61E5945A2B1690E8D7C37BBBCB6B2CC9FDF5FBA37A; pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; ceshi3.com=201; _tp=gqQK%2Btk9D0R7L9OT%2B7S%2Bsw%3D%3D; logining=1; _pst=tolsw; __jda=148612534.1594431679963119011653.1594431679.1595310338.1595466844.6; __jdc=148612534; 3AB9D23F7A4B3C9B=3TBAW6ZHHWSJU4EGOD4UMNJHUOC4LSYR32TLJRZNOVYMO6JFTTHPOSDVASOURI4PVTSWK62X5WQJ4K45EWTNNDPYXI; __jdb=148612534.7.1594431679963119011653|6.1595466844"
	token    = "ru1q8p6iacywm3nf7cy15953103379671sao~NmZeSyVEbFNSdH54cFdaAntxAgtiRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaIx8/FV4dZWEYQwtub35rGjVRZEIFJH4ufAJUVn8mUVliBzMXUHgkfSUHDFcsfQAIPQ4jRj9jaxFmCB5fEWYNZHMANx0QJBtvaD1PWj4waxprOnQBAig3LC1PB1ZjJ1hVDUR6LUExKRFmWzEQYiVCWyUPOR9OJSI5JQgBHX92BQ1mXmVJVh1lYRhDHUYRZg1kcw4iBRMyfRFmTTEQPS1rGms6dENTbXBhck1dHn1oAAANRCtTT2MiKSASChB3Zg4IOwd0XUEiNG9+QwpYeSAHSDgRZBASKC0hPQ0LBSo8Ul83CmYLUy03eycAB0E8IxUUcwV0S0ElN30mUwFROypFUmUOZhYTKy8kdVRUAnRzAwBoXmBEUXk9ITJDQRApN1kaa0QwGVY7LCkrCU8eby1EGmtEZ1NPYy0kJ0NXEHR9Dg9zGw==|~1595467089346~1~20200318~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIxIiwibGUiOiIxMDAiLCJjdCI6IjAiLCJkdCI6ImkifX0=~2~1392~uii5|1d2z-l4,bd,l4,bd;1dc2-l0,b2,gi,3z;1da-kg,9y,1a,r;1dp-j4,7t,1c,a;1d34-gg,3h,c5,5;1dc-gg,3e,c5,3;1dh-gg,3a,ca,2g;1dg-gg,33,ca,2a;1dh-gg,31,ca,27;1d1u-gf,2n,ca,1u;1df-ge,2m,c9,1t;1d41-g8,2o,c2,1u;1d6-ft,2p,bo,1v;1dl-ff,2o,ba,1u;1dd-f1,2m,aw,1s;1dh-en,2j,44,t;1dh-e6,2g,3n,p;1dh-do,2a,34,k;1d3f-co,1x,24,6;1dl-9t,k,59,k;1d2q-9k,6,4z,7;1d2-9g,2,4w,2;1d1f3-ee,4,9u,4;1dh-ej,i,9z,j;1dg-em,t,a2,t;1dh-eq,16,al,d;1dh-eu,1e,ap,l;1dg-ez,1o,au,4;1dh-f1,1u,4i,4;1dh-f6,24,4n,d;1dx-fk,2v,bf,22;1d2s-hk,5k,d9,y;1d2s-l3,8f,1b,x;1d2s-ne,8d,5u,w;cwe1-1088,588;1db2v-tl,9n,5,2o;1d10-os,9q,a9,d;1d2s-n9,66,aa,g;1d2s-qy,q,qy,q;1drr-z0,1m,1c,2;cw0-1536,818;1di-z1,1m,1d,2;1d2k-z1,1n,1d,3;1d2j-z3,2a,e,m;1d2s-zb,41,1m,p;1d2s-102,6d,1c,m;1dpx-104,73,24,5;cwz-1144,818;1d2d-v3,82,1d,u;1d1e-oe,9p,9v,c;1dm5-f6,ib,n,6;bdq6-g5,hu,36,6u;0d2j-g5,hu,36,6u;1d3o-g6,hu,36,6t;bdfw-g8,h6,2z,w;doei:,1,1,0,0,1,1000,-1000,1000,-1000;dmei:,1,1,1,1000,-1000,1000,-1000,1000,-1000;emc:,d:242;emmm:,d:30-0;emcf:,d:242;ivli:;iivl:;ivcvj:;scvje:;ewhi:;1595467065531,1595467089346,0,0,55,55,0,193,0,0,0;44n6"
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
			for price < maxPrice {
				go bidPrice(price)
				price += addPrice
				time.Sleep(time.Millisecond * 2)
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
