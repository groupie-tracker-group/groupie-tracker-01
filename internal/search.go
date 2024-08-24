package data 

import "fmt"
import "net/http"
import "encoding/json"
import "strconv"
// run this function befor lunching the server
func init(){
	InitializeDataBases()
	Decoder()
	FillOutDataForms()
	FillOutSearchKeys()
}

type DataBase struct { 
	Base map[string][]int
}

// add element to the data
// if the id exist scip
func (d *DataBase) Add(key string, value int) {
    values := (*d).Base[key]  // 
    for _, v := range values { // 
        if v == value {
            return
        }
    }
    (*d).Base[key] = append((*d).Base[key], value) // Line 28
}
// fetch an element
func (d *DataBase)Fetch(key string)[]int{
	return (*d).Base[key]
}

// return all keys in a slice
func (d *DataBase)Retrieve()[]string{
	var sl []string
	for key := range (*d).Base {
		sl = append(sl, key)
	}
	return sl
}

// get the data fromt the aip
type TemplateData struct{
	Artists []ArtistData
	SearchSlice []string
}

var Template_data TemplateData


func Decoder(){
	resp , err := http.Get(Urls.ArtistsApi)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK{
		fmt.Println("Error: Status code %d\n", resp.StatusCode)
		return
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&Template_data.Artists)
	if err != nil {
		fmt.Println("err decoding JSON:" , err)
		return
	}
}
/////////////////////////////////////////////////////
// put the data into separate data bases
func InitializeDataBases() {
    By_Name = &DataBase{Base: make(map[string][]int)}
    By_Member = &DataBase{Base: make(map[string][]int)}
    By_First_Album = &DataBase{Base: make(map[string][]int)}
    By_Creation_Date = &DataBase{Base: make(map[string][]int)}
	By_Location = &DataBase{Base: make(map[string][]int)}
}
var By_Name *DataBase
var By_Member *DataBase
var By_First_Album *DataBase
var By_Creation_Date *DataBase
var By_Location *DataBase

func FillOutDataForms(){
	for  _,artist := range Template_data.Artists {
		var id = artist.ID
		By_Name.Add(artist.Name , artist.ID)
		By_First_Album.Add(artist.FirstAlbum, artist.ID)
		By_Creation_Date.Add(strconv.Itoa(artist.CreationDate), artist.ID)
		// need a range 
		for _, mem := range artist.Members{
			By_Member.Add( mem , id)
		}
	}
}
var All_Search_Keys struct{
	SearchKyes []string
}

func FillOutSearchKeys(){
	Template_data.SearchSlice = append(Template_data.SearchSlice, By_Name.Retrieve()...)
	Template_data.SearchSlice = append(Template_data.SearchSlice, By_Creation_Date.Retrieve()...)
	Template_data.SearchSlice = append(Template_data.SearchSlice, By_First_Album.Retrieve()...)
	Template_data.SearchSlice = append(Template_data.SearchSlice, By_Location.Retrieve()...)
}
