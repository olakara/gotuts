
## Maps in Go
- Maps are built-in data structures that represent key value pairs
- In other languages, they are known as associative arrays, dictionaries, or hash maps
- In .NET, they are called dictionaries
- They are un ordered collections of key value pairs
- When defining a map, you must specify the type of the key and the type of the value
- The key must be unique
- The value can be duplicated
- Maps are reference types
- You can use the make function to create a map
- You can use the delete function to remove a key value pair from a map
- You can use the len function to get the number of key value pairs in a map
- You can use the range keyword to iterate over a map

### Declaring a map:
```go
   users := map[string]string{
      "abdel": "abdel@gmail.com",
      "jane": "jane@gmail.com",
   }
```

### Accessing a value:
```go
   fmt.Println(users["abdel"])
```

### Adding a key value pair:
```go
   users["mario"] = "mario@gmail.com"
```

### Deleting a key value pair:
```go
   delete(users, "mario")
```