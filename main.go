
package main

import (
    "fmt"
    "github.com/karshPrime/uni-dc_test/cmd"

    "github.com/wy3148/dc/date" // Program Under Test
    "github.com/karshPrime/uni-dc_test/mutated" // Mutated Program Source
)

const DEFAULT = 0;
var TEST_PURPOSE = [3]string {
    "Inverted Start & End",
    "Same Date, Different Years",
    "Different Date, Same Month & Year",
}

func metamorphicTests( aCase, aCondition int ) {
    fmt.Printf( "Metamorphic Test %v: %v\n", aCase, TEST_PURPOSE[aCase-1] )
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

func main() {
    // Original Code Runs
    metamorphicTests(1, DEFAULT)
    metamorphicTests(2, DEFAULT)
    metamorphicTests(3, DEFAULT)

    // Mutation Test
    for i := 0; i < 30; i++ {
        fmt.Printf("\n|==[ Checking Mutant %02d ", i+1 )
        cmd.PrintRepeat( "]", "=", "|", 66, true )

        metamorphicTests(1, i)
        metamorphicTests(2, i)
        metamorphicTests(3, i)
    }
}

