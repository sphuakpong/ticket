package db

type StatusWithRelated struct {
	Status
	Tickets []Ticket `json:"tickets"`
}

type BoardWithRelated struct {
	Board
	Statuses []StatusWithRelated `json:"statuses"`
}

func NewBoardWithRelated(b Board, s []Status, t []Ticket) BoardWithRelated {
	bw := BoardWithRelated{
		Board:    b,
		Statuses: []StatusWithRelated{},
	}

	for _, status := range s {
		bw.Statuses = append(bw.Statuses, NewStatusWithRelated(status, t))
	}

	return bw
}

func NewStatusWithRelated(s Status, t []Ticket) StatusWithRelated {
	sw := StatusWithRelated{
		Status:  s,
		Tickets: []Ticket{},
	}

	for _, ticket := range t {
		if uint32(ticket.StatusID) == s.ID {
			sw.Tickets = append(sw.Tickets, ticket)
		}
	}

	return sw
}
