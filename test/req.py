import requests

userid = ""
r = requests.get(f"http://localhost:8080/api/v1/activities/{userid}")
print(r.text)