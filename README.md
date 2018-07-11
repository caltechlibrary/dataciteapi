
# dataciteapi

This is a go package for working with the DataCite API. It is inspired by
work my colleague Tom Morrel has done in Python.  This package is meant 
to follow the guidelines for interacting with the public API at 
api.datacite.org. It also follows the same form as the golang 
[CrossRef API](https://github.com/caltechlibrary/crossrefapi) 
developed previously at Caltech Library.

## Go package example

```go
    client, err := dataciteapi.NewDataCiteClient("jane.doe@library.example.edu")
    if err != nil {
        // handle error...
    }
    works, err := client.Works("10.1037/0003-066x.59.1.29")
   
    if err != nil {
        // handle error...
    }
    // continue processing your "works" result...
```

## Command line example

```
    dataciteapi -mailto="jane.doe@library.example.edu" works "10.1037/0003-066x.59.1.29"
```

## Reference

+ [DataCite API Docs](https://support.datacite.org/docs/api)
+ [DataCite Metadata Schema v4.1](http://schema.datacite.org/meta/kernel-4.1/)
