package models

type IMovie interface {
	GetMovieName() string
	SetMovieName(string) IMovie
	GetMovieCast() []string            //todo: change to caat
	SetMovieCast([]string) IMovie      //todo: change to caat
	GetMovieLanguages() []string       //todo: change to language
	SetMovieLanguages([]string) IMovie //todo: change to language
}

type Movie struct {
	movieName      string
	movieCast      []string //todo: change to cast
	movieLanguages []string //todo: change to language
}

func (m *Movie) GetMovieName() string {
	return m.movieName
}
func (m *Movie) SetMovieName(name string) IMovie {
	m.movieName = name
	return m
}
func (m *Movie) GetMovieCast() []string {
	return m.movieCast
} //todo: change to caat

func (m *Movie) SetMovieCast(cast []string) IMovie {
	m.movieCast = cast
	return m
} //todo: change to caat

func (m *Movie) GetMovieLanguages() []string { //todo: change to language
	return m.movieLanguages
}

func (m *Movie) SetMovieLanguages(languages []string) IMovie {
	m.movieLanguages = languages
	return m
	//todo: change to language
}
