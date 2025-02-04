### CURL Commands

- Create Collection

        curl -X POST "http://localhost:8000/api/v1/collection/create" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"name":"testcollection"}'


- Load Data file into collection
  
      curl -X POST -F "file=@test.txt" -F "CollectionName=testcollection" "http://localhost:8000/api/v1/collection/insert"

- Query the collection

        curl -X POST "http://localhost:8000/api/v1/collection/query" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"Query":"When the disk connection spike up to 100% ?" , "CollectionName":"testcollection"}'

