package client

import (
	"fmt"
	"net/http"
	"net/url"
	"paper-summarizer/internal/domain"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type (
	ArxivClient struct{}
)

var _ domain.PaperFetcher = (*ArxivClient)(nil)

func NewArxivClient() *ArxivClient {
	return &ArxivClient{}
}

func (c *ArxivClient) Fetch(req domain.RequestFetch) (domain.Papers, error) {

	baseURL := "http://export.arxiv.org"
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u.Path = "/api/query"

	queryParams := url.Values{}
	query := url.QueryEscape(fmt.Sprintf("all:%s", req.Keyword))
	queryParams.Add("search_query", query)
	queryParams.Add("start", "0")
	queryParams.Add("max_results", fmt.Sprint(req.MaxNum))
	u.RawQuery = queryParams.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var papers domain.Papers
	doc.Find("entry").Each(func(_ int, entry *goquery.Selection) {
		title := strings.TrimSpace(entry.Find("title").Text())
		summary := strings.TrimSpace(entry.Find("summary").Text())
		url := entry.Find("id").Text()

		papers = append(papers, domain.Paper{
			Title:   title,
			Summary: summary,
			URL:     url,
		})
	})

	return papers, nil
}
