// Interface Segregation Principle (ISP)

// The Interface Segregation Principle states that clients 
// should not be forced to depend on interfaces they do not use. 
// This principle encourages creating smaller, 
// more focused interfaces rather than large, monolithic ones.





package main

type Reader interface {
	 Read() string
}

type Writer interface {
	 Write(content string)
}

type Printer interface {
	 Print() string
}

type TextDocument struct {
	 content string
}

// Implement Reader, Writer, and Printer for TextDocument

type ReadOnlyDocument struct {
	 content string
}



func main(){}
