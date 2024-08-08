# Take home assignment

# Simple Bank Transaction

## USER service

user service is portal to interact with user and account
- register API

to register the user
```curl
curl --location 'localhost:8080/api/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test",
    "email": "test@test.com",
    "password": "test"
}'
```

- login API

when login is successful you will receive jwt token and the token must be use to authorize the API
```curl
curl --location 'localhost:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email":"test@test.com",
    "password":"test"
}'
```

- currency API

this api is a list of currency
```curl
curl --location 'localhost:8080/api/currencies'
```

- product API

this api is a list of products
```curl
curl --location 'localhost:8080/api/products'
```

- create Account API

this api is creating account based on product, currency and user. It also required to send the token in Authorization header
```curl
curl --location 'localhost:8080/api/accounts' \
--header 'Authorization: bearer {{ TOKEN }}' \
--header 'Content-Type: application/json' \
--data '{
    "currency_code":"SGD",
    "product_id": 1
}'
```

- get accounts API

this api is listing all of the registered accounts. It also required to send the token in Authorization header
```
curl --location 'localhost:8080/api/accounts/me' \
--header 'Authorization: bearer {{ TOKEN }}'
``` 

- get accounts balance and history transaction API

this api is getting the account balance into payment service with GRPC. It also required to send the token in Authorization header
```
curl --location 'localhost:8080/api/accounts/10020' \
--header 'Authorization: bearer {{ TOKEN }}'
``` 

## Payment service

payment service is portal to interact with transactional data and transfer from one account to another

- Send Payment API

there's two condition to operate this API

1. to send to another account with body request 

the `account_id` keys is *origin* of account_id, 
the `to_account_id` keys is *destination* of account_id and
the amount is amount of the transaction. It also required to send the token in Authorization header, in the backend the token will be checked by GRPC calls.

```curl
curl --location 'localhost:8081/api/payments/send' \
--header 'Authorization: bearer {{ TOKEN }}' \
--header 'Content-Type: application/json' \
--data '{
    "account_id":10019,
    "to_account_id":10020,
    "amount": 10
}'
```



2. to top up the account / simulation of ATM top-up

the `account_id` keys and the `to_account_id` keys is same, so the amount is will be added into balance directly


```curl
curl --location 'localhost:8081/api/payments/send' \
--header 'Authorization: bearer {{ TOKEN }}' \
--header 'Content-Type: application/json' \
--data '{
    "account_id":10020,
    "to_account_id":10020,
    "amount": 10
}'
```


- Withdraw Payment API

this API is used to get to amount from the current balance. It also required to send the token in Authorization header, in the backend the token will be checked by GRPC calls.

```
curl --location 'localhost:8081/api/payments/withdraw' \
--header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJuZ3VqaS1hdXRoIiwiZXhwIjoxNzIzMTcxMTA0LCJpYXQiOjE3MjMwODQ3MDQsImp0aSI6IjAxOTEyZmRhLTA5MzItNzIzMy1hMGRiLWZhZjNhZmM0MmNiNiJ9.B_Bd_x9BPlXn7LOODryE028O_RHJosxF6kIBLp8kjBg' \
--header 'Content-Type: application/json' \
--data '{
    "account_id":10020,
    "amount": 20
}'
```



## Description:
Build 2 Backend services which manages user’s accounts and transactions (send/withdraw). 

In Account Manager service, we have:
- User: Login with Id/Password
- Payment Account: One user can have multiple accounts like credit, debit, loan...
- Payment History: Records of transactions

In Payment Manager service, we have:
- Transaction: Include basic information like amount, timestamp, toAddress, status...
- We have a core transaction process function, that will be executed by `/send` or `/withdraw` API:

```js
function processTransaction(transaction) {
    return new Promise((resolve, reject) => {
        console.log('Transaction processing started for:', transaction);

        // Simulate long running process
        setTimeout(() => {
            // After 30 seconds, we assume the transaction is processed successfully
            console.log('transaction processed for:', transaction);
            resolve(transaction);
        }, 30000); // 30 seconds
    });
}

// Example usage
let transaction = { amount: 100, currency: 'USD' }; // Sample transaction input
processTransaction(transaction)
    .then((processedTransaction) => {
        console.log('transaction processing completed for:', processedTransaction);
    })
    .catch((error) => {
        console.error('transaction processing failed:', error);
    });
```

Features:
- Users need to register/log in and then be able to call APIs.
- APIs for 2 operations send/withdraw. Account statements will be updated after the transaction is successful.
- APIs to retrieve all accounts and transactions per account of the user.
- Write Swagger docs for implemented APIs (Optional)
- Auto Debit/Recurring Payments: Users should be able to set up recurring payments. These payments will automatically be processed at specified intervals. (Optional)

### Tech-stack:
- Recommend using authentication 3rd party: Supertokens, Supabase...
- `NodeJs/Golang` for API server (`Fastify/Gin` framework is the best choices)
- `PostgreSQL/MongoDB` for Database. Recommend using `Prisma` for ORM.
- `Docker` for containerization. Recommend using `docker-compose` for running containers.
 
## Target:
- Good document/README to describe your implementation.
- Make sure app functionality works as expected. Run and test it well.
- Containerized and run the app using Docker.
- Using `docker-compose` or any automation script to run the app with single command is a plus.
- Job schedulers utilization is a plus
