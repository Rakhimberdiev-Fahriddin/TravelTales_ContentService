package postgres

import (
	"database/sql"
	pb "my_module/generated/content_service"
)

type DestinationsRepo struct{
	DB *sql.DB
}

func NewDestinationsRepo(db *sql.DB) *DestinationsRepo{
	return &DestinationsRepo{DB: db}
}

func (d *DestinationsRepo) GetsDestinations(req *pb.GetsDestinationsRequest)(*pb.GetsDestinationsResponce,error){
	rows,err := d.DB.Query(`
	SELECT
		id,
		name,
		country,
		description
	FROM
		Destinations
	WHERE
		name = $1
	OFFSET $2
	LIMIT $3
	`,req.Query,(req.Page-1)*req.Limit,req.Limit)

	if err != nil{
		return nil,err
	}

	var destinations []*pb.Destinat
	for rows.Next(){
		var destination pb.Destinat
		err = rows.Scan(&destination.Id,&destination.Name,&destination.Country,&destination.Description)
		if err != nil{
			return nil,err
		}

		destinations = append(destinations, &destination)
	}
	var total int32
	err = d.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Destinations
	WHERE
		name = $1
	`,req.Query).Scan(&total)

	if err != nil{
		return nil,err
	}

	return &pb.GetsDestinationsResponce{
		Destinations: destinations,
		Total: total,
		Page: req.Page,
		Limit: req.Limit,
	},nil
}

func (d *DestinationsRepo) GetDestinations(id *pb.GetDestinationsRequest)(*pb.GetDestinationsResponce,error){
	res := pb.GetDestinationsResponce{}
	err := d.DB.QueryRow(`
	SELECT
		id,
		name,
		country,
		description,
		best_time_to_visit,
		average_cost_per_day,
		currency,
		language
	FROM
		Destinations
	WHERE
		id = $1
	`,id.Id).Scan(
		&res.Id,
		&res.Name,
		&res.Country,
		&res.Description,
		&res.BestTimeToVisit,
		&res.AverageCostPerDay,
		&res.Currency,
		&res.Language,
	)

	if err != nil{
		return nil,err
	}

	return &res,nil
}