```go
func main() {
    var m = make(map[string]int) // Initialize the map
    m["key"] = 10
    fmt.Println(m["key"])
    // Alternative using map literal:
    m2 := map[string]int{"key": 20}
    fmt.Println(m2["key"])
}
```