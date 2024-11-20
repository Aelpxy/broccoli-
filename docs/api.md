# Overview

Fresh provides a simple API to manage resources.

# Endpoints

## Bucket API

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

## Objects API

### **`GET /api/<bucket-name>/objects`**

Retrieve a list of all objects within a specified bucket.

**Response**:

```json
{
  "success": true,
  "data": [
    {
      "object_id": "1",
      "object_name": "object-1.jpg",
      "size": 1024,
      "created_at": "2024-11-01T12:00:00Z"
    },
    {
      "object_id": "2",
      "object_name": "object-2.png",
      "size": 2048,
      "created_at": "2024-11-02T12:00:00Z"
    }
  ]
}
```

### **`GET /api/<bucket-name>/objects/<object-id>`**

Retrieve detailed information about a specific object within the bucket.

**Response**:

```json
{
  "success": true,
  "data": {
    "object_id": "1",
    "object_name": "object-1.jpg",
    "size": 1024,
    "created_at": "2024-11-01T12:00:00Z",
    "metadata": {
      "type": "image",
      "format": "jpg"
    }
  }
}
```

### **`POST /api/<bucket-name>/objects`**

Create a new object in the specified bucket.

**Request Body**:

```json
// TODO
```

**Response**:

```json
{
  "success": true,
  "data": {
    "object_id": "3",
    "object_name": "new-object.txt",
    "size": 256,
    "created_at": "2024-11-20T12:00:00Z"
  }
}
```

### **`UPDATE /api/<bucket-name>/objects/<object-id>`**

Update an existing object in the bucket. This can involve changing the object’s name, data, or metadata.

**Request Body**:

```json
// TODO
```

**Response**:

```json
{
  "success": true,
  "data": {
    "object_id": "3",
    "object_name": "updated-object.txt",
    "size": 512,
    "created_at": "2024-11-20T12:00:00Z"
  }
}
```

### **`DELETE /api/<bucket-name>/objects/<object-id>`**

Delete a specific object from the bucket.

**Response**:

```json
{
  "success": true,
  "data": {
    "object_id": "3",
    "object_name": "updated-object.txt"
  }
}
```
