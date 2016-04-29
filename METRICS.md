# snap collector plugin - haproxy

## Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type | Description
----------|-----------|-----------------------
/intel/haproxy/info/CompressBpsIn | int64 | bytes per second, before http compression
/intel/haproxy/info/CompressBpsOut | int64 |  bytes per second, after http compression
/intel/haproxy/info/CompressBpsRateLim | int64 | the HTTP compression rate limit
/intel/haproxy/info/ConnRate | int64 | the per-process number of connections per second
/intel/haproxy/info/ConnRateLimit | int64 | the connection rate limit
/intel/haproxy/info/CumConns | int64 | the number of terminated sessions
/intel/haproxy/info/CumReq | int64 | the request counter (HTTP or TCP session) for logs and unique_id
/intel/haproxy/info/CumSslConns | int64 |  the request counter for SSL connections
/intel/haproxy/info/CurrConns | int64 | the number of active sessions
/intel/haproxy/info/CurrSslConns | int64 |  the number of active SSL sessions
/intel/haproxy/info/Hard_maxconn | int64 | the hard max per-process number of concurrent connections
/intel/haproxy/info/Idle_pct | int64 | the idle time in percentage
/intel/haproxy/info/MaxConnRate | int64 | the max per-process number of connections per second
/intel/haproxy/info/MaxSessRate | int64 | the max per-process number of sessions per second
/intel/haproxy/info/MaxSslConns | int64 | the max per-process number of concurrent SSL connections
/intel/haproxy/info/MaxSslRate | int64 | the max per-process number of SSL sessions per second
/intel/haproxy/info/MaxZlibMemUsage | int64 | the max amount of RAM in megabytes per process usable by the zlib
/intel/haproxy/info/Maxconn |  int64| the max per-process number of concurrent connections
/intel/haproxy/info/Maxpipes | int64 | the max per-process number of pipes
/intel/haproxy/info/Maxsock | int64 | the max number of sockets
/intel/haproxy/info/Memmax_MB | int64 | the all-process memory limit in MB
/intel/haproxy/info/Name | string | the product name
/intel/haproxy/info/Nbproc | int64 | the number of processes which is created when going daemon
/intel/haproxy/info/Pid | int64 | the current process id
/intel/haproxy/info/PipesFree | int64 |  the number of pipes in use
/intel/haproxy/info/PipesUsed | int64 | the number of pipes unused
/intel/haproxy/info/Process_num | int64 | the process number
/intel/haproxy/info/Release_date | string | the release date
/intel/haproxy/info/Run_queue | int64 | the run queue size
/intel/haproxy/info/SessRate | int64 | the per-process number of sessions per second
/intel/haproxy/info/SessRateLimit | int64 |  the session rate limit
/intel/haproxy/info/SslBackendKeyRate | int64 |
/intel/haproxy/info/SslBackendMaxKeyRate | int64 |
/intel/haproxy/info/SslCacheLookups | int64 |
/intel/haproxy/info/SslCacheMisses | int64 |
/intel/haproxy/info/SslFrontendKeyRate | int64 |
/intel/haproxy/info/SslFrontendMaxKeyRate | int64 |
/intel/haproxy/info/SslFrontendSessionReuse_pct | int64 |
/intel/haproxy/info/SslRate | int64 | the number of SSL sessions per second
/intel/haproxy/info/SslRateLimit | int64 |  the SSL rate limit
/intel/haproxy/info/Tasks | int64 | tasks count
/intel/haproxy/info/Ulimit-n | int64 | the max number of per-process file-descriptors
/intel/haproxy/info/Uptime | string | the up time
/intel/haproxy/info/Uptime_sec | int64 | the up time in second
/intel/haproxy/info/Version | string | the version of HAProxy
/intel/haproxy/info/ZlibMemUsage | int64 | the memory used by zlib
/intel/haproxy/info/description | string | the text that describes the instance
/intel/haproxy/info/node | string | the node name
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/qcur | int64 | the number of current queued requests
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/qmax | int64 | the max value of queued requests
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/scur | int64 | the number of current sessions
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/smax | int64 | the max number of sessions
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/slim |  int64 | the configured session limit
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/stot | int64 | the cumulative number of connections
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/bin | int64| bytes in
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/bout | int64 | bytes out
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/dreq | int64 | the number of requests denied because of security concerns.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/dresp | int64 | the number of responses denied because of security concerns.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/ereq | int64 | the number of request errors
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/econ | int64 | the number of requests that encountered an error trying to  connect to a backend server
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/eresp | int64| the number of response errors
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/wretr | int64 | the number of times a connection to a server was retried.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/wredis | int64 | the number of times a request was redispatched to another  server
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/status | string | the status (UP/DOWN/NOLB/MAINT/MAINT(via)...)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/weight | int64 | the total weight (backend), server weight (server)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/act | int64 | the number of active servers (backend), server is active (server)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/bck | int64 | the number of backup servers (backend), server is backup (server)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/chkfail | int64 | the number of failed checks
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/chkdown | int64 | the number of UP->DOWN transitions
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/lastchg | int64 | the number of seconds since the last UP<->DOWN transition
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/downtime |  int64 | the total downtime (in seconds)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/qlimit | int64 | the configured maxqueue for the server, or nothing in the  value is 0 (default, meaning no limit)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/pid | int64 | the process id (0 for first instance, 1 for second, ...)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/iid | in64 | the unique proxy id
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/sid | int64 | the server id (unique inside a proxy)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/throttle | int64 | the current throttle percentage for the server, when slowstart is active, or no value if not in slowstart.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/lbtot | int64 | the total number of times a server was selected, either for new sessions, or when re-dispatching
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/tracked | int64 | the id of proxy/server if tracking is enabled.
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/type | int64 | (0=frontend, 1=backend, 2=server, 3=socket/listener)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/rate | int64 | the number of sessions per second over last elapsed second
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/rate_lim | int64 | the configured limit on new sessions per second
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/rate_max | int64 | the max number of new sessions per second
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/check_status | string | the status of last health check
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/check_code | int64 | the layer5-7 code, if available
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/check_duration | int64 | the  time in ms took to finish last health check
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_1xx | int64 | the number of http response with 1xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_2xx | int64 | the number of http response with 2xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_3xx | int64 | the number of http response with 3xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_4xx | int64 | the number of http response with 4xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_5xx | int64 | the number of http response with 5xx code
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hrsp_other | int64 | the number of  http responses with other codes (protocol error)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/hanafail | string | failed health checks details
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/req_rate | int64 |  the number of HTTP requests per second over last elapsed second
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/req_rate_max | int64 | the max number of HTTP requests per second observed
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/req_tot | int64 | the total number of HTTP requests received
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/cli_abrt | int64 | tje number of data transfers aborted by the client
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/srv_abrt | int64| the number of data transfers aborted by the server (inc. in eresp)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_in | int64 | the number of HTTP response bytes fed to the compressor
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_out | int64 | the number of HTTP response bytes emitted by the compressor
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_byp | int64 | the number of bytes that bypassed the HTTP compressor (CPU/BW limit)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/comp_rsp | int64 | the number of HTTP responses that were compressed
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/lastsess | int64 | the number of seconds since last session assigned to server/backend
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/last_chk | string | the last health check contents or textual error
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/last_agt | string | the last agent check contents or textual error
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/qtime | int64 | the average queue time in ms over the 1024 last requests
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/ctime | int64 | the average connect time in ms over the 1024 last requests
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/rtime | int64 | the average response time in ms over the 1024 last requests (0 for TCP)
/intel/haproxy/stat/\<service_name\>/\<proxy_name\>/ttime | int64 | the average total session time in ms over the 1024 last requests

**Important Note: Above metric list is a full list. Actual metrics are dependant on HAProxy service configuration.** 