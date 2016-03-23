# go-ask-awesomely

[![GoDoc](https://godoc.org/github.com/gagliardetto/go-ask-awesomely?status.svg)](https://godoc.org/github.com/gagliardetto/go-ask-awesomely)
[![GitHub license](https://img.shields.io/github/license/gagliardetto/go-ask-awesomely.svg)](https://github.com/gagliardetto/go-ask-awesomely/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/gagliardetto/go-ask-awesomely)](https://goreportcard.com/report/github.com/gagliardetto/go-ask-awesomely)

## Description

Go client for Typeform API.

## Installation

Run the following command to install the package:

```
go get -u github.com/gagliardetto/go-ask-awesomely
```

## Examples

#### Get info

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	client.Config.APIKey = os.Getenv("TYPEFORM_API_KEY")

	baseInfo, err := client.BaseInfo()
	if err != nil {
		fmt.Println("baseInfo error: ", err)
		return
	}

	fmt.Printf("\nbaseInfo: %v\n", beautify(baseInfo))
}
```

#### Create a form

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	client.Config.APIKey = os.Getenv("TYPEFORM_API_KEY")

	newForm := tf.Form{
		Title:    "new form",
		Branding: true,
		Fields: []tf.Field{
			tf.Field{
				Type:          tf.ShortText,
				Question:      "what is that?",
				Tags:          []string{"something"},
				MaxCharacters: 3,
			},

			tf.Field{
				Type:          tf.LongText,
				Question:      "what is the story of your life?",
				Tags:          []string{"something"},
				MaxCharacters: 3000,
			},

			tf.Field{
				Type:                    tf.MultipleChoice,
				Question:                "what is the story of your life?",
				Description:             "some description",
				AllowMultipleSelections: true,
				Randomize:               false,
				VerticalAlignment:       false,
				AddOtherChoice:          true,
				Tags:                    []string{"something"},
				Choices: []tf.Choice{
					tf.Choice{
						Label: "this",
					},
					tf.Choice{
						Label: "that",
					},
					tf.Choice{
						Label: "third",
					},
				},
			},

			/*
				tf.Field{
					Type:                    tf.PictureChoice,
					Question:                "Choose images",
					Description:             "some description",
					ShowLabels:              true,
					Supersize:               true,
					AllowMultipleSelections: true,
					Randomize:               false,
					AddOtherChoice:          true,
					Tags:                    []string{"something"},
					Required:                true,
					Choices: []tf.Choice{
						tf.Choice{
							ImageID: "HNdAk47LS",
							Label:   "this",
						},
						tf.Choice{
							ImageID: "L2DsjN8JS",
							Label:   "that",
						},
						tf.Choice{
							ImageID: "DLs2d43NS",
							Label:   "third",
						},
					},
				},
			*/

			tf.Field{
				Type:       tf.Statement,
				Question:   "This is a statement",
				Tags:       []string{"something"},
				ButtonText: "Ok",
				HideMarks:  false,
			},

			tf.Field{
				Type:     tf.Dropdown,
				Question: "Choose from dropdown",
				Tags:     []string{"something"},
				Choices: []tf.Choice{
					tf.Choice{
						Label: "Europe",
					},
					tf.Choice{
						Label: "Asia",
					},
					tf.Choice{
						Label: "USA",
					},
				},
			},

			tf.Field{
				Type:     tf.YesNo,
				Question: "Yes or no?",
				Tags:     []string{"something"},
				Required: true,
			},

			tf.Field{
				Type:        tf.Number,
				Question:    "How many cats do you have?",
				Description: "some description",
				Tags:        []string{"something"},
				MinValue:    0,
				MaxValue:    99999,
			},

			tf.Field{
				Type:        tf.Rating,
				Question:    "Rate",
				Description: "some description",
				Tags:        []string{"something"},
				Steps:       5,
				Shape:       "star", // \"star\", \"heart\", \"user\", \"up\", \"crown\", \"cat\", \"dog\", \"circle\", \"flag\", \"droplet\", \"tick\", \"lightbulb\", \"trophy\", \"cloud\", \"thunderbolt\", \"pencil\", \"skull\"
			},

			tf.Field{
				Type:        tf.OpinionScale,
				Question:    "Opinion scale",
				Description: "some description",
				Tags:        []string{"something"},
				Labels: &tf.Labels{
					Left:   "Forms suck",
					Center: "Who cares",
					Right:  "I love you",
				},
			},

			tf.Field{
				Type:        tf.Email,
				Question:    "What is your email?",
				Description: "some description",
				Tags:        []string{"something"},
				Required:    true,
			},

			tf.Field{
				Type:        tf.Website,
				Question:    "What is your website?",
				Description: "some description",
				Tags:        []string{"something"},
				Required:    false,
			},

			tf.Field{
				Type:        tf.Legal,
				Question:    "Do you agree to our terms?",
				Description: "some description",
				Tags:        []string{"something"},
				Required:    true,
			},
		},
	}

	resp, err := client.CreateForm(newForm)
	if err != nil {
		fmt.Println("Do error: ", err)
		return
	}
	fmt.Printf("\nDo: %v\n", beautify(resp))

}

func beautify(object interface{}) string {
	out, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(out)
}

```
