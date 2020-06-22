package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
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
	id       = "218571604"
	name     = "佳明935"
	maxPrice = 69
	addPrice = 3 // 加价间隔
	cookie   = "shshshfpa=6339f04d-9797-cc8a-68fb-005de9e57b3d-1588983003; shshshfpb=b%20c6z31wQQyH2M1qE0bE2aQ%3D%3D; user-key=99d8cbba-aa16-46c6-b03a-1743a1866528; __jdu=15889830008982044370215; pinId=O6R_B8s_J3E; pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; _tp=gqQK%2Btk9D0R7L9OT%2B7S%2Bsw%3D%3D; _pst=tolsw; cn=187; areaId=16; ipLoc-djd=16-1303-1305-53259; unpl=V2_ZzNtbUMFQkVwChVSKE5bUWIFFggSXhBBJwESAH0QXQMzVBdbclRCFnQUR1NnGV0UZgsZXkNcQRJFCEdkexhdBGYCEFpBU3NMJVZGV3lFFzVXABJtQlZzFXEBT1xyG10EZgoaVEFWRhB3DkdXex1sNWAzIm1DZ0IldDgNOnpUXAFuChpUQFZCFHwAT1d6HFkHYQIRXUZnQiV2; __jdv=76161171|www.hao123.com|t_1000003625_hao123mz|tuiguang|0c0a43b7bf6e475da8bec8ee7807ef47|1591752035913; shshshfp=6e6b6581defc453d0d3eee429ef5ae0b; __tak=725360ce8f25a81142c533d1bf62895748669df9e0622680547ab747ab272ee31031fc7c64a446ed34e4fc0ca441d254f335682e501c63a5cca6d1e7e756f7368fdf6cda1d55314c67fdb885d0dfc642; wlfstk_smdl=9q4umu5kylsub67r8zedtmpt70t6whr5; TrackID=1g062xhpg7oneV-71eGRHB68Ds6ziQzIbnMyBX2fD36Rxd0pOcnt7x6_jlZ8yS3olekH_CcVufGXWf2jR-UIcQYAnvJ4VHye966we2RZfM84; thor=A6CD0F6C27D13A138EBD4D18B8A97E3CA15CCD78A63A95F690E8A7BD4A99C977FE2CA366E2A2A4E6555723A1A57E6D15DF35BAC32AF5C9F03E20772DF7B2C0770C84503207525DE4C978E5615FC55BB834FE432EBDCA30FB46F3C9EA2308ECA628FBB43B740D566C0F0D74592675BDA925C4E62F00526B162A11B79ED0396C12; ceshi3.com=201; logining=1; __jda=148612534.15889830008982044370215.1588983000.1591752036.1591785075.29; __jdc=148612534; 3AB9D23F7A4B3C9B=3TBAW6ZHHWSJU4EGOD4UMNJHUOC4LSYR32TLJRZNOVYMO6JFTTHPOSDVASOURI4PVTSWK62X5WQJ4K45EWTNNDPYXI; __jdb=148612534.5.15889830008982044370215|29.1591785075"
	token    = "8fo1qp2d29nkt0g6jkq1591785074681rf80~NmZeSyVEbFNSdH58c1lYA353BwpoRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaIx8/FV4dZWEYQwtub35rGjJTZxNXeHF1IgdYBX11AAlmA2YUASJyL3JUXlcpcVIJYhQ0CD9jaxFmCB5fEWYNZHMANx0QJBtvaD1PWj4waxprOnQBAig3LC1PB1ZjJ1hVDUR6LUExKRFmWzEQYiVCWyUPOR9OJSI5JQgBHX91AQ9oX2VAVB1lYRhDHUYRZg1kcw4iBRMyfRFmTTEQPS1rGms6dENTbXBhck1dHn1oAAANRCtTT2MiKSASChB3Zg4IOxB0XUEiNG9+QxkGInFFSmFWbxZQMTIgIBYYQSRyU147UGAdDywwez4LGlErKxUUcwV0S0ElN30mUwFROypFUmUOZhYTKy8kdVRUAnRzAwBoXmBEUXk9ITJDQRApN1kaa0Q4HFY7LCkvVk8eby1EGmtEZ1NPYy0kJ0NXEHR9Dg9zGw==|~1591785137468~1~20200318~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIxIiwibGUiOiIxMDAiLCJjdCI6IjAiLCJkdCI6ImkifX0=~2~1645~5iiz|1d2s-qc,4c,qc,4d;1d33-sb,2v,sc,2w;1dp-s8,2u,s9,2v;1d27-ro,2o,ro,2o;1d25-qp,r,o2,r;1da-ql,l,ny,l;1d13-qg,1,nt,2;2d1d6;cw9c-838,657;1d6f-n7,p,n8,q;1d7-n6,p,n7,q;1d8-n5,r,n5,r;1d8-n3,s,n3,s;1dh-my,w,mz,w;1d6-mw,z,mx,z;1d8-mt,12,mu,9;1d8-mq,17,mr,d;1dh-mk,1j,26,0;1d8-mh,1o,23,5;1d8-me,1w,21,c;1d8-ma,21,1x,i;1dg-m6,2f,1t,w;1d8-m4,2m,m4,1s;1d8-m2,2s,m2,1y;1d7-m0,2y,m1,25;1da-lz,32,lz,28;1df-lv,3d,v,2;1d9-lu,3i,v,7;1d6-lu,3l,v,9;1dh-lu,3v,v,j;1dn-m1,3w,11,k;1d5e-ma,3n,1a,b;1d9-ma,3p,1a,d;1d2x-j2,9y,4k,4;1d2o-hf,g1,j,3w;1d2o-h6,gx,3w,6;1d2y-g5,i2,g5,i3;1d6i-fs,i4,ft,i5;1d1q-fu,hl,2l,u;1d6r-fv,hd,2m,m;1d1l-fv,h7,2m,g;bd43-fv,h4,2m,d;doei:,1,1,0,0,1,1000,-1000,1000,-1000;dmei:,1,1,1,1000,-1000,1000,-1000,1000,-1000;emc:,d:120;emmm:,d:0-0;emcf:,d:120;ivli:;iivl:;ivcvj:;scvje:;ewhi:;1591785133025,1591785137468,0,0,42,42,0,81,0,0,0;0egp"
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
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}

func getPrice() (*Data, error) {
	start := time.Now()
	defer func() {
		fmt.Println("查询价格用时：", time.Since(start))
	}()
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
