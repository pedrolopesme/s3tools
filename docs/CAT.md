<h1 align="center">
  S3 Tools 
</h1>

#### CAT
The cat utility reads files sequentially, writing them to the standard output.  
The file operands are processed in command-line order. 

Basic usage:

```
$ s3tools cat --help
The cat utility reads files sequentially, writing them to the standard output.  
The file operands are processed in command-line order. For example:

	s3tools cat bucket-name file1 file2 file3

will print the contents of both file1, file2 and file3 to the standard output.
```

Ex1: Cat all files whose name start with server_log and error_log
```
s3tools cat bucket-test server_log* error_log*
```

Ex2: Cat all files whose name start with log, following by a number
```
s3tools grep bucket-test "log[0-9"
```

By default, s3tools cat command does not store any file into your local storage.