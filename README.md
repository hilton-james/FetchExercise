# FetchTake-HomeExercise

## How to Run 
To run the program perform:
```bash
make run 
```
You can test the endpoints:
```bash
# /receipts/{id}/points
curl -v -X GET "localhost:5001/receipts/1111/
points"


# /receipts/process
curl -v -X POST "localhost:5001/receipts/proc
ess"

```