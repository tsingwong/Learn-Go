/*
 * @Description:
 * @Author: Tsingwong
 * @Date: 2021-11-09 11:11:17
 * @LastEditors: Tsingwong
 * @LastEditTime: 2021-11-12 18:00:41
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Question struct {
	Difficulty string `json:"difficulty"`
	Title      string `json:"title"`
	TitleSlug  string `json:"titleSlug"`
}

type TodayRecord struct {
	Date     string   `json:"date"`
	Question Question `json:"question"`
}

type TodayData struct {
	TodayRecord []TodayRecord `json:"todayRecord"`
}
type TodayRes struct {
	Data TodayData `json:"data"`
}

type QuestionData struct {
	Title   string   `json:"translatedTitle"`
	Content string   `json:"translatedContent"`
	Id      string   `json:"questionId"`
	Hints   []string `json:"hints"`
}
type QuestionContent struct {
	Question QuestionData `json:"question"`
}

type QuestionRes struct {
	Data QuestionContent `json:"data"`
}

func main() {
	data, err := getTodayName()
	if err != nil {
		log.Fatal(fmt.Sprintln("发生了一些错误，具体信息是: ", err))
	}
	log.Printf("每日一题为：%s", data.Title)
	log.Printf("难度为：%s", data.Difficulty)
	log.Printf("地址为：%s", "https://leetcode-cn.com/problems/"+data.TitleSlug)
	fmt.Println("=====================")
	content, err := getTodayContent(data.TitleSlug)
	if err != nil {
		log.Fatal(fmt.Sprintln("发生了一些错误，具体信息是: ", err))
	}
	log.Printf("当前题目的 HTML 是：%s", content)

	preReg := regexp.MustCompile(`<pre>[\s\S]*?</pre>`)
	reg := regexp.MustCompile(`<[^>p]*>`)
	content = preReg.ReplaceAllStringFunc(content, func(src string) string {
		return reg.ReplaceAllString(src, ``)
	})
	removePattern := []string{"</?p>", "</?ul>", "</?ol>", "</li>", "</sup>"}
	for _, pattern := range removePattern {
		reg := regexp.MustCompile(pattern)
		content = reg.ReplaceAllString(content, "")
	}
	replacePattern := [][]string{
		[]string{"&nbsp;", " "},
		[]string{"&quot;", `"`},
		[]string{"&lt;", "<"},
		[]string{"&gt;", ">"},
		[]string{"&le;", "≤"},
		[]string{"&ge;", "≥"},
		[]string{"<sup>", "^"},
		[]string{"&#39", "'"},
		[]string{"&times;", "x"},
		[]string{"&ldquo;", "“"},
		[]string{"&rdquo;", "”"},
		[]string{" *<strong> *", " **"},
		[]string{" *</strong> *", "** "},
		[]string{" *<code> *", " `"},
		[]string{" *</code> *", "` "},
		[]string{"<pre>", "```\n"},
		[]string{"</pre>", "\n```\n"},
		[]string{"<em> *</em>", ""},
		[]string{" *<em> *", " *"},
		[]string{" *</em> *", "* "},
		[]string{"</?div.*?>", ""},
		[]string{"	*</?li>", "- "},
	}
	for _, pattern := range replacePattern {
		reg := regexp.MustCompile(pattern[0])
		content = reg.ReplaceAllString(content, pattern[1])
	}
	fmt.Println(content)
}

func getTodayName() (Question, error) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"
	url := "https://leetcode-cn.com/graphql"

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(`{"query":"\n    query questionOfToday {\n  todayRecord {\n    date\n    userStatus\n    question {\n      questionId\n      frontendQuestionId: questionFrontendId\n      difficulty\n      title\n      titleCn: translatedTitle\n      titleSlug\n      paidOnly: isPaidOnly\n      freqBar\n      isFavor\n      acRate\n      status\n      solutionNum\n      hasVideoSolution\n      topicTags {\n        name\n        nameTranslated: translatedName\n        id\n      }\n      extra {\n        topCompanyTags {\n          imgUrl\n          slug\n          numSubscribed\n        }\n      }\n    }\n    lastSubmission {\n      id\n    }\n  }\n}\n    ","variables":{},"operationName":"questionOfToday"}`)))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://leetcode-cn.com/problemset/all/")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return Question{}, err
	}
	defer req.Body.Close()
	// statuscode := res.StatusCode
	body, _ := ioutil.ReadAll(res.Body)
	var data TodayRes
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Question{}, err
	}
	return data.Data.TodayRecord[0].Question, nil
}

func getTodayContent(slug string) (string, error) {
	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"
	url := "https://leetcode-cn.com/graphql"
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(`{"operationName":"questionData","variables":{"titleSlug":"guess-number-higher-or-lower-ii"},"query":"query questionData($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    categoryTitle\n    boundTopicId\n    title\n    titleSlug\n    content\n    translatedTitle\n    translatedContent\n    isPaidOnly\n    difficulty\n    likes\n    dislikes\n    isLiked\n    similarQuestions\n    contributors {\n      username\n      profileUrl\n      avatarUrl\n      __typename\n    }\n    langToValidPlayground\n    topicTags {\n      name\n      slug\n      translatedName\n      __typename\n    }\n    companyTagStats\n    codeSnippets {\n      lang\n      langSlug\n      code\n      __typename\n    }\n    stats\n    hints\n    solution {\n      id\n      canSeeDetail\n      __typename\n    }\n    status\n    sampleTestCase\n    metaData\n    judgerAvailable\n    judgeType\n    mysqlSchemas\n    enableRunCode\n    envInfo\n    book {\n      id\n      bookName\n      pressName\n      source\n      shortDescription\n      fullDescription\n      bookImgUrl\n      pressImgUrl\n      productUrl\n      __typename\n    }\n    isSubscribed\n    isDailyQuestion\n    dailyRecordStatus\n    editorType\n    ugcQuestionId\n    style\n    exampleTestcases\n    __typename\n  }\n}\n"}
	`)))
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Referer", "https://leetcode-cn.com/problemset/all/")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	var data QuestionRes
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}
	return data.Data.Question.Content, nil
}
