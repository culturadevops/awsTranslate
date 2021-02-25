package awsTranslate

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

const (
	Frances = "fr"
	Espanol = "es"
	Ingles  = "en"
)

type Awstranslate struct {
	translate          *translate.Translate
	SourceLanguageCode string
	TargetLanguageCode string
	Region             string
}

func (t *Awstranslate) Init() {
	if t.Region == "" {
		t.Region = "us-east-1"
	}
	t.translate = translate.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(t.Region),
	})))
}

func (t *Awstranslate) SetRegion(region string) {
	t.Region = region
}
func (t *Awstranslate) SetLanguageCode(source string, target string) {
	t.SourceLanguageCode = source
	t.TargetLanguageCode = target
}
func (t *Awstranslate) TranslateByte(text []byte) (string, error) {
	response, err := t.translate.Text(&translate.TextInput{
		SourceLanguageCode: aws.String(t.SourceLanguageCode),
		TargetLanguageCode: aws.String(t.TargetLanguageCode),
		Text:               aws.String(string(text)),
	})
	if err != nil {
		print(err)
		return "", err
	}
	return *response.TranslatedText, nil
}
