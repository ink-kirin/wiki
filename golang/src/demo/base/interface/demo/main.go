package main

import "fmt"

type User struct {
	Name  string
	Phone int
}

func main() {
	var info = make(map[string]interface{})
	info["name"] = "simba"
	info["hobby"] = []string{"吃饭", "睡觉"}
	info["user"] = User{
		Name:  "lalala",
		Phone: 13402901129,
	}
	fmt.Println(info)

	hobby, _ := info["hobby"].([]string)
	fmt.Println(hobby[1])

	user, _ := info["user"].(User)
	fmt.Println(user.Name, user.Phone)
}
