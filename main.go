
package main

import (
    "fmt"
    "github.com/karshPrime/uni-dc_test/cmd"

    "github.com/wy3148/dc/date" // Program Under Test
)

func metamorphicTests( aCase int ) {
    cmd.TableHead()

    for i := 0; i < 5; i++ {

        lTestResult := cmd.TestResult {
        }

        cmd.TableData( cmd.MetamorphicValues[aCase-1][i], lTestResult )
    }
    cmd.TableEnd()
}

func main() {
    fmt.Println(" Metamorphic Test 1: Inverted Start & End")
    fmt.Println("\n Metamorphic Test 2: Same Date, Different Years")

    fmt.Println("\n Metamorphic Test 3: Different Date, Same Month & Year")
}

