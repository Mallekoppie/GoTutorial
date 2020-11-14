curl --request POST \
  --url http://localhost:8888/point \
  --header 'content-type: application/json' \
  --cookie 'REVEL_FLASH=; THE_HIVE_SESSION=eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7InVzZXJuYW1lIjoiYXV0b21hdGlvbiIsImV4cGlyZSI6IjE1ODc5OTQ3MDQ5MzYiLCJhdXRoTWV0aG9kIjoibG9jYWwifSwibmJmIjoxNTg3OTkxMjg2LCJpYXQiOjE1ODc5OTEyODZ9.hd9hpLTCCggDQfR_pMaBP4Lo0CkhEaSUHpqPRIdGwQY' \
  --data '{
	"car":"A002",
	"scantime" :"2018-09-22T16:35:08+07:00",
	"user":"drikkie",
	"points":123,
	"checkpoint":"api checkpoint",
	"method":"start"
}'

curl --request POST \
  --url http://localhost:8888/scan \
  --header 'content-type: application/json' \
  --cookie 'REVEL_FLASH=; THE_HIVE_SESSION=eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7InVzZXJuYW1lIjoiYXV0b21hdGlvbiIsImV4cGlyZSI6IjE1ODc5OTQ3MDQ5MzYiLCJhdXRoTWV0aG9kIjoibG9jYWwifSwibmJmIjoxNTg3OTkxMjg2LCJpYXQiOjE1ODc5OTEyODZ9.hd9hpLTCCggDQfR_pMaBP4Lo0CkhEaSUHpqPRIdGwQY' \
  --data '{
	"car":"A002",
	"scantime" :"2018-09-22T16:35:08+07:00",
	"user":"drikkie",
	"points":123,
	"checkpoint":"api checkpoint",
	"method":"start"
}'

curl --request POST \
  --url http://localhost:8888/stopwatch \
  --header 'content-type: application/json' \
  --cookie 'REVEL_FLASH=; THE_HIVE_SESSION=eyJhbGciOiJIUzI1NiJ9.eyJkYXRhIjp7InVzZXJuYW1lIjoiYXV0b21hdGlvbiIsImV4cGlyZSI6IjE1ODc5OTQ3MDQ5MzYiLCJhdXRoTWV0aG9kIjoibG9jYWwifSwibmJmIjoxNTg3OTkxMjg2LCJpYXQiOjE1ODc5OTEyODZ9.hd9hpLTCCggDQfR_pMaBP4Lo0CkhEaSUHpqPRIdGwQY' \
  --data '{
	"car":"A002",
	"scantime" :"2018-09-22T16:35:08+07:00",
	"user":"drikkie",
	"method":"start",
	"lap":"start"
}'