package repository

import (
	"zkfmapf123/crud-api/src/base"
)

type IMovie interface{
	RetrieveAll() ([]base.Movie, error)
	Retrieve(id string) (base.Movie, error)
	MovieUpdate(id string, m base.Movie) error
	MovieDelete(id string) error
}

type MovieRepository struct{}

func (movie MovieRepository) RetrieveAll() ([]base.Movie, error) {
	movies := make([]base.Movie,5)
	movies = append(movies, base.InitMovie("1", "asdfadsf","avatar","one","donggyu"))
	movies = append(movies, base.InitMovie("2", "asldknf","titanic","two","donggyu"))
	movies = append(movies, base.InitMovie("3", "aslgnlq","the_glory","three","donggyu"))
	movies = append(movies, base.InitMovie("4", "qwlkenr","back_to_the_future","four","donggyu"))
	movies = append(movies, base.InitMovie("5", "abalnks","i_am_legend","five","donggyu"))

	return movies, nil
}

func (movie MovieRepository) Retrieve(id string) (base.Movie, error) {

	return base.Movie{},nil
}

func (movie MovieRepository) MovieUpdate(id string, m base.Movie) error {

	return nil
}

func (movie MovieRepository) MovieDelete(id string) error {

	return nil
}