package telegram

const NbrOfButtonsPerRow = 3

const subscriptionsMsg = `Select the projects that you want to follow. You will receive notifications about new governance proposals once they enter the voting period.`

const startBroadcastInfoMsg = `Send the broadcast that you want to send to all participants

**Format**
<b>bold</b>
<a href='https://telegram.org'>Telegram</a>
<i>italic</i>
<code>code</code>
<s>strike</s>
<u>underline</u>
`
const confirmBroadcastMsg = "Are you sure you want to send this message to %v users?\nyes/**no**/abort"
const abortBroadcastMsg = "Abort broadcasting message."
const successBroadcastMsg = "Successfully sent message to %v users."

const chainStatisticHeaderMsg = `
      Chain ┆ Props ┆ Subs 
 -----------┼-------┼--------`
const chainStatisticRowMsg = `
%11.11s ┆ %5d ┆ %5d`
const chainStatisticFooterMsg = `
 -----------┼-------┼--------
  %s ┆ %5d ┆ %5d`
const userStatisticMsg = `
       Time ┆   Users (change)
 -----------┼------------------
   All time ┆ %7d
     1 week ┆ %7d %+7.2f%%
      1 day ┆ %7d %+7.2f%%`

const proposalsMsg = "This are all ongoing proposals for your subscriptions.\n\n"
const noSubscriptionsMsg = "You are not subscribed to any project.\nType /subscriptions to select the projects that you want to follow."
const noProposalsMsg = "There are currently no proposals in voting period."
const helpMsg = `<b>Commands List</b>
/subscriptions - Manage your subscriptions
/proposals - Show proposals in voting period
/support - Show how you can support this bot
/help - Show bot commands`

const adminHelpMsg = `<b>Admin Commands</b>
/stats - Show statistics
/chains - Manage chains
/broadcast - Broadcast message to everyone`

const supportMsg = `We would like to continue developing this bot and other products that improve the Cosmos ecosystem.

You can support us by staking to my validator service <b>DeCrypto</b> on <a href='https://ping.pub/dig/staking/digvaloper1fhp54fwlfmpwwgrnfwk3v47v53yjtp8fw6nelw'>Dig</a>.

You have a good idea, feedback or want to contribute in other ways? Shoot a message to @rapha_decrypto`

const newChainsMsg = `Select the chains that should be enabled for everyone`

const errMsg = `There was an unknown error`
