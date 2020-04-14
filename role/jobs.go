package role

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"MonsterIndia.com/proxy"
	"MonsterIndia.com/response"
	"github.com/PuerkitoBio/goquery"
)

type JobResult struct {
	ResultRank        int
	ResultJob         string
	ResultCompnay     string
	ResultLocation    string
	ResultIncome      string
	ResultDescription string
	ResultSkills      string
}

func FindJobsByRole() {

	jobtitle := bufio.NewReader(os.Stdin)
	fmt.Println("Please Provide the Job Title..")
	input, _ := jobtitle.ReadString('\n')
	title := strings.Replace(input, " ", "-", -1)
	title = strings.Replace(title, "\n", "", -1)
	title = strings.ToLower(title)
	BaseUrl := "https://www.monsterindia.com/search/" + title
	resp, err := proxy.MonsterDomain(BaseUrl)
	if err != nil {
		fmt.Println("Proxy server Failed!! ", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	SecretIdForRedirectionForJobByRole := ""
	re := regexp.MustCompile(`"applicationID":"(.*)"`)
	ApplicationID := re.FindAllStringSubmatch(string(body), -1)

	for _, url := range ApplicationID {
		id := strings.Split(url[1], ",")
		newid := id[0]
		newid = strings.Replace(newid, `"`, "", -1)
		SecretIdForRedirectionForJobByRole = newid
	}

	SecretIdKey := ""
	reg := regexp.MustCompile(`"licenseKey":"(.*)"`)
	ApplicationKey := reg.FindAllStringSubmatch(string(body), -1)

	for _, url := range ApplicationKey {
		id := strings.Split(url[1], ",")
		newid := id[0]
		newid = strings.Replace(newid, `"`, "", -1)
		secret1 := newid[0:4]
		secret2 := newid[6:10]
		secret3 := newid[11:15]
		secret4 := newid[16:]
		SecretIdKey = secret1 + "-" + secret2 + "-" + secret3 + "-" + secret4

	}

	// https://www.monsterindia.com/search/accountant-jobs?searchId=53847430-NRJS-d928-22b4-067d8977e54
	respAgain, err := proxy.MonsterDomain(BaseUrl + "?searchId=" + SecretIdForRedirectionForJobByRole + SecretIdKey)
	if err != nil {
		log.Println("Error failed proxy with secret id!!", err)
	}
	output, err := response.MonsterResponsePage("div.lft-content strong.fs-24.normal.ffm-arial", respAgain)
	if err != nil {
		log.Println("Error to find Total jobs!!.", err)
		// return nil, 0, err
	}
	output = strings.TrimSpace(output)
	fmt.Printf("Total: %s %s\n", output, input)
	fmt.Println("How many pages You want to scrape?")
	var count int
	fmt.Scanln(&count)
	for i := 1; i <= count; i++ {
		num := strconv.Itoa(i)
		resp, err := proxy.MonsterDomain(BaseUrl + "-" + num + "?searchId=" + SecretIdForRedirectionForJobByRole + SecretIdKey)
		if err != nil {
			log.Println("Error failed proxy with secret id!!", err)
		}
		// body, _ := ioutil.ReadAll(resp.Body)
		// fmt.Println(string(body))
		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			// return nil, 0, err
			fmt.Println(err)
		}
		self := doc.Find("h3.medium > a")
		if self == nil {
			break
		}
		results := []JobResult{}
		sel := doc.Find("div.card-apply-content")
		rank := 1
		var totalCompanies int

		for i := range sel.Nodes {
			item := sel.Eq(i)
			profileName := item.Find("h3.medium > a")
			compnayName := item.Find("span.company-name > a")
			activelocation := item.Find("div.col-xxs-12.col-sm-5.text-ellipsis > span.loc > small")
			activeIncome := item.Find("div.package.col-xxs-12.col-sm-4.text-ellipsis > span.loc > small")
			activeDescription := item.Find("p.job-descrip")
			Skills := item.Find("p.descrip-skills > label.grey-link > a")

			profile := profileName.Text()
			// profile = strings.TrimSpace(profile)
			profile = strings.Replace(profile, " ", "", -1)

			name := compnayName.Text()
			// name = strings.TrimSpace(name)
			name = strings.Replace(name, " ", "", -1)

			location := activelocation.Text()
			// location = strings.TrimSpace(location)
			location = strings.Replace(location, " ", "", -1)

			Income := activeIncome.Text()
			// Income = strings.TrimSpace(Income)
			Income = strings.Replace(Income, " ", "", -1)

			Description := activeDescription.Text()
			Description = strings.TrimSpace(Description)
			// Description = strings.Replace(Description, " ", "", -1)

			skills := Skills.Text()
			skills = strings.TrimSpace(skills)
			skills = strings.Replace(skills, " ", "", -1)

			result := JobResult{
				rank,
				profile,
				name,
				location,
				Income,
				Description,
				skills,
			}
			results = append(results, result)
			rank += 1
		}
		for _, item := range results {
			totalCompanies = item.ResultRank
			fmt.Println(item.ResultDescription)
		}
		fmt.Println(totalCompanies)

	}

}
