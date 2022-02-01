package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/xwzy/godan/time_meter"
)

func main() {
	Test01()
	Test02()
	Test03()
	Test04()
}

func Test01() {
	expression, err := govaluate.NewEvaluableExpression("10 > 0")
	result, err := expression.Evaluate(nil)
	fmt.Println(result)
	if err != nil {
		panic(err)
	}
}

func Test02() {
	expression, err := govaluate.NewEvaluableExpression("foo > 0")

	parameters := make(map[string]interface{}, 8)
	parameters["foo"] = -1

	result, err := expression.Evaluate(parameters)
	fmt.Println(result)
	if err != nil {
		panic(err)
	}
}

func Test03() {
	expression, err := govaluate.NewEvaluableExpression("(requests_made * requests_succeeded / 100) >= 90")

	parameters := make(map[string]interface{}, 8)
	parameters["requests_made"] = 100
	parameters["requests_succeeded"] = 80

	result, err := expression.Evaluate(parameters)

	fmt.Println(result)
	if err != nil {
		panic(err)
	}
}

func Test04() {
	expression, err := govaluate.NewEvaluableExpression("http_response_body == 'service is ok'")

	parameters := make(map[string]interface{}, 8)
	parameters["http_response_body"] = "service is ok"

	result, err := expression.Evaluate(parameters)
	fmt.Println(result)
	if err != nil {
		panic(err)
	}
}

func Test05() {

	
}