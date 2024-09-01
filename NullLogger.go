package logger

type NullLoggerimpl struct{}

func (lgr *NullLoggerimpl) init() {}

func (logger *NullLoggerimpl) WriteErr_Request(err error, request string) (errnum int) {
	if err != nil {
		errnum = 1
	}
	return errnum
}

func (logger *NullLoggerimpl) Write_Request(message string, request string) {}

func (logger *NullLoggerimpl) PrintSection(string) {}

func (logger *NullLoggerimpl) PrintSection_DEBUG(string) {}

func (logger *NullLoggerimpl) PrintDate() {
}

func (logger *NullLoggerimpl) PrintDate_DEBUG() {
}

func (logger *NullLoggerimpl) WriteErr(err error) (errnum int) {
	if err != nil {
		errnum = 1
	}
	return errnum
}

func (logger *NullLoggerimpl) WriteErr_DEBUG(err error) (errnum int) {
	if err != nil {
		errnum = 1
	}
	return errnum
}

func (logger *NullLoggerimpl) Write(message string) {
}

func (logger *NullLoggerimpl) Write_DEBUG(message string) {
}

func (logger *NullLoggerimpl) WriteNoNewLine(message string) {
}

func (logger *NullLoggerimpl) WriteNoNewLine_DEBUG(message string) {
}

func (logger *NullLoggerimpl) WriteWithoutDate(message string) {
}

func (logger *NullLoggerimpl) WriteWithoutDate_DEBUG(message string) {
}

func (logger *NullLoggerimpl) StartLogger() {
}
