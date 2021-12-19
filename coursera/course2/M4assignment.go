package main
import ("fmt"
"strings"
"os"
"bufio")

type Animal interface{
	Eat()
	Move()
	Speak()
	fetchName() string
	}
type Cow struct {name string}
type Bird struct {name string}
type Snake struct {name string}

func (c Cow) Eat(){
	    fmt.Println("grass")
		}
func (c Cow) Move(){    
	fmt.Println("walk")
}
func (c Cow) Speak(){
	fmt.Println("moo")
}
func (c Cow) fetchName() string{
	return c.name
}
func (b Bird) Eat(){    
	fmt.Println("worms")
}
func (b Bird) Move(){
	fmt.Println("fly")
}
func (b Bird) Speak(){
	fmt.Println("peep")
}
func (b Bird) fetchName() string{
	return b.name
}
func (s Snake) Eat(){
	fmt.Println("mice")
}
func (s Snake) Move(){    
	fmt.Println("slither")
}
func (s Snake) Speak(){
	fmt.Println("hsss")
}
func (s Snake) fetchName() string{
	return s.name
}
func main(){
inputReader := bufio.NewReader(os.Stdin)
animalList := []Animal{}
animalList = append(animalList, Cow{name: "c1"})
animalList = append(animalList, Bird{name: "b1"})
animalList = append(animalList, Snake{name: "s1"})
for{        
	fmt.Println(">")        
	fmt.Println("Animal list",animalList)
	fmt.Println("Enter your command")
	comm, _:= inputReader.ReadString('\n')
	command_type := strings.Split(comm," ")[0]
	animal_name := strings.Split(comm," ")[1]
	switch command_type{
		case"newanimal":
			animal_type := strings.Trim(strings.Split(comm," ")[2]," ")
			fmt.Println("animal type is",animal_type)
			if strings.Contains(animal_type,"cow") {			
					fmt.Println("case cow")
					animalList = append(animalList, Cow{name: animal_name})
					fmt.Println("Created a new animal "+animal_name+" of the type "+animal_type)
			} else if strings.Contains(animal_type,"bird"){
					fmt.Println("case bird")
					animalList = append(animalList, Bird{name: animal_name})
					fmt.Println("Created a new animal "+animal_name+" of the type "+animal_type)
			} else if strings.Contains(animal_type,"snake") {
					fmt.Println("case snake")
					animalList = append(animalList, Snake{name: animal_name})   
					fmt.Println("Created a new animal "+animal_name+" of the type "+animal_type)
			}else{
					fmt.Println("Incorrect animal type")		                
			}                                  
			fmt.Println("Length of animal array is ",len(animalList))                
			fmt.Println("List of animals",animalList)
		case"query":
			
			info_type := strings.Trim(strings.Split(comm," ")[2]," ")
			fmt.Println("Information type is",info_type)
			fmt.Println("Animal name is",animal_name)
			for _, animal := range animalList{
				if animal.fetchName() == animal_name{
					if strings.Contains(info_type,"eat"){
							animal.Eat()
					}else if strings.Contains(info_type,"move"){
							animal.Move()
					}else if strings.Contains(info_type,"speak"){
							animal.Speak()
					}else{
							fmt.Println("Incorrect information requested")
					}
				}         
			}               
		default: fmt.Println("Incorrect input type given")        
		}       
    }
}