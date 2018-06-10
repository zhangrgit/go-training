package go_training

import "log"

func Println(err error)  {
	if err!=nil {
		println(err)
	}
}

func LogPrintln(err error)  {
	if err!=nil {
		log.Println(err)
		return
	}
}