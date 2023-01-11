package merror

import (
	"fmt"
	"log"
)

func PrintReturn(service string, reason string) {
	fmt.Println("[" + service + "] " + reason)
	return
}

func DisGoClientError(error error) {
	if error != nil {
		fmt.Println("[DiscordGo] Error loading Client", error)
		return
	}

}

func LogFatal(service string, reason string) {
	log.Fatal("[" + service + "] " + reason)
}

func CheckFatal(err error, service string, reason string) {
	if err != nil {
		LogFatal(service, reason)
		return
	}
}

func CheckPrintReturn(err error, service string, reason string, printError error) {
	if err != nil {
		fmt.Println(err, service, reason, printError)
		return
	}
}

func CheckReturn(err error) {
	if err != nil {
		return
	}
}
