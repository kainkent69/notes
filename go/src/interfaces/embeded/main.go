package main

import (
	"fmt"
	)

type Namer interface {
    GetName() string
}
// cursing
type Curser interface  {
    Curse(n Namer, word string) 
}

// the talkative interface
type Talkative interface {
	Talk(words string)
	Joke(joke string)
}

type Greeter interface {
    Namer
	Greet(to string)
}
// person
type Person struct {
	Greeter
	Name string
}
// talkative person
type TalkativePerson struct {
	*Person
	Talkative
}
// crazy 
type Crazy struct {
    *Person
    Curser
    Rude bool
}

// Mad 
type Mad struct {
   *Crazy
    Talkative
}


// method for greeter
func (p *Person) Greet(n Namer) {
	fmt.Printf("%q greets %s\n", p.Name, n.GetName())
}

func (p Person) GetName() string {
    return p.Name
}
// talkative methods
func (p TalkativePerson) Talk(word string) {
	fmt.Printf("%s says %q\n", p.Name, word)
}
func (p *TalkativePerson) Joke(word string) {
	fmt.Printf("%s joked %q\n", p.Name, word)
}

// mad talk
func (p Mad) Talk(word string) {
	fmt.Printf("%s says %q\n", p.Name, word)
}
// mad jokes
func (p Mad) Joke(word string) {
	fmt.Printf("%s joked %q\n", p.Name, word)
}
func (c Crazy) Curse(n Namer, word string){
    fmt.Printf("crazy %s  hey %q!, %s\n", c.Name, n.GetName(), word);
}

func main() {
	fmt.Println("Starting: First scenario")

	kim := Person{Name: "Kim"}
	rimmy := TalkativePerson{Person: &Person{Name: "Rimmy"}}
	// greering
	kim.Greet(rimmy)
	rimmy.Talk("What a great day!!!")
	// next part
	cdave := Crazy{Person:&Person{Name: "Dave"}, Rude: true}
	cdave.Curse(kim, "Fuck you!!");
	// forth part 
	james := Mad{
	    Crazy: &Crazy{Person: &Person{Name: "james"}},
	}
	_ = james
	
	// james doing weird things
	james.Greet(cdave)
	james.Talk("What a shame")
	james.Joke("This is a joke")
	// curse them all
	james.Curse(rimmy, "I hate you")
	james.Curse(cdave, "I hate you")
	james.Curse(kim, "I hate you")
	james.Curse(james, "I hate you")
}
