# Complete Go File I/O Functions Guide

## 📚 Function Categories & When to Use Each

### 🔍 **READING FUNCTIONS**

#### **1. os.ReadFile()**
- **What**: Reads entire file into memory
- **Returns**: `[]byte, error`
- **Best for**: Small files, configuration files, one-time reads
- **Memory**: Loads entire file into memory
- **Performance**: Fast for small files, memory-intensive for large files

```go
data, err := os.ReadFile("config.txt")
```

#### **2. File.Read()**
- **What**: Reads into provided buffer
- **Returns**: `(n int, err error)`
- **Best for**: Streaming, large files, custom buffer management
- **Memory**: Uses provided buffer size
- **Performance**: Efficient for large files

```go
f, _ := os.Open("large_file.txt")
buf := make([]byte, 1024)
n, err := f.Read(buf)
```

#### **3. io.ReadAll()**
- **What**: Reads all data from any reader
- **Returns**: `([]byte, error)`
- **Best for**: Network streams, pipes, any io.Reader
- **Memory**: Loads all data into memory
- **Performance**: Good for unknown-size streams

```go
data, err := io.ReadAll(httpResponse.Body)
```

#### **4. io.ReadAtLeast()**
- **What**: Guarantees minimum bytes read
- **Returns**: `(n int, err error)`
- **Best for**: Network protocols, binary formats, when partial data is invalid
- **Memory**: Uses provided buffer
- **Performance**: Blocks until minimum is met

```go
header := make([]byte, 16)
n, err := io.ReadAtLeast(conn, header, 16)
```

#### **5. io.ReadFull()**
- **What**: Reads exactly buffer size
- **Returns**: `(n int, err error)`
- **Best for**: Fixed-size records, binary formats
- **Memory**: Uses provided buffer
- **Performance**: Fails if can't fill buffer

```go
record := make([]byte, 64)
n, err := io.ReadFull(file, record)
```

#### **6. io.Copy()**
- **What**: Copies from reader to writer
- **Returns**: `(written int64, err error)`
- **Best for**: File copying, network streaming, data transformation
- **Memory**: Streams data (no buffering)
- **Performance**: Very efficient for large data

```go
written, err := io.Copy(destFile, srcFile)
```

### ✍️ **WRITING FUNCTIONS**

#### **1. os.WriteFile()**
- **What**: Writes entire data to file
- **Returns**: `error`
- **Best for**: Small files, configuration, one-time writes
- **Memory**: Loads all data into memory
- **Performance**: Fast for small files

```go
err := os.WriteFile("output.txt", data, 0644)
```

#### **2. File.Write()**
- **What**: Writes bytes to file
- **Returns**: `(n int, err error)`
- **Best for**: Streaming, large files, custom buffer management
- **Memory**: Uses provided buffer
- **Performance**: Efficient for large files

```go
f, _ := os.Create("output.txt")
n, err := f.Write(data)
```

#### **3. File.WriteString()**
- **What**: Writes string to file
- **Returns**: `(n int, err error)`
- **Best for**: Text files, logging, string data
- **Memory**: Converts string to bytes
- **Performance**: Convenient for strings

```go
n, err := f.WriteString("Hello World\n")
```

#### **4. File.WriteAt()**
- **What**: Writes at specific position
- **Returns**: `(n int, err error)`
- **Best for**: Random access, binary formats, database-like operations
- **Memory**: Uses provided buffer
- **Performance**: Direct positioning

```go
n, err := f.WriteAt(data, 100) // Write at position 100
```

### 🎯 **SEEKING FUNCTIONS**

#### **1. SeekStart**
- **What**: Seek from beginning of file
- **Use**: `f.Seek(offset, io.SeekStart)`
- **Best for**: Absolute positioning, jumping to known locations
- **Example**: `f.Seek(100, io.SeekStart)` - go to position 100

#### **2. SeekCurrent**
- **What**: Seek relative to current position
- **Use**: `f.Seek(offset, io.SeekCurrent)`
- **Best for**: Parsing, skipping sections, relative navigation
- **Example**: `f.Seek(5, io.SeekCurrent)` - skip 5 bytes forward

#### **3. SeekEnd**
- **What**: Seek relative to end of file
- **Use**: `f.Seek(offset, io.SeekEnd)`
- **Best for**: Reading from end, finding file size, tail operations
- **Example**: `f.Seek(-10, io.SeekEnd)` - go 10 bytes from end

### 🚀 **BUFFERED OPERATIONS**

#### **1. bufio.Reader**
- **What**: Buffered reading with extra features
- **Best for**: Small frequent reads, text processing, network I/O
- **Benefits**: 10-50x faster for small reads, peek functionality
- **Memory**: Internal buffering

```go
reader := bufio.NewReader(file)
line, _ := reader.ReadLine()
peek, _ := reader.Peek(10)
```

#### **2. bufio.Writer**
- **What**: Buffered writing
- **Best for**: Small frequent writes, logging, network I/O
- **Benefits**: Reduces system calls, must call Flush()
- **Memory**: Internal buffering

```go
writer := bufio.NewWriter(file)
writer.WriteString("Hello\n")
writer.Flush() // Important!
```

#### **3. bufio.Scanner**
- **What**: Line-by-line reading
- **Best for**: Text files, log processing, configuration parsing
- **Benefits**: Handles different line endings, memory efficient
- **Memory**: Reads line by line

```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
}
```

### 🌊 **STREAMING OPERATIONS**

#### **1. io.Pipe**
- **What**: In-memory pipe between goroutines
- **Best for**: Producer-consumer patterns, data transformation
- **Benefits**: No disk I/O, goroutine communication
- **Memory**: Streams data

```go
reader, writer := io.Pipe()
go func() { writer.Write(data); writer.Close() }()
io.Copy(output, reader)
```

#### **2. io.TeeReader**
- **What**: Copies data to multiple destinations
- **Best for**: Logging, debugging, data duplication
- **Benefits**: Transparent copying
- **Memory**: Streams data

```go
tee := io.TeeReader(source, logger)
io.Copy(dest, tee)
```

#### **3. io.MultiReader**
- **What**: Combines multiple readers
- **Best for**: Concatenating files, merging streams
- **Benefits**: Seamless reading from multiple sources
- **Memory**: Streams data

```go
multi := io.MultiReader(file1, file2, file3)
io.Copy(output, multi)
```

#### **4. io.LimitReader**
- **What**: Limits reading to specified bytes
- **Best for**: Reading partial files, rate limiting
- **Benefits**: Prevents reading too much
- **Memory**: Streams data

```go
limited := io.LimitReader(file, 1024)
io.Copy(output, limited)
```

## 🎯 **DECISION MATRIX**

| Use Case | Best Function | Why |
|----------|---------------|-----|
| **Small config file** | `os.ReadFile()` | Simple, fast for small files |
| **Large file processing** | `File.Read()` + buffer | Memory efficient |
| **Network protocol** | `io.ReadAtLeast()` | Guarantees minimum data |
| **Binary file format** | `File.Read()` + `Seek()` | Random access needed |
| **Text file line-by-line** | `bufio.Scanner` | Convenient, handles line endings |
| **Log file processing** | `bufio.Scanner` | Memory efficient, line-based |
| **File copying** | `io.Copy()` | Most efficient for large files |
| **Configuration parsing** | `bufio.Scanner` | Line-based processing |
| **Network streaming** | `bufio.Reader` | Buffered for efficiency |
| **Data transformation** | `io.Pipe()` | Producer-consumer pattern |
| **Debugging/logging** | `io.TeeReader` | Transparent copying |
| **Multiple file merge** | `io.MultiReader` | Seamless concatenation |

## ⚡ **PERFORMANCE TIPS**

1. **Use buffered readers for small, frequent reads** (10-50x faster)
2. **Use `io.Copy()` for large file operations** (most efficient)
3. **Use `os.ReadFile()` for small files** (simplest)
4. **Use `File.Read()` with large buffers for large files** (memory efficient)
5. **Use `bufio.Scanner` for text processing** (convenient and efficient)
6. **Always call `Flush()` on buffered writers** (critical!)

## 🚨 **COMMON PITFALLS**

1. **Forgetting to close files** - Use `defer f.Close()`
2. **Not calling `Flush()` on buffered writers** - Data won't be written
3. **Using `os.ReadFile()` for large files** - Memory issues
4. **Not handling EOF errors** - Infinite loops possible
5. **Mixing buffered and unbuffered operations** - Unexpected behavior
6. **Not checking return values** - Silent failures

## 🔧 **REAL-WORLD EXAMPLES**

### **Log File Processor**
```go
file, _ := os.Open("app.log")
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    if strings.Contains(line, "ERROR") {
        fmt.Println("Error found:", line)
    }
}
```

### **Configuration Parser**
```go
config := make(map[string]string)
file, _ := os.Open("config.ini")
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    if strings.Contains(line, "=") {
        parts := strings.SplitN(line, "=", 2)
        config[parts[0]] = parts[1]
    }
}
```

### **Binary File Reader**
```go
file, _ := os.Open("data.bin")
header := make([]byte, 16)
io.ReadFull(file, header)
file.Seek(32, io.SeekStart) // Skip to data section
data := make([]byte, 1024)
n, _ := file.Read(data)
```

### **Network Stream Processor**
```go
conn, _ := net.Dial("tcp", "server:8080")
reader := bufio.NewReader(conn)
for {
    line, err := reader.ReadString('\n')
    if err != nil { break }
    processLine(line)
}
```

This comprehensive guide covers all the major Go file I/O functions with their benefits, use cases, and performance characteristics!
