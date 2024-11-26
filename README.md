# FetchTake-HomeExercise

## How to Run 
To run the program perform:
```bash
make run 
```
You can test the endpoints:
### /receipts/{id}/points

```bash
# Request
curl -X POST localhost:5001/receipts/proce
ss -H "Content-Type: application/json" -d '{"id": "1", "retailer": "M&M Corner Market", "purch
aseDate": "2022-01-01", "purchaseTime": "13:01", "items": [{"shortDescription": "Mountain Dew 
12PK", "price": "6.49"}], "total": "6.49"}'

# Response
{"id":"89d77822-6e10-4476-8716-40b78fcd7254"}⏎  


```
### /receipts/process

```bash
# Request
curl -X GET "localhost:5001/receipts/e7663b11-1a66-47bb-9faf-297486809d5c/points"

# Response
{"points":20}⏎
```
## Docker

## Build
```bash
sudo docker build -f ./build/package/Dockerfile -t receipt:latest .
```
## Run
```bash
sudo docker run --rm -p 5001:5001 -e RECEIPT_PORT="5001" -e RECEIPT_DEBUG="TRUE" receipt-app
```