package logger

import (
	"fmt"
	"time"
)

type ConsoleLoggerimpl struct {
	messages chan string
}

func (lgr *ConsoleLoggerimpl) init() {
	lgr.messages = make(chan string, logbuffersize)
}

func (logger *ConsoleLoggerimpl) WriteErr_Request(err error, request string) (errnum int) {
	if err != nil {
		logger.messages <- request + " : " + time.Now().Format(time.UnixDate) + " : Error: " + err.Error() + "\n"
		errnum = 1
	}
	return errnum
}

func (logger *ConsoleLoggerimpl) Write_Request(message string, request string) {
	logger.messages <- request + " : " + time.Now().Format(time.UnixDate) + " : " + message + "\n"
}

func (logger *ConsoleLoggerimpl) PrintSection(message string) {
	logger.messages <- message
	logger.messages <- " : "
}

func (logger *ConsoleLoggerimpl) PrintSection_DEBUG(message string) {
	if DEBUG {
		logger.PrintSection(message)
	}
}

func (logger *ConsoleLoggerimpl) PrintDate() {
	logger.messages <- time.Now().Format(time.UnixDate) + " : "
}

func (logger *ConsoleLoggerimpl) PrintDate_DEBUG() {
	if DEBUG {
		logger.PrintDate()
	}
}

func (logger *ConsoleLoggerimpl) WriteErr(err error) (errnum int) {
	if err != nil {
		logger.Write("Error: " + err.Error())
		errnum = 1
	}
	return errnum
}

func (logger *ConsoleLoggerimpl) WriteErr_DEBUG(err error) (errnum int) {
	if err != nil {
		logger.Write_DEBUG("Error: " + err.Error())
		errnum = 1
	}
	return errnum
}

func (logger *ConsoleLoggerimpl) Write(message string) {
	logger.PrintDate()
	logger.messages <- message
	logger.messages <- "\n"
}

func (logger *ConsoleLoggerimpl) Write_DEBUG(message string) {
	if DEBUG {
		logger.Write(message)
	}
}

func (logger *ConsoleLoggerimpl) WriteNoNewLine(message string) {
	logger.PrintDate()
	logger.messages <- message
}

func (logger *ConsoleLoggerimpl) WriteNoNewLine_DEBUG(message string) {
	if DEBUG {
		logger.WriteNoNewLine(message)
	}
}

func (logger *ConsoleLoggerimpl) WriteWithoutDate(message string) {
	logger.messages <- message
}

func (logger *ConsoleLoggerimpl) WriteWithoutDate_DEBUG(message string) {
	if DEBUG {
		logger.WriteWithoutDate(message)
	}
}

func (logger *ConsoleLoggerimpl) StartLogger() {
	fmt.Println("Starting Logger")
	loggerlogonce.Do(func() {
		for msg := range logger.messages {
			fmt.Print(msg)
		}
	})
}
