package main 

import "fmt"

const appName = "Vibe Queue"
var djName = "Bubbles"

func main() {
    greetUser()
    
    for { 
        userName, songName, artistName := requestSong()
        fmt.Printf("Your Song has Been queued!,\n")
        isValidUser, isValidSong, isValidArtist := validateSongRequest(userName, songName, artistName)
            if !isValidUser{
                fmt.Printf("Please enter a valid user name...\n")
                break  
            }
            if !isValidSong{
                fmt.Printf("Please enter a valid song name...\n")
                break
            }
            if !isValidArtist{
                fmt.Printf("Please enter a valid artist name...\n")
                break
            }
        }
}

func greetUser()(){
    fmt.Printf("Welcome to %v!\nThis is DJ %v, ready to make you dance! \n", appName, djName)
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

func validateSongRequest(userName string, songName string, artistName string)(bool, bool, bool){
    isValidUser := len(userName) >= 2
	isValidSong := len(songName) >= 2
	isValidArtist := len(artistName) >= 2
    return isValidUser, isValidSong, isValidArtist

}