package parser

import (
	"crawler/engine"
	"crawler/model"
	"regexp"
	"strconv"
)

//个人信息 按照顺序
//是否离异、年龄、星座、身高、体重、工作地、月收入、职业、学历
//民族、籍贯、体型、是否抽烟、是否喝酒、住址、是否购车、有无孩子、是否想要孩子、何时结婚
const profileRe = `<div class="m-btn purple"[^>]*>([^<]+)</div>`

const genderNanRe = `<a href="http://www.zhenai.com/zhenghun/[a-z]+/nan">`
const genderNvRe = `<a href="http://www.zhenai.com/zhenghun/[a-z]+/nv">`

var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

//<div class="m-btn purple" data-v-bff6f798>未婚</div>

//择偶信息

func ProfileParser(contents []byte, url string, name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name

	re := regexp.MustCompile(genderNanRe)
	matches := re.FindAllSubmatch(contents, -1)
	if len(matches) > 0 {
		profile.Gender = "男"
	}

	re = regexp.MustCompile(genderNvRe)
	matches = re.FindAllSubmatch(contents, -1)
	if len(matches) > 0 {
		profile.Gender = "女"
	}

	re = regexp.MustCompile(profileRe)
	matches = re.FindAllSubmatch(contents, -1)

	if len(matches) > 1 {
		profile.Age, _ = strconv.Atoi(string(matches[1][1]))
	}
	if len(matches) > 3 {
		profile.Height, _ = strconv.Atoi(string(matches[3][1]))
	}
	if len(matches) > 4 {
		profile.Weight, _ = strconv.Atoi(string(matches[4][1]))
	}
	if len(matches) > 6 {
		profile.Income = string(matches[6][1])
	}
	if len(matches) > 0 {
		profile.Marriage = string(matches[0][1])
	}
	if len(matches) > 8 {
		profile.Education = string(matches[8][1])
	}
	if len(matches) > 7 {
		profile.Occupation = string(matches[7][1]) //工作
	}
	if len(matches) > 10 {
		profile.HuKou = string(matches[10][1])
	}
	if len(matches) > 2 {
		profile.Constellation = string(matches[2][1])
	}
	if len(matches) > 14 {
		profile.House = string(matches[14][1])
	}
	if len(matches) > 15 {
		profile.Car = string(matches[15][1])
	}

	result := engine.ParserResult{
		Items: []engine.Item{{
			Url:     url,
			Type:    "zhenai",
			Id:      extractString([]byte(url), idUrlRe),
			Payload: profile,
		}},
	}
	return result
}

func extractString(
	contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
