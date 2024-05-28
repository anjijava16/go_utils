In Go, `map[string]interface{}` is a map where the keys are strings and the values can be of any type. This type of map is often used to represent JSON objects or to store heterogeneous data, where the values can be of different types.

### Example Usage

Here is an example that demonstrates how to use `map[string]interface{}`:

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Creating a map with string keys and interface{} values
	data := make(map[string]interface{})

	// Adding different types of values to the map
	data["name"] = "John Doe"
	data["age"] = 30
	data["isMember"] = true
	data["address"] = map[string]interface{}{
		"street": "123 Main St",
		"city":   "Anytown",
		"state":  "CA",
	}

	// Printing the map
	fmt.Println("Map:", data)

	// Converting the map to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	// Printing the JSON representation of the map
	fmt.Println("JSON:", string(jsonData))
}
```

### Explanation

1. **Creating the Map**:
   ```go
   data := make(map[string]interface{})
   ```
   This line creates a new map where keys are strings and values can be of any type (`interface{}`).

2. **Adding Values to the Map**:
   ```go
   data["name"] = "John Doe"
   data["age"] = 30
   data["isMember"] = true
   data["address"] = map[string]interface{}{
       "street": "123 Main St",
       "city":   "Anytown",
       "state":  "CA",
   }
   ```
   Here, different types of values are added to the map. The `address` field itself is a nested map with string keys and `interface{}` values.

3. **Printing the Map**:
   ```go
   fmt.Println("Map:", data)
   ```
   This prints the map in a readable format.

4. **Marshalling the Map to JSON**:
   ```go
   jsonData, err := json.Marshal(data)
   if err != nil {
       fmt.Println("Error marshalling to JSON:", err)
       return
   }
   ```
   This converts the map to a JSON string using `json.Marshal`.

5. **Printing the JSON**:
   ```go
   fmt.Println("JSON:", string(jsonData))
   ```
   This prints the JSON representation of the map.

### Practical Use Case: Parsing JSON into `map[string]interface{}`

A common use case for `map[string]interface{}` is parsing JSON data whose structure is not known at compile time.

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// JSON data
	jsonStr := `{"name": "John Doe", "age": 30, "isMember": true, "address": {"street": "123 Main St", "city": "Anytown", "state": "CA"}}`

	// Unmarshal the JSON into a map
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Accessing values from the map
	fmt.Println("Name:", data["name"])
	fmt.Println("Age:", data["age"])
	fmt.Println("Is Member:", data["isMember"])
	fmt.Println("Address:", data["address"])

	// Type assertion to access nested map
	if address, ok := data["address"].(map[string]interface{}); ok {
		fmt.Println("Street:", address["street"])
		fmt.Println("City:", address["city"])
		fmt.Println("State:", address["state"])
	}
}
```

### Explanation

1. **JSON Data**:
   ```go
   jsonStr := `{"name": "John Doe", "age": 30, "isMember": true, "address": {"street": "123 Main St", "city": "Anytown", "state": "CA"}}`
   ```
   This is a JSON string containing various data types.

2. **Unmarshalling the JSON**:
   ```go
   var data map[string]interface{}
   if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
       fmt.Println("Error unmarshalling JSON:", err)
       return
   }
   ```
   The JSON string is unmarshalled into a `map[string]interface{}`.

3. **Accessing Values**:
   ```go
   fmt.Println("Name:", data["name"])
   fmt.Println("Age:", data["age"])
   fmt.Println("Is Member:", data["isMember"])
   fmt.Println("Address:", data["address"])
   ```
   Values are accessed directly from the map.

4. **Type Assertion**:
   ```go
   if address, ok := data["address"].(map[string]interface{}); ok {
       fmt.Println("Street:", address["street"])
       fmt.Println("City:", address["city"])
       fmt.Println("State:", address["state"])
   }
   ```
   Type assertions are used to access nested maps.

### Summary

- `map[string]interface{}` is a versatile way to store heterogeneous data.
- It is commonly used for handling JSON data where the structure is not known at compile time.
- The `interface{}` type allows storing values of any type, making it flexible but requiring type assertions when accessing the values.
