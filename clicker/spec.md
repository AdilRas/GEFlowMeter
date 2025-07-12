# Required Functions

## Step 1: Login

### Request
```curl
curl --request POST \
  --url https://clicker.iiindia.org/api/auth/login \
  --header 'Accept: */*' \
  --header 'Accept-Encoding: gzip, deflate, br, zstd' \
  --header 'Accept-Language: en-US,en;q=0.5' \
  --header 'Connection: keep-alive' \
  --header 'Content-Type: application/json' \
  --header 'Priority: u=0' \
  --header 'Sec-Fetch-Dest: empty' \
  --data '{"username":"cc01","password":"cc01"}'
```
- The username and password should be retrieved from environment variables

### Response
```json
{
	"success": true,
	"token": "EXAMPLE.TOKEN",
	"user": {
		"id": 4,
		"username": "cc01",
		"role": "admin",
		"created_at": "2025-07-10T18:38:20.933022Z",
		"updated_at": "2025-07-10T18:38:20.933022Z"
	}
}
```
- Note that the token needs to be saved so we can use it for requests later.

## Step 2: Get Areas

### Request
```curl
curl --request GET \
  --url https://clicker.iiindia.org/api/proxy/areas \
  --header 'Accept: */*' \
  --header 'Connection: keep-alive' \
  --header 'authorization: Bearer EXAMPLE.TOKEN'
```
- Using the token we saved prior.

### Response
```json
{"areas":[{"id":1,"name":"IAG","capacity":1000,"current_count":27,"status":"ACTIVE","isEnabled":true,"created_at":"2025-07-10T17:27:57.358802Z","updated_at":"2025-07-12T16:10:07.716429Z"},{"id":3,"name":"IFF","capacity":500,"current_count":30,"status":"ACTIVE","isEnabled":true,"created_at":"2025-07-10T18:48:25.370717Z","updated_at":"2025-07-12T16:10:07.716429Z"}],"success":true}
```