import requests

# 目标URL
url = "http://localhost:8080/api/students/1"

# 设置Cookie
cookies = {
    'token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTc5MjcxNTEsInVzZXJfaWQiOjF9.DxDx1vFdVoPAypHiP75th0yPmRinWLOP-SMnBmzEkl4'
}

# 发送GET请求
response = requests.get(url, cookies=cookies)

# 打印响应状态码
print("Status Code:", response.status_code)

# 打印响应内容
print("Response Body:", response.text)