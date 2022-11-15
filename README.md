# portal-custom-auth
Example demonstrating customAuth usage plan in Gloo Portal

## Instructions
1. Deploy the custom passthrough auth service
```
kubectl apply -f auth-service.yaml
```
2. Deploy the petstore example API
```
kubectl apply -f petstore.yaml
```
3. Deploy the petstore portal
```
kubectl apply -f petstore-portal.yaml
```
4. Login to the portal at ```petstore.example.com:31500``` with userid ```dev1``` and password ```Pa$$w0rd```
5. Generate an API Key
6. Using the Try It Out page for the UI, authorize using your API Key and execute a GET request for ```/api/pets```
7. Note the following output in the ```extauth-httpservice``` pod which indicates the custom auth service is being invoked through customAuth configuration in the usage plan:
```
Listening on port 9001 for auth requests
received request with url: /, with headers map[Accept-Encoding:[gzip] Api-Key:[ZDdkNzRlMTQtODNmOC1lNTM5LTdlMTUtOTJkNWUyYzE4Yzkw] Content-Length:[0] User-Agent:[Go-http-client/1.1]]
exchanging API Key for access token ...
```
8. Verify that an Authorization header has been added to the upstream HTTP request by view the logs for the ```petstore``` pod
```
Authorization: Bearer abc123xyz890
```

## Notes
- Source for the custom auth service can be found in the ```auth-service``` directory. Be sure to change the image repo name if you plan on customizing along with the associate container def in the Deployment
- I'm running locally in kind and mapping ```:31500``` on my machine to the Kube ingress. You'll need to manage the DNS resolution (I added petstore.example.com and api.example.com directly to my /etc/hosts) based on your own setup. If your hostname or port information are different then make sure you update the Environment and Portal resources in ```petstore-portal.yaml```.
