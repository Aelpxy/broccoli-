# Fresh

## Technical Specifications

When you create a bucket, a record is created in SQLite, which is running in WAL mode for the best possible performance. The same happens when an object is uploaded — a record is created that maps the object to the bucket. Behind the scenes, it’s actually a folder (the core bucket), which can contain other folders and files (which we call objects).

- A bucket is a container of files called objects.

- An object is a file with metadata attached.

Each bucket has a `metadata.msgpack` file, which you should not delete, as it stores metadata for faster lookups to retrieve and send back objects. If this file is not found, a new one will be created based on the files in the bucket and the SQLite database.

MessagePack was chosen for its efficiency. Initially, I wanted to store lookups in JSON, but I decided against it because JSON uses significantly more space. The main reason for using SQLite, however, was simplicity.

I plan to load the entire `metadata.msgpack` file into a Bloom filter (which will be updated whenever a new object is uploaded) and then perform lookups.

### Is this actually fast enough?

For the most part, it should be fast enough in my opinion (this wasn't really meant for production use).
