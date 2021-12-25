package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
)

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://zenmod.shop/ec/odnorazovie_sigareti/"},
		ParseFunc: ZenmodParse,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}

func ZenmodParse(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.prdc__body").Each(func(i int, s *goquery.Selection) {
		g.Exports <- map[string]interface{}{
			"prdl__list":        s.Find("a.prdc__title").Text(),
			"prdc__attribute":   s.Find("span.prdc__attribute-text").Text(),
			"prdc__discription": s.Find("prdc__title").Find("href="),
		}

	})

	if href, ok := r.HTMLDoc.Find("li.next > a").Attr("href"); ok {
		g.Get(r.JoinURL(href), ZenmodParse)
	}
}
