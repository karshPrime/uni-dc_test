
package cmd

import (
    "fmt"
)

func TableHead() {
}

func TableEnd() {
}

func TableData( aValueSet ProgramIO, aTestResult TestResult ) {
    fmt.Printf( "| %v | %v |", aValueSet.IStart, aValueSet.IEnd ) 
}

