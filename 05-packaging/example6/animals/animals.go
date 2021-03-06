// Package animals provides support for animals.
package animals

// animal represents information about all animals.
type animal struct {
	Name string
	Age  int
}

// Dog represents information about dogs.
type Dog struct {
	animal       // Make the embeeded type unexported.
	BarkStrength int
}
