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
	"us":  "com/search?q=intitle:",
	"ac":  "ac/search?q=intitle:",
	"ad":  "ad/search?q=intitle:",
	"ae":  "ae/search?q=intitle:",
	"af":  "com.af/search?q=intitle:",
	"ag":  "com.ag/search?q=intitle:",
	"ai":  "com.ai/search?q=intitle:",
	"al":  "al/search?q=intitle:",
	"am":  "am/search?q=intitle:",
	"ao":  "co.ao/search?q=intitle:",
	"ar":  "com.ar/search?q=intitle:",
	"as":  "as/search?q=intitle:",
	"at":  "at/search?q=intitle:",
	"au":  "com.au/search?q=intitle:",
	"az":  "az/search?q=intitle:",
	"ba":  "ba/search?q=intitle:",
	"bd":  "com.bd/search?q=intitle:",
	"be":  "be/search?q=intitle:",
	"bf":  "bf/search?q=intitle:",
	"bg":  "bg/search?q=intitle:",
	"bh":  "com.bh/search?q=intitle:",
	"bi":  "bi/search?q=intitle:",
	"bj":  "bj/search?q=intitle:",
	"bn":  "com.bn/search?q=intitle:",
	"bo":  "com.bo/search?q=intitle:",
	"br":  "com.br/search?q=intitle:",
	"bs":  "bs/search?q=intitle:",
	"bt":  "bt/search?q=intitle:",
	"bw":  "co.bw/search?q=intitle:",
	"by":  "by/search?q=intitle:",
	"bz":  "com.bz/search?q=intitle:",
	"ca":  "ca/search?q=intitle:",
	"kh":  "com.kh/search?q=intitle:",
	"cc":  "cc/search?q=intitle:",
	"cd":  "cd/search?q=intitle:",
	"cf":  "cf/search?q=intitle:",
	"cat": "cat/search?q=intitle:",
	"cg":  "cg/search?q=intitle:",
	"ch":  "ch/search?q=intitle:",
	"ci":  "ci/search?q=intitle:",
	"ck":  "co.ck/search?q=intitle:",
	"cl":  "cl/search?q=intitle:",
	"cm":  "cm/search?q=intitle:",
	"cn":  "cn/search?q=intitle:",
	"co":  "com.co/search?q=intitle:",
	"cr":  "co.cr/search?q=intitle:",
	"cu":  "com.cu/search?q=intitle:",
	"cv":  "cv/search?q=intitle:",
	"cy":  "com.cy/search?q=intitle:",
	"cz":  "cz/search?q=intitle:",
	"de":  "de/search?q=intitle:",
	"dj":  "dj/search?q=intitle:",
	"dk":  "dk/search?q=intitle:",
	"dm":  "dm/search?q=intitle:",
	"do":  "com.do/search?q=intitle:",
	"dz":  "dz/search?q=intitle:",
	"ec":  "com.ec/search?q=intitle:",
	"ee":  "ee/search?q=intitle:",
	"eg":  "com.eg/search?q=intitle:",
	"es":  "es/search?q=intitle:",
	"et":  "com.et/search?q=intitle:",
	"fi":  "fi/search?q=intitle:",
	"fj":  "com.fj/search?q=intitle:",
	"fm":  "fm/search?q=intitle:",
	"fr":  "fr/search?q=intitle:",
	"ga":  "ga/search?q=intitle:",
	"gb":  "co.uk/search?q=intitle:",
	"ge":  "ge/search?q=intitle:",
	"gf":  "gf/search?q=intitle:",
	"gg":  "gg/search?q=intitle:",
	"gh":  "com.gh/search?q=intitle:",
	"gi":  "com.gi/search?q=intitle:",
	"gl":  "gl/search?q=intitle:",
	"gm":  "gm/search?q=intitle:",
	"gp":  "gp/search?q=intitle:",
	"gr":  "gr/search?q=intitle:",
	"gt":  "com.gt/search?q=intitle:",
	"gy":  "gy/search?q=intitle:",
	"hk":  "com.hk/search?q=intitle:",
	"hn":  "hn/search?q=intitle:",
	"hr":  "hr/search?q=intitle:",
	"ht":  "ht/search?q=intitle:",
	"hu":  "hu/search?q=intitle:",
	"id":  "co.id/search?q=intitle:",
	"iq":  "iq/search?q=intitle:",
	"ie":  "ie/search?q=intitle:",
	"il":  "co.il/search?q=intitle:",
	"im":  "im/search?q=intitle:",
	"in":  "co.in/search?q=intitle:",
	"io":  "io/search?q=intitle:",
	"is":  "is/search?q=intitle:",
	"it":  "it/search?q=intitle:",
	"je":  "je/search?q=intitle:",
	"jm":  "com.jm/search?q=intitle:",
	"jo":  "jo/search?q=intitle:",
	"jp":  "co.jp/search?q=intitle:",
	"ke":  "co.ke/search?q=intitle:",
	"ki":  "ki/search?q=intitle:",
	"kg":  "kg/search?q=intitle:",
	"kr":  "co.kr/search?q=intitle:",
	"kw":  "com.kw/search?q=intitle:",
	"kz":  "kz/search?q=intitle:",
	"la":  "la/search?q=intitle:",
	"lb":  "com.lb/search?q=intitle:",
	"lc":  "com.lc/search?q=intitle:",
	"li":  "li/search?q=intitle:",
	"lk":  "lk/search?q=intitle:",
	"ls":  "co.ls/search?q=intitle:",
	"lt":  "lt/search?q=intitle:",
	"lu":  "lu/search?q=intitle:",
	"lv":  "lv/search?q=intitle:",
	"ly":  "com.ly/search?q=intitle:",
	"ma":  "co.ma/search?q=intitle:",
	"md":  "md/search?q=intitle:",
	"me":  "me/search?q=intitle:",
	"mg":  "mg/search?q=intitle:",
	"mk":  "mk/search?q=intitle:",
	"ml":  "ml/search?q=intitle:",
	"mm":  "com.mm/search?q=intitle:",
	"mn":  "mn/search?q=intitle:",
	"ms":  "ms/search?q=intitle:",
	"mt":  "com.mt/search?q=intitle:",
	"mu":  "mu/search?q=intitle:",
	"mv":  "mv/search?q=intitle:",
	"mw":  "mw/search?q=intitle:",
	"mx":  "com.mx/search?q=intitle:",
	"my":  "com.my/search?q=intitle:",
	"mz":  "co.mz/search?q=intitle:",
	"na":  "com.na/search?q=intitle:",
	"ne":  "ne/search?q=intitle:",
	"nf":  "com.nf/search?q=intitle:",
	"ng":  "com.ng/search?q=intitle:",
	"ni":  "com.ni/search?q=intitle:",
	"nl":  "nl/search?q=intitle:",
	"no":  "no/search?q=intitle:",
	"np":  "com.np/search?q=intitle:",
	"nr":  "nr/search?q=intitle:",
	"nu":  "nu/search?q=intitle:",
	"nz":  "co.nz/search?q=intitle:",
	"om":  "com.om/search?q=intitle:",
	"pa":  "com.pa/search?q=intitle:",
	"pe":  "com.pe/search?q=intitle:",
	"ph":  "com.ph/search?q=intitle:",
	"pk":  "com.pk/search?q=intitle:",
	"pl":  "pl/search?q=intitle:",
	"pg":  "com.pg/search?q=intitle:",
	"pn":  "pn/search?q=intitle:",
	"pr":  "com.pr/search?q=intitle:",
	"ps":  "ps/search?q=intitle:",
	"pt":  "pt/search?q=intitle:",
	"py":  "com.py/search?q=intitle:",
	"qa":  "com.qa/search?q=intitle:",
	"ro":  "ro/search?q=intitle:",
	"rs":  "rs/search?q=intitle:",
	"ru":  "ru/search?q=intitle:",
	"rw":  "rw/search?q=intitle:",
	"sa":  "com.sa/search?q=intitle:",
	"sb":  "com.sb/search?q=intitle:",
	"sc":  "sc/search?q=intitle:",
	"se":  "se/search?q=intitle:",
	"sg":  "com.sg/search?q=intitle:",
	"sh":  "sh/search?q=intitle:",
	"si":  "si/search?q=intitle:",
	"sk":  "sk/search?q=intitle:",
	"sl":  "com.sl/search?q=intitle:",
	"sn":  "sn/search?q=intitle:",
	"sm":  "sm/search?q=intitle:",
	"so":  "so/search?q=intitle:",
	"st":  "st/search?q=intitle:",
	"sv":  "com.sv/search?q=intitle:",
	"td":  "td/search?q=intitle:",
	"tg":  "tg/search?q=intitle:",
	"th":  "co.th/search?q=intitle:",
	"tj":  "com.tj/search?q=intitle:",
	"tk":  "tk/search?q=intitle:",
	"tl":  "tl/search?q=intitle:",
	"tm":  "tm/search?q=intitle:",
	"to":  "to/search?q=intitle:",
	"tn":  "tn/search?q=intitle:",
	"tr":  "com.tr/search?q=intitle:",
	"tt":  "tt/search?q=intitle:",
	"tw":  "com.tw/search?q=intitle:",
	"tz":  "co.tz/search?q=intitle:",
	"ua":  "com.ua/search?q=intitle:",
	"ug":  "co.ug/search?q=intitle:",
	"uk":  "co.uk/search?q=intitle:",
	"uy":  "com.uy/search?q=intitle:",
	"uz":  "co.uz/search?q=intitle:",
	"vc":  "com.vc/search?q=intitle:",
	"ve":  "co.ve/search?q=intitle:",
	"vg":  "vg/search?q=intitle:",
	"vi":  "co.vi/search?q=intitle:",
	"vn":  "com.vn/search?q=intitle:",
	"vu":  "vu/search?q=intitle:",
	"ws":  "ws/search?q=intitle:",
	"za":  "co.za/search?q=intitle:",
	"zm":  "co.zm/search?q=intitle:",
	"zw":  "co.zw/search?q=intitle:",
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
