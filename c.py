import requests

# 目标URL
url = "http://localhost:8080/api/matches"

# 设置Cookie
cookies = {
    'token': "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3MTc5MzIwNTMsImlhdCI6MTcxNzg0NTY1M30.5JtezJMqdPIiccF8ralxdGjQU-g_MolWjI8ZrTi1ZaM"
}

# 请求体数据
payload = {
  "role": "tutor"
}

# 发送POST请求
response = requests.post(url, json=payload, cookies=cookies)
# response = requests.get(url + '/1', cookies=cookies)

# 打印响应状态码
print("Status Code:", response.status_code)

# 打印响应内容
print("Response Body:", response.text)