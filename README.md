# nginxLogParse

nginxLogParse 是一个 nginx 日志分析工具，涵盖大部分解析 nginx 脚本需求

由于日志传输和分析需求，nginx 日志格式由行日志格式更换为 json 格式，方便反序列化解析。

在使用 json 格式带来的问题就是，shell 脚本切割的不便利性，

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
                
                
         