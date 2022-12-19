import requests
import sys
import config

args = sys.argv

print(args[1])


# ---------------- OPENAI -----------------------

url = "https://api.openai.com/v1/images/generations"


payload = {
  "prompt": args[1],
  "n": 1,
  "size": '512x512'
}
# size must be one ogf ['256x256', '512x512', '1024x1024']

headers = {
    'Content-Type': 'application/json',
    'Authorization': 'Bearer {}'.format(config.OPENAI_API_KEY)
}

response = requests.post(url, headers=headers, json=payload)

print(response.text)

# image_url = response['data'][0]['url']
# print(image_url)

# {
#   "prompt": "A cute baby sea otter",
#   "n": 2,
#   "size": "1024x1024"
# }