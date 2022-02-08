package dtos

import (
	"time"
)

type ProposalContent struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Proposal struct {
	ProposalId      uint64          `json:"proposal_id,string"`
	Content         ProposalContent `json:"content"`
	VotingStartTime time.Time       `json:"voting_start_time"`
	VotingEndTime   time.Time       `json:"voting_end_time"`
	Status          string          `json:"status"`
}

type Proposals struct {
	Proposals []Proposal `json:"proposals"`
}

type Chain struct {
	Name        string
	DisplayName string
	Notify      bool
}
