### CURL Commands

- Create Collection

        curl -X POST "http://localhost:8000/api/v1/collection/create" -H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwdGhwenlSNmg3WVpmTlJjcFhMeWZLS293WmZUUmF1MjlQS1BfcUFoa0RJIn0.eyJleHAiOjE3NDA1ODgwNjAsImlhdCI6MTc0MDU4ODAwMCwiYXV0aF90aW1lIjoxNzQwNTg3OTU2LCJqdGkiOiI3NmRkMzg0OS1jOWU3LTQ1MDYtODFjNC05MzRmNmNmOTVmMDUiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvcmVhbG1zL3FkcmFudC1nby1yZWFsbSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiIyOGYyMGE4OS1lYTc0LTQ3N2UtYjE2YS05YTRiOTFlODRkMjciLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ2ZWN0b3ItYWJhYyIsInNlc3Npb25fc3RhdGUiOiI1MTMxZDY5YS0yOGNkLTRlMDctOTgxOS1kMzE3MTlhODIyMTYiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIi8qIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsImRlZmF1bHQtcm9sZXMtcWRyYW50LWdvLXJlYWxtIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ2ZWN0b3ItYWJhYyI6eyJyb2xlcyI6WyJtYW5hZ2VyIiwiYWRtaW4iLCJlbXBsb3llZSJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIHB1YmxpYyBhdWQiLCJzaWQiOiI1MTMxZDY5YS0yOGNkLTRlMDctOTgxOS1kMzE3MTlhODIyMTYiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6InVzZXIxIHVzZXIxIiwicHJlZmVycmVkX3VzZXJuYW1lIjoidXNlcjFAZW1haWwuY29tIiwiZ2l2ZW5fbmFtZSI6InVzZXIxIiwiZmFtaWx5X25hbWUiOiJ1c2VyMSIsImVtYWlsIjoidXNlcjFAZW1haWwuY29tIn0.mMBH2i5VIveXr93H1nFlnpQM4YsbZhuYDttxSvUunQJvnDrnOLvsAz4rfyPHWb6W3CVlg_oCe5kUlcin5f1_Q8hiT7H8KsKYKJ2lgKd-IIn66HaU70WIHYWk6Si16T6VmF49XaOeeL9uo9EnykIJVFGuF0WUIy3zY-C-cZNJpYrpQRjyQts5taZ-5ntPRxAd0CZ07GMIarHH9ZQrvMCQ6gkgBi2vhts3PxVaSg0Su3q_uxSL-LuTrzxkVAyziZq83rYh1tndY7f3BDWdHmNFod4N5cnp0CYPTVKNZJqMqTWlxGVMjKizW8hsBz0mUJc7Yo0QOLJBc5FHJ5GnQuHb4Q" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"name":"testcollection"}'


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
      
       curl -X POST "http://localhost:8000/api/v1/auth/token" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"authorization_code":"dd85e623-b6ce-4222-931c-d9c20510a8c0.5131d69a-28cd-4e07-9819-d31719a82216.94faa181-b1cb-4d45-8d29-a93f4345595f"}'  

- Decode Token

       curl -X POST "http://localhost:8000/api/v1/auth/token-decode" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwdGhwenlSNmg3WVpmTlJjcFhMeWZLS293WmZUUmF1MjlQS1BfcUFoa0RJIn0.eyJleHAiOjE3NDAxODEzMjMsImlhdCI6MTc0MDE4MTI2MywiYXV0aF90aW1lIjoxNzQwMTgxMjIxLCJqdGkiOiI3OWQzNjc5Mi1kYjJkLTRhOTYtYjIwNS05YzY0YzNlMTI2ZjIiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvcmVhbG1zL3FkcmFudC1nby1yZWFsbSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiIyOGYyMGE4OS1lYTc0LTQ3N2UtYjE2YS05YTRiOTFlODRkMjciLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ2ZWN0b3ItYWJhYyIsInNlc3Npb25fc3RhdGUiOiI3Yjc0ZmI1My1kNmI5LTQ4MGItYTdhMi1iZDQyNzdhMmE3ZTgiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbIi8qIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsImRlZmF1bHQtcm9sZXMtcWRyYW50LWdvLXJlYWxtIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ2ZWN0b3ItYWJhYyI6eyJyb2xlcyI6WyJtYW5hZ2VyIiwiYWRtaW4iLCJlbXBsb3llZSJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIHB1YmxpYyBhdWQiLCJzaWQiOiI3Yjc0ZmI1My1kNmI5LTQ4MGItYTdhMi1iZDQyNzdhMmE3ZTgiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6InVzZXIxIHVzZXIxIiwicHJlZmVycmVkX3VzZXJuYW1lIjoidXNlcjFAZW1haWwuY29tIiwiZ2l2ZW5fbmFtZSI6InVzZXIxIiwiZmFtaWx5X25hbWUiOiJ1c2VyMSIsImVtYWlsIjoidXNlcjFAZW1haWwuY29tIn0.Wj3FXayNv0LL9ww6pU1h3ZA-mpaSf6cboxFEhE9Vxy8m7lnEae4uuHjAcfCjd4hR52p8k1IOeEh_br_72ULf7y7Tm_b1HNOjNPCnXRt6Uxo2TOegZl9CMyGkZyitzPiFdgOOVHGHhXgUmgaxxc0hoKCxraf56bS6H6yCuy7DvHpaB4fvpHCHwQFB1TctPadC2uLsZDQTbGTTdqm3bOZWQ4VVc6TuAvD_XtMBgC0Q4Ua_aocRuJbAPLfr_stwJ0_n2KI6d15W8WrGfWaPM9FSUrpBG4yGLN23GsT4DDlN0vK7xiR-axSr0q_HLnrG6gwNk9-0IhvLjC3dA-fRsYvB7g"}'  

- Login via keycloak login page to get authorization code

      http://localhost:8080/realms/qdrant-go-realm/protocol/openid-connect/auth?client_id=vector-abac&redirect_uri=http://localhost:8081/callback&response_type=code&scope=email profile

