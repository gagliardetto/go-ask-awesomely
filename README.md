# go-ask-awesomely

[![GoDoc](https://godoc.org/github.com/gagliardetto/go-ask-awesomely?status.svg)](https://godoc.org/github.com/gagliardetto/go-ask-awesomely)
[![GitHub license](https://img.shields.io/github/license/gagliardetto/go-ask-awesomely.svg)](https://github.com/gagliardetto/go-ask-awesomely/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/gagliardetto/go-ask-awesomely)](https://goreportcard.com/report/github.com/gagliardetto/go-ask-awesomely)

## Description

Golang client for Typeform API.

Unofficial.

Complete API v0.4 implementation.

With test suite.

# TODO

- Implement https://www.typeform.com/help/data-api/ (API to retrieve form responses)

## How to get an API key

You can get an API key from http://typeform.io/

## How to get a throw-away API key to run tests (and not pollute the actual account)

- Go to http://docs.typeform.io/page/sandbox

- Open Chrome Developer Tools

- In the page http://docs.typeform.io/page/sandbox, click on the "Build" button. This will fire a request to the Typeform API.

- In the Developer Tools, go to the `Network` tab.

- Look for a `POST` request sent to `https://api.typeform.io/v0.4/forms`

- In the `Request Headers` section, look for the `X-API-TOKEN` header, and copy the value; this is a valid API token you can use to run tests.

- To use this token for running this golang package's tests, export the `TYPEFORM_TEST_API_KEY` environment variable, e.g. `$  export TYPEFORM_TEST_API_KEY=0a000aa00aa00a0aa00aaa0a00a0000a`

## Installation

Run the following command to install the package:

```
go get -u github.com/gagliardetto/go-ask-awesomely
```

# Testing

```bash
$  export TYPEFORM_TEST_API_KEY=0a000cd00ae00a0dd00caa0c00e0000d
$ go test -v
```

## API Usage Examples (complete)

#### Get API info

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	baseInfo, err := client.BaseInfo()
	if err != nil {
		fmt.Println("baseInfo error: ", err)
		return
	}

	fmt.Printf("\nAPI info: %#v\n", baseInfo)
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
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	newForm := tf.Form{
		Title:    "My amazing new form",
		Branding: true,
		//Tags:             []string{},
		//DesignID:         "<design ID>",
		//WebhookSubmitURL: "<webhook submit URL>",
		//URLIDs:           []string{},		

		LogicJumps: []tf.LogicJump{
			tf.LogicJump{
				From: "decisive-question",
				To:   "jump-here",
				If:   true,
			},
		},

		Fields: []tf.Field{
			tf.Field{
				Type:          tf.ShortText,
				Question:      "What are your favorite 3 characters?",
				Tags:          []string{"some-tag"},
				MaxCharacters: 3,
			},

			tf.Field{
				Type:          tf.LongText,
				Question:      "what is the story of your life?",
				Tags:          []string{"some-tag"},
				MaxCharacters: 3000,
			},

			tf.Field{
				Type:                    tf.MultipleChoice,
				Question:                "Please select a few choices",
				Description:             "some description",
				AllowMultipleSelections: true,
				Randomize:               false,
				VerticalAlignment:       false,
				AddOtherChoice:          true,
				Tags:                    []string{"some-tag"},
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
					Tags:                    []string{"some-tag"},
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
				Tags:       []string{"some-tag"},
				ButtonText: "Ok",
				HideMarks:  false,
			},

			tf.Field{
				Type:     tf.Dropdown,
				Question: "Choose from dropdown",
				Tags:     []string{"some-tag"},
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
				Question: "Do you wanna jump?",
				Tags:     []string{"some-tag"},
				Ref:      "decisive-question",
				Required: true,
			},

			tf.Field{
				Type:        tf.Number,
				Question:    "How many cats do you have?",
				Description: "some description",
				Tags:        []string{"some-tag"},
				MinValue:    0,
				MaxValue:    99999,
			},

			tf.Field{
				Type:        tf.Rating,
				Question:    "Rate",
				Description: "You probably jumped here from yes/no question",
				Tags:        []string{"some-tag"},
				Ref:         "jump-here",
				Steps:       5,
				Shape:       "star", // Alternatives: "star", "heart", "user", "up", "crown", "cat", "dog", "circle", "flag", "droplet", "tick", "lightbulb", "trophy", "cloud", "thunderbolt", "pencil", "skull"
			},

			tf.Field{
				Type:        tf.OpinionScale,
				Question:    "Opinion scale",
				Description: "some description",
				Tags:        []string{"some-tag"},
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
				Tags:        []string{"some-tag"},
				Required:    true,
			},

			tf.Field{
				Type:        tf.Website,
				Question:    "What is your website?",
				Description: "some description",
				Tags:        []string{"some-tag"},
				Required:    false,
			},

			tf.Field{
				Type:        tf.Legal,
				Question:    "Do you agree to our terms?",
				Description: "some description",
				Tags:        []string{"some-tag"},
				Required:    true,
			},
		},
	}

	resp, err := client.CreateForm(newForm)
	if err != nil {
		fmt.Println("CreateForm error: ", err)
		return
	}

	fmt.Printf("\nNew form: %v\n", beautify(resp))
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

#### Get info about a form

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	formInfo, err := client.GetForm("<form ID>")
	if err != nil {
		fmt.Println("GetForm error: ", err)
		return
	}

	fmt.Printf("\nForm info: %#v\n", formInfo)
}
```

#### Add image to typeform

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	newImage, err := client.CreateImage("https://www.google.it/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png")
	if err != nil {
		fmt.Println("CreateImage error: ", err)
		return
	}

	fmt.Printf("\nNew image info: %#v\n", newImage)
}
```

#### Get info about an image

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	imageInfo, err := client.GetImage("t7iWgPKgEk")
	if err != nil {
		fmt.Println("GetImage error: ", err)
		return
	}

	fmt.Printf("\nImage info: %#v\n", imageInfo)
}
```

#### Create a new design

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	newDesign := tf.Design{
		Colors: tf.Colors{
			Question:   "#3D3D3D",
			Button:     "#4FB0AE",
			Answer:     "#4FB0AE",
			Background: "#FFFFFF",
		},
		Font: "Source Sans Pro",
	}

	newDesignInfo, err := client.CreateDesign(newDesign)
	if err != nil {
		fmt.Println("CreateDesign error: ", err)
		return
	}

	fmt.Printf("\nNew design info: %#v\n", newDesignInfo)
}
```

#### Get info about a design

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	designInfo, err := client.GetDesign("<design ID>")
	if err != nil {
		fmt.Println("GetDesign error: ", err)
		return
	}

	fmt.Printf("\nDesign info: %#v\n", designInfo)
}
```

#### Create a new URL for a form

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	newFormURL, err := client.CreateURL("<form ID>")
	if err != nil {
		fmt.Println("CreateURL error: ", err)
		return
	}

	fmt.Printf("\nNew form URL info: %#v\n", newFormURL)
}
```

#### Get info about an URL

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	URLInfo, err := client.GetURL("<URL ID>")
	if err != nil {
		fmt.Println("GetURL error: ", err)
		return
	}

	fmt.Printf("\nURL info: %#v\n", URLInfo)
}
```

#### Modify URL

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	modifiedURLInfo, err := client.ModifyURL("<URL ID>", "<form ID>")
	if err != nil {
		fmt.Println("ModifyURL error: ", err)
		return
	}

	fmt.Printf("\nModified URL info: %#v\n", modifiedURLInfo)
}
```

#### Delete URL

```go
package main

import (
	"encoding/json"
	"fmt"
	tf "github.com/gagliardetto/go-ask-awesomely"
)

func main() {
	client, err := tf.NewClient(tf.Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}		

	token := os.Getenv("TYPEFORM_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}

	err := client.DeleteURL("<URL ID>")
	if err != nil {
		fmt.Println("DeleteURL error: ", err)
		return
	}
}
```
