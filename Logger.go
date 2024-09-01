package logger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"sync"
)

// Use this in the init() function to initialize the size of the buffered channel
const logbuffersize int32 = 200

var DEBUG bool = false

// Use _DEBUG prints to strip them out of release builds
type LGRImpl interface {
	init()                                          // Private, call it once to initialize inner members
	StartLogger()                                   // !only call it once! starts an infinite loop to write out messages from a channel
	Write(message string)                           // Writes: 'date : message "\n"' (3 messages in channel)
	Write_Request(message string, request string)   // Writes: 'request : date : message "\n"' !(1 message in channel)!
	Write_DEBUG(message string)                     // Writes: 'date : message "\n"' (3 messages in channel)
	PrintSection(message string)                    // Writes: 'message : ' (takes 2 messages in channel)
	PrintSection_DEBUG(message string)              // Writes: 'message : ' (takes 2 messages in channel)
	PrintDate()                                     // Writes: 'date : ' (takes 1 message(s) in channel)
	PrintDate_DEBUG()                               // Writes: 'date : ' (takes 1 message(s) in channel)
	WriteErr(error) int                             // Writes: 'date : error.Error()"\n"' and returns 1 OR Writes: nothing and returns 0 (Uses Write(message) internally)
	WriteErr_Request(err error, request string) int // Writes: 'request : date : error.Error()"\n"' and returns 1 OR Writes: nothing and returns 0 !(1 message in channel)!
	WriteErr_DEBUG(err error) (errnum int)          // Writes: 'date: error.Error()"\n"' and returns 1 OR Writes: nothing and returns 0 (Uses Write(message) internally)
	WriteNoNewLine(message string)                  // Writes: 'date : message'  (takes 2 messages in channel)
	WriteNoNewLine_DEBUG(message string)            // Writes: 'date : message'  (takes 2 messages in channel)
	WriteWithoutDate(message string)                // Writes: 'message' (takes 1 message in channel)
	WriteWithoutDate_DEBUG(message string)          // Writes: 'message' (takes 1 message in channel)
}

// debug prints && print req ids

var (
	loggerInstance LGRImpl
	loggeronce     sync.Once
	loggerlogonce  sync.Once
)

func Create(instance LGRImpl) {
	loggeronce.Do(func() {
		loggerInstance = instance
		loggerInstance.init()
	})
}

func Logger() LGRImpl {
	return loggerInstance
}

func PrintJson[T any](entity *T) string {
	typename := reflect.TypeFor[T]().Name()
	outputStringJson, err := json.MarshalIndent((*entity), "", "     ")
	if err != nil {
		return "Error parsing json data"
	} else {
		return typename + ":\n" + string(outputStringJson) + "\n"
	}
}

func ExampleLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger().PrintDate()
		Logger().WriteWithoutDate(fmt.Sprintf("\tRequest from: %s to: Host: %s URL: %s\n", r.RemoteAddr, r.Host, r.URL))
		Logger().WriteWithoutDate(fmt.Sprintf("\tWith HEADERS: %s\n", r.Header))
		Logger().WriteWithoutDate(fmt.Sprintf("\tWith BODY: %s\n", r.Body))
		next.ServeHTTP(w, r)
		// Here we can log the full response

		Logger().Write("Response done!")
	})
}

// Ensure all methods from LGRImpl are implemented ccompile time
var (
	_ LGRImpl = (*NullLoggerimpl)(nil)
	_ LGRImpl = (*ConsoleLoggerimpl)(nil)
	_ LGRImpl = (*FileLoggerimpl)(nil)
)
