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
      
       curl -X POST "http://localhost:8000/api/v1/auth/token" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"authorization_code":"72d1971d-2765-4acb-9383-d66ea421f4b3.87b6cf30-0d68-4b10-a1c9-226665dcb401.94faa181-b1cb-4d45-8d29-a93f4345595f"}'  

- Decode Token

       curl -X POST "http://localhost:8000/api/v1/auth/token-decode" -H "Content-Type: application/json" -H "Accept: application/json" -d '{"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwdGhwenlSNmg3WVpmTlJjcFhMeWZLS293WmZUUmF1MjlQS1BfcUFoa0RJIn0.eyJleHAiOjE3NDAxNzEwNTAsImlhdCI6MTc0MDE2OTU1MCwiYXV0aF90aW1lIjoxNzQwMTY4MDMxLCJqdGkiOiI0ZDAzZDI1My0zY2U4LTQ2NmEtYjMxNS0xZDE3Mjk5N2Y4ZWMiLCJpc3MiOiJodHRwOi8vbG9jYWxob3N0OjgwODAvcmVhbG1zL3FkcmFudC1nby1yZWFsbSIsImF1ZCI6ImFjY291bnQiLCJzdWIiOiIyOGYyMGE4OS1lYTc0LTQ3N2UtYjE2YS05YTRiOTFlODRkMjciLCJ0eXAiOiJCZWFyZXIiLCJhenAiOiJ2ZWN0b3ItYWJhYyIsInNlc3Npb25fc3RhdGUiOiI4N2I2Y2YzMC0wZDY4LTRiMTAtYTFjOS0yMjY2NjVkY2I0MDEiLCJhY3IiOiIwIiwiYWxsb3dlZC1vcmlnaW5zIjpbIi8qIl0sInJlYWxtX2FjY2VzcyI6eyJyb2xlcyI6WyJvZmZsaW5lX2FjY2VzcyIsImRlZmF1bHQtcm9sZXMtcWRyYW50LWdvLXJlYWxtIiwidW1hX2F1dGhvcml6YXRpb24iXX0sInJlc291cmNlX2FjY2VzcyI6eyJ2ZWN0b3ItYWJhYyI6eyJyb2xlcyI6WyJtYW5hZ2VyIiwiYWRtaW4iLCJlbXBsb3llZSJdfSwiYWNjb3VudCI6eyJyb2xlcyI6WyJtYW5hZ2UtYWNjb3VudCIsIm1hbmFnZS1hY2NvdW50LWxpbmtzIiwidmlldy1wcm9maWxlIl19fSwic2NvcGUiOiJlbWFpbCBwcm9maWxlIHB1YmxpYyBhdWQiLCJzaWQiOiI4N2I2Y2YzMC0wZDY4LTRiMTAtYTFjOS0yMjY2NjVkY2I0MDEiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6InVzZXIxIHVzZXIxIiwicHJlZmVycmVkX3VzZXJuYW1lIjoidXNlcjFAZW1haWwuY29tIiwiZ2l2ZW5fbmFtZSI6InVzZXIxIiwiZmFtaWx5X25hbWUiOiJ1c2VyMSIsImVtYWlsIjoidXNlcjFAZW1haWwuY29tIn0.d211SG90XxvUukj98TrskH_kWnfkIT8q55O92UboSf7gHBLYlnoCSeyxVprb6N_QjCDXC1ABo33guoIuPQdphfuHhf9MSzUB26XyNQc43pNluVY-GSoTIzwQJYn4AaprUJqoUiw_gK0RyZ_7AFBadJOiL2RvnP3evyBcM09fjot4zyi3NSlk6oYnWLIOsJx4VZgP_dMvG5MN_RYcW1Bx2AP81Lu_pFF2QcWXiZZ3PyJilxs91EJEJ4k0k-3wIDF-QCr-_vMGOqC2t4_g-dnodLrTvb3MCMNcTe98Z6AUWK-HNLlSsrbkdhLvXm8EuYYKMNAIecENUp9u5PiDooxF9g"}'  

- Login via keycloak login page to get authorization code

      http://localhost:8080/realms/qdrant-go-realm/protocol/openid-connect/auth?client_id=vector-abac&redirect_uri=http://localhost:8081/callback&response_type=code&scope=email profile

