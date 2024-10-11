
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


//- Metamorphic ----------------------------------------------------------------

func MetamorphicTableHead() {
    fmt.Printf("|--------------------------------------|")
    fmt.Println("|--------------------------------------||----------|")

    fmt.Printf("|");
    padPrint(38, "I N I T I A L"); fmt.Printf("||");
    padPrint(38, "U P D A T E D"); fmt.Printf("||");
    padPrint(10, "Test"); fmt.Printf("|\n");

    fmt.Printf("| Start Date |  End Date  | Days Count |")
    fmt.Printf("| Start Date |  End Date  | Days Count ||  Status  |\n")

    MetamorphicTableEnd()
}

func MetamorphicTableEnd() {
    fmt.Printf("|------------|------------|------------|")
    fmt.Println("|------------|------------|------------||----------|")
}

func MetamorphicTableData( aCase, aIndex, aIResult, aMResult int ) {
    lExpected := TestValues[aIndex]
    lTesting  := MetamorphicExpected[aCase][aIndex]

    fmt.Printf( "| %v | %v |", lExpected.Start, lExpected.End)
    padPrint( 12, aIResult)

    fmt.Printf( "|| %v | %v |", lTesting.Start, lTesting.End )
    padPrint( 12, aMResult )

    fmt.Printf( "||" )
    padPrint( 10, ( (aIResult + lTesting.Difference) == aMResult ) )
    fmt.Println( "|" )
}

//- Comparison -----------------------------------------------------------------

func ComparisonTableHead() {
    fmt.Printf("|-------------------------||-------------------------|")
    fmt.Println("|------------|")

    fmt.Printf("|");
    padPrint(25, "Input Dates"); fmt.Printf("||");
    padPrint(25, "Counted Days"); fmt.Printf("||");
    padPrint(12, "Test"); fmt.Printf("|\n");

    fmt.Printf("|    Start   |    End     ||  Original  |  Modified  |")
    fmt.Println("|   Result   |")

    ComparisonTableEnd() 
}

func ComparisonTableEnd() {
    fmt.Printf("|------------|------------||------------|------------|")
    fmt.Println("|------------|")
}

func ComparisonTableData( aIndex , aIResult, aMResult int ) {
    lValues := TestValues[aIndex]

    fmt.Printf( "| %v | %v ||", lValues.Start, lValues.End )
    padPrint( 12, aIResult )

    fmt.Printf( "|" )
    padPrint( 12, aIResult )

    fmt.Printf( "||" )
    padPrint( 12, ( (aIResult + lValues.Difference) == aMResult ) )
    fmt.Println( "|" )
}

