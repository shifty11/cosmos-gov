package telegram

const NbrOfButtonsPerRow = 3

const menuInfoMsg = "Select the projects that you want to follow. You will receive notifications about new governance proposals once they enter the voting period."
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
      Chain ┆ Subscriptions
 -----------┼---------------`
const chainStatisticRowMsg = `
%11.11s ┆ %6d`
const chainStatisticFooterMsg = `
 -----------┼---------------
  %s ┆ %6d`
const userStatisticMsg = `
       Time ┆   Users (change)
 -----------┼------------------
   All time ┆ %7d
     1 week ┆ %7d %+7.2f%%
      1 day ┆ %7d %+7.2f%%`

const showProposalMessage = "This are all ongoing proposals for your subscriptions.\n\n"
