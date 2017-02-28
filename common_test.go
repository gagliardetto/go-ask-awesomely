package typeform

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	APIKey string
	client *Client
)

func init() {
	APIKey = getTestKey()

	var err error
	client, err = NewClient(Latest)
	if err != nil {
		fmt.Println("client setup error: ", err)
		return
	}

	token := os.Getenv("TYPEFORM_TEST_API_KEY")
	err = client.SetAPIToken(token)
	if err != nil {
		fmt.Println("token error: ", err)
		return
	}
}

func getTestKey() string {
	key := os.Getenv("TYPEFORM_TEST_API_KEY")

	if len(key) == 0 {
		panic("TYPEFORM_TEST_API_KEY environment variable is not set, but is needed to run the tests!\n")
	}

	return key
}

func beautify(object interface{}) string {
	out, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(out)
}

func TestFetchAndReturnPage(t *testing.T) {
	testBody := `{some:"json"}`

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, testBody)
	}))
	defer testServer.Close()

	path := fmt.Sprintf(testServer.URL)
	method := http.MethodGet
	headers := http.Header{}
	queryParameters := url.Values{}
	var bodyPayload interface{}

	// need to make the request with no URL
	APIDomain = testServer.URL
	defer func() {
		APIDomain = "https://api.typeform.io/"
	}()

	response, _, err := client.fetchAndReturnPage(path, method, headers, queryParameters, bodyPayload)
	assert.Nil(t, err, "no error should occur")

	assert.Equal(t, testBody+"\n", string(response), "the two bodies should be equal")
}

func TestBaseInfo(t *testing.T) {
	baseInfo, err := client.BaseInfo()
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, baseInfo, "returned response object pointer should not be nil")

	fmt.Printf("\nAPI info: %#v\n", baseInfo)
}

func TestCreateForm(t *testing.T) {
	newForm := Form{
		Title:    "My amazing new form",
		Branding: true,
		//Tags:             []string{},
		//DesignID:         "<design ID>",
		//WebhookSubmitURL: "<webhook submit URL>",
		//URLIDs:           []string{},

		LogicJumps: []LogicJump{
			LogicJump{
				From: "decisive-question",
				To:   "jump-here",
				If:   true,
			},
		},

		Fields: []Field{
			Field{
				Type:          ShortText,
				Question:      "What are your favorite 3 characters?",
				Tags:          []string{"some-tag"},
				MaxCharacters: 3,
			},

			Field{
				Type:          LongText,
				Question:      "what is the story of your life?",
				Tags:          []string{"some-tag"},
				MaxCharacters: 3000,
			},

			Field{
				Type:                    MultipleChoice,
				Question:                "Please select a few choices",
				Description:             "some description",
				AllowMultipleSelections: true,
				Randomize:               false,
				VerticalAlignment:       false,
				AddOtherChoice:          true,
				Tags:                    []string{"some-tag"},
				Choices: []Choice{
					Choice{
						Label: "this",
					},
					Choice{
						Label: "that",
					},
					Choice{
						Label: "third",
					},
				},
			},

			/*
			   Field{
			       Type:                    PictureChoice,
			       Question:                "Choose images",
			       Description:             "some description",
			       ShowLabels:              true,
			       Supersize:               true,
			       AllowMultipleSelections: true,
			       Randomize:               false,
			       AddOtherChoice:          true,
			       Tags:                    []string{"some-tag"},
			       Required:                true,
			       Choices: []Choice{
			           Choice{
			               ImageID: "HNdAk47LS",
			               Label:   "this",
			           },
			           Choice{
			               ImageID: "L2DsjN8JS",
			               Label:   "that",
			           },
			           Choice{
			               ImageID: "DLs2d43NS",
			               Label:   "third",
			           },
			       },
			   },
			*/

			Field{
				Type:       Statement,
				Question:   "This is a statement",
				Tags:       []string{"some-tag"},
				ButtonText: "Ok",
				HideMarks:  false,
			},

			Field{
				Type:     Dropdown,
				Question: "Choose from dropdown",
				Tags:     []string{"some-tag"},
				Choices: []Choice{
					Choice{
						Label: "Europe",
					},
					Choice{
						Label: "Asia",
					},
					Choice{
						Label: "USA",
					},
				},
			},

			Field{
				Type:     YesNo,
				Question: "Do you wanna jump?",
				Tags:     []string{"some-tag"},
				Ref:      "decisive-question",
				Required: true,
			},

			Field{
				Type:        Number,
				Question:    "How many cats do you have?",
				Description: "some description",
				Tags:        []string{"some-tag"},
				MinValue:    0,
				MaxValue:    99999,
			},

			Field{
				Type:        Rating,
				Question:    "Rate",
				Description: "You probably jumped here from yes/no question",
				Tags:        []string{"some-tag"},
				Ref:         "jump-here",
				Steps:       5,
				Shape:       "star", // Alternatives: "star", "heart", "user", "up", "crown", "cat", "dog", "circle", "flag", "droplet", "tick", "lightbulb", "trophy", "cloud", "thunderbolt", "pencil", "skull"
			},

			Field{
				Type:        OpinionScale,
				Question:    "Opinion scale",
				Description: "some description",
				Tags:        []string{"some-tag"},
				Labels: &Labels{
					Left:   "Forms suck",
					Center: "Who cares",
					Right:  "I love you",
				},
			},

			Field{
				Type:        Email,
				Question:    "What is your email?",
				Description: "some description",
				Tags:        []string{"some-tag"},
				Required:    true,
			},

			Field{
				Type:        Website,
				Question:    "What is your website?",
				Description: "some description",
				Tags:        []string{"some-tag"},
				Required:    false,
			},

			Field{
				Type:        Legal,
				Question:    "Do you agree to our terms?",
				Description: "some description",
				Tags:        []string{"some-tag"},
				Required:    true,
			},
		},
	}

	resp, err := client.CreateForm(newForm)
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, resp, "returned response object pointer should not be nil")

	fmt.Printf("\nNew form: %v\n", beautify(resp))

	_TestGetForm(t, resp.ID)
	_TestCreateURL(t, resp.ID)
}

func _TestGetForm(t *testing.T, formID string) {
	formInfo, err := client.GetForm(formID)
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, formInfo, "returned response object pointer should not be nil")

	fmt.Printf("\nForm info: %#v\n", formInfo)
}

func TestCreateImage(t *testing.T) {
	newImage, err := client.CreateImage("https://www.google.it/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png")
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, newImage, "returned response object pointer should not be nil")

	fmt.Printf("\nNew image info: %#v\n", newImage)

	_TestGetImage(t, newImage.ID)
}

func _TestGetImage(t *testing.T, imageID string) {
	imageInfo, err := client.GetImage(imageID)
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, imageInfo, "returned response object pointer should not be nil")

	fmt.Printf("\nImage info: %#v\n", imageInfo)
}

func TestCreateDesign(t *testing.T) {
	newDesign := Design{
		Colors: Colors{
			Question:   "#3D3D3D",
			Button:     "#4FB0AE",
			Answer:     "#4FB0AE",
			Background: "#FFFFFF",
		},
		Font: "Source Sans Pro",
	}

	newDesignInfo, err := client.CreateDesign(newDesign)
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, newDesignInfo, "returned response object pointer should not be nil")

	fmt.Printf("\nNew design info: %#v\n", newDesignInfo)

	_TestGetDesign(t, newDesignInfo.ID)
}

func _TestGetDesign(t *testing.T, designID string) {
	designInfo, err := client.GetDesign(designID)
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, designInfo, "returned response object pointer should not be nil")

	fmt.Printf("\nDesign info: %#v\n", designInfo)
}

func _TestCreateURL(t *testing.T, formID string) {
	newFormURL, err := client.CreateURL(formID)
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, newFormURL, "returned response object pointer should not be nil")

	fmt.Printf("\nNew form URL info: %#v\n", newFormURL)

	_TestGetURL(t, newFormURL.ID)
	_TestModifyURL(t, newFormURL.ID, formID)
}

func _TestGetURL(t *testing.T, URLID string) {
	URLInfo, err := client.GetURL(URLID)
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, URLInfo, "returned response object pointer should not be nil")

	fmt.Printf("\nURL info: %#v\n", URLInfo)
}

func _TestModifyURL(t *testing.T, URLID, formID string) {
	modifiedURLInfo, err := client.ModifyURL(URLID, formID)
	assert.Nil(t, err, "no error should occur")
	assert.NotNil(t, modifiedURLInfo, "returned response object pointer should not be nil")

	fmt.Printf("\nModified URL info: %#v\n", modifiedURLInfo)

	_TestDeleteURL(t, modifiedURLInfo.ID)
}

func _TestDeleteURL(t *testing.T, URLID string) {
	err := client.DeleteURL(URLID)
	assert.Nil(t, err, "no error should occur")
}
