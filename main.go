package main 

import "fmt"

const appName = "Vibe Queue"
var djName = "Bubbles"

func main() {
    greetUser()
    requestSong()
}

func greetUser()(string, string){
    fmt.Printf("Welcome to %v!\nThis is DJ %v, ready to make you dance! \n", appName, djName)
    return appName, djName
}

func requestSong()(string, string, string){
    var userName string
    var songName string
    var artistName string

    fmt.Printf("Please enter your name: ")
    fmt.Scanln(&userName)

    fmt.Printf("Please enter the name of the song: ")
    fmt.Scanln(&songName)

    fmt.Printf("Please enter the artist of the song: ")
    fmt.Scanln(&artistName)

    fmt.Printf("Hi %v, you have requested for %v by %v\n", userName, songName, artistName)
    return userName, songName, artistName

}