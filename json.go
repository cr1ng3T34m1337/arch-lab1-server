package json 

func Stringify(str string) ([]byte, error){
    response := map[string]string{
        "time": str,
    }
    return json.MarshalIndent(response, "", "\t")
}