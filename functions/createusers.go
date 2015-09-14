package functions

import (
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/humankeyboard/awsmagic-tools/types"
	"io"
	"os"
)

func CreateUsers(profile, srcFile string) {
	fmt.Printf("Reading users from %s\n", srcFile)

	// check if a data file not provided
	if srcFile == "" {
		fmt.Printf("srcFile=%s\n", srcFile)
		fmt.Printf("provide a file with users -create-users -file={your file} \n")
		return
	}

	var users = []types.User{}

	// read configuration file provided
	csvFile, err := os.Open(srcFile)

	if err != nil {
		fmt.Printf("Can't open file: %s\n", srcFile)
		fmt.Println(err)
		return
	}

	defer csvFile.Close()

	// create User struct from csv file
	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			return
		}
		if len(line) != 3 {
			fmt.Printf("Line stated below doesn't have 3 values\n%s\n", line)
			return
		}
		users = append(users, types.User{line[0], line[1], line[2]})
	}

	// get IAM user to use in order to create other users
	c := types.NewCliConfig()

	creds := credentials.NewStaticCredentials(c.GetKeyID(profile), c.GetAccessKey(profile), "")

	svc := iam.New(&aws.Config{
		Credentials: creds,
	})

	for _, u := range users {
		if err := u.CreateUser(svc); err == nil {
			if err = u.CreateLoginProfile(svc); err == nil {
				if err = u.AddUserToGroup(svc); err == nil {
					fmt.Printf("Created user \"%s\" with password \"%s\" in group \"%s\"\n", u.UserName, u.Password, u.GroupName)
				}
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}
