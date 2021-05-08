package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)
type passedAPIResponse struct{
	Name string
	Age int
	Passed bool
	Skills []string
}



const passedUrl = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func linearBackoff(retry int) time.Duration {
	return time.Duration(retry) * time.Second
}

func writingToFile(f *os.File, decodedContent []passedAPIResponse){
	for i := range decodedContent{
		if decodedContent[i].Passed{
			var found = false
			//var knownSkills []string
			var knownSkillText = strings.Join(decodedContent[i].Skills, ", ")
			for _,val := range decodedContent[i].Skills{
				//knownSkills = append(knownSkills, val)

				if val == "Java" || val == "Go"{
					found = true
				}
			}
			if found{
				defer f.WriteString(fmt.Sprintf("%s - %v\n", decodedContent[i].Name, knownSkillText))
			}

		}

	}
}

func main() {
	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	httpResponse, err := httpClient.Get(passedUrl)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards passed API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of passed API response"),
		)
	}

	var decodedContent []passedAPIResponse
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	log.Printf("Response from passed: %v", decodedContent)

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	defer f.Close()
	writingToFile(f, decodedContent)
	f.Sync()


}