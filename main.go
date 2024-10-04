package main

import (
	
	"fmt"
	"os"
	"strings"
	"asciiartcolor/functions"
)

func main() {
	args := os.Args[1:]


if len(args)==1{
	
	input:=args[0]
	banner:="standard"
	result:=asciiartcolor.AsciiNormal(input,banner)
	fmt.Print(result)
	return
}else if len(args)==2{
    
if strings.HasPrefix(args[0], "--color="){

	result := asciiartcolor.AsciiColor(args)
	fmt.Print(result)
    return

}else if strings.HasPrefix(args[0], "--output=") {

	asciiartcolor.AsciiOutput(args)

}else if args[1]=="standard"||args[1]=="thinkertoy"||args[1]=="shadow"{
	input:=args[0]
    banner:=args[1]
    result:=asciiartcolor.AsciiFs(input,banner)
    fmt.Print(result)
    return
}else {

	fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something' ")

    return
}


}else if len(args)==3{

	if strings.HasPrefix(args[0], "--color="){
		result := asciiartcolor.AsciiColor(args)
		fmt.Print(result)
		return
	}else if strings.HasPrefix(args[0], "--output=") {

		asciiartcolor.AsciiOutput(args)
	
	}else {
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something' ")
		return
	}

}else if len(args)==4{
	if strings.HasPrefix(args[0], "--color="){
		result := asciiartcolor.AsciiColor(args)
		fmt.Print(result)
		return
	}else{
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something' ")
		return
	}
}else {
	fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <substring to be colored> 'something'")

		return
	}
}
