package postgres

import (
	"database/sql"
	pb "my_module/generated/content_service"
)

type MessageRepo struct{
	DB *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo{
	return &MessageRepo{DB: db}
}

func (m *MessageRepo) Message(message *pb.MessageRequest)(*pb.MessageResponce,error){
	res := pb.MessageResponce{}
	err := m.DB.QueryRow(`
	INSERT INTO
		Messages(
			recipient_id,
			content,
			sender_id
		)
		VALUES(
			$1,
			$2,
			$3
		)
	Returning
		id,
		sender_id,
		recipient_id,
		content,
		created_at
	`,message.RecipientId,message.Content,message.SenderId).Scan(
		&res.Id,
		&res.SenderId,
		&res.RecipientId,
		&res.Content,
		&res.CreatedAt,
	)

	if err != nil{
		return nil,err
	}
	return &res,nil
}

func (m *MessageRepo) GetMessage(req *pb.GetMessageRequest)(*pb.GetMessageResponce,error){
	rows,err := m.DB.Query(`
	SELECT
		id,
		content,
		created_at,
		sender_id
	FROM
		Messages
	OFFSET $1
	LIMIT $2
	`,(req.Page-1)*req.Limit,req.Limit)

	if err != nil{
		return nil,err
	}

	var messages []*pb.Message
	for rows.Next(){
		var senderId string
		var message pb.Message
		err = rows.Scan(&message.Id,&message.Content,&message.CreatedAt,&senderId)
		
		if err != nil{
			return nil,err
		}

		var sender pb.Sender

		m.DB.QueryRow(`
		SELECT
			id,
			user_name
		FROM
			Users
		WHERE
			id = $1
		`,senderId).Scan(&sender.Id,&sender.UserName)
		message.Sender = &sender

		messages = append(messages, &message)
	}
	
	var total int32

	m.DB.QueryRow(`
	SELECT
		COUNT(*)
	FROM
		Messages
	`).Scan(&total)

	return &pb.GetMessageResponce{
		Messages: messages,
		Total: total,
		Page: req.Page,
		Limit: req.Limit,
	},nil
}