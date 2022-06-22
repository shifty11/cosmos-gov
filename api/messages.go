package messages

const SubscriptionCmd = "subscriptions"
const SubscriptionsMsg = `🔔 *Subscriptions*

Select the projects that you want to follow. You will receive notifications about new governance proposals once they enter the voting period.
`

const ProposalsCmd = "proposals"
const ProposalsMsg = `

This are all ongoing proposals for your subscriptions.

`
const NoSubscriptionsMsg = `

You are not subscribed to any project.
Type /subscriptions to select the projects that you want to follow.
`
const NoProposalsMsg = `

There are currently no proposals in voting period.`

const SupportCmd = "support"
const SupportMsg = "💰 *Support*\n\nI would like to continue developing this bot and other products that improve the Cosmos ecosystem.\n\n" +
	"You have a good idea, feedback or want to contribute in other ways? Shoot a message to %v."
