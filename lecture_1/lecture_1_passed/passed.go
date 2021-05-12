package main

import (
	//imports
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
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

//checking if the specific skill is in a slice which contains all skills
func contains(skillsToCheck []string, searchedSkill string) bool {
	for _, skill := range skillsToCheck {
		if skill == searchedSkill{
			return true
		}
	}
	return false
}

func writeApplicationsToFile(f *os.File, applications []passedAPIResponse){
	//used to store all strings so we can write in a file only once
	var outputForTheFile string

	for _, applicant := range applications{

		if applicant.Passed{
			var foundJavaOrGoInSkills = false

			//checking if there is "Java" or "Go" in Skills
			if contains(applicant.Skills, "Java") || contains(applicant.Skills, "Go") {
				foundJavaOrGoInSkills = true
			}

			//each line looks like {Name} - {List of skill separated by a ","} with a newline at the end
			if foundJavaOrGoInSkills{
				outputForTheFile += applicant.Name + " - " + strings.Join(applicant.Skills, ", ") + "\n"
			}
		}
	}

	//error handling when writing to file
	_, err := f.WriteString(fmt.Sprintf("%v", outputForTheFile))
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "writing to file"),
		)
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

	var applications []passedAPIResponse
	err = json.Unmarshal(bodyContent, &applications)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	log.Printf("Response from passed API: %v", applications)

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	defer f.Close()
	writeApplicationsToFile(f, applications)
	f.Sync()

}
