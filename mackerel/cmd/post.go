package cmd

import (
	"fmt"
	mkr "github.com/mackerelio/mackerel-client-go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func init() {
	postCmd.Flags().StringVarP(&title, "title", "s", "", "required: annotation title")
	postCmd.Flags().StringVar(&description, "description", "", "[optional] annotation details")
	postCmd.Flags().StringVar(&service, "service", os.Getenv("MACKEREL_SERVICE_NAME"), "required: service name, when it's empty value then will use enviroment variable of 'MACKEREL_SERVICE_NAME''")
	RootCmd.AddCommand(postCmd)
}

var title string
var description string
var service string

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Creating graph annotations",
	Long:  "Tne post command creates graph annotations via graph-annotations API.",
	Example: `graph-annotations post -s 'deploy application' ExampleRole1 ExampleRole2
graph-annotations post --service ExampleService -s 'deploy application' ExampleRole1 ExampleRole2`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if title == "" {
			return errors.New("an annotation title is required")
		}

		if service == "" {
			return errors.New("a service name is required")
		}

		if len(args) == 0 {
			return errors.New("a role is required")
		}

		if os.Getenv("MACKEREL_API_KEY") == "" {
			return errors.New("a enviroment variable that named 'MACKEREL_API_KEY' is required")
		}

		time := time.Now().Unix()
		client := mkr.NewClient(os.Getenv("MACKEREL_API_KEY"))

		annotation := &mkr.GraphAnnotation{
			Service:     service,
			Roles:       args,
			From:        time,
			To:          time,
			Title:       title,
			Description: description,
		}

		err := client.CreateGraphAnnotation(annotation)

		if err != nil {
			return errors.Wrap(err, "client error.")
		}

		fmt.Printf("completed. params title:%s, from:%d to:%d, service:%s, roles:%s", title, time, time, service, args)
		return nil
	},
}
