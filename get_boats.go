func get_boats() int {
	fmt.Println(">>> nbrBateau")
	var URL_SERVER = "0.0.0.0:3001"
	fmt.Println(URL_SERVER + "/boats")
	resp, err := http.Get(URL_SERVER + "/boats")
	if err != nil {
		panic(err)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	type BoatResponse struct {
		NbrBateau int
	}
	var responseObject BoatResponse
	json.Unmarshal(responseData, &responseObject)
	fmt.Println("nbrBateau = ", responseObject.NbrBateau)
	return responseObject.NbrBateau
}