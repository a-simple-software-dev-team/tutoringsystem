import requests

# 目标URL
url = "http://localhost:8080/api/tutors"

# 设置Cookie
cookies = {
    'token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTc5MjcxNTEsInVzZXJfaWQiOjF9.DxDx1vFdVoPAypHiP75th0yPmRinWLOP-SMnBmzEkl4'
}

# 请求体数据
payload = {
  "subjects_needed": ["Math", "Science"],
  "grades": ["A", "B"],
  "available_time": [
    {
      "day": "Monday",
      "time_range": "09:00 - 11:00"
    },
    {
      "day": "Wednesday",
      "time_range": "14:00 - 16:00"
    }
  ]
}

# 发送POST请求
response = requests.post(url, json=payload, cookies=cookies)

# 打印响应状态码
print("Status Code:", response.status_code)

# 打印响应内容
print("Response Body:", response.text)