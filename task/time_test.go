package task

import (
	"github.com/tidwall/gjson"
	"testing"
)

const json2 = `{
  "widget": {
    "debug": "on",
    "window": {
      "title": "Sample Konfabulator Widget",
      "name": "main_window",
      "width": 500,
      "height": 500
    },
    "image": { 
      "src": "Images/Sun.png",
      "hOffset": 250,
      "vOffset": 250,
      "alignment": "center"
    },
    "text": {
      "data": "Click Here",
      "size": 36,
      "style": "bold",
      "vOffset": 100,
      "alignment": "center",
      "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
    }
	"programmers": [
		{
		  "firstName": "Janet", 
		  "lastName": "McLaughlin", 
		}, {
		  "firstName": "Elliotte", 
		  "lastName": "Hunter", 
		}, {
		  "firstName": "Jason", 
		  "lastName": "Harold", 
		}
	]
  }
}`

func TestSend(t *testing.T) {
	//value := gjson.GetMany(json2, "widget.window.title", "widget.programmers.#.firstName")

	name := gjson.Get(json2, `widget.programmers`)

	name.ForEach(func(key, value gjson.Result) bool {
		println(value.String())
		return true // keep iterating
	})
}
