
package cmd

import (
    "fmt"
)

func TableHead() {
    fmt.Printf("|--------------------------------------|")
    fmt.Printf("|--------------------------------------||----------|\n")

    TableEnd()
}

func TableEnd() {
    fmt.Printf("|------------|------------|------------|")
    fmt.Printf("|------------|------------|------------||----------|\n")
}

func TableData( aValueSet ProgramIO, aTestResult TestResult ) {
    fmt.Printf( "| %v | %v |", aValueSet.IStart, aValueSet.IEnd ) 
}

