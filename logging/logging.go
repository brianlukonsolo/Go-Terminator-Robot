package logging

import "fmt"

func LOG(moduleName string, message string){
	fmt.Print("[" + moduleName + "]: " + message + "\n")
}
