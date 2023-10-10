package models

type IMovie interface {
	GetMovieName() string
	SetMovieName(string) IMovie
	GetMovieCast() []IActor
	SetMovieCast([]IActor) IMovie
	GetMovieLanguages() []Language
	SetMovieLanguages([]Language) IMovie
}

type Movie struct {
	movieName      string
	movieCast      []IActor
	movieLanguages []Language
}

func (m *Movie) GetMovieName() string {
	return m.movieName
}
func (m *Movie) SetMovieName(name string) IMovie {
	m.movieName = name
	return m
}
func (m *Movie) GetMovieCast() []IActor {
	return m.movieCast
}

func (m *Movie) SetMovieCast(cast []IActor) IMovie {
	m.movieCast = cast
	return m
}

func (m *Movie) GetMovieLanguages() []Language {
	return m.movieLanguages
}

func (m *Movie) SetMovieLanguages(languages []Language) IMovie {
	m.movieLanguages = languages
	return m
}
