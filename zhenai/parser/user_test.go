package parser

import (
	"crawler/engine"
	"crawler/model"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestProfileParser(t *testing.T) {
	contents, err := ioutil.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	type args struct {
		contents []byte
		url      string
		name     string
	}
	tests := []struct {
		name string
		args args
		want engine.ParserResult
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				contents: contents,
				url:      "http://album.zhenai.com/u/108906739",
				name:     "HiSiri",
			},
			want: engine.ParserResult{
				Requests: nil,
				Items: []engine.Item{
					{
						Url:  "http://album.zhenai.com/u/108906739",
						Type: "zhenai",
						Id:   "108906739",
						Payload: model.Profile{
							Name:          "HiSiri",
							Gender:        "女",
							Age:           28,
							Height:        163,
							Weight:        100,
							Income:        "3001-5000元",
							Marriage:      "未婚",
							HuKou:         "内蒙古赤峰",
							Constellation: "金牛座",
							House:         "自住",
							Car:           "未购车",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProfileParser(tt.args.contents, tt.args.url, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProfileParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
