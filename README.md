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
      
       curl -X POST "http://localhost:8000/api/v1/auth/token" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"authorization_code":"92281b75-4766-48cd-a8d0-3ae4f7ffdeae.7b74fb53-d6b9-480b-a7a2-bd4277a2a7e8.94faa181-b1cb-4d45-8d29-a93f4345595f"}'  

- Decode Token

       curl -X POST "http://localhost:8000/api/v1/auth/token-decode" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwdGhwenlSNmg3WVpmTlJjcFhMeWZLS293WmZUUmF1MjlQS1BfcUFoa0RJIn0.eyJleHAiOjE3NDAxODEzMjMsImlhdCI6MTc0MDE4MTI2MywiYXV0aF90aW1lIjoxNzQwMTgxMjIxLCJqdGkiOiI3OWQzNjc5Mi1kYjJkLTRhOTYtYjIwNS05YzY0YzNlMTI2ZjIiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvcmVhbG1zL3FkcmFudC1nby1yZWFsbSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiIyOGYyMGE4OS1lYTc0LTQ3N2UtYjE2YS05YTRiOTFlODRkMjciLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ2ZWN0b3ItYWJhYyIsInNlc3Npb25fc3RhdGUiOiI3Yjc0ZmI1My1kNmI5LTQ4MGItYTdhMi1iZDQyNzdhMmE3ZTgiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIi8qIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsImRlZmF1bHQtcm9sZXMtcWRyYW50LWdvLXJlYWxtIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ2ZWN0b3ItYWJhYyI6eyJyb2xlcyI6WyJtYW5hZ2VyIiwiYWRtaW4iLCJlbXBsb3llZSJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIHB1YmxpYyBhdWQiLCJzaWQiOiI3Yjc0ZmI1My1kNmI5LTQ4MGItYTdhMi1iZDQyNzdhMmE3ZTgiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6InVzZXIxIHVzZXIxIiwicHJlZmVycmVkX3VzZXJuYW1lIjoidXNlcjFAZW1haWwuY29tIiwiZ2l2ZW5fbmFtZSI6InVzZXIxIiwiZmFtaWx5X25hbWUiOiJ1c2VyMSIsImVtYWlsIjoidXNlcjFAZW1haWwuY29tIn0.Wj3FXayNv0LL9ww6pU1h3ZA-mpaSf6cboxFEhE9Vxy8m7lnEae4uuHjAcfCjd4hR52p8k1IOeEh_br_72ULf7y7Tm_b1HNOjNPCnXRt6Uxo2TOegZl9CMyGkZyitzPiFdgOOVHGHhXgUmgaxxc0hoKCxraf56bS6H6yCuy7DvHpaB4fvpHCHwQFB1TctPadC2uLsZDQTbGTTdqm3bOZWQ4VVc6TuAvD_XtMBgC0Q4Ua_aocRuJbAPLfr_stwJ0_n2KI6d15W8WrGfWaPM9FSUrpBG4yGLN23GsT4DDlN0vK7xiR-axSr0q_HLnrG6gwNk9-0IhvLjC3dA-fRsYvB7g"}'  

- Login via keycloak login page to get authorization code

      http://localhost:8080/realms/qdrant-go-realm/protocol/openid-connect/auth?client_id=vector-abac&redirect_uri=http://localhost:8081/callback&response_type=code&scope=email profile

