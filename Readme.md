```bash
Install dependencies
curl 'go mod tidy'
```

```bash
Install Mongodb Drivers
curl 'go get go.mongodb.org/mongo-driver/mongo'
curl 'go get go.mongodb.org/mongo-driver/mongo/options'
```

```bash
----------------APIS-----------------
```
```bash

Create Person
curl --location 'http://localhost:8000/person/create' \
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

```bash
GET Person
curl --location 'http://localhost:8000/person/getone/673430c6d7d801e65a3a54cb'
```

```bash
Update Person
curl --location --request PATCH 'http://localhost:8000/person/update?id=673430c6d7d801e65a3a54cb' \
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

```bash
Get Age
curl --location 'http://localhost:8000/person/getage?id=673430c6d7d801e65a3a54cb'
```

```bash
Delete Person
curl --location --request DELETE 'http://localhost:8000/person/delete?id=673430c6d7d801e65a3a54cb'
```
