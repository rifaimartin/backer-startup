package campaign

import (
	"backer-startup/user"
	"time"

	"github.com/leekchan/accounting"
)

// Campaign Struct
type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

// GoalAmountFormatIDR format goal amount
func (c Campaign) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.GoalAmount)
}

//CurrentAmountFormatIDR format current amount

func (c Campaign) CurrentAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "IDR ", Precision: 0, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.CurrentAmount)
}

// CampaignImage this is Struct
type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
