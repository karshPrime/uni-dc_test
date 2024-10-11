
package cmd

type ProgramInputs struct {
    Start   string
    End     string
    Difference int
}

var TestValues = [5]ProgramInputs {
    { "18/02/2015", "27/03/2016", 0 }, // covers 1 leap day
    { "15/01/2012", "05/02/2013", 0 },
    { "11/04/2000", "20/05/2003", 0 },
    { "08/09/2023", "17/10/2024", 0 }, // covers 1 leap day
    { "08/08/2024", "27/09/2026", 0 },
    // difference field doesn't make logical sense for the test struct
    // but is here to comply with the struct definition
}

//- Metamorphic Tests ----------------------------------------------------------

var MetamorphicExpected = [3][5]ProgramInputs {
    {// Start becomes End. End becomes Start
        { "27/03/2016", "18/02/2015", 0 },
        { "05/02/2013", "15/01/2012", 0 },
        { "20/05/2003", "11/04/2000", 0 },
        { "17/10/2024", "08/09/2023", 0 },
        { "27/09/2026", "08/08/2024", 0 },
        // expected difference is 0, since dates are just swapped
    },

    {// Changed Years for Start and End with same constant
        { "18/02/2007", "27/03/2008", 0 }, // -8
        { "15/01/1912", "05/02/1913", 0 }, // -100
        { "11/04/2024", "20/05/2027", 0 }, // +24
        { "08/09/2055", "17/10/2056", 0 }, // +32
        { "08/08/2124", "27/09/2126", 0 }, // +100
        // expected difference is 0, since years difference is kept consistent
    },

    {// Different dates for Same Month and Same Year
        { "08/02/2015", "27/03/2016", +10 },
        { "25/01/2012", "05/02/2013", -10 },
        { "01/04/2000", "10/05/2003",   0 }, // rise in one balances other's fall
        { "01/09/2023", "30/10/2024", +20 },
        { "28/08/2024", "07/09/2026", -40 },
        // increasing start = decreases days count
        // increasing end   = increases days count
    },
}

//------------------------------------------------------------------------------

