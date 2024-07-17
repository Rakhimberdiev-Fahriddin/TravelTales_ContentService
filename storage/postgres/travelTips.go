package postgres

import (
	"database/sql"
	pb "my_module/generated/content_service"
)

type TravelTipsRepo struct{
	DB *sql.DB
}

func NewTravelTipsRepo(db *sql.DB)*TravelTipsRepo{
	return &TravelTipsRepo{DB: db}
}

func (t *TravelTipsRepo) AddTravelTips(travelTips *pb.AddTravelTipsRequest)(*pb.AddTravelTipsResponce,error){
	res := pb.AddTravelTipsResponce{}
	err := t.DB.QueryRow(`
	INSERT INTO
		Travel_tips(
			author_id,
			title,
			content,
			category
		)
		VALUES(
			$1,
			$2,
			$3,
			$4	
		)
	Returning
		id,
		title,
		content,
		category,
		author_id,
		created_at
	`,travelTips.AuthorId,travelTips.Title,travelTips.Content,travelTips.Category).Scan(
		&res.Id,
		&res.Title,
		&res.Content,
		&res.Category,
		&res.AuthorId,
		&res.CreatedAt,
	)

	if err != nil{
		return nil,err
	}

	return &res,nil
}

func (t *TravelTipsRepo) GetTravelTips(req *pb.GetTravelTipsRequest)(*pb.GetTravelTipsResponce,error){
	rows,err := t.DB.Query(`
	SELECT
		id,
		title,
		category,
		created_at,
		author_id
	FROM
		Travel_tips
	WHERE
		category = $1
	OFFSET $2
	LIMIT $3
	`,req.Category,(req.Page-1)*req.Limit,req.Limit)

	if err != nil{
		return nil,err
	}

	var tips []*pb.Tip
	for rows.Next(){
		var userId string
		var tip pb.Tip
		err = rows.Scan(&tip.Id,&tip.Title,&tip.Category,&tip.CreatedAt,&userId)
		
		if err != nil{
			return nil,err
		}

		var author pb.Author
		err = t.DB.QueryRow(`
		SELECT 
			id,
			username
		FROM
			Users
		WHERE
			id = $1
		`,userId).Scan(&author.Id,&author.UserName)

		if err != nil{
			return nil,err
		}

		tip.Author = &author

		tips = append(tips, &tip)
	}

	var total int32

	err = t.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Travel_tips
	`).Scan(&total)

	if err != nil{
		return nil,err
	}

	return &pb.GetTravelTipsResponce{
		Tips: tips,
		Total: total,
		Page: req.Page,
		Limit: req.Limit,
	},nil
}

func (t *TravelTipsRepo) StatisticsUser(req *pb.StatisticsUserRequest)(*pb.StatisticsUserResponce,error){
	var totalStories int32
	err := t.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Stories
	WHERE
		author_id = $1
	`,req.UserId).Scan(&totalStories)

	if err != nil{
		return nil,err
	}

	var totalItineraries int32
	err = t.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Itineraries
	WHERE
		author_id = $1
	`,req.UserId).Scan(&totalItineraries)

	if err != nil{
		return nil,err
	}

	var totalLikes int32
	err = t.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Likes
	WHERE
		user_id = $1
	`,req.UserId).Scan(&totalLikes)

	if err != nil{
		return nil,err
	}

	var totalComments int32

	err = t.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Messages
	WHERE
		sender_id = $1
	`,req.UserId).Scan(&totalComments)

	if err != nil{
		return nil,err
	}

	var mostStory pb.MostPopularStory
	err = t.DB.QueryRow(`
	SELECT
		id,
		title
	FROM
		Stories
	WHERE
		author_id = $1
	LIMIT $2
	`,req.UserId,1).Scan(&mostStory.Id,&mostStory.Title)
	mostStory.LikesCount = totalLikes

	if err != nil{
		return nil,err
	}

	var mostItinerary pb.MostPopularItinerary
	err = t.DB.QueryRow(`
	SELECT
		id,
		title
	FROM
		Itineraries
	WHERE
		author_id = $1
	LIMIT $2
	`,req.UserId,1).Scan(&mostItinerary.Id,&mostItinerary.Title)
	mostItinerary.LikesCount = totalLikes

	if err != nil{
		return nil,err
	}

	return &pb.StatisticsUserResponce{
		UserId: req.UserId,
		TotalStories:  totalStories,
		TotalItineraries: totalItineraries,
		TotalCountriesVisited: req.TotalCountriesVisited,
		TotalLikesReceived: totalLikes,
		TotalCommentsReceived: totalComments,
		MostPopularStory: &mostStory,
		MostPopularItinerary: &mostItinerary,
	},nil
}

func (t *TravelTipsRepo) GetTrendingDestinations(req *pb.GetTrendingDestinationsRequest)(*pb.GetTrendingDestinationsResponce,error){
	rows,err := t.DB.Query(`
	SELECT
		id,
		name,
		country,
		popularity_score
	FROM
		Destinations
	ORDER BY popularity_score  DESC
	`,)
	
	if err != nil{
		return nil,err
	} 

	var destinations []*pb.Destinations
	for rows.Next(){
		var destination pb.Destinations
		
		err = rows.Scan(&destination.Id,&destination.Name,&destination.Country,&destination.PopularityScore)
		
		if err != nil{
			return nil,err
		}

		destinations = append(destinations, &destination)
	}
	var total int32
	err = t.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Destinations
	WHERE
		popularity_score > $1
	`,80).Scan(&total)
	
	if err != nil{
		return nil,err
	}

	return &pb.GetTrendingDestinationsResponce{
		Destinations: destinations,
		Total: total,
	},nil
}

