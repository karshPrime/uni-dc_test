
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


//- Datasets -------------------------------------------------------------------

// start becomes end, end becomes start
var MetamorphicValuesSet1 = [5]ProgramIO {
    // InitialStart, InitialEnd , ModifiedStart, ModifiedEnd, Difference
    { "05/01/2020", "15/05/2022", "15/05/2022", "05/01/2020", 0 },
    { "12/08/2018", "22/09/2019", "22/09/2019", "12/08/2018", 0 },
    { "03/11/2021", "14/12/2020", "14/12/2020", "03/11/2021", 0 },
    { "25/04/2015", "30/06/2016", "30/06/2016", "25/04/2015", 0 },
    { "09/02/2022", "18/03/2023", "18/03/2023", "09/02/2022", 0 },
}

// changed years with same difference
var MetamorphicValuesSet2 = [5]ProgramIO {
    // InitialStart, InitialEnd , ModifiedStart, ModifiedEnd, Difference
    { "10/10/2018", "20/11/2019", "10/10/2009", "20/11/2010", 0 },
    { "01/01/2021", "12/12/2022", "01/01/1990", "12/12/1991", 0 },
    { "15/07/2017", "25/08/2018", "15/07/2041", "25/08/2042", 0 },
    { "22/03/2020", "30/04/2021", "22/03/2037", "30/04/2038", 0 },
    { "05/06/2016", "14/07/2017", "05/06/1902", "14/07/1903", 0 },
}

// Different dates for Same Month and Same Year
var MetamorphicValuesSet3 = [5]ProgramIO {
    // InitialStart, InitialEnd , ModifiedStart, ModifiedEnd, Difference
    { "18/02/2015", "27/03/2016", "28/02/2015", "27/03/2016", -10 },
    { "08/09/2023", "17/10/2024", "08/09/2023", "27/10/2024", +10 },
    { "15/01/2012", "05/02/2013", "25/01/2012", "15/02/2013",   0 },
    { "11/04/2019", "20/05/2020", "01/04/2019", "30/05/2020", +20 },
    { "08/08/2025", "27/09/2026", "28/08/2025", "07/09/2026", -40 },
}

// Array of [5]ProgramIO arrays
var MetamorphicValues = [3][5]ProgramIO {
    MetamorphicValuesSet1,
    MetamorphicValuesSet2,
    MetamorphicValuesSet3,
}

