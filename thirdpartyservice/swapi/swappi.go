package swapi

import (
	"context"
	"fmt"
	"net/url"

	"github.com/psbernardo/graphql/graph/model"
)

func (c *Client) Search(ctx context.Context, filter string) (*model.QueryResult, error) {
	peopleData := &PeopleData{}

	params := url.Values{

		"search": {filter},
	}
	url := fmt.Sprintf("https://swapi.dev/api/people/?%s", params.Encode())
	if err := c.GetData(ctx, url, peopleData); err != nil {
		return &model.QueryResult{}, err
	}

	queryResult := model.QueryResult{
		Count:    peopleData.Count,
		Next:     peopleData.Next,
		Previous: peopleData.Previous,
		Results:  []*model.People{},
	}

	for _, item := range peopleData.ResultList {

		people := model.People{
			Name:      item.Name,
			Height:    &item.Height,
			Mass:      &item.Mass,
			HairColor: &item.HairColor,
			SkinColor: &item.SkinColor,
			BirthYear: &item.Gender,
			Created:   &item.Created,
			Edited:    &item.Edited,
			URL:       item.URL,
		}

		VehicleList, err := c.GetVehicle(ctx, item.Vehicles...)
		if err != nil {
			return &model.QueryResult{}, err
		}
		people.Vehicles = VehicleList

		filmList, err := c.GetFilmList(ctx, item.Films...)
		if err != nil {
			return &model.QueryResult{}, err
		}
		people.Films = filmList

		queryResult.Results = append(queryResult.Results, &people)

	}

	return &queryResult, nil
}

func (c *Client) PageResult(ctx context.Context, url string) (*model.QueryResult, error) {
	peopleData := &PeopleData{}
	if err := c.GetData(ctx, url, peopleData); err != nil {
		return &model.QueryResult{}, err
	}

	queryResult := model.QueryResult{
		Count:    peopleData.Count,
		Next:     peopleData.Next,
		Previous: peopleData.Previous,
		Results:  []*model.People{},
	}

	for _, item := range peopleData.ResultList {

		people := model.People{
			Name:      item.Name,
			Height:    &item.Height,
			Mass:      &item.Mass,
			HairColor: &item.HairColor,
			SkinColor: &item.SkinColor,
			BirthYear: &item.Gender,
			Created:   &item.Created,
			Edited:    &item.Edited,
			URL:       item.URL,
		}

		VehicleList, err := c.GetVehicle(ctx, item.Vehicles...)
		if err != nil {
			return &model.QueryResult{}, err
		}
		people.Vehicles = VehicleList

		filmList, err := c.GetFilmList(ctx, item.Films...)
		if err != nil {
			return &model.QueryResult{}, err
		}
		people.Films = filmList

		queryResult.Results = append(queryResult.Results, &people)

	}

	return &queryResult, nil
}

func (c Client) GetVehicle(ctx context.Context, urlLIst ...string) (vehicleList []*model.Vehicle, err error) {

	vehicleChan := make(chan Vehicle)
	defer close(vehicleChan)
	chanErr := make(chan error)
	defer close(chanErr)

	for _, url := range urlLIst {
		go func() {
			var vehicle Vehicle
			if err := c.GetData(ctx, url, &vehicle); err != nil {
				chanErr <- err
			}
			vehicleChan <- vehicle
		}()
	}

	for range urlLIst {
		select {
		case err = <-chanErr:
			return
		case vehicle := <-vehicleChan:
			vehicleList = append(vehicleList, &model.Vehicle{
				Name:  vehicle.Name,
				Model: vehicle.Model,
			})
		}
	}
	return
}

func (c Client) GetFilmList(ctx context.Context, urlLIst ...string) (filmList []*model.Film, err error) {

	filmChan := make(chan Film)
	defer close(filmChan)
	chanErr := make(chan error)
	defer close(chanErr)

	for _, url := range urlLIst {
		go func(urlParam string) {
			var film Film
			if err := c.GetData(ctx, urlParam, &film); err != nil {
				chanErr <- err
			}
			filmChan <- film
		}(url)
	}

	for range urlLIst {
		select {
		case err = <-chanErr:
			return
		case film := <-filmChan:
			filmList = append(filmList, &model.Film{
				Title:     film.Title,
				EpisodeID: film.EpisodeID,
				Director:  film.Director,
				Producer:  film.Producer,
			})
		}
	}
	return
}
