# User Service API Documentation

## Endpoints
### 1. **Login User**

**POST** `/login`

**Request Body**:
```json
{
  "username": "string",
  "password": "string"
}
```
**OR**
```json
{
  "email": "string",
  "password": "string"
}
```

**Response**:
- 200 OK:
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "user123",
  "user_fullname": "John Doe",
  "user_birthday": "1990-01-01",
  "user_gender": 1,
  "user_profile": "profile_data"
}
```
- 400 Bad Request:
```plaintext
  username or email is empty || invalid password
```
- 404 Not Found
```plaintext
  user not found
```
##
### 2. **Register User**

**POST** `/register`

**Request Body**:
```json
{
  "username": "string",
  "email": "example@mail.com",
  "password": "string8chars",
  "user_fullname": "string",
  "user_birthday": "YYYY-MM-DD",
  "user_gender": 0 // 0 for female, 1 for male, 2 for other
}
```

**Response**:
- 200 OK:
```plaintext
1
```
- 400 Bad Request:
```plaintext
  missing or invalid fields
```
- 409 Conflict:
```plaintext
  email or username already exists
```
##
### 3. **Get User**

**GET** `/secure/getuser/{userId}`

**Headers:**
```plaintext
  Authorization: Brarer <token>
```

**Response**
- 200 OK:
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "user123",
  "user_fullname": "User Name",
  "user_birthday": "2000-01-01",
  "user_gender": 1,
  "user_profile": "profile_data"
}
```
- 401 Unauthorized:
```plaintext
  unauthorized
```
- 403 Forbidden:
```plaintext
  forbidden
```
##
### 4. **Update User Information**

**POST** `/secure/updateuser/{userId}/info`

**Headers:**
```plaintext
  Authorization: Brarer <token>
```

**Request Body**:
```json
{
  "user_fullname": "string",
  "user_birthday": "YYYY-MM-DD",
  "user_gender": 0 // 0 for female, 1 for male, 2 for other
}
```

**Response**
- 200 OK:
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "user123",
  "user_fullname": "User Name",
  "user_birthday": "2000-01-01",
  "user_gender": 1,
  "user_profile": "profile_data"
}
```
- 401 Unauthorized:
```plaintext
  unauthorized
```
- 403 Forbidden:
```plaintext
  forbidden
```
##
### 5. **Update User Password**

**POST** `/secure/updateuser/{userId}/password`

**Headers:**
```plaintext
  Authorization: Brarer <token>
```

**Request Body**:
```json
{
  "password": "string"
}
```

**Response**
- 200 OK:
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "user123",
  "user_fullname": "User Name",
  "user_birthday": "2000-01-01",
  "user_gender": 1,
  "user_profile": "profile_data"
}
```
- 401 Unauthorized:
```plaintext
  unauthorized
```
- 403 Forbidden:
```plaintext
  forbidden
```
##
### 6. **Update User Email**

**POST** `/secure/updateuser/{userId}/email`

**Headers:**
```plaintext
  Authorization: Brarer <token>
```

**Request Body**:
```json
{
  "email": "string"
}
```

**Response**
- 200 OK:
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "user123",
  "user_fullname": "User Name",
  "user_birthday": "2000-01-01",
  "user_gender": 1,
  "user_profile": "profile_data"
}
```
- 401 Unauthorized:
```plaintext
  unauthorized
```
- 403 Forbidden:
```plaintext
  forbidden
```
##
### 7. **Update User Username**

**POST** `/secure/updateuser/{userId}/username`

**Headers:**
```plaintext
  Authorization: Brarer <token>
```

**Request Body**:
```json
{
  "username": "string"
}
```

**Response**
- 200 OK:
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "user123",
  "user_fullname": "User Name",
  "user_birthday": "2000-01-01",
  "user_gender": 1,
  "user_profile": "profile_data"
}
```
- 401 Unauthorized:
```plaintext
  unauthorized
```
- 403 Forbidden:
```plaintext
  forbidden
```
##
### 8. **Update User Profile**

**POST** `/secure/updateuser/{userId}/username`

**Headers:**
```plaintext
  Authorization: Brarer <token>
```

**Request Body**:
```json
{
  "user_profile": "string"
}
```

**Response**
- 200 OK:
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "user123",
  "user_fullname": "User Name",
  "user_birthday": "2000-01-01",
  "user_gender": 1,
  "user_profile": "profile_data"
}
```
- 401 Unauthorized:
```plaintext
  unauthorized
```
- 403 Forbidden:
```plaintext
  forbidden
```
##
### 9. **Delete User**

**Delete** `/secure/deleteuser/{userId}`

**Headers:**
```plaintext
  Authorization: Brarer <token>
```

**Response**
- 200 OK:
```plaintext
  1
```
- 401 Unauthorized:
```plaintext
  unauthorized
```
- 403 Forbidden:
```plaintext
  forbidden
```