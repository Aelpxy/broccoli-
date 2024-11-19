# Fresh

Fresh is a modern, efficient file server powered by Go.

## Overview

Fresh provides a simple API to manage resources.

## Endpoints

### Bucket API

### **`POST /api/<bucket-name>`**

Create a new bucket.  
**Request Body**:

```json
{ "bucket_name": "a-fresh-object-bucket" }
```

**Response**:

```json
{
  "success": true,
  "data": {
    "bucket_name": "a-fresh-object-bucket"
  }
}
```

### **`PATCH /api/<bucket-name>`**

Update an existing bucket.

**Request Body**:

```json
{ "bucket_name": "a-fresh-updated-object-bucket" }
```

**Response**:

```json
{
  "success": true,
  "data": {
    "bucket_name": "a-fresh-updated-object-bucket"
  }
}
```

### **`GET /api/<bucket-name>`**

Retrieve information about a specific bucket.

**Response**:

```json
{
  "success": true,
  "data": {
    "bucket_name": "a-fresh-updated-object-bucket",
    "objects": 2
  }
}
```

### **`DELETE /api/<bucket-name>`**

Delete a bucket and its objects.

**Response**:

```json
{
  "success": true,
  "data": {
    "bucket_name": "a-fresh-updated-object-bucket",
    "objects": 2
  }
}
```

### Objects API

## License

This project is licensed under the [MIT License](./LICENSE).
