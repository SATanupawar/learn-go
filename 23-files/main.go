package main

import (

	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	// 	f , err := os.Open("example.txt")
	// 	if err != nil {

	// 		panic(err)
	// 	}

	// fileInfo , err := f.Stat()

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println("File name:", fileInfo.Name())
	// 	fmt.Println("File size:", fileInfo.Size())
	// 	fmt.Println("File mode:", fileInfo.Mode())
	// 	fmt.Println("File is directory:", fileInfo.IsDir())

	// data , err := os.ReadFile("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// dir , err := os.Open(".")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo , err := dir.ReadDir(20)

	// for _, file := range fileInfo {
	// 	fmt.Println(file.Name())
	// }

	// f , err := os.Create("example.txt2")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// // f.WriteString("hello go !")
	// // f.WriteString("hello go lang  !")

	// bytes := []byte("hello go lang ")
	// f.Write(bytes)

	// sourceFile, err := os.Create("source.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destinationFile, err := os.Create("destination.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destinationFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destinationFile)

	// for {
	// 	b, err := reader.ReadBytes('\n')
	// 	if err != nil {
	// 		if err.Error() == "EOF" {
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	writer.Write(b)

	// }
	// writer.Flush()


	err := os.Remove("destination.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File removed successfully" )

}
