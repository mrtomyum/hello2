package main
import "fmt"
type User struct {
    FName string
    LName string
}

func main() {
    // var u User //จองหน่วยความจำแล้ว หรือ = new Instance Object!
    // var u2 *User
    // fmt.Println("u:", u)
    // u.FName = "Kasem"
    // fmt.Println("FName", u.FName)
    // u.LName = "Arnontavilas"
    // fmt.Println("LName:", u.LName)

    //จองหน่วยความจำแล้ว หรือ = new Instance Object!
    u := User{
        FName: "Kasem",
        LName: "Arnontavilas"
    }
    var u2 *User
    fmt.Println("u:", u)

    u2 = &u
    fmt.Println("u2:", u2)
    u.FName = "Kittithat"
    fmt.Println(u2)
    fmt.Println("LName:", u2.FName)

    
}
