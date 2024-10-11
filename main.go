
package main

import (
    "fmt"
    "github.com/karshPrime/uni-dc_test/cmd"

    "github.com/wy3148/dc/date" // Program Under Test
)

func metamorphicTests( aCase int ) {
    cmd.TableHead()

    for i := 0; i < 5; i++ {
        lIResult, _ := date.Elapsed(
            cmd.MetamorphicValues[aCase-1][i].IStart,
            cmd.MetamorphicValues[aCase-1][i].IEnd )

        lMResult, _ := date.Elapsed(
            cmd.MetamorphicValues[aCase-1][i].MStart,
            cmd.MetamorphicValues[aCase-1][i].MEnd )


        lTestResult := cmd.TestResult {
            IResult: lIResult,
            MResult: lMResult,

            TestCase: lMResult ==
                ( lIResult + cmd.MetamorphicValues[aCase-1][i].Difference ),
        }

        cmd.TableData( cmd.MetamorphicValues[aCase-1][i], lTestResult )
    }
    cmd.TableEnd()
}

func main() {
    // fmt.Println(" Metamorphic Test 1: Inverted Start & End")
    // metamorphicTests(1)

    // fmt.Println(" Metamorphic Test 2: Same Date, Different Years")
    // metamorphicTests(2)
    //
    fmt.Println("Metamorphic Test 3: Different Date, Same Month & Year")
    metamorphicTests(3)
}

