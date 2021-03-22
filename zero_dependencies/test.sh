curl localhost:8080/admin -u admin:secret
curl localhost:8080/coasters -X POST -d '{name:Taron,inPark: Phantsialand,height: 30,manufacturer: Intamin}' -H Content-Type: application/json
