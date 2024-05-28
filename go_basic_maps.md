In Go, maps are used to store key-value pairs, where each key is unique. The `map` type provides efficient lookups, additions, and deletions. Here is a simple example of how to create and use a map with string keys and string values.

### Basic Example

Let's create a map to store country codes and their corresponding country names.

```go
package main

import (
	"fmt"
)

func main() {
	// Create a map to store country codes and country names
	countries := make(map[string]string)

	// Add some entries to the map
	countries["US"] = "United States"
	countries["CA"] = "Canada"
	countries["MX"] = "Mexico"

	// Retrieve and print a value from the map
	fmt.Println("US:", countries["US"])

	// Iterate over all entries in the map
	for code, name := range countries {
		fmt.Printf("%s: %s\n", code, name)
	}

	// Check if a key exists in the map
	if name, ok := countries["CA"]; ok {
		fmt.Println("CA exists in the map with value:", name)
	} else {
		fmt.Println("CA does not exist in the map")
	}

	// Delete an entry from the map
	delete(countries, "MX")

	// Try to retrieve the deleted entry
	if _, ok := countries["MX"]; !ok {
		fmt.Println("MX has been deleted from the map")
	}
}
```

### Explanation

1. **Creating a Map**:
   ```go
   countries := make(map[string]string)
   ```
   This creates a new map called `countries` where both the keys and values are of type `string`.

2. **Adding Entries**:
   ```go
   countries["US"] = "United States"
   countries["CA"] = "Canada"
   countries["MX"] = "Mexico"
   ```
   This adds three entries to the `countries` map.

3. **Retrieving a Value**:
   ```go
   fmt.Println("US:", countries["US"])
   ```
   This prints the value associated with the key `"US"`.

4. **Iterating Over a Map**:
   ```go
   for code, name := range countries {
       fmt.Printf("%s: %s\n", code, name)
   }
   ```
   This loop iterates over all key-value pairs in the map and prints them.

5. **Checking for Existence**:
   ```go
   if name, ok := countries["CA"]; ok {
       fmt.Println("CA exists in the map with value:", name)
   } else {
       fmt.Println("CA does not exist in the map")
   }
   ```
   This checks if the key `"CA"` exists in the map and prints its value if it does.

6. **Deleting an Entry**:
   ```go
   delete(countries, "MX")
   ```
   This deletes the entry with the key `"MX"` from the map.

7. **Checking for Deletion**:
   ```go
   if _, ok := countries["MX"]; !ok {
       fmt.Println("MX has been deleted from the map")
   }
   ```
   This checks if the key `"MX"` has been deleted from the map.

### Full Example with Custom Functions

Here’s a more comprehensive example that includes functions to manipulate the map:

```go
package main

import (
	"fmt"
)

// Function to add a country to the map
func addCountry(countries map[string]string, code, name string) {
	countries[code] = name
}

// Function to delete a country from the map
func deleteCountry(countries map[string]string, code string) {
	delete(countries, code)
}

// Function to print all countries
func printCountries(countries map[string]string) {
	for code, name := range countries {
		fmt.Printf("%s: %s\n", code, name)
	}
}

func main() {
	// Create a map to store country codes and country names
	countries := make(map[string]string)

	// Add some entries to the map using a function
	addCountry(countries, "US", "United States")
	addCountry(countries, "CA", "Canada")
	addCountry(countries, "MX", "Mexico")

	// Print all countries
	fmt.Println("All countries:")
	printCountries(countries)

	// Check if a key exists in the map
	if name, ok := countries["CA"]; ok {
		fmt.Println("CA exists in the map with value:", name)
	} else {
		fmt.Println("CA does not exist in the map")
	}

	// Delete an entry from the map using a function
	deleteCountry(countries, "MX")

	// Try to retrieve the deleted entry
	if _, ok := countries["MX"]; !ok {
		fmt.Println("MX has been deleted from the map")
	}

	// Print all countries after deletion
	fmt.Println("Countries after deletion:")
	printCountries(countries)
}
```

This example demonstrates how to create, manipulate, and print a map with custom functions, providing a more modular approach to working with maps in Go.
