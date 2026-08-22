package loadbench

import (
	"baguette/tools/pathResolver"
	"fmt"
	"log"
	"runtime"
	"strings"
)

func InitializeConfig() {

	var target string
	var clientType writer.ClientType
	var socketPath string
	var isSocketPathGood bool
	var errResolvePath error

	target = runtime.GOOS

	fmt.Printf("You're on %s \n\n", target)

	providers := writer.GetAllProviders()
	var cliProviderMessage strings.Builder

	for i, provider := range providers {
		cliProviderMessage.WriteString(fmt.Sprintf(" %d: %s \n", i, provider))
	}
	fmt.Println("Now, What is the client you are using?")
	fmt.Printf(cliProviderMessage.String())
	fmt.Printf("> ")
	if _, errClientType := fmt.Scan(&clientType); errClientType != nil {
		log.Fatal(errClientType)
	}

	socketPath, errResolvePath = pathResolver.GetPathFromTargetAndClientType(target, clientType)

	if errResolvePath != nil {
		log.Fatal(errResolvePath)
	}

	fmt.Printf("Great ! You're using %s \n", clientType.String())
	fmt.Printf("Now, Is that path your socket engine OS path ? %s \n", socketPath)
	fmt.Println("Yes: 1, No: 0 ")
	fmt.Printf("> ")

	if _, errSocketPathCheck := fmt.Scan(&isSocketPathGood); errSocketPathCheck != nil {
		panic("Something went wrong")
		return
	}

	if !isSocketPathGood {
		fmt.Println("Ok so, share me the good path")
		fmt.Printf("> ")

		if _, errSocketPath := fmt.Scan(&socketPath); errSocketPath != nil {
			panic("Something went wrong")
			return
		}
	}

	fmt.Printf("You're using %s \n", socketPath)
	fmt.Println("Well done ! You can now use the config")

}
