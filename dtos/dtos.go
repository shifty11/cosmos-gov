package dtos

import (
	"time"
)

type Proposal struct {
	ProposalId int `json:"proposal_id,string"`
	Content    struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	VotingStartTime time.Time `json:"voting_start_time"`
	VotingEndTime   time.Time `json:"voting_end_time"`
	Status          string    `json:"status"`
}

type Proposals struct {
	Proposals []Proposal `json:"proposals"`
}

type Chain struct {
	Name        string
	DisplayName string
	Notify      bool
}
