package postgres

import (
	"database/sql"
	pb "my_module/generated/content_service"
	"time"
)

type ItinerariesRepo struct {
	DB *sql.DB
}

func NewItinerariesRepo(db *sql.DB) *ItinerariesRepo {
	return &ItinerariesRepo{
		DB: db,
	}
}

func (i *ItinerariesRepo) CreateItineraries(reqItineraries *pb.CreateItinerariesRequest) (*pb.CreateItinerariesResponce, error) {
	resItineraries := pb.CreateItinerariesResponce{}
	err := i.DB.QueryRow(`
	INSERT INTO
		Itineraries(
			title,
			description,
			start_date,
			end_date,
			author_id
		)
		VALUES(
			$1,
			$2,
			$3,
			$4,
			$5
		)
		Returning
			id,
			title,
			description,
			start_date,
			end_date,
			author_id,
			created_at
	`,
		reqItineraries.Title,
		reqItineraries.Description,
		reqItineraries.StartDate,
		reqItineraries.EndDate,
		reqItineraries.AuthorId,
	).Scan(
		&resItineraries.Id,
		&resItineraries.Title,
		&resItineraries.Description,
		&resItineraries.StartDate,
		&resItineraries.EndDate,
		&resItineraries.AuthorId,
		&resItineraries.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	for _, destination := range reqItineraries.Destinations {
		var destinationId string
		err := i.DB.QueryRow(`
		INSERT INTO
			Itinerary_destinations(
				name,
				start_date,
				end_date
			)
			VALUES(
				$1,
				$2,
				$3
			)
			Returning
				id
		`, destination.Name, destination.StartDate, destination.EndDate).Scan(
			&destinationId,
		)
		if err != nil {
			return nil, err
		}

		for _, activiti := range destination.Activities {
			_, err = i.DB.Exec(`
				INSERT INTO
					Itinerary_activities(
						destination_id,
						activity
					)
					VALUES(
						$1,
						$2
					)
			`, destinationId, activiti)

			if err != nil {
				return nil, err
			}
		}
	}
	return &resItineraries, nil
}

func (i *ItinerariesRepo) UpdateItineraries(update *pb.UpdateItinerariesRequest) (*pb.UpdateItinerariesResponce, error) {
	resUpdate := pb.UpdateItinerariesResponce{}
	err := i.DB.QueryRow(`
	UPDATE 
		Itineraries
	SET
		title = $1,
		description = $2,
		updated_at = $3
	WHERE
		id = $4
	Returning
		id,
		title,
		description,
		start_date,
		end_date,
		author_id,
		updated_at
	`, update.Title, update.Description, time.Now(), update.Id).Scan(
		&resUpdate.Id,
		&resUpdate.Title,
		&resUpdate.Description,
		&resUpdate.StartDate,
		&resUpdate.EndDate,
		&resUpdate.AuthorId,
		&resUpdate.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &resUpdate, nil
}

func (i *ItinerariesRepo) DeleteItineraries(id *pb.DeleteItinerariesRequest) (*pb.DeleteItinerariesResponce, error) {
	_, err := i.DB.Exec(`
	DELETE FROM
		Itineraries
	WHERE
		id = $1
	`, id.Id)
	if err != nil {
		return &pb.DeleteItinerariesResponce{
			Message: "Failed deleted Itineraries",
		}, err
	}
	return &pb.DeleteItinerariesResponce{
		Message: "Succesfully deleted Itineraries",
	}, nil
}

func (i *ItinerariesRepo) GetsItineraries(req *pb.GetsItinerariesRequest) (*pb.GetsItinerariesResponce, error) {
	rows, err := i.DB.Query(`
	SELECT
		id,
		title,
		author_id,
		start_date,
		end_date,
		likes_count,
		comments_count,
		created_at
	FROM
		Itineraries
	OFFSET $1
	LIMIT $2;
	`, (req.Page-1)*req.Limit, req.Limit)

	if err != nil {
		return nil, err
	}

	var itineraries []*pb.Itinerarie
	for rows.Next() {
		var authorId string
		var itinerarie pb.Itinerarie
		err = rows.Scan(&itinerarie.Id, &itinerarie.Title, &authorId, &itinerarie.StartDate, &itinerarie.EndDate, &itinerarie.LikesCount, &itinerarie.CommentsCount, &itinerarie.CreatedAt)
		if err != nil {
			return nil, err
		}
		var author pb.Author
		err := i.DB.QueryRow(`
		SELECT
			username
		FROM
			users
		WHERE
			id = $1
		`, authorId).Scan(&author.UserName)
		author.Id = authorId

		if err != nil {
			return nil, err
		}
		itinerarie.Author = &author
		itineraries = append(itineraries, &itinerarie)
	}
	var total int32

	err = i.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Itineraries
	`).Scan(&total)

	if err != nil {
		return nil, err
	}

	return &pb.GetsItinerariesResponce{
		Itineraries: itineraries,
		Total:       total,
		Page:        req.Page,
		Limit:       req.Limit,
	}, nil
}

func (i *ItinerariesRepo) GetItineraries(id *pb.GetItinerariesRequest) (*pb.GetItinerariesResponce, error) {
	res := pb.GetItinerariesResponce{}

	var authorId string

	err := i.DB.QueryRow(`
	SELECT
		id,
		title,
		description,
		start_date,
		end_date,
		likes_count,
		comments_count,
		created_at,
		updated_at,
		author_id
	FROM
		Itineraries
	WHERE
		id = $1
	`, id.Id).Scan(
		&res.Id,
		&res.Title,
		&res.Description,
		&res.StartDate,
		&res.EndDate,
		&res.LikesCount,
		&res.CommentsCount,
		&res.CreatedAt,
		&res.UpdatedAt,
		&authorId,
	)

	if err != nil {
		return nil, err
	}

	var author pb.ItinerarieAuthor

	err = i.DB.QueryRow(`
	SELECT
		id,
		username,
		full_name
	FROM
		Users
	WHERE
		id = $1
	`, authorId).Scan(
		&author.Id,
		&author.UserName,
		&author.FullName,
	)

	if err != nil {
		return nil, err
	}

	res.Author = &author

	rows, err := i.DB.Query(`
	SELECT
		name,
		start_date,
		end_date,
		destination_id
	FROM
		Itinerary_destinations
	WHERE
		itinerary_id = $1
	`, id.Id)

	if err != nil {
		return nil, err
	}

	var destinations []*pb.Destination

	for rows.Next() {
		var destination pb.Destination
		var destinationId string
		rows.Scan(
			&destination.Name,
			&destination.StartDate,
			&destination.EndDate,
			&destinationId,
		)

		rows, err := i.DB.Query(`
		SELECT
			activity
		FROM
			Itinerary_activities
		WHERE
			destination_id = $1
		`, destinationId)
		if err != nil {
			return nil, err
		}

		var activities []string

		for rows.Next() {
			var activitie string
			err := rows.Scan(&activitie)
			if err != nil {
				return nil, err
			}
			activities = append(activities, activitie)
		}
		destination.Activities = activities

		destinations = append(destinations, &destination)
	}
	res.Destinations = destinations

	return &res, nil
}
