1. 配置config.json文件

2. windows在cmd下执行: `jcqsign.exe config.json`
   linux下执行: `./jcqsign config.json`
   mac下执行: `./jcqsign config.json`
   
3. config.json

   ```json
   {
       "accessKey": "",
       "secretKey": "",
       "endpoint": "jcq-shared-004.cn-north-1.jdcloud.com",
       "dateTime": "",
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

   