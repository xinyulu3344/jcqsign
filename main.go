package main

import (
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "sort"
    "time"
)

type JCQHttpProcessor struct {
    AccessKey       string `json:"accessKey"`
    SecretKey       string `json:"secretKey"`
    Endpoint        string `json:"endpoint"`
    DateTime        string `json:"dateTime"`
    Scheme          string `json:"scheme"`
    *ConsumerParams `json:"consumerParams"`
}

func NewJCQHttpProcessor(paramsJson []byte) *JCQHttpProcessor {
    jcqHttpProcessor := &JCQHttpProcessor{}
    err := json.Unmarshal(paramsJson, jcqHttpProcessor)
    if err != nil {
        panic(err)
    }
    // 如果配置文件中 dateTime 为空字符串，则默认用系统UTC时间
    if jcqHttpProcessor.DateTime == "" {
        jcqHttpProcessor.DateTime = time.Now().UTC().Format(time.RFC3339)
    }
    return jcqHttpProcessor
}

// 获取签名
func (j *JCQHttpProcessor) GetSignature() string {
    mac := hmac.New(sha1.New, []byte(j.SecretKey))
    mac.Write([]byte(j.GetSignSourceStr()))
    signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))
    log.Println("签名: ", signature)
    return signature
}

func (j *JCQHttpProcessor) getSignSource() map[string]interface{} {
    signSource := make(map[string]interface{})
    signSource["accessKey"] = j.AccessKey
    signSource["dateTime"] = j.DateTime
    signSource["topic"] = j.ConsumerParams.Topic
    signSource["consumerGroupId"] = j.ConsumerParams.ConsumerGroupId
    signSource["size"] = j.ConsumerParams.Size
    signSource["consumerId"] = j.ConsumerParams.ConsumerId
    signSource["consumeFromWhere"] = j.ConsumerParams.ConsumeFromWhere
    signSource["filterExpressionType"] = j.ConsumerParams.FilterExpressionType
    signSource["filterExpression"] = j.ConsumerParams.FilterExpression
    signSource["ack"] = j.ConsumerParams.Ack

    keys := make([]string, 0)
	// 将零值的key放到切片里
    for key := range signSource {
        if valueIsEmpty(signSource[key]) {
            keys = append(keys, key)
        }
    }
	// 遍历keys，将零值得元素从map中删除
    for _, key := range keys {
        delete(signSource, key)
    }

    return signSource
}

// 获取signsource字符串(最终签名所需的data)
func (j *JCQHttpProcessor) GetSignSourceStr() string {
    signSource := j.getSignSource()
    keys := make([]string, 0)
    for key := range signSource {
        keys = append(keys, key)
    }
    sort.Strings(keys)
    log.Println("排序后keys: ", keys)
    var signSourceStr string
    for _, key := range keys {
        signSourceStr = signSourceStr + fmt.Sprintf("%v=%v&", key, signSource[key])
    }
    signSourceStr = signSourceStr[:len([]rune(signSourceStr))-1]
    log.Println("signSource: ", signSourceStr)
    return signSourceStr
}

type ConsumerParams struct {
    Topic                string `json:"topic"`
    ConsumerGroupId      string `json:"consumerGroupId"`
    Size                 int    `json:"size"`
    ConsumerId           string `json:"consumerId"`
    ConsumeFromWhere     string `json:"consumeFromWhere"`
    FilterExpressionType string `json:"filterExpressionType"`
    FilterExpression     string `json:"filterExpression"`
    Ack                  string `json:"ack"`
}


// 判断元素是否为零值
func valueIsEmpty(v interface{}) bool {
    switch value := v.(type) {
    case string:
        if value == "" {
            return true
        } else {
            return false
        }
    case int:
        if value == 0 {
            return true
        } else {
            return false
        }
    }
    return false
}

func main() {
    if len(os.Args) != 2 {
        log.Fatal("参数数量错误, 执行: jcqsign config.json")
    }
    configPath := os.Args[1]
    configBytes, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("config.json: ", string(configBytes))

    jcqHttpProcessor := NewJCQHttpProcessor(configBytes)
    jcqHttpProcessor.GetSignature()
}
