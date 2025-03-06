package main

type Person struct {
	Name string
	Age  int
}

func (p *Person) Grow() {
	p.Age++ // This will impact the original object
}

func (p Person) DoesNotGrow() {
	p.Age++ // This will not impact the original object
}

func main() {

	p := Person{Name: "John", Age: 20}
	println("Starting Age", p.Age)
	p.Grow()
	println("After calling Grow from p", p.Age)
	p.DoesNotGrow()
	println("After calling DoesNotGrow with p", p.Age)
	ptr := &p
	println("Created Pointer..")
	ptr.Grow()
	println("After calling Grow from ptr", ptr.Age)
	ptr.DoesNotGrow()
	println("After calling DoesNotGrow with ptr", ptr.Age)

}
