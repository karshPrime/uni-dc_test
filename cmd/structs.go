
package cmd

type ProgramIO struct {
    IStart  string
    IEnd    string

    MStart  string
    MEnd    string

    Difference int
}

type TestResult struct {
    IResult     int
    MResult     int
    TestCase    bool
}

