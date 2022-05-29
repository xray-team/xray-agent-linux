# Nginx
## Description
The Nginx collector picks the basic status information about nginx server from ngx_http_stub_status_module.  
The following status information metrics are available:

| Name     | Type    | Data source                 | Description                                                                               |
|----------|---------|-----------------------------|-------------------------------------------------------------------------------------------|
| Active   | gauge   | ngx_http_stub_status_module | The current number of active client connections including Waiting connections.            |
| Accepts  | counter | ngx_http_stub_status_module | The total number of accepted client connections.                                          |
| Handled  | counter | ngx_http_stub_status_module | The total number of handled connections.                                                  |
| Requests | counter | ngx_http_stub_status_module | The total number of client requests.                                                      |
| Reading  | gauge   | ngx_http_stub_status_module | The current number of connections where nginx is reading the request header.              |
| Writing  | gauge   | ngx_http_stub_status_module | The current number of connections where nginx is writing the response back to the client. |
| Waiting  | gauge   | ngx_http_stub_status_module | The current number of idle client connections waiting for a request.                      |

## Limitations
Before using the Nginx collector, you must first enable ngx_http_stub_status_module.  
More details in the nginx documentation: [https://nginx.org/en/docs/http/ngx_http_stub_status_module.html](https://nginx.org/en/docs/http/ngx_http_stub_status_module.html)

## Configuration
```json
{
  "collectors": {
    "Nginx": {
      "enabled": true,
      "endpoint": "http://127.0.0.1/basic_status",
      "timeout": 5
    }
  }
}
```
* **"enabled"**
    * **true** - enable collector
    * **false** - disable collector
* **"endpoint"** - ngx_http_stub_status_module url (location)
* **"timeout"** - timeout in seconds