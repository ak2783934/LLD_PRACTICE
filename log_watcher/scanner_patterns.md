# bufio.Scanner Real-World Code Examples

## 🎯 **10 Common Scanner Patterns**

### **1. Basic Line Reading**
```go
func basicLineReading() {
    file, err := os.Open("file.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineNumber := 0

    for scanner.Scan() {
        lineNumber++
        line := scanner.Text()
        fmt.Printf("Line %d: %s\n", lineNumber, line)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}
```

### **2. Log File Processing**
```go
func logFileProcessing() {
    file, _ := os.Open("app.log")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    errorCount := 0
    var errors []string

    for scanner.Scan() {
        line := scanner.Text()
        if strings.Contains(line, "ERROR") {
            errorCount++
            errors = append(errors, line)
            fmt.Printf("🚨 ERROR: %s\n", line)
        }
    }

    fmt.Printf("Total errors: %d\n", errorCount)
}
```

### **3. Configuration File Parsing**
```go
func configFileParsing() {
    file, _ := os.Open("config.ini")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    config := make(map[string]string)

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        
        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }
        
        // Parse key=value pairs
        if strings.Contains(line, "=") {
            parts := strings.SplitN(line, "=", 2)
            key := strings.TrimSpace(parts[0])
            value := strings.TrimSpace(parts[1])
            config[key] = value
        }
    }

    // Use config map
    for key, value := range config {
        fmt.Printf("%s = %s\n", key, value)
    }
}
```

### **4. CSV Data Processing**
```go
func csvDataProcessing() {
    file, _ := os.Open("data.csv")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var headers []string
    var records []map[string]string
    lineNumber := 0

    for scanner.Scan() {
        lineNumber++
        line := scanner.Text()
        
        if lineNumber == 1 {
            // Header row
            headers = strings.Split(line, ",")
        } else {
            // Data row
            fields := strings.Split(line, ",")
            record := make(map[string]string)
            for i, field := range fields {
                if i < len(headers) {
                    record[headers[i]] = strings.TrimSpace(field)
                }
            }
            records = append(records, record)
        }
    }

    // Process records
    for _, record := range records {
        fmt.Printf("Record: %+v\n", record)
    }
}
```

### **5. JSON-like Data Processing**
```go
func jsonLikeDataProcessing() {
    file, _ := os.Open("data.json")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var users []map[string]string
    currentUser := make(map[string]string)
    inUser := false

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        
        // Skip empty lines and braces
        if line == "" || line == "{" || line == "}" {
            continue
        }
        
        // Check for user object start
        if strings.Contains(line, "\"id\":") {
            if len(currentUser) > 0 {
                users = append(users, currentUser)
            }
            currentUser = make(map[string]string)
            inUser = true
        }
        
        // Parse key-value pairs
        if inUser && strings.Contains(line, ":") {
            line = strings.Trim(line, " \t\n\r,")
            if strings.Contains(line, ":") {
                parts := strings.SplitN(line, ":", 2)
                key := strings.Trim(strings.Trim(parts[0], " \""), " ")
                value := strings.Trim(strings.Trim(parts[1], " \""), " ")
                currentUser[key] = value
            }
        }
    }
    
    // Add last user
    if len(currentUser) > 0 {
        users = append(users, currentUser)
    }
}
```

### **6. Network Log Analysis**
```go
func networkLogAnalysis() {
    file, _ := os.Open("access.log")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    statusCodes := make(map[string]int)
    ipAddresses := make(map[string]int)
    totalRequests := 0

    for scanner.Scan() {
        line := scanner.Text()
        
        // Parse Apache/nginx log format
        parts := strings.Split(line, " ")
        if len(parts) >= 7 {
            ip := parts[0]
            method := strings.Trim(parts[5], "\"")
            path := parts[6]
            status := parts[8]
            
            totalRequests++
            statusCodes[status]++
            ipAddresses[ip]++
            
            fmt.Printf("Request: %s %s %s (Status: %s)\n", ip, method, path, status)
        }
    }

    // Analysis
    fmt.Printf("Total requests: %d\n", totalRequests)
    for status, count := range statusCodes {
        fmt.Printf("Status %s: %d requests\n", status, count)
    }
}
```

### **7. System Monitoring**
```go
func systemMonitoring() {
    file, _ := os.Open("system.log")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var cpuReadings []float64
    var memoryReadings []float64

    for scanner.Scan() {
        line := scanner.Text()
        
        // Parse system metrics
        if strings.Contains(line, "CPU:") && strings.Contains(line, "Memory:") {
            // Extract CPU percentage
            cpuStart := strings.Index(line, "CPU: ") + 5
            cpuEnd := strings.Index(line[cpuStart:], "%")
            if cpuEnd > 0 {
                cpuStr := line[cpuStart : cpuStart+cpuEnd]
                if cpu, err := strconv.ParseFloat(cpuStr, 64); err == nil {
                    cpuReadings = append(cpuReadings, cpu)
                }
            }
            
            // Extract memory usage
            memStart := strings.Index(line, "Memory: ") + 8
            memEnd := strings.Index(line[memStart:], "GB")
            if memEnd > 0 {
                memStr := line[memStart : memStart+memEnd]
                if mem, err := strconv.ParseFloat(memStr, 64); err == nil {
                    memoryReadings = append(memoryReadings, mem)
                }
            }
        }
    }

    // Calculate statistics
    if len(cpuReadings) > 0 {
        cpuSum := 0.0
        for _, cpu := range cpuReadings {
            cpuSum += cpu
        }
        avgCPU := cpuSum / float64(len(cpuReadings))
        fmt.Printf("Average CPU: %.2f%%\n", avgCPU)
    }
}
```

### **8. Data Validation**
```go
func dataValidation() {
    file, _ := os.Open("users.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    validRecords := 0
    invalidRecords := 0
    var validationErrors []string
    lineNumber := 0

    for scanner.Scan() {
        lineNumber++
        line := scanner.Text()
        
        if line == "" {
            continue
        }
        
        // Parse user record: username,email,age,status
        fields := strings.Split(line, ",")
        if len(fields) != 4 {
            invalidRecords++
            validationErrors = append(validationErrors, fmt.Sprintf("Line %d: Invalid field count", lineNumber))
            continue
        }
        
        username := strings.TrimSpace(fields[0])
        email := strings.TrimSpace(fields[1])
        ageStr := strings.TrimSpace(fields[2])
        status := strings.TrimSpace(fields[3])
        
        // Validate each field
        if username == "" {
            invalidRecords++
            validationErrors = append(validationErrors, fmt.Sprintf("Line %d: Empty username", lineNumber))
            continue
        }
        
        if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
            invalidRecords++
            validationErrors = append(validationErrors, fmt.Sprintf("Line %d: Invalid email: %s", lineNumber, email))
            continue
        }
        
        if age, err := strconv.Atoi(ageStr); err != nil || age < 0 || age > 150 {
            invalidRecords++
            validationErrors = append(validationErrors, fmt.Sprintf("Line %d: Invalid age: %s", lineNumber, ageStr))
            continue
        }
        
        if status != "active" && status != "inactive" {
            invalidRecords++
            validationErrors = append(validationErrors, fmt.Sprintf("Line %d: Invalid status: %s", lineNumber, status))
            continue
        }
        
        validRecords++
        fmt.Printf("✓ Valid: %s, %s, %s, %s\n", username, email, ageStr, status)
    }

    fmt.Printf("Valid: %d, Invalid: %d\n", validRecords, invalidRecords)
}
```

### **9. Performance Monitoring**
```go
func performanceMonitoring() {
    file, _ := os.Open("data.txt")
    defer file.Close()

    scanner := bufio.NewScanner(file)
    startTime := time.Now()
    lineCount := 0
    totalChars := 0

    for scanner.Scan() {
        lineCount++
        line := scanner.Text()
        totalChars += len(line)
        
        // Simulate processing
        _ = strings.ToUpper(line)
    }
    
    elapsed := time.Since(startTime)
    
    fmt.Printf("Processing time: %v\n", elapsed)
    fmt.Printf("Lines processed: %d\n", lineCount)
    fmt.Printf("Lines per second: %.2f\n", float64(lineCount)/elapsed.Seconds())
    fmt.Printf("Characters per second: %.2f\n", float64(totalChars)/elapsed.Seconds())
}
```

### **10. Custom Split Functions**
```go
func customSplitFunctions() {
    file, _ := os.Open("config.ini")
    defer file.Close()

    // Split by '=' character
    scanner := bufio.NewScanner(file)
    scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
        if atEOF && len(data) == 0 {
            return 0, nil, nil
        }
        if i := strings.Index(string(data), "="); i >= 0 {
            return i + 1, data[0:i], nil
        }
        if atEOF {
            return len(data), data, nil
        }
        return 0, nil, nil
    })
    
    for scanner.Scan() {
        token := scanner.Text()
        if token != "" {
            fmt.Printf("Token: %q\n", token)
        }
    }
}
```

## 🎯 **Key Scanner Patterns Summary**

### **Common Patterns:**
1. **Line-by-line processing** - Most common use case
2. **Log analysis** - Error detection, filtering, statistics
3. **Configuration parsing** - Key-value pairs, comments
4. **CSV processing** - Record-based data handling
5. **JSON-like parsing** - Structured data extraction
6. **Network log analysis** - Request/response processing
7. **System monitoring** - Metrics extraction and analysis
8. **Data validation** - Field validation and error reporting
9. **Performance monitoring** - Timing and efficiency measurement
10. **Custom splitting** - Non-standard data formats

### **Essential Scanner Techniques:**
- **Error handling**: Always check `scanner.Err()`
- **Line counting**: Track line numbers for error reporting
- **Field parsing**: Use `strings.Split()` for delimited data
- **Validation**: Check field formats and ranges
- **Statistics**: Collect counts, sums, averages
- **Custom splitting**: Use `scanner.Split()` for special formats
- **Performance**: Monitor processing speed and efficiency

### **When to Use Each Pattern:**
- **Basic reading**: Simple text files, logs
- **Log processing**: Error analysis, monitoring
- **Config parsing**: Application settings, environment files
- **CSV processing**: Data import/export, analytics
- **JSON parsing**: API responses, configuration files
- **Network logs**: Security analysis, traffic monitoring
- **System monitoring**: Performance tracking, alerting
- **Data validation**: Input validation, data quality
- **Performance monitoring**: Optimization, benchmarking
- **Custom splitting**: Specialized formats, protocols

These patterns cover 90% of real-world scanner use cases! 🚀
