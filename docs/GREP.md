<h1 align="center">
  S3 Tools 
</h1>

#### Grep
Grep command allows you to make regex powered searches into 
you S3 files. 

Basic usage:

```
$ s3tools grep --help
With Grep you can search for files stored at you bucket recursively. For example:

s3tools grep my-bucket "search string"

$ s3tools grep bucket-test dummy
```

Ex1: Filter all lines containing a string "error"
```
s3tools grep bucket-test error
```

Ex2: Filter all lines containing an URL:
```
s3tools grep bucket-test "(http|https):\/\/([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:\/~+#-]*[\w@?^=%&\/~+#-])?"
```

By default, s3tools grep command does not store any
file into your local storage.