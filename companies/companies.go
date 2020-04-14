package companies

import (
	"fmt"
	"strconv"
	"strings"

	"MonsterIndia.com/proxy"
	"MonsterIndia.com/response"
	"github.com/PuerkitoBio/goquery"
)

// https://my.monsterindia.com/find-companies.html?sh=all
// https://my.monsterindia.com/find-companies.html
func FindCompanies() {
	BaseUrl := "https://my.monsterindia.com/find-companies.html"
	resp, err := proxy.MonsterDomain(BaseUrl)
	if err != nil {
		fmt.Println("Error On Proxy: ", err)
	}
	output, err := response.MonsterResponsePage("body", resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)

}

type CompanyResult struct {
	ResultRank               int
	ResultCompnay            string
	ResultProfileAndLocation string
	ResultActiveJobs         string
	ResultFollowers          string
}

func FindAllCompanies() ([]CompanyResult, int, error) {
	fmt.Println("There is 71 pages..")
	fmt.Println("How many pages You want to scrape? ")
	var count int
	fmt.Scanln(&count)
	BaseUrl := "https://my.monsterindia.com/find-companies.html?sh=all"
	resp, err := proxy.MonsterDomain(BaseUrl)
	if err != nil {
		return nil, 0, err
	}
	output, err := response.MonsterResponsePage(".row.find_rowfrst.mrgnbtm30 > h1", resp)
	if err != nil {
		return nil, 0, err
	}
	output = strings.TrimSpace(output)
	fmt.Println("Total: ", output)
	results := []CompanyResult{}
	var totalCompanies int
	for i := 1; i <= count; i++ {
		// fmt.Println(i)
		num := strconv.Itoa(i)
		resp, err := proxy.MonsterDomain(BaseUrl + "&p=" + num)
		if err != nil {
			return nil, 0, err
		}
		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			return nil, 0, err
		}
		self := doc.Find("span.cmpname")
		if self == nil {
			break
		}
		sel := doc.Find("div.mn-shdcmmn")
		rank := 1
		for i := range sel.Nodes {
			item := sel.Eq(i)
			compnayName := item.Find("span.cmpname")
			ProfileAndLocation := item.Find("div.cmp_txt")
			activeJobs := item.Find("a.mn-lnk1")
			Followers := item.Find("div.fc_btxt > span#183088_cnt")

			name := compnayName.Text()
			name = strings.TrimSpace(name)
			profile := ProfileAndLocation.Text()
			profile = strings.TrimSpace(profile)
			jobs := activeJobs.Text()
			jobs = strings.TrimSpace(jobs)
			follow := Followers.Text()
			follow = strings.TrimSpace(follow)
			result := CompanyResult{
				rank,
				name,
				profile,
				jobs,
				follow,
			}
			results = append(results, result)
			rank += 1
		}
		for _, item := range results {
			totalCompanies = item.ResultRank
		}
	}
	return results, totalCompanies, nil

}

// https: //my.monsterindia.com/find-companies.html?l=D
func FindCompaniesByCharacter(Alphabat string) ([]CompanyResult, int, error) {
	Alphabat = strings.ToUpper(Alphabat)
	fmt.Println("How many pages You want to scrape? ")
	var count int
	fmt.Scanln(&count)
	BaseUrl := "https://my.monsterindia.com/find-companies.html?l=" + Alphabat
	resp, err := proxy.MonsterDomain(BaseUrl)
	if err != nil {
		return nil, 0, err
	}
	output, err := response.MonsterResponsePage(".row.find_rowfrst.mrgnbtm30 > h1", resp)
	if err != nil {
		return nil, 0, err
	}
	output = strings.TrimSpace(output)
	fmt.Println("Total: ", output)
	results := []CompanyResult{}
	var totalCompanies int
	for i := 1; i <= count; i++ {
		// fmt.Println(i)
		num := strconv.Itoa(i)
		resp, err := proxy.MonsterDomain(BaseUrl + "&p=" + num)
		if err != nil {
			return nil, 0, err
		}
		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			return nil, 0, err
		}
		self := doc.Find("span.cmpname")
		if self == nil {
			break
		}
		sel := doc.Find("div.mn-shdcmmn")
		rank := 1
		for i := range sel.Nodes {
			item := sel.Eq(i)
			compnayName := item.Find("span.cmpname")
			ProfileAndLocation := item.Find("div.cmp_txt")
			activeJobs := item.Find("a.mn-lnk1")
			Followers := item.Find("div.fc_btxt > span#183088_cnt")

			name := compnayName.Text()
			name = strings.TrimSpace(name)
			profile := ProfileAndLocation.Text()
			profile = strings.TrimSpace(profile)
			jobs := activeJobs.Text()
			jobs = strings.TrimSpace(jobs)
			follow := Followers.Text()
			follow = strings.TrimSpace(follow)
			result := CompanyResult{
				rank,
				name,
				profile,
				jobs,
				follow,
			}
			results = append(results, result)
			rank += 1
		}
		for _, item := range results {
			totalCompanies = item.ResultRank
		}
	}
	return results, totalCompanies, nil

}
