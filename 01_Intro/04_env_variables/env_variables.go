/*An Environment Variable is a dynamic object pair on the Operating System.
These value pairs can be manipulated with the help of the operating system.
These value pairs can be used to store file path, user profile, authentication keys, execution mode, etc.
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Setenv("website", "golang.org")
	os.Setenv("DB_USERNAME", "GO")
	// os.Setenv()
	println(os.Getenv("website"))
	println(os.LookupEnv("website"))
	for _, env := range os.Environ() {
		println(env)
	}
	os.Unsetenv("website")
	println(os.LookupEnv("website"))

	dbURL := os.ExpandEnv("postgres://$DB_USERNAME:$DB_PASSWORD@DB_HOST:$DB_PORT/$DB_NAME")
	fmt.Println("DB URL: ", dbURL)
	os.Unsetenv("DB_USERNAME")

}
