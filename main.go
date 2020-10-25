package main

import (
 "fmt"
 "go/build"
)

//Handler
type Handler interface {
 handle(message string)
}

type ConsoleHandler struct {

}

func (consoleHandler ConsoleHandler) handle(message string)  {
 fmt.Println(message)
}
//
// Formatter
type Formatter interface {
 format(message string, context build.Context) string
}

type LineFormatter struct {

}

type JsonFormatter struct {

}

func (lineFormatter LineFormatter) format(message string, context build.Context) string  {
 return "Log message formatted by amqp formatter: " + message
}

func (jsonFormatter JsonFormatter ) format(message string, context build.Context) string  {
 return "Log message formatted by json formatter : " + message
}
//
// Logger
type Logger struct {
 formatter Formatter
 handlers []Handler
}

func (logger *Logger) log(message string, context build.Context)  {
 var formattedMessage = logger.formatter.format(message, context)

 var handlers = logger.handlers

 for _,handler := range handlers {
  handler.handle(formattedMessage)
 }
}

// use case
func main() {
 //var lineFormatter LineFormatter
 var jsonFormatter JsonFormatter
 var consoleHandler = ConsoleHandler{}
 var handlers = []Handler{consoleHandler}

 var logger = Logger{
  jsonFormatter,
  handlers,
 }

 logger.log("Hello World !", build.Context{})
}