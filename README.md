#### RUN
````
go run main.go
````

####API
````
- convert decimal to binary
localhost:8088/binary?output=binary
    
    body
    {
        "input": int
    }
    example
    {
        "input": 29
    }

- convert binary to decimal
localhost:8088/binary?output=decimal
    
    body
    {
        "input": int
    }
    example
    {
        "input": 1001
    }

- find longest palindrome
localhost:8088/palindrome
    
    body
    {
        "word": string
    }
    example
    {
        "word": "aku suka makan nasi"
    }

````