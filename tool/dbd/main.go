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
	id       = "216651701"
	name     = "小爱同学"
	maxPrice = 52
	addPrice = 3 // 加价间隔
	cookie   = "shshshfpa=6339f04d-9797-cc8a-68fb-005de9e57b3d-1588983003; shshshfpb=b%20c6z31wQQyH2M1qE0bE2aQ%3D%3D; user-key=99d8cbba-aa16-46c6-b03a-1743a1866528; __jdu=15889830008982044370215; pinId=O6R_B8s_J3E; pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; _tp=gqQK%2Btk9D0R7L9OT%2B7S%2Bsw%3D%3D; _pst=tolsw; areaId=16; ipLoc-djd=16-1303-1305-53259; cn=187; unpl=V2_ZzNtbRVXE0UlXEZQfx1UUGJTFFwRUxMSdQFEXSwdXQMyVhZUclRCFnQUR1NnGFgUZwQZWERcRhxFCEdkexhdBGYCEFpBU3NMJVZGV3lFFzVXABJtQlZzFXEBT1xyG10EZgoaVEFWRhB3DkdXex1sNWAzIm1DZ0IldDgNOnpUXAFuChpUQFZCFHwAT1d6HFkHYQIRXUZnQiV2; __jdv=76161171|www.hao123.com|t_1000003625_hao123mz|tuiguang|f1aaae15559d4a70b5a61838f507dd58|1591346868275; shshshfp=6e6b6581defc453d0d3eee429ef5ae0b; __tak=52a1f8d0ec6aacddd42c72d92e14c3c37fb1e8580dd98f58f2e0205374d08d3f6278857bca233e60ef641ce129d187279e2abd695519728a04747c04bf8c742ced116cc8b11bc6af79325709cece7000; wlfstk_smdl=zvx0p7jhourgctsg1g568w4k6m8fmnnl; TrackID=1o-n_4ERKNTjtYRrpGfq5BehxzGpJlSGDiLIteLdl1Gp6mTnqs0eTfEtW9ic0YITuluuv_9rfe3ldY71AoKK8Y3Oxkd2C5W4A6EMsYbUwxts; thor=A6CD0F6C27D13A138EBD4D18B8A97E3C794CD84061A9306C05955B0D2FE298141F06482621DDDE6E305F5CB09D8C7D8ABA2AC2A15207997BEFA36118C08DC52B9286C77973C837B8F3A4C6C6F00BB030441693340750D1E8A5C579782FD9983394F733A885D8F6B54C0796219798DE20B4A1F61CC7215EE6027B47CBFA4FA430; ceshi3.com=201; logining=1; __jda=148612534.15889830008982044370215.1588983000.1591350306.1591576031.25; __jdc=148612534; 3AB9D23F7A4B3C9B=3TBAW6ZHHWSJU4EGOD4UMNJHUOC4LSYR32TLJRZNOVYMO6JFTTHPOSDVASOURI4PVTSWK62X5WQJ4K45EWTNNDPYXI; __jdb=148612534.5.15889830008982044370215|25.1591576031"
	token    = "958w6zukc8rs9g5c7ql1591576029781vcf5~NmZeSyVEbFNSdH58cVZbA313Dg9iRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaIx8/FV4dZWEYQwtub35rGjBSNRBUJ3R0fFYJAHRyUgtkAmZIAiJ0fCYFVQF8IQZZMF9lGz9jaxFmCB5fEWYNZHMANx0QJBtvaD1PWj4waxprOnQBAig3LC1PB1ZjJ1hVDUR6LUExKRFmWzEQYiVCWyUPOR9OJSI5JQgBHX91AQ5lUWBGVB1lYRhDHUYRZg1kcw4iBRMyfRFmTTEQPS1rGms6dENTbXBhck1dHn1oAAANRCtTT2MiKSASChB3Zg5eOxd0XUEiNG9+Q1VdKz0PXWRXJEcSKTInLlEHWS9xBgooXm8UCycjJHAZGwV4cBUUcwV0S0ElN30mUwFROypFUmUOZhYTKy8kdVRUAnRzAwBoXmBEUXk9ITJDQRApN1kaa0QgR1Y7KyQrGU8eby1EGmtEZ1NPYy0kJ0NXEHR9Dg9zGw==|~1591576106531~1~20200318~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIxIiwibGUiOiIxMDAiLCJjdCI6IjAiLCJkdCI6ImkifX0=~2~1077~8ii5|1d3c-ja,do,jb,dp;1d2y-jh,do,gx,3i;1dn-jj,do,gy,3i;2dn9;1dab-ji,do,4z,c;cw1-838,744;1dax-jj,dq,50,f;1dg-jh,du,14,h;1d9-jf,dw,12,k;1d7-je,dz,11,3;1da-jc,e4,y,8;1dd-j5,ec,r,g;1da-j2,ee,p,j;1d8-iw,el,4e,5;1d9-is,ep,49,9;1d8-il,ev,43,g;1df-id,f2,3u,n;1d8-i8,f8,3p,s;1dc-i2,fb,52,20;1d4-hr,fj,4r,4j;1dg-he,ft,4e,4s;1d8-h6,fy,47,4y;1d8-h1,g4,41,53;1d8-gy,g5,3y,54;1dg-gl,ge,3c,4;1d8-gg,gh,37,8;1d8-gc,gm,32,c;1d8-g8,gn,2y,e;1dg-g2,gs,2s,i;1d8-fy,gt,2p,j;1dh-fu,gw,2k,m;bdad-fg,h5,26,v;doei:,1,1,0,0,1,1000,-1000,1000,-1000;dmei:,1,1,1,1000,-1000,1000,-1000,1000,-1000;emc:,d:44;emmm:,d:4-0;emcf:,d:44;ivli:;iivl:;ivcvj:;scvje:;ewhi:;1591576103971,1591576106531,0,0,32,32,0,15,0,0,0;8y8f"
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
			//retCode := 0
			for price < maxPrice {
				go bidPrice(price)
				//retCode = bidPrice(price)
				//if retCode == 302 {
				//	continue // 连续出价不抬价
				//}
				price += addPrice
				time.Sleep(time.Millisecond * 100)
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
