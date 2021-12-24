

## JCQ文档参考

1. 签名算法文档：https://docs.jdcloud.com/cn/message-queue/signature-algorithm
2. 消费消息接口文档：https://docs.jdcloud.com/cn/message-queue/consume-message

## 使用步骤

1. 配置config.json文件

2. windows在cmd下执行: `jcqsign.exe config.json`
   linux下执行: `./jcqsign config.json`
   mac下执行: `./jcqsign config.json`

3. config.json

   `dateTime` 参数如果为空字符串，则默认会传系统UTC时间，时间格式: `2021-12-24T03:28:47Z`

   ```json
   {
       "accessKey": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
       "secretKey": "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY",
       "endpoint": "jcq-shared-004.cn-north-1.jdcloud.com",
       "dateTime": "2021-12-24T03:28:47Z",
       "scheme": "https",
       "consumerParams": {
           "topic": "xinyulu3344",
           "consumerGroupId": "consumer1",
           "size": 32,
           "consumerId": "httpProxyId",
           "consumeFromWhere": "HEAD",
           "filterExpressionType":"",
           "filterExpression": "",
           "ack": "true"
       }
   }
   ```


4. 返回结果

   ```
   2021/12/24 11:36:20 config.json:  {
       "accessKey": "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",    
       "secretKey": "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY",    
       "endpoint": "jcq-shared-004.cn-north-1.jdcloud.com",
       "dateTime": "2021-12-24T03:28:47Z",
       "scheme": "https",
       "consumerParams": {
           "topic": "xinyulu3344",
           "consumerGroupId": "consumer1",
           "size": 32,
           "consumerId": "httpProxyId",
           "consumeFromWhere": "HEAD",
           "filterExpressionType":"",
           "filterExpression": "",
           "ack": "true"
       }
   }
   2021/12/24 11:36:20 排序后keys:  [accessKey ack consumeFromWhere consumerGroupId consumerId dateTime size topic]
   2021/12/24 11:36:20 signSource:  accessKey=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX&ack=true&consumeFromWhere=HEAD&consumerGroupId=consumer1&consumerId=httpProxyId&dateTime=2021-12-24T03:28:47Z&size=32&topic=xinyulu3344
   2021/12/24 11:36:20 签名:  zk51bnoU4nitiM0Cq6BawHWAkDI=
   ```

   > 1. 如果消费接口的可选参数不传，则不要将该参数纳入signSource中进行签名计算