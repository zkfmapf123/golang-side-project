package base

type Movie struct {
	Id string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

func InitMovie(id string, isbn string, title string, firstName string, lastName string) Movie{
	d := Director{FirstName: firstName, LastName: lastName}
	return Movie{Id: id,Isbn: isbn, Title: title, Director: &d}
}

