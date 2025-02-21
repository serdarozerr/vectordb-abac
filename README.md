### CURL Commands

- Create Collection

        curl -X POST "http://localhost:8000/api/v1/collection/create" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"name":"testcollection"}'


- Load Data file into collection
  
      curl -X POST -F "file=@test.txt" -F "CollectionName=testcollection" "http://localhost:8000/api/v1/collection/insert"

- Query the collection

        curl -X POST "http://localhost:8000/api/v1/collection/query" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"Query":"When the disk connection spike up to 100% ?" , "CollectionName":"testcollection"}'

- Get Token with Client Secrets 
        
        curl -X POST "http://localhost:8080/realms/coglex-realm/protocol/openid-connect/token" -H "Content-Type: application/x-www-form-urlencoded" -d "client_id=browser-app" -d "client_secret=JgaeKLEmr6WJljwuIMnsv9ZojwwTZL6l" -d "grant_type=client_credentials"

- Get Token for User that registered

        curl -X POST "http://localhost:8080/realms/coglex-realm/protocol/openid-connect/token" -H "Content-Type: application/x-www-form-urlencoded" -d "client_id=browser-app" -d "client_secret=JgaeKLEmr6WJljwuIMnsv9ZojwwTZL6l" -d "grant_type=password" -d "password=test" -d "username=test@test.com"
        
        curl -X POST "http://localhost:8080/realms/coglex-realm/protocol/openid-connect/token" -H "Content-Type: application/x-www-form-urlencoded" -d "client_id=browser-app" -d "client_secret=JgaeKLEmr6WJljwuIMnsv9ZojwwTZL6l" -d "grant_type=password" -d "password=ceo" -d "username=ceo@ceo.com"

- Get Token with Authorization Code
      
       curl -X POST "http://localhost:8000/api/v1/auth/token" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"authorization_code":"607e3cb2-0986-412e-b43e-37443f19e073.db051bfc-fcc3-465d-a097-01690c44ceed.94faa181-b1cb-4d45-8d29-a93f4345595f"}'  

- Decode Token

       curl -X POST "http://localhost:8000/api/v1/auth/token-decode" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwdGhwenlSNmg3WVpmTlJjcFhMeWZLS293WmZUUmF1MjlQS1BfcUFoa0RJIn0.eyJleHAiOjE3NDAxNjcwNDIsImlhdCI6MTc0MDE2NTU0MiwiYXV0aF90aW1lIjoxNzQwMTY1NDk0LCJqdGkiOiI2NjRkZjJmYS04Y2U5LTQzNTQtYmMxZC0wNTQ5ZDViN2Q3YjYiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvcmVhbG1zL3FkcmFudC1nby1yZWFsbSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiIyOGYyMGE4OS1lYTc0LTQ3N2UtYjE2YS05YTRiOTFlODRkMjciLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ2ZWN0b3ItYWJhYyIsInNlc3Npb25fc3RhdGUiOiJkYjA1MWJmYy1mY2MzLTQ2NWQtYTA5Ny0wMTY5MGM0NGNlZWQiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIi8qIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsImRlZmF1bHQtcm9sZXMtcWRyYW50LWdvLXJlYWxtIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ2ZWN0b3ItYWJhYyI6eyJyb2xlcyI6WyJtYW5hZ2VyIiwiYWRtaW4iLCJlbXBsb3llZSJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIHB1YmxpYyBhdWQiLCJzaWQiOiJkYjA1MWJmYy1mY2MzLTQ2NWQtYTA5Ny0wMTY5MGM0NGNlZWQiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6InVzZXIxIHVzZXIxIiwicHJlZmVycmVkX3VzZXJuYW1lIjoidXNlcjFAZW1haWwuY29tIiwiZ2l2ZW5fbmFtZSI6InVzZXIxIiwiZmFtaWx5X25hbWUiOiJ1c2VyMSIsImVtYWlsIjoidXNlcjFAZW1haWwuY29tIn0.rQO5MTmNiWCXlfq3zbglI0WDSGFrS8Fd-SwpNKJdas49laDOEVXQhyhnXtnpDTHuToX2stf5KDIP1h9hmTHrz35mkXwGR-2v6mbXcsqeqXRXliznnHz5An_p_1x09hINNElDJVdEgeIMgr0XcqhM0pFE85baPptbUiEk6OD1TWUgc5Osqm9VOS4TLV0BCcs5o-TZQECGcbD1Q007L_xs9yLiecqIyxdm-mxrMIT18tOg14bP54aUFW2IL3CIml0xy9NQmZk_PKOlUpZDSULsp9oBp6ePNLCLtCRTv5tlfyARrmwPsb3YCU39Fk2EylZPdKDmYN_JObQ-oERxwwzClw"}'  

- Login via keycloak login page to get authorization code

      http://localhost:8080/realms/qdrant-go-realm/protocol/openid-connect/auth?client_id=vector-abac&redirect_uri=http://localhost:8081/callback&response_type=code&scope=email profile



curl -X POST "http://localhost:8080/realms/qdrant-go-realm/protocol/openid-connect/token" -H "Content-Type: application/x-www-form-urlencoded" -d "client_id=vector-abac" -d "client_secret=1DIS65nKdjupn6ohkRTZ4l3P3ixYQGK5" -d "grant_type=password" -d "password=user3" -d "username=user3@email.com"

http://localhost:8080/realms/qdrant-go-realm/protocol/openid-connect/auth?client_id=vector-abac&redirect_uri=http://localhost:8081/callback&response_type=code&scope=email profile

