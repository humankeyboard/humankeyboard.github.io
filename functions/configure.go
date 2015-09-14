package functions

import (
	"fmt"
	"github.com/humankeyboard/awsmagic-tools/types"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Configure() {
	fmt.Printf("inside configure\n")

	c := types.NewCliConfig()
	fmt.Printf("Modifying \"%s\" \n\n", c.CredentialsFile)

	var Profile string

	if len(c.GetProfiles()) < 1 {
		fmt.Printf("Profile Name [default]: ")
	} else {
		fmt.Printf("Choose an existing profile to update or enter a new one to create: \n")
		fmt.Printf("Available profiles: [ ")
		for i, profile := range c.GetProfiles() {
			fmt.Printf("%s", profile)
			if i+1 < len(c.GetProfiles()) {
				fmt.Printf(", ")
			}
		}
		fmt.Printf(" ]\n")
		fmt.Printf("Profile Name [default]: ")
		fmt.Scanf("%s", &Profile)
		if Profile == "" {
			Profile = "default"
		}
	}

	var newProfile = true
	for _, profile := range c.GetProfiles() {
		if Profile == profile {
			newProfile = false
		}
	}

	var AccessKeyID string
	var SecretAccessKey string

	if newProfile {
		AccessKeyID = ""
		SecretAccessKey = ""
	} else {
		AccessKeyID = c.GetKeyID(Profile)
		SecretAccessKey = c.GetAccessKey(Profile)
	}

	if AccessKeyID == "" {
		AccessKeyID = "None"
	}

	fmt.Printf("AWS Access Key ID [%s]: ", AccessKeyID)
	fmt.Scanf("%s", &AccessKeyID)

	if SecretAccessKey == "" {
		SecretAccessKey = "None"
	}

	fmt.Printf("AWS Secret Access Key [%s]: ", SecretAccessKey)
	fmt.Scanf("%s", &SecretAccessKey)

	// goconfig updates profile or creates new one
	c.WriteProfile(Profile, AccessKeyID, SecretAccessKey)

}
