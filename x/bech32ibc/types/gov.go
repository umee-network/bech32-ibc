package types

import (
	"fmt"
	"strings"
	time "time"

	gov1b1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

const (
	ProposalTypeUpdateHrpIbcChannel = "UpdateHrpIbcChannel"
)

func init() {
	// Deprecated: gov1b1 is deprecated
	gov1b1.RegisterProposalType(ProposalTypeUpdateHrpIbcChannel)
}

var _ gov1b1.Content = &UpdateHrpIbcChannelProposal{}

func NewUpdateHrpIBCRecordProposal(title, description, hrp, sourceChannel string, toHeightOffset uint64, toTimeOffset time.Duration) gov1b1.Content {
	return &UpdateHrpIbcChannelProposal{
		Title:             title,
		Description:       description,
		Hrp:               hrp,
		SourceChannel:     sourceChannel,
		IcsToHeightOffset: toHeightOffset,
		IcsToTimeOffset:   toTimeOffset,
	}
}

func (p *UpdateHrpIbcChannelProposal) GetTitle() string { return p.Title }

func (p *UpdateHrpIbcChannelProposal) GetDescription() string { return p.Description }

func (p *UpdateHrpIbcChannelProposal) ProposalRoute() string { return RouterKey }

func (p *UpdateHrpIbcChannelProposal) ProposalType() string {
	return ProposalTypeUpdateHrpIbcChannel
}

func (p *UpdateHrpIbcChannelProposal) ValidateBasic() error {
	err := gov1b1.ValidateAbstract(p)
	if err != nil {
		return err
	}
	if p.IcsToHeightOffset == 0 && p.IcsToTimeOffset == 0 {
		return ErrInvalidOffsetHeightTimeout
	}
	return ValidateHrp(p.Hrp)
}

func (p UpdateHrpIbcChannelProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Update HRP IBC Channel Proposal:
  Title:          %s
  Description:    %s
  HRP:            %s
  Source Channel: %s
`, p.Title, p.Description, p.Hrp, p.SourceChannel))
	return b.String()
}
