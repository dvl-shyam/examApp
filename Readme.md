
# Install dependencies
```bash
'go mod tidy'
```


Install Mongodb Drivers
```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

 # Set Environment Variables
 ```bash
export MONGO_URI="mongodb://localhost:{PORT}"
export MONGO_USER="your-access-key"
export MONGO_PASS="your-secret-key"
export PORT="{PORT}"
```


 --------------APIS-----------------



Create Person
```bash
curl --location 'http://localhost:{PORT}/person/create' \
--header 'Content-Type: application/json' \
--data '{
    "firstName": "Shyam",
    "middleName": "",
    "lastName": "Kuntal",
    "gender": "MALE",
    "homeDistrict": "Mathura",
    "dob": "2000-06-19",
    "stateOfDomicile": "UP",
    "fatherFirstName": "Bhav",
    "fatherMiddleName": "",
    "fatherLastName": "Singh",
    "boardName": "CBSE",
    "yearOfPassing": "2019",
    "rollNumber": "12345",
    "address": "Vrindavan",
    "houseNoVillage": "38",
    "state": "UP",
    "district": "Mathura",
    "city": "Mathura",
    "pinCode": 281004
}
'
```


GET Person
```bash
curl --location --request GET 'http://localhost:{PORT}/person/getone/673430c6d7d801e65a3a54cb'
```


Update Person
```bash
curl --location --request PUT 'http://localhost:{PORT}/person/update?id=673430c6d7d801e65a3a54cb' \
--header 'Content-Type: application/json' \
--data '{
    "firstName": "Ram",
    "middleName": "Singh",
    "lastName": "Kuntal",
    "gender": "MALE",
    "homeDistrict": "Mathura",
    "dob": "2002-06-19",
    "stateOfDomicile": "UP",
    "fatherFirstName": "Bhav",
    "fatherMiddleName": "",
    "fatherLastName": "Singh",
    "boardName": "CBSE",
    "yearOfPassing": "2019",
    "rollNumber": "12345",
    "address": "Vrindavan",
    "houseNoVillage": "38",
    "state": "UP",
    "district": "Mathura",
    "city": "Mathura",
    "pinCode": 281004
}'
```

Get Age
```bash
curl --location --request GET 'http://localhost:{PORT}/person/getage?id=673430c6d7d801e65a3a54cb'
```

Delete Person
```bash
curl --location --request DELETE 'http://localhost:{PORT}/person/delete?id=673430c6d7d801e65a3a54cb'
```
