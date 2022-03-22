package telegram

const NbrOfButtonsPerRow = 3

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
       Time ┆  Users (%v)
 -----------┼------------------
   All time ┆ %7d
     1 week ┆ %7d %+7.2f%%
      1 day ┆ %7d %+7.2f%%`

const helpMsg = `<b>Commands List</b>
/subscriptions - Manage your subscriptions
/proposals - Show proposals in voting period
/support - Show how you can support this bot
/help - Show bot commands`

const adminHelpMsg = `<b>Admin Commands</b>
/stats - Show statistics
/chains - Manage chains
/broadcast - Broadcast message to everyone`

const newChainsMsg = `Select the chains that should be enabled for everyone`

const errMsg = `There was an unknown error`
