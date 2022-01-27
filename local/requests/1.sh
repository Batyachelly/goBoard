# curl -XPOST -H "Content-Type:application/json" -d \
# '{"boardId": 1, "title": "some title2", "text":"Hello"  }' \
#  localhost:8080/thread

curl -XPOST -H "Content-Type:application/json" -d \
'{"threadID": 1, "title": "some title2", "text":"Hello"  }' \
 localhost:8080/comment
