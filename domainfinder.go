package domainfinder


import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"strings"
)

type Result struct {
	Rank         int     `json:"rank"`
	URL          string `json:"url"`
	Title        string `json:"title"`
}

const BaseUrl = "https://www.google."

var GoogleDomains = map[string]string{
	"us":  "com/search?q=",
	"ac":  "ac/search?q=",
	"ad":  "ad/search?q=",
	"ae":  "ae/search?q=",
	"af":  "com.af/search?q=",
	"ag":  "com.ag/search?q=",
	"ai":  "com.ai/search?q=",
	"al":  "al/search?q=",
	"am":  "am/search?q=",
	"ao":  "co.ao/search?q=",
	"ar":  "com.ar/search?q=",
	"as":  "as/search?q=",
	"at":  "at/search?q=",
	"au":  "com.au/search?q=",
	"az":  "az/search?q=",
	"ba":  "ba/search?q=",
	"bd":  "com.bd/search?q=",
	"be":  "be/search?q=",
	"bf":  "bf/search?q=",
	"bg":  "bg/search?q=",
	"bh":  "com.bh/search?q=",
	"bi":  "bi/search?q=",
	"bj":  "bj/search?q=",
	"bn":  "com.bn/search?q=",
	"bo":  "com.bo/search?q=",
	"br":  "com.br/search?q=",
	"bs":  "bs/search?q=",
	"bt":  "bt/search?q=",
	"bw":  "co.bw/search?q=",
	"by":  "by/search?q=",
	"bz":  "com.bz/search?q=",
	"ca":  "ca/search?q=",
	"kh":  "com.kh/search?q=",
	"cc":  "cc/search?q=",
	"cd":  "cd/search?q=",
	"cf":  "cf/search?q=",
	"cat": "cat/search?q=",
	"cg":  "cg/search?q=",
	"ch":  "ch/search?q=",
	"ci":  "ci/search?q=",
	"ck":  "co.ck/search?q=",
	"cl":  "cl/search?q=",
	"cm":  "cm/search?q=",
	"cn":  "cn/search?q=",
	"co":  "com.co/search?q=",
	"cr":  "co.cr/search?q=",
	"cu":  "com.cu/search?q=",
	"cv":  "cv/search?q=",
	"cy":  "com.cy/search?q=",
	"cz":  "cz/search?q=",
	"de":  "de/search?q=",
	"dj":  "dj/search?q=",
	"dk":  "dk/search?q=",
	"dm":  "dm/search?q=",
	"do":  "com.do/search?q=",
	"dz":  "dz/search?q=",
	"ec":  "com.ec/search?q=",
	"ee":  "ee/search?q=",
	"eg":  "com.eg/search?q=",
	"es":  "es/search?q=",
	"et":  "com.et/search?q=",
	"fi":  "fi/search?q=",
	"fj":  "com.fj/search?q=",
	"fm":  "fm/search?q=",
	"fr":  "fr/search?q=",
	"ga":  "ga/search?q=",
	"gb":  "co.uk/search?q=",
	"ge":  "ge/search?q=",
	"gf":  "gf/search?q=",
	"gg":  "gg/search?q=",
	"gh":  "com.gh/search?q=",
	"gi":  "com.gi/search?q=",
	"gl":  "gl/search?q=",
	"gm":  "gm/search?q=",
	"gp":  "gp/search?q=",
	"gr":  "gr/search?q=",
	"gt":  "com.gt/search?q=",
	"gy":  "gy/search?q=",
	"hk":  "com.hk/search?q=",
	"hn":  "hn/search?q=",
	"hr":  "hr/search?q=",
	"ht":  "ht/search?q=",
	"hu":  "hu/search?q=",
	"id":  "co.id/search?q=",
	"iq":  "iq/search?q=",
	"ie":  "ie/search?q=",
	"il":  "co.il/search?q=",
	"im":  "im/search?q=",
	"in":  "co.in/search?q=",
	"io":  "io/search?q=",
	"is":  "is/search?q=",
	"it":  "it/search?q=",
	"je":  "je/search?q=",
	"jm":  "com.jm/search?q=",
	"jo":  "jo/search?q=",
	"jp":  "co.jp/search?q=",
	"ke":  "co.ke/search?q=",
	"ki":  "ki/search?q=",
	"kg":  "kg/search?q=",
	"kr":  "co.kr/search?q=",
	"kw":  "com.kw/search?q=",
	"kz":  "kz/search?q=",
	"la":  "la/search?q=",
	"lb":  "com.lb/search?q=",
	"lc":  "com.lc/search?q=",
	"li":  "li/search?q=",
	"lk":  "lk/search?q=",
	"ls":  "co.ls/search?q=",
	"lt":  "lt/search?q=",
	"lu":  "lu/search?q=",
	"lv":  "lv/search?q=",
	"ly":  "com.ly/search?q=",
	"ma":  "co.ma/search?q=",
	"md":  "md/search?q=",
	"me":  "me/search?q=",
	"mg":  "mg/search?q=",
	"mk":  "mk/search?q=",
	"ml":  "ml/search?q=",
	"mm":  "com.mm/search?q=",
	"mn":  "mn/search?q=",
	"ms":  "ms/search?q=",
	"mt":  "com.mt/search?q=",
	"mu":  "mu/search?q=",
	"mv":  "mv/search?q=",
	"mw":  "mw/search?q=",
	"mx":  "com.mx/search?q=",
	"my":  "com.my/search?q=",
	"mz":  "co.mz/search?q=",
	"na":  "com.na/search?q=",
	"ne":  "ne/search?q=",
	"nf":  "com.nf/search?q=",
	"ng":  "com.ng/search?q=",
	"ni":  "com.ni/search?q=",
	"nl":  "nl/search?q=",
	"no":  "no/search?q=",
	"np":  "com.np/search?q=",
	"nr":  "nr/search?q=",
	"nu":  "nu/search?q=",
	"nz":  "co.nz/search?q=",
	"om":  "com.om/search?q=",
	"pa":  "com.pa/search?q=",
	"pe":  "com.pe/search?q=",
	"ph":  "com.ph/search?q=",
	"pk":  "com.pk/search?q=",
	"pl":  "pl/search?q=",
	"pg":  "com.pg/search?q=",
	"pn":  "pn/search?q=",
	"pr":  "com.pr/search?q=",
	"ps":  "ps/search?q=",
	"pt":  "pt/search?q=",
	"py":  "com.py/search?q=",
	"qa":  "com.qa/search?q=",
	"ro":  "ro/search?q=",
	"rs":  "rs/search?q=",
	"ru":  "ru/search?q=",
	"rw":  "rw/search?q=",
	"sa":  "com.sa/search?q=",
	"sb":  "com.sb/search?q=",
	"sc":  "sc/search?q=",
	"se":  "se/search?q=",
	"sg":  "com.sg/search?q=",
	"sh":  "sh/search?q=",
	"si":  "si/search?q=",
	"sk":  "sk/search?q=",
	"sl":  "com.sl/search?q=",
	"sn":  "sn/search?q=",
	"sm":  "sm/search?q=",
	"so":  "so/search?q=",
	"st":  "st/search?q=",
	"sv":  "com.sv/search?q=",
	"td":  "td/search?q=",
	"tg":  "tg/search?q=",
	"th":  "co.th/search?q=",
	"tj":  "com.tj/search?q=",
	"tk":  "tk/search?q=",
	"tl":  "tl/search?q=",
	"tm":  "tm/search?q=",
	"to":  "to/search?q=",
	"tn":  "tn/search?q=",
	"tr":  "com.tr/search?q=",
	"tt":  "tt/search?q=",
	"tw":  "com.tw/search?q=",
	"tz":  "co.tz/search?q=",
	"ua":  "com.ua/search?q=",
	"ug":  "co.ug/search?q=",
	"uk":  "co.uk/search?q=",
	"uy":  "com.uy/search?q=",
	"uz":  "co.uz/search?q=",
	"vc":  "com.vc/search?q=",
	"ve":  "co.ve/search?q=",
	"vg":  "vg/search?q=",
	"vi":  "co.vi/search?q=",
	"vn":  "com.vn/search?q=",
	"vu":  "vu/search?q=",
	"ws":  "ws/search?q=",
	"za":  "co.za/search?q=",
	"zm":  "co.zm/search?q=",
	"zw":  "co.zw/search?q=",
}

type SearchOptions struct {
	CountryCode  string
	LanguageCode string
	Limit        int
	Start        int
	UserAgent    string
	OverLimit    bool
	ProxyAddr    string
}

func WebAppear(ctx context.Context, searchTerm string,opts ...SearchOptions) ([]Result,error){
	if ctx == nil {
		ctx = context.Background()

	}
	if err := RateLimit.Wait(ctx); err != nil {
		return nil, err 
	}
	c := colly.NewCollector(colly.MaxDepth(1))
	if len(opts) == 0 {
		opts = append(opts, SearchOptions{})
	}

	if opts[0].UserAgent == "" {
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	} else {
		c.UserAgent = opts[0].UserAgent
	}

	var lc string
	if opts[0].LanguageCode == "" {
		lc = "en"
	} else {
		lc = opts[0].LanguageCode
	}

	results := []Result{}
	var rErr error
	rank := 1

	c.OnRequest(func(r *colly.Request) {
		if err := ctx.Err(); err != nil {
			r.Abort()
			rErr = err
			return
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		rErr = err
	})

	c.OnHTML("div.yuRUbf", func(e *colly.HTMLElement) {
		sel := e.DOM
        
		linkHref, _ := sel.Find("a").Attr("href")
		linkText := strings.TrimSpace(linkHref)
		titleText := strings.TrimSpace(sel.Find("a > h3").Text())

		if linkText != "" && linkText != "#" && titleText != "" {
			result := Result{
				Rank : rank, 
				URL:   linkText,
				Title: titleText,
			}
			results = append(results, result)
			rank += 1
		}
	})

	limit := opts[0].Limit
	if opts[0].OverLimit {
		limit = int(float64(opts[0].Limit) * 1.5)
	}
	/*
	 make the intitle (dorking) +  searchTerm
	*/
    front := "intitle:\""
	end := "\""
	searchTerm = front + searchTerm + end 
	
	url := url(searchTerm, opts[0].CountryCode, lc, limit, opts[0].Start)
	if opts[0].ProxyAddr != "" {
		rp, err := proxy.RoundRobinProxySwitcher(opts[0].ProxyAddr)
		if err != nil {
			return nil, err
		}
		c.SetProxyFunc(rp)
	}
	c.Visit(url)

	if rErr != nil {
		if strings.Contains(rErr.Error(), "Too many requests") {
			return nil, ErrBlocked
		}
		return nil, rErr
	}
	if opts[0].Limit != 0 && len(results) > opts[0].Limit {
		return results[:opts[0].Limit], nil
	}

	return results, nil
}

func Search(ctx context.Context, searchTerm string, opts ...SearchOptions) ([]Result, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := RateLimit.Wait(ctx); err != nil {
		return nil, err
	}

	c := colly.NewCollector(colly.MaxDepth(1))
	if len(opts) == 0 {
		opts = append(opts, SearchOptions{})
	}

	if opts[0].UserAgent == "" {
		c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
	} else {
		c.UserAgent = opts[0].UserAgent
	}

	var lc string
	if opts[0].LanguageCode == "" {
		lc = "en"
	} else {
		lc = opts[0].LanguageCode
	}

	results := []Result{}
	var rErr error
	rank := 1

	c.OnRequest(func(r *colly.Request) {
		if err := ctx.Err(); err != nil {
			r.Abort()
			rErr = err
			return
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		rErr = err
	})

	c.OnHTML("div.yuRUbf", func(e *colly.HTMLElement) {
		sel := e.DOM
        
		linkHref, _ := sel.Find("a").Attr("href")
		linkText := strings.TrimSpace(linkHref)
		titleText := strings.TrimSpace(sel.Find("a > h3").Text())

		if linkText != "" && linkText != "#" && titleText != "" {
			result := Result{
				Rank : rank, 
				URL:   linkText,
				Title: titleText,
			}
			results = append(results, result)
			rank += 1
		}
	})

	limit := opts[0].Limit
	if opts[0].OverLimit {
		limit = int(float64(opts[0].Limit) * 1.5)
	}

	url := url(searchTerm, opts[0].CountryCode, lc, limit, opts[0].Start)

	if opts[0].ProxyAddr != "" {
		rp, err := proxy.RoundRobinProxySwitcher(opts[0].ProxyAddr)
		if err != nil {
			return nil, err
		}
		c.SetProxyFunc(rp)
	}
	c.Visit(url)

	if rErr != nil {
		if strings.Contains(rErr.Error(), "Too many requests") {
			return nil, ErrBlocked
		}
		return nil, rErr
	}
	if opts[0].Limit != 0 && len(results) > opts[0].Limit {
		return results[:opts[0].Limit], nil
	}

	return results, nil
}


func base(url string) string {
	if strings.HasPrefix(url, "http") {
		return url
	} else {
		return BaseUrl + url
	}
}

func url(searchTerm string, countryCode string, languageCode string, limit int, start int) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	countryCode = strings.ToLower(countryCode)

	var url string

	if googleBase, found := GoogleDomains[countryCode]; found {
		if start == 0 {
			url = fmt.Sprintf("%s%s&hl=%s", base(googleBase), searchTerm, languageCode)
    
		} else {
			url = fmt.Sprintf("%s%s&hl=%s&start=%d", base(googleBase), searchTerm, languageCode, start)
      

    }
	} else {
		if start == 0 {
			url = fmt.Sprintf("%s%s&hl=%s", BaseUrl+GoogleDomains["us"], searchTerm, languageCode)
       

		} else {
			url = fmt.Sprintf("%s%s&hl=%s&start=%d", BaseUrl+GoogleDomains["us"], searchTerm, languageCode, start)
       
    
		}
	}

	if limit != 0 {
		url = fmt.Sprintf("%s&num=%d",url,limit)
	}
	return url
}
