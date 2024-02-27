package main

func main(){
	tolleMap := map[string]string{
		"test": "test",
		"test2" : "test2",
		"test3" : "test3",
		"test4" : "test4",
		"test5" : "test5",
	}

	for k, v := range tolleMap {
		println("k:", k, "v:", v)
	}
}