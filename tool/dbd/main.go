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
	id       = "224404803"
	name     = "B&O音响"
	maxPrice = 599
	addPrice = 3 // 加价间隔
	cookie   = "pinId=O6R_B8s_J3E; shshshfpa=d9396e71-77e5-0d91-f212-08f43a6787fb-1594431682; __jdu=1594431679963119011653; shshshfpb=pCb1F%20bWkBQmteQYOsYxeqw%3D%3D; user-key=7f4db9d7-1636-495f-b16a-75f63b45a938; cn=0; pin=tolsw; unick=%E6%B0%91%E9%97%B4%E5%8D%97%E5%AF%92%E5%B8%A6; _tp=gqQK%2Btk9D0R7L9OT%2B7S%2Bsw%3D%3D; _pst=tolsw; unpl=V2_ZzNsbUAASxBwW04GZxhVBG4fFwlKBV8ddAgTSHIYVVdgBRZbQwVLQWlJKFRzEVQZJkB8XkdTSgklTShUeB1dAmIzEVxBVl8UcxREVGoZVA5mAhlfSmdDJXUJR1V6Gl4HbgcibXJXQSV0OEZQcxxZB2cEFFVKUEEVdQBHVHseWAdXVkYJEVFBFncUFld%2FTEAANwpHQUsDQRZpABYGKx0OBGNXEllEZ0UT; __jdv=122270672|kong|t_1000023385_125064|zssc|3f944b9c-0808-4e9c-901d-808c67570c9e-p_1999-pr_2458-at_125064|1596193239921; PCSYCityID=CN_350000_350100_350102; shshshfp=f29d17d857ae68206626f19fb892a433; __tak=fc842d2fdaa57db822e1fe7792c43f03c1cbe715bf5313b67797f203c2d7d3586d3726d75300e90ddf03457ae1d1922514f73d7044a3c425bed3903813407e67ae7be4837c31808743257e8bc4268c83; wlfstk_smdl=xbn7yovxiawz4sbcetr3kvbkm4sdz0cw; TrackID=1po1X5Ud-I6JPtU87ZHW2ra_6A15cegLq0Q5Uf1VZ7cAoSve2R_I_ojxSn8NO4qsa8TTpT6BP8Kpof_JkBNGpaC5oinvlbjGfI3EZ9ABgwjc; thor=A6CD0F6C27D13A138EBD4D18B8A97E3C4DA6609DA28A55CDC5CD3A5EEAF7FCE2F4280774E6202F65CDE234C1535ADC61EF003C6B6059C87BB5C47114BC67BB5307324AF9EA00EF08170FE9987BFFA6D3D74CA99CD0E53B50D9F760C649BEFFB31E3A8FBED42D0612359CFD16EA05B04A50206289433729F5352F7FE87B1FCDBD; ceshi3.com=201; logining=1; __jda=148612534.1594431679963119011653.1594431679.1596193240.1596592095.12; __jdb=148612534.5.1594431679963119011653|12.1596592095; __jdc=148612534; 3AB9D23F7A4B3C9B=3TBAW6ZHHWSJU4EGOD4UMNJHUOC4LSYR32TLJRZNOVYMO6JFTTHPOSDVASOURI4PVTSWK62X5WQJ4K45EWTNNDPYXI"
	token    = "5kywbk328kkhlqopo6x1596592094034yehd~NmZeSyVEbFNSdH57cVhfA3p3BgFlRHpTBiUjb35DFm5vLUROOBEzLUF7G28iAAFBKBgVFA1EPwIVKDclGENXbm8iVlQiAwpTTx1lKSsTCG5vfmsaYlNvQVMdZWEYQwtub35rGjVRZEIFJH4ufAJUVn8mUVliBzMXUHgkfSUHDFcsfQAIPQ4jRj9jaxFmCB5fEWYNZHMANx0QJBtvaD1PWj4waxprOnQBAig3LC1PB1ZjJ1hVDUR6LUExKRFmWzEQYiVCWyUPOR9OJSI5JQgBHX92Aw1iU29BUx1lYRhDHUYRZg1kcw4iBRMyfRFmTTEQPS1rGms6dENTbXBhck1dHn1oAAANRCtTT2MiKSASChB3ZgIIOx90XUEiNG9+QxpdOnJGWilUYRlaeXU6MFUXVXwmQFowAzRHACkxdSgVWFQ9IhUUcwV0S0ElN30mUwFROypFUmUOZhYTKy8kdVRUAnRzAwBoXmBEUXk9ITJDQRApN1kaa0QjQVY3LCkgBk8eby1EGmtEZ1NPYy0kJ0NXEHR9Dg9zGw==|~1596592178575~1~20200318~eyJ2aXdlIjoiMCIsImJhaW4iOnsiaWMiOiIxIiwibGUiOiIxMDAiLCJjdCI6IjAiLCJkdCI6ImkifX0=~2~-1031~fii0|1d3g-og,bj,oh,bk;1d3m-of,bj,jw,1b;1ddd-oi,bh,3c,g;1da-oj,bg,3d,g;1d6-ol,bf,3f,f;1dd-oo,be,3i,d;1di-os,bc,3l,c;1de-p7,b7,41,7;1dc-pc,b5,46,4;1de-pu,ay,8k,59;1dh-q8,au,8x,54;1dp-qy,aj,9o,1r;1d9-sa,9y,b0,16;1dg-st,9q,a0,d;1dj-tu,97,ck,f;1dg-uo,8v,de,3;1de-vl,8n,p,17;1dh-w1,8l,15,14;1d1i-wk,8l,1p,14;1d5-wl,8m,1p,15;1do-wl,8o,1p,17;1dh-wk,8r,1o,1b;1dh-wh,8u,f7,2;1dg-wc,92,f2,9;1dh-w7,96,ex,d;1dh-w0,9c,19,-1;1dg-vx,9e,16,2;1dh-vs,9i,10,6;1dh-vn,9m,w,9;1dg-vi,9p,r,c;1dh-vc,9r,l,f;1d2s-v1,9y,dr,15;1d2s-ui,a4,d8,1b;1d2s-uk,9y,d9,16;1d2s-wl,8i,1p,12;1d42-xf,89,t5,2j;1d1l-xc,87,t1,2i;2daz;1d6j-xb,87,t1,2i;1d1v-wo,8f,1s,z;cwx-1019,722;1dn4-s0,9k,s0,9k;1dw-pt,9z,cu,17;1d2y-kd,ay,3d,0;1d2p-id,cb,4,6;1d2n-h4,e2,7,2e;1d2z-g6,f7,1n,s;1d2v-ft,fx,2u,4w;1d2o-fk,gg,2b,7;bd10e-fk,gi,2a,9;doei:,1,1,0,0,1,1000,-1000,1000,-1000;dmei:,1,1,1,1000,-1000,1000,-1000,1000,-1000;emc:,d:148;emmm:,d:37-0;emcf:,d:148;ivli:;iivl:;ivcvj:;scvje:;ewhi:;1596592173191,1596592178575,0,0,50,50,0,101,0,0,0;uku3"
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
