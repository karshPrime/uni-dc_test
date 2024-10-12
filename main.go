
package main

import (
    "os"
    "fmt"
    "bufio"
    "github.com/karshPrime/uni-dc_test/cmd"

    "github.com/wy3148/dc/date"                 // Program Under Test
    "github.com/karshPrime/uni-dc_test/mutated" // Mutated Program Source
)

const DEFAULT = 0;               // for better clarity with function calls
var TEST_PURPOSE = [3]string {   // Purpose for each Metamorphic Test
    "Inverted Start & End",
    "Same Date, Different Years",
    "Different Date, Same Month & Year",
}

// Run Metamorphic Tests on the Program (either original code or mutated)
func metamorphicTests( aCase, aCondition int ) {
    fmt.Printf( " Metamorphic Test %v: %v\n", aCase, TEST_PURPOSE[aCase-1] )
    cmd.MetamorphicTableHead()

    /* Got identical loop bellow to not have if-else check on every loop
     * Sure, this reduces code reusability, but in this case i reckon its
     * justified.
     *
     * Could have used a functional variable for Data.Elapsed(...) and set that
     * with 'if' before loop, though that'd require making changes to wy3148/dc
     * since my mutated dc takes in an extra parameter.
     *
     * 'aCondition' is used to select specific change in dc, instead of editing
     * source code again and again for every test.
     */

    if aCondition == DEFAULT {
        // original function at wy3148/dc
        for i := 0; i < 5; i++ {
            lIResult, _ := date.Elapsed(
                cmd.TestValues[i].Start,
                cmd.TestValues[i].End )

            lMResult, _ := date.Elapsed(
                cmd.MetamorphicExpected[aCase-1][i].Start,
                cmd.MetamorphicExpected[aCase-1][i].End )

            cmd.MetamorphicTableData( (aCase-1), i, lIResult, lMResult )
        }
    } else {
        // modified
        for i := 0; i < 5; i++ {
            lIResult, _ := mutated.Elapsed( aCondition,
                cmd.TestValues[i].Start,
                cmd.TestValues[i].End )

            lMResult, _ := mutated.Elapsed( aCondition,
                cmd.MetamorphicExpected[aCase-1][i].Start,
                cmd.MetamorphicExpected[aCase-1][i].End )

            cmd.MetamorphicTableData( (aCase-1), i, lIResult, lMResult )
        }
    }

    cmd.MetamorphicTableEnd()
    fmt.Println("")
}

// Compare results of Mutated Code against Original
func compareOriginal( aCondition int ) {
    fmt.Println( " Comparing Mutated Code Against Original Code" )
    cmd.ComparisonTableHead()

    for i := 0; i < 5; i++ {
        lIResult, _ := date.Elapsed(
            cmd.TestValues[i].Start,
            cmd.TestValues[i].End )

        lMResult, _ := mutated.Elapsed( aCondition,
            cmd.TestValues[i].Start,
            cmd.TestValues[i].End )

        cmd.ComparisonTableData( i, lIResult, lMResult )
    }

    cmd.ComparisonTableEnd()
}


func main() {
    // Original Code Runs
    // metamorphicTests(1, DEFAULT)
    // metamorphicTests(2, DEFAULT)
    // metamorphicTests(3, DEFAULT)

    lScanner := bufio.NewScanner(os.Stdin)

    // Mutation Test
    for i := 0; i < 30; i++ {
        fmt.Print("\033[H\033[2J") // clear screen [macos/linux]

        fmt.Printf("\n|==[ Checking Mutant %02d ", i+1 )
        cmd.PrintRepeat( "]", "=", "|", 66, true )

        // Running Metamorphic Tests
        metamorphicTests(1, i)
        metamorphicTests(2, i)
        metamorphicTests(3, i)

        // Comparing Against OG
        compareOriginal( i )

        // pause for [Y] before continuing [N] quits
        fmt.Print("\nContinue? (Y/n): ")
        lScanner.Scan()
        input := lScanner.Text()
        if input == "n" { break }
    }
}

