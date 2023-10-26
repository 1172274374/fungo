package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readJson(filePath string) {
	// Let's first read the `config.json` fileos
	content, err := ioutil.ReadFile(filePath)
	checkErr(err)

	// Now let's unmarshall the data into `payload`
	payload := make(map[string]interface{})
	err = json.Unmarshal(content, &payload)
	checkErr(err)

	// Let's print the unmarshalled data!
	fmt.Println(payload["article"].(map[string]interface{})["name"])
	fmt.Println(payload["hobby"].([]interface{})[0])

	ran := append(payload["hobby"].([]interface{}), "打飞机")
	fmt.Println(ran)

}

func writeJson(filePath string) {
	type UserInfo struct {
		Name  string   `json:"name"`
		Age   int      `json:"age"`
		Hobby []string `json:"hobby"`
	}
	user := &UserInfo{
		Name:  "GoGin",
		Age:   25,
		Hobby: []string{"看书", "学习新技术", "编程"},
	}
	data, err := json.MarshalIndent(user, "", "    ")
	checkErr(err)
	err = ioutil.WriteFile(filePath, data, 0755)
	checkErr(err)
}

func updateJson(filePath string) {
	// 读取json数据
	data := make(map[string]interface{})
	file, err := ioutil.ReadFile(filePath)
	checkErr(err)
	err = json.Unmarshal(file, &data)
	checkErr(err)

	// 读取爱好列表, 新增新的爱好
	hobbyList := make([]string, 0)
	for _, hobby := range data["hobby"].([]interface{}) {
		hobbyList = append(hobbyList, hobby.(string))
	}

	hobbyList = append(hobbyList, "玩游戏")
	data["hobby"] = hobbyList

	// 写入 data.json
	dataJson, err := json.MarshalIndent(data, "", "    ")
	checkErr(err)
	err = ioutil.WriteFile(filePath, dataJson, 0755)
	checkErr(err)
}

func removeJson(filePath string) {
	// 读取json数据
	data := make(map[string]interface{})
	file, err := ioutil.ReadFile(filePath)
	checkErr(err)
	err = json.Unmarshal(file, &data)
	checkErr(err)

	// 删除 key
	delete(data, "age")

	// 写入 data.json
	dataJson, err := json.MarshalIndent(data, "", "    ")
	checkErr(err)
	err = ioutil.WriteFile(filePath, dataJson, 0755)
	checkErr(err)
}
