package main

func main() {
	// var age int //declare an int
	// fmt.Println("My age is", age)
	// age = 29
	// fmt.Println("My age is", age)
	// var newAge int = 32
	// fmt.Println("My newAge is", newAge)

	// var name = "blake" //type is inferred like cpp auto
	// fmt.Println("My name is", name)

	//----------two variables at once----------
	// var width, height int = 100, 50
	// fmt.Println("width", width, "height", height)

	// var newWidth, newHeight = 100, 50 //can infer these too

	//----------different types----------
	// var (
	// 	newName   = "Blake"
	// 	newAge    = 29
	// 	newHeight int
	// )
	// fmt.Println("name", newName, "age", newAge, "height", newHeight)

	//-----------Shorthand----------
	//shorthand vs using var variableName type...
	//Note: using shorthand requires variables to be defined
	// name := "Blake"
	// age := 29

	//can do both simultaneously
	//name, age := "Blake", 23
	// fmt.Println(name, age)

	//can be used even if one variable is already declared
	//name, car := "John", "Mercedes"

	// a, b := 23.35, 45.12
	// c := math.Min(a, b)
	// fmt.Println(c)

	//-----------Type Casting-----------
	// i:=3.2
	// fmt.Println(int(i))
}
