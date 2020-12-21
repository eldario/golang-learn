$ cat loop.sh
#!/bin/bash

count=10

while [ 1 = 1 ]; do
  echo $count

  curl -X POST -H "Content-Type: application/json" \
    -d '{"number": 22, 	"text": "Scarlett OHara was not beautiful, but men did not realize this when caught by her charm as the Tarleton twins were. Her eyes were green, and her skin was that soft white skin which Southern women valued so highly, and covered so carefully from the hot Georgia sun with hats and gloves." }' \
    127.0.0.1:4001/text
  curl -X POST -H "Content-Type: application/json" \
    -d '{"number": 42, 	"text": "Scarlett OHara was not beautiful, but men did not realize this when caught by her charm as the Tarleton twins were. Her eyes were green, and her skin was that soft white skin which Southern women valued so highly, and covered so carefully from the hot Georgia sun with hats and gloves." }' \
    127.0.0.1:4001/text
  curl -X POST -H "Content-Type: application/json" \
    -d '{"number": 26, 	"text": "Scarlett OHara was not beautiful, but men did not realize this when caught by her charm as the Tarleton twins were. Her eyes were green, and her skin was that soft white skin which Southern women valued so highly, and covered so carefully from the hot Georgia sun with hats and gloves." }' \
    127.0.0.1:4001/text
  curl -X POST -H "Content-Type: application/json" \
    -d '{"number": 32, 	"text": "Scarlett OHara was not beautiful, but men did not realize this when caught by her charm as the Tarleton twins were. Her eyes were green, and her skin was that soft white skin which Southern women valued so highly, and covered so carefully from the hot Georgia sun with hats and gloves." }' \
    127.0.0.1:4001/text
  curl 127.0.0.1:4001/stat/2
  curl 127.0.0.1:4001/stat/3
  curl 127.0.0.1:4001/stat/4
  curl 127.0.0.1:4001/stat/5
  curl 127.0.0.1:4001/stat/2
  curl 127.0.0.1:4001/stat/3
  curl 127.0.0.1:4001/stat/4
  curl 127.0.0.1:4001/stat/5
  curl -X POST -H "Content-Type: application/json" \
    -d '{"number": 22, 	"text": "Scarlett OHara was not beautiful, but men did not realize this when caught by her charm as the Tarleton twins were. Her eyes were green, and her skin was that soft white skin which Southern women valued so highly, and covered so carefully from the hot Georgia sun with hats and gloves." }' \
    127.0.0.1:4001/text
  curl -X POST -H "Content-Type: application/json" \
    -d '{"number": 42, 	"text": "Scarlett OHara was not beautiful, but men did not realize this when caught by her charm as the Tarleton twins were. Her eyes were green, and her skin was that soft white skin which Southern women valued so highly, and covered so carefully from the hot Georgia sun with hats and gloves." }' \
    127.0.0.1:4001/text
  curl -X POST -H "Content-Type: application/json" \
    -d '{"number": 26, 	"text": "Scarlett OHara was not beautiful, but men did not realize this when caught by her charm as the Tarleton twins were. Her eyes were green, and her skin was that soft white skin which Southern women valued so highly, and covered so carefully from the hot Georgia sun with hats and gloves." }' \
    127.0.0.1:4001/text
  curl -X POST -H "Content-Type: application/json" \
    -d '{"number": 32, 	"text": "Scarlett OHara was not beautiful, but men did not realize this when caught by her charm as the Tarleton twins were. Her eyes were green, and her skin was that soft white skin which Southern women valued so highly, and covered so carefully from the hot Georgia sun with hats and gloves." }' \
    127.0.0.1:4001/text
  curl 127.0.0.1:4001/stat/2
  curl 127.0.0.1:4001/stat/3
  curl 127.0.0.1:4001/stat/4
  curl 127.0.0.1:4001/stat/5
  curl 127.0.0.1:4001/stat/2
  curl 127.0.0.1:4001/stat/3
  curl 127.0.0.1:4001/stat/4
  curl 127.0.0.1:4001/stat/5

done
