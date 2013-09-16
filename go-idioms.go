//http://starp-germany.de/blog/go-specials/

    /*
    Represent the last assignment
    
    */
    type TypeAssignable struct {
    
        lastAssertionFrom *reflect.Type
    	lastAssertionTo *reflect.Type
    
    }
    /*
    Compares to types whether type1 is assignable to type2
    
    */
    func (t *TypeAssignable) AssignableTo(data interface{}, Type interface{} ) (bool,error) {
    	
    	if data == nil || Type == nil {
    
    		return false,errors.New("Invalid type passed!")
    	}
    
    	typ := reflect.TypeOf(data)
    	t.lastAssertionFrom = &typ
    	typ2 := reflect.TypeOf(Type)
    	t.lastAssertionTo = &typ2
    
    	if typ.AssignableTo(typ2) {
    
    
    	} else {
    
    		return false, errors.New("type assignment from '" + typ.String() +"' to '"+typ2.String()+"' failed")
    	}
    
    
    	return true , nil
    	
    
    }
    
    //Subtyping through Interfaces
    
    type Animal interface {
     
         Run()
    }
    type Dog struct {
        
    }
    func (d Dog) Run() {
        
     
    }
    dog := Dog{}
    animal  := Animal(dog) //or var animal Animal = dog
    animal.Run()
    fmt.Println("Type: ",reflect.TypeOf(animal)) //Prints Dog
    
    //Subclassing through ‘type embedding’
    
    type Person struct {
        
        username string
        password string
     
    }
    type Admin struct {
        Person //type embedding
     
    }
    func (d *Admin) Login() {
        
        fmt.Println("yeah access! with" + d.username)
    }
    admin := Admin{Person{ username : "lisa", password: "14de43" } }
    fmt.Println(admin.username) //prints 14de43
    fmt.Println("Type: ",reflect.TypeOf(admin)) //prints admin

    //Underlying type and Conversion

    var assign TypeAssignable
 
    /*
 
    SPECIAL - its always depends on equality of the underlying types
 
    */
    //because runes has int32 in the underlying system so rune == int32
    dynATest,err := assign.AssignableTo([]rune{},[]int32{})
    fmt.Println(dynATest,err) //true
 
    //The same type but different memory struture so it isn´t assignable
    dynATest,err = assign.AssignableTo([]string{},[5]string{})
    fmt.Println(dynATest,err) //false
    /*
 
    Cannot be assigned to one another without a conversion
 
 
    */
    type MyInt int
 
    var i int
    var j MyInt
 
    dynATest, err = assign.AssignableTo(i, j)
    fmt.Println(dynATest,err) //false 
    dynATest, err = assign.AssignableTo(i, int(j))
    fmt.Println(dynATest,err) //true



    /*
 
    SPECIAL - No implicit numeric conversion
 
    */
    //Go doens´t provide implicit numeric conversions
    dynATest,err = assign.AssignableTo(5,5.0)
    fmt.Println(dynATest,err) //false
    dynATest,err = assign.AssignableTo(5.0,5)
    fmt.Println(dynATest,err) //false
 
    /*
    //##############------Interface a dynamic way--------######################
    /*

    That´s the only “dynamic” mechansim in Go. A interface store a value of any type with a method set.
    The reason that every type have at least zero methods allows an empty interface to store every type without explicit type declaration.
    The type is encapsulated under the interface and the type information are never get lost.
    
    */
    /*
 
    SPECIAL - variable of interface type can store a value of any type
 
    */
    // Valid declaration with init value nil
    var d interface{}
    d = 44    //int
    d = "333" //string
    d = assign //struct
 
    dType := reflect.TypeOf(d)
    fmt.Println("Type: ",dType) //prints Type:  main.TypeAssignable

    /*
 
    SPECIAL - Allocation of an empty interface pointer
 
    */
    d2 := new (interface{})
 
    fmt.Println(d2)

    /*
 
    SPECIAL - Invalid with Short variable declaration because an interface has no initial expression
 
    */
    //d3 := interface{}{}

    /*



    Arrays call by value

    In lots of other languages like Java, C# arrays are handled as references but in Go not. The reason:
    */
    
    /*
    SPECIAL - Arrays are values
	

	*/

	array := [...]string{ "test" }
	

	test := func(test [1]string) {
		
		fmt.Println(test) //prints [test]
		//change value
		test[0] = "ttttt"

	}

	test(array)
	fmt.Println(array[0]) //prints test


    /*
    
	SPECIAL - Maps, slices, and channels are references

	*/

	
	maps := map[string]string{ "t" : "test"}

	test2 := func(test2 map[string]string) {
		
		fmt.Println(test2) //prints map[t:test]
		//change value
		test2["t"] = "ttttt"

	}

	test2(maps)
	fmt.Println(maps["t"]) //prints ttttt
