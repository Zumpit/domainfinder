package domainfinder


import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/proxy"
	"strings"
)

type Result struct {
	URL string `json:"url"`

	Title string `json:"title"`
}

const BaseUrl = "https://www.google."

var GoogleDomains = map[string]string{
	"us":  "com/search?q=title:",
	"ac":  "ac/search?q=title:",
	"ad":  "ad/search?q=title:",
	"ae":  "ae/search?q=title:",
	"af":  "com.af/search?q=title:",
	"ag":  "com.ag/search?q=title:",
	"ai":  "com.ai/search?q=title:",
	"al":  "al/search?q=title:",
	"am":  "am/search?q=title:",
	"ao":  "co.ao/search?q=title:",
	"ar":  "com.ar/search?q=title:",
	"as":  "as/search?q=title:",
	"at":  "at/search?q=title:",
	"au":  "com.au/search?q=title:",
	"az":  "az/search?q=title:",
	"ba":  "ba/search?q=title:",
	"bd":  "com.bd/search?q=title:",
	"be":  "be/search?q=title:",
	"bf":  "bf/search?q=title:",
	"bg":  "bg/search?q=title:",
	"bh":  "com.bh/search?q=title:",
	"bi":  "bi/search?q=title:",
	"bj":  "bj/search?q=title:",
	"bn":  "com.bn/search?q=title:",
	"bo":  "com.bo/search?q=title:",
	"br":  "com.br/search?q=title:",
	"bs":  "bs/search?q=title:",
	"bt":  "bt/search?q=title:",
	"bw":  "co.bw/search?q=title:",
	"by":  "by/search?q=title:",
	"bz":  "com.bz/search?q=title:",
	"ca":  "ca/search?q=title:",
	"kh":  "com.kh/search?q=title:",
	"cc":  "cc/search?q=title:",
	"cd":  "cd/search?q=title:",
	"cf":  "cf/search?q=title:",
	"cat": "cat/search?q=title:",
	"cg":  "cg/search?q=title:",
	"ch":  "ch/search?q=title:",
	"ci":  "ci/search?q=title:",
	"ck":  "co.ck/search?q=title:",
	"cl":  "cl/search?q=title:",
	"cm":  "cm/search?q=title:",
	"cn":  "cn/search?q=title:",
	"co":  "com.co/search?q=title:",
	"cr":  "co.cr/search?q=title:",
	"cu":  "com.cu/search?q=title:",
	"cv":  "cv/search?q=title:",
	"cy":  "com.cy/search?q=title:",
	"cz":  "cz/search?q=title:",
	"de":  "de/search?q=title:",
	"dj":  "dj/search?q=title:",
	"dk":  "dk/search?q=title:",
	"dm":  "dm/search?q=title:",
	"do":  "com.do/search?q=title:",
	"dz":  "dz/search?q=title:",
	"ec":  "com.ec/search?q=title:",
	"ee":  "ee/search?q=title:",
	"eg":  "com.eg/search?q=title:",
	"es":  "es/search?q=title:",
	"et":  "com.et/search?q=title:",
	"fi":  "fi/search?q=title:",
	"fj":  "com.fj/search?q=title:",
	"fm":  "fm/search?q=title:",
	"fr":  "fr/search?q=title:",
	"ga":  "ga/search?q=title:",
	"gb":  "co.uk/search?q=title:",
	"ge":  "ge/search?q=title:",
	"gf":  "gf/search?q=title:",
	"gg":  "gg/search?q=title:",
	"gh":  "com.gh/search?q=title:",
	"gi":  "com.gi/search?q=title:",
	"gl":  "gl/search?q=title:",
	"gm":  "gm/search?q=title:",
	"gp":  "gp/search?q=title:",
	"gr":  "gr/search?q=title:",
	"gt":  "com.gt/search?q=title:",
	"gy":  "gy/search?q=title:",
	"hk":  "com.hk/search?q=title:",
	"hn":  "hn/search?q=title:",
	"hr":  "hr/search?q=title:",
	"ht":  "ht/search?q=title:",
	"hu":  "hu/search?q=title:",
	"id":  "co.id/search?q=title:",
	"iq":  "iq/search?q=title:",
	"ie":  "ie/search?q=title:",
	"il":  "co.il/search?q=title:",
	"im":  "im/search?q=title:",
	"in":  "co.in/search?q=title:",
	"io":  "io/search?q=title:",
	"is":  "is/search?q=title:",
	"it":  "it/search?q=title:",
	"je":  "je/search?q=title:",
	"jm":  "com.jm/search?q=title:",
	"jo":  "jo/search?q=title:",
	"jp":  "co.jp/search?q=title:",
	"ke":  "co.ke/search?q=title:",
	"ki":  "ki/search?q=title:",
	"kg":  "kg/search?q=title:",
	"kr":  "co.kr/search?q=title:",
	"kw":  "com.kw/search?q=title:",
	"kz":  "kz/search?q=title:",
	"la":  "la/search?q=title:",
	"lb":  "com.lb/search?q=title:",
	"lc":  "com.lc/search?q=title:",
	"li":  "li/search?q=title:",
	"lk":  "lk/search?q=title:",
	"ls":  "co.ls/search?q=title:",
	"lt":  "lt/search?q=title:",
	"lu":  "lu/search?q=title:",
	"lv":  "lv/search?q=title:",
	"ly":  "com.ly/search?q=title:",
	"ma":  "co.ma/search?q=title:",
	"md":  "md/search?q=title:",
	"me":  "me/search?q=title:",
	"mg":  "mg/search?q=title:",
	"mk":  "mk/search?q=title:",
	"ml":  "ml/search?q=title:",
	"mm":  "com.mm/search?q=title:",
	"mn":  "mn/search?q=title:",
	"ms":  "ms/search?q=title:",
	"mt":  "com.mt/search?q=title:",
	"mu":  "mu/search?q=title:",
	"mv":  "mv/search?q=title:",
	"mw":  "mw/search?q=title:",
	"mx":  "com.mx/search?q=title:",
	"my":  "com.my/search?q=title:",
	"mz":  "co.mz/search?q=title:",
	"na":  "com.na/search?q=title:",
	"ne":  "ne/search?q=title:",
	"nf":  "com.nf/search?q=title:",
	"ng":  "com.ng/search?q=title:",
	"ni":  "com.ni/search?q=title:",
	"nl":  "nl/search?q=title:",
	"no":  "no/search?q=title:",
	"np":  "com.np/search?q=title:",
	"nr":  "nr/search?q=title:",
	"nu":  "nu/search?q=title:",
	"nz":  "co.nz/search?q=title:",
	"om":  "com.om/search?q=title:",
	"pa":  "com.pa/search?q=title:",
	"pe":  "com.pe/search?q=title:",
	"ph":  "com.ph/search?q=title:",
	"pk":  "com.pk/search?q=title:",
	"pl":  "pl/search?q=title:",
	"pg":  "com.pg/search?q=title:",
	"pn":  "pn/search?q=title:",
	"pr":  "com.pr/search?q=title:",
	"ps":  "ps/search?q=title:",
	"pt":  "pt/search?q=title:",
	"py":  "com.py/search?q=title:",
	"qa":  "com.qa/search?q=title:",
	"ro":  "ro/search?q=title:",
	"rs":  "rs/search?q=title:",
	"ru":  "ru/search?q=title:",
	"rw":  "rw/search?q=title:",
	"sa":  "com.sa/search?q=title:",
	"sb":  "com.sb/search?q=title:",
	"sc":  "sc/search?q=title:",
	"se":  "se/search?q=title:",
	"sg":  "com.sg/search?q=title:",
	"sh":  "sh/search?q=title:",
	"si":  "si/search?q=title:",
	"sk":  "sk/search?q=title:",
	"sl":  "com.sl/search?q=title:",
	"sn":  "sn/search?q=title:",
	"sm":  "sm/search?q=title:",
	"so":  "so/search?q=title:",
	"st":  "st/search?q=title:",
	"sv":  "com.sv/search?q=title:",
	"td":  "td/search?q=title:",
	"tg":  "tg/search?q=title:",
	"th":  "co.th/search?q=title:",
	"tj":  "com.tj/search?q=title:",
	"tk":  "tk/search?q=title:",
	"tl":  "tl/search?q=title:",
	"tm":  "tm/search?q=title:",
	"to":  "to/search?q=title:",
	"tn":  "tn/search?q=title:",
	"tr":  "com.tr/search?q=title:",
	"tt":  "tt/search?q=title:",
	"tw":  "com.tw/search?q=title:",
	"tz":  "co.tz/search?q=title:",
	"ua":  "com.ua/search?q=title:",
	"ug":  "co.ug/search?q=title:",
	"uk":  "co.uk/search?q=title:",
	"uy":  "com.uy/search?q=title:",
	"uz":  "co.uz/search?q=title:",
	"vc":  "com.vc/search?q=title:",
	"ve":  "co.ve/search?q=title:",
	"vg":  "vg/search?q=title:",
	"vi":  "co.vi/search?q=title:",
	"vn":  "com.vn/search?q=title:",
	"vu":  "vu/search?q=title:",
	"ws":  "ws/search?q=title:",
	"za":  "co.za/search?q=title:",
	"zm":  "co.zm/search?q=title:",
	"zw":  "co.zw/search?q=title:",
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
