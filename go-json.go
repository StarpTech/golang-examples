//http://starp-germany.de/blog/unmarshal-of-json-strings-in-go/

package main
 
import (
    "encoding/json" //import the json package
    "fmt"
    
)
 
//Json Type with a underlying interface as data
type Json struct {
 
    data interface{}
}
 
//Represent "result"
type Result struct {
 
    //Represent our Person object
    Person struct {
 
        Name string `json:"name"`
        Id float64  `json:"id"` //numbers are always float64 look on the bottom link
 
    } `json:"result"` //more details http://golang.org/pkg/encoding/json/
}
 
func main() {
 
    /*
 
    Method 1 unmarshaling of JSON string
 
 
    */
    var jsonBlob2 = []byte(`{ "result" : { "name" : "mark zuckerberg", "id" : 4 } }`)
    //create type object
    var personData Result
    
 
    if err_feed := json.Unmarshal(jsonBlob2, &personData); err_feed != nil  {
 
        panic("Unmarshal invalid")
 
    } else {
 
        //Print whole result and data consistency
        fmt.Println("ID: ",personData.Person.Id," consistency: ",
                        personData.Person.Id == float64(4) , 
                        " - Name: ",personData.Person.Name ," consistency: ",
                        personData.Person.Name == string("mark zuckerberg"))
    }
 
    //##################################################################################
    
 
    /*
 
    Method 2 unmarshaling of JSON string
 
 
    */
    var jsonBlob = []byte(`{ "result" : { "name" : "mark zuckerberg", "id" : 4 } }`)
    j := new(Json)
    err := json.Unmarshal(jsonBlob, &j.data)
    if err != nil {
 
        panic("Unmarshal invalid");
    }
 
    //Type asserts to `map` with string as key to interfaces as value
    if m, ok := (j.data).(map[string]interface{}); ok {
        
 
        //Key exist?
        if val, ok := m["result"]; ok {
            
            //create new Json Object
            result := &Json{val}
 
            //Type asserts to `map` with string as key to interfaces as value
            if innerResult,ok := (result.data).(map[string]interface{}); ok {
 
                //print whole result and data consistency
                fmt.Println("ID: ",innerResult["id"]," consistency: ",
                            innerResult["id"] == float64(4) , 
                            " - Name: ",innerResult["name"] ," consistency: ",
                            innerResult["name"] == string("mark zuckerberg"))
            } else {
 
                fmt.Println("Assertion to (map[string]interface{}) invalid!")
            }
            
        } else { //Wrong index 
 
            fmt.Println("Wrong index")
        }
    
    } else { //Key doesn´t exist
 
        fmt.Println("Assertion to (map[string]interface{}) invalid!")
 
    }
 
 
 
 
 
}
/*
 
OUTPUT:
 
ID:  4  consistency:  true  - Name:  mark zuckerberg  consistency:  true
ID:  4  consistency:  true  - Name:  mark zuckerberg  consistency:  true
 
*/
