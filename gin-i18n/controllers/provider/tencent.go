package provider

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

type TencentClient struct {
	SecretId  string
	SecretKey string
	Region    string
}

func (s *TencentClient) Translate(originalTexts map[string]string) (targetTexts map[string]string, err error) {

	// todo chunk
	targetTexts = make(map[string]string)
	// Set your Tencent Cloud credentials (replace with your actual credentials)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tmt.tencentcloudapi.com"

	client, err := tmt.NewClient(credential, s.Region, cpf)
	if err != nil {
		panic(fmt.Sprintf("Failed to create client: %v", err))
	}

	// 创建批量翻译请求
	request := tmt.NewTextTranslateBatchRequest()

	// 设置源语言和目标语言
	// 源语言，例如：zh（中文），en（英文）
	request.Source = common.StringPtr("zh")
	// 目标语言，例如：en（英文），ja（日语）
	request.Target = common.StringPtr("en")
	// 项目 ID，可选，默认是 0
	request.ProjectId = common.Int64Ptr(0)

	sourceTextList := []string{}
	for _, text := range originalTexts {
		sourceTextList = append(sourceTextList, text)
	}

	// 设置要翻译的文本数组
	request.SourceTextList = common.StringPtrs(sourceTextList)

	// 调用批量翻译 API
	response, err := client.TextTranslateBatch(request)
	if err != nil {
		panic(fmt.Sprintf("Failed to translate: %v", err))
	}

	responseMap := make(map[string]string)
	// 输出翻译结果
	fmt.Println("Translation Results:")
	for idx, translatedText := range response.Response.TargetTextList {
		responseMap[sourceTextList[idx]] = *translatedText
	}
	for key, text := range originalTexts {
		if _, ok := responseMap[text]; ok {
			targetTexts[key] = responseMap[text]
		}
	}
	return
}
