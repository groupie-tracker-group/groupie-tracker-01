package data 

import "fmt"
import "net/http"
import "encoding/json"

func init(){
	Decoder()
}

type DataBase map[string][]string
// define a new data base 
func NewDataBase()*DataBase{
	var New DataBase
	return &New
}
// fetch an element
func (d *DataBase)Fetch(key string)[]string{
	return (*d)[key]
}

// return all keys in a slice
func (d *DataBase)Retrieve()[]string{
	var sl []string
	for key := range *d {
		sl = append(sl, key)
	}
	return sl
}

// get the data
type TemplateData struct{
	Artists []ArtistData
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
