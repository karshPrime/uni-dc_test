
package cmd

import (
    "fmt"
)

func padPrint[ T any ]( aWidth int, aMsg T ) {
    lMsgString := fmt.Sprintf("%v", aMsg)

    lLeftPad  := (aWidth - len( lMsgString )) / 2
    lRightPad := aWidth - len( lMsgString ) - lLeftPad

    fmt.Printf("%*s%s%*s", lLeftPad, "", lMsgString, lRightPad, "")
}

func TableHead() {
    fmt.Printf("|--------------------------------------|")
    fmt.Printf("|--------------------------------------||----------|\n")

    fmt.Printf("|");
    padPrint(38, "I N I T I A L"); fmt.Printf("||");
    padPrint(38, "U P D A T E D"); fmt.Printf("||");
    padPrint(10, "Test"); fmt.Printf("|\n");

    fmt.Printf("| Start Date |  End Date  | Days Count |")
    fmt.Printf("| Start Date |  End Date  | Days Count ||  Status  |\n")

    TableEnd()
}

func TableEnd() {
    fmt.Printf("|------------|------------|------------|")
    fmt.Printf("|------------|------------|------------||----------|\n")
}

func TableData( aValueSet ProgramIO, aTestResult TestResult ) {
    fmt.Printf( "| %v | %v |", aValueSet.IStart, aValueSet.IEnd )
    padPrint( 12, aTestResult.IResult )

    fmt.Printf( "|| %v | %v |", aValueSet.MStart, aValueSet.MEnd )
    padPrint( 12, aTestResult.MResult )

    fmt.Printf( "||" )
    padPrint( 10, aTestResult.TestCase )
    fmt.Println( "|" )
}

