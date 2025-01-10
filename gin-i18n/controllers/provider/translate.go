package provider

import "fmt"

type TransProvider interface {
	Translate(map[string]string) (map[string]string, error)
}

func GetTransProvider(vendor string) (provider TransProvider, err error) {
	switch vendor {
	case "tencent":
		provider = &TencentClient{
			Region: "ap-shanghai",
		}
	default:
		err = fmt.Errorf("unknown translate provider: %v", vendor)
	}

	return
}
