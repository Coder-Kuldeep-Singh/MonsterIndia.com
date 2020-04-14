package response

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func MonsterResponsePage(selector string, resp *http.Response) (string, error) {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		fmt.Println("Go query failed!!!", err)
		return "", err
	}
	// res, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(res))
	output := ""
	doc.Find(selector).Each(func(index int, item *goquery.Selection) {
		contents, _ := item.Html()
		output = contents
	})
	if output == "" {
		log.Println("WARNING: output string is blank!!!")
	}
	// fmt.Println(output)
	return output, nil

}
