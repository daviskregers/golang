package main

import (
	"fmt"
    "reflect"
)

type Doctor struct {
    number int
    actorName string
    companions []string
}

func main() {

    // Maps
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)

    // You can use arrays as keys
    m := map[[3]int]string{}
    fmt.Println(m)

    // make function
    stpop := make(map[string]int)
    stpop = map[string]int{"California": 39250017}
    fmt.Println(stpop)

    // 

    fmt.Println(statePopulations["California"])
    statePopulations["Georgia"] = 10310371
    fmt.Println(statePopulations)
    delete(statePopulations, "Georgia")
    fmt.Println(statePopulations)
    fmt.Println(statePopulations["Georgia"]) // 0
    pop, ok := statePopulations["Georgia"]
    fmt.Println(pop, ok) // 0, false

    // underlying data is passed by references

    sp := statePopulations
    delete(sp, "Ohio")
    fmt.Println(sp)
    fmt.Println(statePopulations)

    // structs

    aDoctor := Doctor {
        number: 3,
        actorName:  "Jon Pertwee",
        companions: []string {
            "Liz Shaw",
            "Jo Grant",
            "Sarah Jane Smith",
        },
    }

    fmt.Println(aDoctor)
    fmt.Println(aDoctor.companions)
    fmt.Println(aDoctor.companions[1])

    // you can also initialize the struct without
    // providing field names (positional structs), but later on
    // if there are any changes in the code, this might break.
    // not recommended.

    bDoctor := Doctor{3, "Jon Pertwee", []string{"1", "2"}}
    fmt.Println(bDoctor)

    // anonymous structs
    doc := struct{name string}{name: "John Pertwee"}
    fmt.Println(doc)

    // unlike maps, we create copies
    anotherDoctor := doc
    anotherDoctor.name = "Tom Baker"
    fmt.Println(doc)
    fmt.Println(anotherDoctor)

    // If we like to to reference, we use a pointer
    yetAnotherDoc := &doc
    yetAnotherDoc.name = "Thomas Baker"
    fmt.Println(doc)
    fmt.Println(yetAnotherDoc)

    // Composition, similar concept to inheritance

    type Animal struct {
        Name string
        Origin string
    }

    type Bird struct {
        Animal
        SpeedKPH    float32
        CanFly      bool
    }

    b := Bird{}
    b.Name = "Emu"
    b.Origin = "Australia"
    b.SpeedKPH = 48
    b.CanFly = false

    fmt.Println(b) // {{Emu Australia} 48 false}
    fmt.Println(b.Name) // Emu

    c := Bird {
        Animal:     Animal{Name: "Emu", Origin: "Australia"},
        SpeedKPH: 48,
        CanFly: false,
    }
    fmt.Println(c)

    // tags, using reflect package

    type AnimalWithTags struct {
        Name        string     `required max:"100"`
        Origin      string
    }

    t := reflect.TypeOf(AnimalWithTags{})
    field, _ := t.FieldByName("Name")
    fmt.Println(field.Tag) // required max:"100"

}

