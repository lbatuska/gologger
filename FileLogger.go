package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type FileLoggerimpl struct {
	messages chan string
	mutex    *sync.Mutex
	logFile  *os.File
}

var filepath string = "./log"

func (lgr *FileLoggerimpl) init() {
	envfp, envexist := os.LookupEnv("LOGFILE_GO_LOGGER")
	if envexist {
		if len(envfp) > 0 {
			filepath = envfp
		} else {
			fmt.Printf("LOGFILE_GO_LOGGER env exist but has an empty value using default value: %s !\n", filepath)
		}
	} else {
		fmt.Printf("LOGFILE_GO_LOGGER env doesn't exist using default value: %s !\n", filepath)
	}
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		fmt.Printf("Error opening or creating file: %s\n", err.Error())
		panic(fmt.Sprintf("Error opening or creating file: %s", err.Error()))
	}
	lgr.messages = make(chan string, logbuffersize)
	lgr.logFile = f
	lgr.mutex = &sync.Mutex{}
}

func (logger *FileLoggerimpl) WriteErr_Request(err error, request string) (errnum int) {
	if err != nil {
		logger.messages <- request + " : " + time.Now().Format(time.UnixDate) + " : Error: " + err.Error() + "\n"
		errnum = 1
	}
	return errnum
}

func (logger *FileLoggerimpl) Write_Request(message string, request string) {
	logger.messages <- request + " : " + time.Now().Format(time.UnixDate) + " : " + message + "\n"
}

func (logger *FileLoggerimpl) PrintSection(message string) {
	logger.messages <- message
	logger.messages <- " : "
}

func (logger *FileLoggerimpl) PrintSection_DEBUG(message string) {
	if DEBUG {
		logger.PrintSection(message)
	}
}

func (logger *FileLoggerimpl) PrintDate() {
	logger.messages <- time.Now().Format(time.UnixDate) + " : "
}

func (logger *FileLoggerimpl) PrintDate_DEBUG() {
	if DEBUG {
		logger.PrintDate()
	}
}

func (logger *FileLoggerimpl) WriteErr(err error) (errnum int) {
	if err != nil {
		logger.Write("Error: " + err.Error())
		errnum = 1
	}
	return errnum
}

func (logger *FileLoggerimpl) WriteErr_DEBUG(err error) (errnum int) {
	if err != nil {
		logger.Write_DEBUG("Error: " + err.Error())
		errnum = 1
	}
	return errnum
}

func (logger *FileLoggerimpl) Write(message string) {
	logger.PrintDate()
	logger.messages <- message
	logger.messages <- "\n"
}

func (logger *FileLoggerimpl) Write_DEBUG(message string) {
	if DEBUG {
		logger.Write(message)
	}
}

func (logger *FileLoggerimpl) WriteNoNewLine(message string) {
	logger.PrintDate()
	logger.messages <- message
}

func (logger *FileLoggerimpl) WriteNoNewLine_DEBUG(message string) {
	if DEBUG {
		logger.WriteNoNewLine(message)
	}
}

func (logger *FileLoggerimpl) WriteWithoutDate(message string) {
	logger.messages <- message
}

func (logger *FileLoggerimpl) WriteWithoutDate_DEBUG(message string) {
	if DEBUG {
		logger.WriteWithoutDate(message)
	}
}

func (logger *FileLoggerimpl) StartLogger() {
	fmt.Println("Starting FileLogger")
	loggerlogonce.Do(func() {
		for msg := range logger.messages {

			logger.mutex.Lock()
			_, err := logger.logFile.WriteString(msg)
			if err != nil {
				fmt.Println(err.Error())
				logger.logFile.Close()
				logger.mutex.Unlock()
				panic("Failed to write to file")
			}
			err = logger.logFile.Sync()
			if err != nil {
				logger.logFile.Close()
				logger.mutex.Unlock()

				panic("Failed to write to file")
			}
			logger.mutex.Unlock()
		}
	})
	// Technically we should do this but this will never run
	// logger.mutex.Lock()
	// logger.logFile.Close()
	// logger.mutex.Unlock()
}
