package main 

import "fmt"

const appName = "Vibe Queue"
var djName = "Bubbles"

func main() {
    greetUser()
}

func greetUser()(string, string){
    fmt.Printf("Welcome to %v!\nThis is DJ %v, ready to make you dance! \n", appName, djName)
    return appName, djName
}