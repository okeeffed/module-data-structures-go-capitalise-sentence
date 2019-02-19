package caps

type testCase struct {
	description string
	input       string
	expected    string
}

var testCases = []testCase{
	{
		description: "should return capitalised sentence",
		input:       "hello, world!",
		expected:    "Hello, World!",
	},
	{
		description: "should return capitalised sentence",
		input:       "hello, ## @world!",
		expected:    "Hello, ## @World!",
	},
	{
		description: "should return capitalised sentence",
		input:       "    ?wh , 324i up",
		expected:    "    ?Wh , 324I Up",
	},
	{
		description: "should return capitalised sentence",
		input:       "",
		expected:    "",
	},
}
