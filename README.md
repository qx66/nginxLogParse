# nginxLogParse

nginxLogParse 是一个 nginx 日志分析工具，用于日常 nginx 日志文件临时分析需求

个人认为 json 格式日志相对于行日志更方便日志数据解析，所以将 nginx 日志格式由行日志格式更换为 json 格式，方便后续解析需求。

在使用 json 格式带来的问题就是，不利于 shell 脚本截取解析

```
log_format main escape=json     '{ "timestamp": $msec, "request_id": "$request_id", "hostname": "$hostname",'
                '"http_x_forwarded_for": "$http_x_forwarded_for", "remote_addr": "$remote_addr",'
                '"remote_port": $remote_port, "request_method": "$request_method", "http_host": "$http_host",'
                '"request_uri": "$request_uri", "request_body": "$request_body", "body_bytes_sent": $body_bytes_sent,'
                '"status": "$status", "request_time": $request_time, "upstream_addr": "$upstream_addr",'
                '"upstream_response_time": "$upstream_response_time", "upstream_connect_time": "$upstream_connect_time",'
                '"upstream_cache_status": "$upstream_cache_status", "upstream_status": "$upstream_status",'
                '"ssl_session_id": "$ssl_session_id", "ssl_cipher": "$ssl_cipher", "ssl_session_reused": "$ssl_session_reused",'
                '"http_user_agent": "$http_user_agent", "http_referer": "$http_referer" }';
```
 
# feature

常用统计需求:

    事后统计:
        某段时间通常是指 某天/某小时/某月
        
        1. 某段时间独立 IP 访问量
        2. 某段时间总请求次数
        3. 某段时间 HTTP CODE 分布
        4. 某段时间每秒的平均相应时长，以及响应时长 >= x秒的请求占比
        5. 某段时间请求量接口排名 (使用 "?" 切割 $request_uri)
        
    实时:
        1. 过滤 request_time >= x 秒的请求
        2. 
        
        
实时统计需求:

    固定输出:
        Time: 2022-05-23 11:15:04, 
        RequestCount: 12, 
        TotalBodyByteSize: 2(KB), 
        AvgRequestTime: 0.004167, 
        AvgResponseTime: 0.000000, 
        RemoteAddrCount: 4.
    
    --printRemoteAddCount=true
    
        remote_addr: x.x.x.x,	count: n.

    -tail=true
        
        从最后一个 byte 开始 tail
        
    --printHttpCodeCount=true
        200: n, 301: y, 502: z

    --printUpstreamDistribute=true
        x.x.x.x:yy => n
        x.x.x.z:yy => p
    
    
事后统计需求:

    支持多文件合并统计






